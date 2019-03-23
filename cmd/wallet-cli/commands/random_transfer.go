package commands

import (
	"context"
	"fmt"
	"github.com/coschain/cobra"
	"github.com/coschain/contentos-go/cmd/wallet-cli/commands/utils"
	"github.com/coschain/contentos-go/cmd/wallet-cli/wallet"
	"github.com/coschain/contentos-go/prototype"
	"github.com/coschain/contentos-go/rpc/pb"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const (
	randAccounts = 10000
)

var RandomTransferCmd = func() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "randtran",
		Short:   fmt.Sprintf("create %d random accounts and transfer randomly among them", randAccounts),
		Example: "randtran [creator] [#threads]",
		Args:    cobra.ExactArgs(2),
		Run:     randomTransfer,
	}
	return cmd
}

func randomTransfer(cmd *cobra.Command, args []string) {
	var (
		ok bool
	)
	rand.Seed(time.Now().UnixNano())
	rt := new(RandTransfer)
	rt.client = cmd.Context["rpcclient"].(grpcpb.ApiServiceClient)
	rt.wallet = cmd.Context["wallet"].(wallet.Wallet)
	rt.creator, ok = rt.wallet.GetUnlockedAccount(args[0])
	if !ok {
		fmt.Println(fmt.Sprintf("creator: %s should be loaded or created first", args[0]))
		return
	}
	rt.threads, _ = strconv.Atoi(args[1])
	if rt.threads < 1 {
		rt.threads = 1
	}
	rt.prefix = randStr(5)

	rt.do()
}

func randStr(size int) string {
	chars := make([]byte, size)
	for i := range chars {
		chars[i] = byte(97 + rand.Intn(26))
	}
	return string(chars)
}

type RandTransfer struct {
	client grpcpb.ApiServiceClient
	wallet wallet.Wallet
	creator *wallet.PrivAccount
	threads int
	prefix string
}

func (rt *RandTransfer) do() {
	var (
		trxs []*prototype.SignedTransaction
		keys = make(map[string]*prototype.PrivateKeyType)
		wg sync.WaitGroup
	)

	fmt.Printf("create/fund %d random accounts: %s%d - %s%d\n", randAccounts, rt.prefix, 0, rt.prefix, randAccounts-1)
	groupSize := 100
	for i := 0; i < randAccounts; i+=groupSize {
		d := randAccounts - i
		if d > groupSize {
			d = groupSize
		}
		trxs = trxs[:0]
		for j := 0; j < d; j++ {
			name := fmt.Sprintf("%s%d", rt.prefix, i + j)
			err, tx, pk := rt.createAndFundAccount(name)
			if err != nil {
				fmt.Printf("failed generating trx for create/fund %s: %s\n", name, err.Error())
				return
			}
			trxs = append(trxs, tx)
			keys[name] = pk
		}
		if !rt.waitTrxs(trxs) {
			fmt.Println("failed create/fund accounts")
			return
		}
	}

	fmt.Printf("random transfers with %d threads\n", rt.threads)
	wg.Add(rt.threads)
	for i := 0; i < rt.threads; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				a, b := 0, 0
				for a == b {
					a, b = rand.Intn(randAccounts), rand.Intn(randAccounts)
				}
				from := fmt.Sprintf("%s%d", rt.prefix, a)
				to := fmt.Sprintf("%s%d", rt.prefix, b)
				err, tx := rt.transfer(from, to, keys[from])
				if err != nil {
					fmt.Printf("failed generating transfer trx %s -> %s: %s\n", from, to, err.Error())
					break
				}
				res, err := rt.client.BroadcastTrx(context.Background(), &grpcpb.BroadcastTrxRequest{
					Transaction: tx,
					OnlyDeliver: true,
				})
				if err != nil {
					fmt.Printf("%s->%s: %v\n", from, to, err.Error())
					break
				}
				if res.Status != prototype.StatusSuccess {
					fmt.Printf("%s->%s: %v\n", from, to, res)
					break
				}
			}
		}()
	}
	wg.Wait()
}

func (rt *RandTransfer) waitTrxs(trxs []*prototype.SignedTransaction) bool {
	success := int64(1)
	var wg sync.WaitGroup
	wg.Add(len(trxs))
	for i := range trxs {
		go func(idx int) {
			defer wg.Done()

			s := fmt.Sprintf("error of trx #%d", idx)
			res, err := rt.client.BroadcastTrx(context.Background(), &grpcpb.BroadcastTrxRequest{
				Transaction: trxs[idx],
				OnlyDeliver: false,
			})
			if err != nil {
				atomic.StoreInt64(&success, 0)
				fmt.Printf("%s: %s\n", s, err.Error())
			}
			if res.Invoice == nil {
				atomic.StoreInt64(&success, 0)
				fmt.Printf("%s: invoice is nil\n", s)
			}
			if res.Invoice.Status != prototype.StatusSuccess {
				atomic.StoreInt64(&success, 0)
				fmt.Printf("%s: %v\n", s, res)
			}
		}(i)
	}
	wg.Wait()
	return success != 0
}

func (rt *RandTransfer) createAndFundAccount(account string) (error, *prototype.SignedTransaction, *prototype.PrivateKeyType) {
	prvKey, _ := prototype.GenerateNewKey()
	pubKey, _ := prvKey.PubKey()
	opCreate := &prototype.AccountCreateOperation{
		Fee:            prototype.NewCoin(10000),
		Creator:        &prototype.AccountName{Value: rt.creator.Name},
		NewAccountName: &prototype.AccountName{Value: account},
		Owner:          prototype.NewAuthorityFromPubKey(pubKey),
	}
	opFund := &prototype.TransferOperation{
		From:   &prototype.AccountName{Value: rt.creator.Name},
		To:     &prototype.AccountName{Value: account},
		Amount: prototype.NewCoin(50000),
		Memo:   "",
	}
	trx, err := utils.GenerateSignedTxAndValidate2(rt.client, []interface{}{opCreate, opFund}, rt.creator)
	if err != nil {
		return err, nil, nil
	}
	return nil, trx, prvKey
}

func (rt *RandTransfer) transfer(from, to string, fromKey *prototype.PrivateKeyType) (error, *prototype.SignedTransaction) {
	op := &prototype.TransferOperation{
		From:   &prototype.AccountName{Value: from},
		To:     &prototype.AccountName{Value: to},
		Amount: prototype.NewCoin(1),
		Memo:   randStr(8),
	}
	trx, err := utils.GenerateSignedTxAndValidate3(rt.client, []interface{}{op}, fromKey)
	if err != nil {
		return err, nil
	}
	return nil, trx
}
