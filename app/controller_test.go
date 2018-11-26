package app

import (
	"errors"
	"fmt"
	"github.com/coschain/contentos-go/app/table"
	"github.com/coschain/contentos-go/common"
	"github.com/coschain/contentos-go/common/constants"
	"github.com/coschain/contentos-go/prototype"
	"github.com/golang/protobuf/proto"
	"testing"
	"time"
)

const (
	accountName = "bob"
	pubKey = "COS6oLVaFEtHZmPDuCvuB48NpSKytjyavPk5MwtN4HqKG16oSA2wS"
	priKey = "EpgwWxboEdaWfEBdWswobsBt8pBF6xoYQPayBs4eVysMGGGYL"

)


func createSigTrx(op interface{}) (*prototype.SignedTransaction,error) {

	privKey, err := prototype.PrivateKeyFromWIF(constants.INITMINER_PRIKEY)
	if err != nil {
		return nil, err
	}

	tx := &prototype.Transaction{RefBlockNum: 0, RefBlockPrefix: 0,
	Expiration: &prototype.TimePointSec{UtcSeconds: uint32(time.Now().Second()+20)}}
	tx.AddOperation(op)

	signTx := prototype.SignedTransaction{Trx: tx}

	res := signTx.Sign(privKey, prototype.ChainId{Value: 0})
	signTx.Signatures = append(signTx.Signatures, &prototype.SignatureType{Sig: res})

	return &signTx, nil
}

func makeCreateAccountOP() (*prototype.AccountCreateOperation,error) {
	pub,err := prototype.PublicKeyFromWIF(pubKey)
	if err != nil {
		return nil,errors.New("PublicKeyFromWIF error")
	}
	acop := &prototype.AccountCreateOperation{
		Fee:            prototype.NewCoin(1),
		Creator:        &prototype.AccountName{Value: constants.COS_INIT_MINER},
		NewAccountName: &prototype.AccountName{Value: accountName},
		Owner: &prototype.Authority{
			Cf:              prototype.Authority_owner,
			WeightThreshold: 1,
			AccountAuths: []*prototype.KvAccountAuth{
				&prototype.KvAccountAuth{
					Name:    &prototype.AccountName{Value: constants.COS_INIT_MINER},
					Weight: 3,
				},
			},
			KeyAuths: []*prototype.KvKeyAuth{
				&prototype.KvKeyAuth{
					Key: pub,	// owner key
					Weight: 1,
				},
			},
		},
		Active: &prototype.Authority{
		},
		Posting: &prototype.Authority{
		},
	}

	return acop,nil
}

func Test_PushTrx(t *testing.T) {
	clearDB()

	acop,err := makeCreateAccountOP()
	if err != nil {
		t.Error("makeCreateAccountOP error:",err)
	}

	signedTrx, err := createSigTrx(acop)
	if err != nil {
		t.Error("createSigTrx error:",err)
	}

	// set up controller
	db := startDB()
	defer db.Close()
	c := startController(db)

	invoice := c.PushTrx(signedTrx)
	if invoice.Status != 200 {
		t.Error("PushTrx return status error:",invoice.Status)
	}

	bobName := &prototype.AccountName{Value:accountName}
	bobWrap := table.NewSoAccountWrap(db,bobName)
	if !bobWrap.CheckExist() {
		t.Error("create account failed")
	}
}

func Test_PushBlock(t *testing.T) {
	clearDB()

	createOP,err := makeCreateAccountOP()
	if err != nil {
		t.Error("makeCreateAccountOP error:",err)
	}
	signedTrx,err := createSigTrx(createOP)
	if err != nil {
		t.Error("createSigTrx error:",err)
	}

	// set up controller
	db := startDB()
	defer db.Close()
	c := startController(db)

	sigBlk := new(prototype.SignedBlock)

	// add trx wraper
	trxWraper := &prototype.TransactionWrapper{
		SigTrx:signedTrx,
		Invoice:&prototype.TransactionInvoice{Status:200},
	}
	sigBlk.Transactions = append(sigBlk.Transactions,trxWraper)

	// calculate merkle
	id := sigBlk.CalculateMerkleRoot()

	// write signed block header
	sigBlkHdr := new(prototype.SignedBlockHeader)

	sigBlkHdr.Header = new(prototype.BlockHeader)
	sigBlkHdr.Header.Previous = c.dgpo.GetHeadBlockId()
	sigBlkHdr.Header.Timestamp = &prototype.TimePointSec{UtcSeconds:20}
	sigBlkHdr.Header.Witness = &prototype.AccountName{Value:constants.INIT_MINER_NAME}
	sigBlkHdr.Header.TransactionMerkleRoot = &prototype.Sha256{Hash:id.Data[:]}
	sigBlkHdr.WitnessSignature = &prototype.SignatureType{}

	// signature
	pri,err := prototype.PrivateKeyFromWIF(constants.INITMINER_PRIKEY)
	if err != nil {
		t.Error("PrivateKeyFromWIF error")
	}
	sigBlkHdr.Sign(pri)

	sigBlk.SignedHeader = sigBlkHdr

	fmt.Println("block size:",proto.Size(sigBlk))

	c.PushBlock(sigBlk,prototype.Skip_nothing)

	bobName := &prototype.AccountName{Value:accountName}
	bobWrap := table.NewSoAccountWrap(db,bobName)
	if !bobWrap.CheckExist() {
		t.Error("create account failed")
	}
}

