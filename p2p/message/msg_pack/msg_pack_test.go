package msgpack

import (
	"testing"

	"github.com/coschain/contentos-go/p2p/common"
	msgTypes "github.com/coschain/contentos-go/p2p/message/types"
	"github.com/coschain/contentos-go/prototype"
	"github.com/stretchr/testify/assert"
)

var msgBefore msgTypes.Message
var msgAfter *msgTypes.TransferMsg
var sink *common.ZeroCopySink
var source *common.ZeroCopySource
var err error

func TestMsgPack(t *testing.T) {
	var addrStr []*msgTypes.PeerAddr
	msgBefore = NewAddrs(addrStr)
	processAndCheck(t)

	msgBefore = NewAddrReq()
	processAndCheck(t)

	sigBlk := new(prototype.SignedBlock)
	sigBlkHdr := new(prototype.SignedBlockHeader)
	sigBlkHdr.Header = new(prototype.BlockHeader)
	sigBlkHdr.Header.Witness = new(prototype.AccountName)
	sigBlkHdr.Header.Witness.Value = "alice"
	sigBlk.SignedHeader = sigBlkHdr
	msgBefore = NewSigBlkIdMsg(sigBlk)
	processAndCheck(t)

	msgBefore = NewSigBlk(sigBlk)
	processAndCheck(t)

	msgBefore = NewPingMsg(0)
	processAndCheck(t)

	msgBefore = NewPongMsg(0)
	processAndCheck(t)

	trx := &prototype.Transaction{
		RefBlockNum:    1,
		RefBlockPrefix: 2,
	}
	sigtrx := new(prototype.SignedTransaction)
	sigtrx.Trx = trx
	msgBefore = NewTxn(sigtrx)
	processAndCheck(t)

	msgBefore = NewVerAck(true)
	processAndCheck(t)
}

func processAndCheck(t *testing.T) {
	sink = common.NewZeroCopySink(nil)
	err = msgBefore.Serialization(sink)
	assert.Nil(t, err)
	source = common.NewZeroCopySource(sink.Bytes())
	msgAfter = new(msgTypes.TransferMsg)
	err = msgAfter.Deserialization(source)
	assert.Nil(t, err)
}