package commands


import (
	"context"
	"fmt"
	"github.com/coschain/cobra"
	"github.com/coschain/contentos-go/cmd/wallet-cli/commands/utils"
	"github.com/coschain/contentos-go/cmd/wallet-cli/wallet"
	"github.com/coschain/contentos-go/prototype"
	"github.com/coschain/contentos-go/rpc/pb"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var CntIdx uint64 = 0
var StressCreAccountCmd = func() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "stressaccount",
		Short:   "create a new random account ",
		Long:    "use thread-count thread to stress test create new random account by a exist creator(initminer)",
		Example: "stressaccount 2 initminer",
		Args:    cobra.MinimumNArgs(2),
		Run:     stressCreAccount,
	}
	return cmd
}

func stressCreAccount(cmd *cobra.Command, args []string) {
	c := cmd.Context["rpcclient"]
	client := c.(grpcpb.ApiServiceClient)
	w := cmd.Context["wallet"]
	mywallet := w.(wallet.Wallet)
	creator := args[1]
	tCount, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	creatorAccount, ok := mywallet.GetUnlockedAccount(creator)
	if !ok {
		fmt.Println(fmt.Sprintf("account: %s should be loaded or created first", creator))
		return
	}
	wg := &sync.WaitGroup{}
	for i := 0; i < tCount; i++  {
		tid := i
		wg.Add(1)
		go func(){
			s := time.Now()
			defer func() {
				fmt.Println("stress create account cost: ", time.Now().Sub(s), ", thread-number: ", tid)
			}()
			for index:=0; index < 1000; index ++ {
				pubKeyStr, _, err := mywallet.GenerateNewKey()
				pubkey, err := prototype.PublicKeyFromWIF(pubKeyStr)
				if err != nil {
					fmt.Println(err)
					break
				}
				keys := prototype.NewAuthorityFromPubKey(pubkey)
				newAccountName := getNewAccountName()
				acop := &prototype.AccountCreateOperation{
					Fee:            prototype.NewCoin(1),
					Creator:        &prototype.AccountName{Value: creator},
					NewAccountName: &prototype.AccountName{Value: newAccountName},
					Owner:          keys,
				}
				signTx, err := utils.GenerateSignedTxAndValidate([]interface{}{acop}, creatorAccount)
				if err != nil {
					fmt.Println(err)
					return
				}

				req := &grpcpb.BroadcastTrxRequest{Transaction: signTx}
				res, err := client.BroadcastTrx(context.Background(), req)

				if err != nil {
					fmt.Println(err)
					break
				}

				if !res.Invoice.IsSuccess() {
					fmt.Println(res)
					break
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func getNewAccountName() string {
	idx := atomic.AddUint64(&CntIdx,2)
	atomic.StoreUint64(&CntIdx,idx)
	name := "account" + strconv.FormatUint(idx,10)
	return name
}