func TestController_GenerateBlock(t *testing.T) {
	clearDB()
	createOP,err := makeCreateAccountOP()
	if err != nil {
		t.Error("makeCreateAccountOP error:",err)
	}
	signedTrx,err := createSigTrx(createOP)
	if err != nil {
		t.Error("createSigTrx error:",err)
	}

	// set up controller
	db := startDB()
	defer db.Close()
	c := startController(db)

	// set reference
	id := &common.BlockID{}
	sha256ID := c.dgpo.GetHeadBlockId()
	copy(id.Data[:],sha256ID.Hash[:])
	signedTrx.Trx.SetReferenceBlock(id)

	invoice := c.PushTrx(signedTrx)
	if invoice.Status != 200 {
		t.Error("PushTrx return status error:",invoice.Status)
	}

	bobName := &prototype.AccountName{Value:accountName}
	bobWrap := table.NewSoAccountWrap(db,bobName)
	if !bobWrap.CheckExist() {
		t.Error("create account failed")
	}

	pri,err := prototype.PrivateKeyFromWIF(constants.INITMINER_PRIKEY)
	if err != nil {
		t.Error("PrivateKeyFromWIF error")
	}

	c.GenerateBlock(constants.INIT_MINER_NAME,18,pri,0)
}

func Test_list(t *testing.T) {
	clearDB()

	// set up controller
	db := startDB()
	defer db.Close()

	// make trx
	acop,err := makeCreateAccountOP()
	if err != nil {
		t.Error("makeCreateAccountOP error:",err)
	}

	signedTrx, err := createSigTrx(acop)
	if err != nil {
		t.Error("createSigTrx error:",err)
	}
	id,err := signedTrx.Id()

	// insert trx into DB unique table
	transactionObjWrap := table.NewSoTransactionObjectWrap(db, id)
	if transactionObjWrap.CheckExist() {
		panic("Duplicate transaction check failed")
	}

	cErr := transactionObjWrap.Create(func(tInfo *table.SoTransactionObject) {
		tInfo.TrxId = id
		tInfo.Expiration = &prototype.TimePointSec{UtcSeconds: 100}
	})
	if cErr != nil {
		panic("create transactionObject failed")
	}

	// check and delete

	sortWrap := table.STransactionObjectExpirationWrap{Dba: db}
	itr := sortWrap.QueryListByOrder(nil, nil) // query all
	if itr != nil {
		for itr.Next() {

			subPtr := sortWrap.GetSubVal(itr)
			if subPtr == nil {
				break
			}

			k := sortWrap.GetMainVal(itr)
			objWrap := table.NewSoTransactionObjectWrap(db, k)
			if !objWrap.RemoveTransactionObject() {
				panic("RemoveTransactionObject error")
			}

		}
		sortWrap.DelIterater(itr)
	}
}

func TestController_GetWitnessTopN(t *testing.T) {
	clearDB()

	// set up controller
	db := startDB()
	defer db.Close()
	c := startController(db)

	name := &prototype.AccountName{Value:"wit1"}
	witnessWrap := table.NewSoWitnessWrap(db,name)
	mustNoError(witnessWrap.Create(func(tInfo *table.SoWitness) {
		tInfo.Owner = name
		tInfo.WitnessScheduleType = &prototype.WitnessScheduleType{Value: prototype.WitnessScheduleType_miner}
		tInfo.CreatedTime = &prototype.TimePointSec{UtcSeconds: 0}
		tInfo.SigningKey = &prototype.PublicKeyType{Data:[]byte{1}}
		tInfo.LastWork = &prototype.Sha256{Hash: []byte{0}}
	}), "Witness Create Error")

	name2 := &prototype.AccountName{Value:"wit2"}
	witnessWrap2 := table.NewSoWitnessWrap(db,name2)
	mustNoError(witnessWrap2.Create(func(tInfo *table.SoWitness) {
		tInfo.Owner = name2
		tInfo.WitnessScheduleType = &prototype.WitnessScheduleType{Value: prototype.WitnessScheduleType_miner}
		tInfo.CreatedTime = &prototype.TimePointSec{UtcSeconds: 0}
		tInfo.SigningKey = &prototype.PublicKeyType{Data:[]byte{2}}
		tInfo.LastWork = &prototype.Sha256{Hash: []byte{0}}
	}), "Witness Create Error")

	witnesses := c.GetWitnessTopN(10)

	for _,wit := range witnesses {
		fmt.Println(wit)
	}
}
