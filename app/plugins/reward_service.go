package plugins

import (
	"github.com/asaskevich/EventBus"
	"github.com/coschain/contentos-go/app/table"
	"github.com/coschain/contentos-go/iservices"
	"github.com/coschain/contentos-go/node"
	"github.com/coschain/contentos-go/prototype"
)

var RewardServiceName = "rewardsrv"

type RewardService struct {
	node.Service
	db  iservices.IDatabaseService
	ev  EventBus.Bus
	pool iservices.ITrxPool
	ctx *node.ServiceContext
}

// service constructor
func NewRewardService(ctx *node.ServiceContext) (*RewardService, error) {
	return &RewardService{ctx:ctx}, nil
}

func (p *RewardService) Start(node *node.Node) error {
	db, err := p.ctx.Service(iservices.DbServerName)
	if err != nil {
		return err
	}
	p.db = db.(iservices.IDatabaseService)

	pool, err := p.ctx.Service(iservices.TxPoolServerName)
	if err != nil {
		return err
	}
	p.pool = pool.(iservices.ITrxPool)
	p.ev = node.EvBus

	p.hookEvent()
	return nil
}

func (p *RewardService) hookEvent() {
	p.ev.Subscribe("rewards", p.onReward)
}

// author, realReward, globalProps.GetHeadBlockNumber(), globalProps.GetTime())
func (p *RewardService) unhookEvent() {
	p.ev.Unsubscribe("rewards", p.onReward)
}

func (p *RewardService) onReward(name string, reward uint64, blockHeight uint64) {
	exRewardWrap := table.NewSoExtCashoutWrap(p.db, &prototype.RewardCashoutId{Account:&prototype.AccountName{Value:name}, BlockHeight:blockHeight})
	if exRewardWrap != nil {
		if !exRewardWrap.CheckExist() {
			_ = exRewardWrap.Create(func(tInfo *table.SoExtCashout) {
				tInfo.Id = &prototype.RewardCashoutId{
					Account:     &prototype.AccountName{Value: name},
					BlockHeight: blockHeight,
				}
				tInfo.BlockHeight = blockHeight
				tInfo.Reward = &prototype.Vest{Value: reward}
			})
		} else {
			r := exRewardWrap.GetReward()
			newReward := &prototype.Vest{Value: reward + r.Value}
			exRewardWrap.MdReward(newReward)
		}
	}
}

func (p *RewardService) Stop() error {
	p.unhookEvent()
	return nil
}
