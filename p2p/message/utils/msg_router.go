package utils

import (
	"github.com/ontio/ontology-eventbus/actor"
	"github.com/coschain/contentos-go/p2p/depend/common/log"
	msgCommon "github.com/coschain/contentos-go/p2p/common"
	"github.com/coschain/contentos-go/p2p/message/types"
	"github.com/coschain/contentos-go/p2p/net/protocol"
)

// MessageHandler defines the unified api for each net message
type MessageHandler func(data *types.MsgPayload, p2p p2p.P2P, pid *actor.PID, args ...interface{})

// MessageRouter mostly route different message type-based to the
// related message handler
type MessageRouter struct {
	msgHandlers  map[string]MessageHandler // Msg handler mapped to msg type
	RecvSyncChan chan *types.MsgPayload    // The channel to handle sync msg
	RecvConsChan chan *types.MsgPayload    // The channel to handle consensus msg
	stopSyncCh   chan bool                 // To stop sync channel
	stopConsCh   chan bool                 // To stop consensus channel
	p2p          p2p.P2P                   // Refer to the p2p network
	pid          *actor.PID                // P2P actor
}

// NewMsgRouter returns a message router object
func NewMsgRouter(p2p p2p.P2P) *MessageRouter {
	msgRouter := &MessageRouter{}
	msgRouter.init(p2p)
	return msgRouter
}

// init initializes the message router's attributes
func (this *MessageRouter) init(p2p p2p.P2P) {
	this.msgHandlers = make(map[string]MessageHandler)
	this.RecvSyncChan = p2p.GetMsgChan(false)
	this.RecvConsChan = p2p.GetMsgChan(true)
	this.stopSyncCh = make(chan bool)
	this.stopConsCh = make(chan bool)
	this.p2p = p2p

	// Register message handler
	this.RegisterMsgHandler(msgCommon.VERSION_TYPE, VersionHandle)
	this.RegisterMsgHandler(msgCommon.VERACK_TYPE, VerAckHandle)
	this.RegisterMsgHandler(msgCommon.GetADDR_TYPE, AddrReqHandle)
	this.RegisterMsgHandler(msgCommon.ADDR_TYPE, AddrHandle)
	this.RegisterMsgHandler(msgCommon.PING_TYPE, PingHandle)
	this.RegisterMsgHandler(msgCommon.PONG_TYPE, PongHandle)
	this.RegisterMsgHandler(msgCommon.HASH_TYPE, HashMsgHandle)
	this.RegisterMsgHandler(msgCommon.BLOCK_TYPE, BlockHandle)
	this.RegisterMsgHandler(msgCommon.NOT_FOUND_TYPE, NotFoundHandle)
	this.RegisterMsgHandler(msgCommon.TX_TYPE, TransactionHandle)
	this.RegisterMsgHandler(msgCommon.DISCONNECT_TYPE, DisconnectHandle)
}

// RegisterMsgHandler registers msg handler with the msg type
func (this *MessageRouter) RegisterMsgHandler(key string,
	handler MessageHandler) {
	this.msgHandlers[key] = handler
}

// UnRegisterMsgHandler un-registers the msg handler with
// the msg type
func (this *MessageRouter) UnRegisterMsgHandler(key string) {
	delete(this.msgHandlers, key)
}

// SetPID sets p2p actor
func (this *MessageRouter) SetPID(pid *actor.PID) {
	this.pid = pid
}

// Start starts the loop to handle the message from the network
func (this *MessageRouter) Start() {
	go this.hookChan(this.RecvSyncChan, this.stopSyncCh)
	go this.hookChan(this.RecvConsChan, this.stopConsCh)
	log.Debug("[p2p]MessageRouter start to parse p2p message...")
}

// hookChan loops to handle the message from the network
func (this *MessageRouter) hookChan(channel chan *types.MsgPayload,
	stopCh chan bool) {
	for {
		select {
		case data, ok := <-channel:
			if ok {
				msgType := data.Payload.CmdType()

				handler, ok := this.msgHandlers[msgType]
				if ok {
					go handler(data, this.p2p, this.pid)
				} else {
					log.Warn("unknown message handler for the msg: ",
						msgType)
				}
			}
		case <-stopCh:
			return
		}
	}
}

// Stop stops the message router's loop
func (this *MessageRouter) Stop() {

	if this.stopSyncCh != nil {
		this.stopSyncCh <- true
	}
	if this.stopConsCh != nil {
		this.stopConsCh <- true
	}
}
