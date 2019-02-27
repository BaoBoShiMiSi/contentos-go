package economist

import (
	"github.com/coschain/contentos-go/app/table"
	"github.com/coschain/contentos-go/common/constants"
	"github.com/coschain/contentos-go/iservices"
	"github.com/coschain/contentos-go/node"
	"github.com/coschain/contentos-go/prototype"
	"github.com/pkg/errors"
	"math"
)

var (
	// fixme: the single id should be share with service
	SINGLE_ID int32 = 1
)

type Economist struct {
	ctx               *node.ServiceContext
	db                iservices.IDatabaseService
	rewardAccumulator uint64 // reward accumulator
	vpAccumulator     uint64 // vote power accumulator
	globalProps       *prototype.DynamicProperties
	rewardsKeeper     *prototype.InternalRewardsKeeper
}

func (e *Economist) getDb() (iservices.IDatabaseService, error) {
	s, err := e.ctx.Service(iservices.DbServerName)
	if err != nil {
		return nil, err
	}
	db := s.(iservices.IDatabaseService)
	return db, nil
}

func New(ctx *node.ServiceContext) (*Economist, error) {

	return &Economist{ctx: ctx}, nil
}

func (e *Economist) Start(node *node.Node) error {
	db, err := e.getDb()
	if err != nil {
		return errors.New("Economist fetch db service error")
	}
	e.db = db
	dgpWrap := table.NewSoGlobalWrap(e.db, &SINGLE_ID)
	if !dgpWrap.CheckExist() {
		return errors.New("the mainkey is already exist")
	}
	e.globalProps = dgpWrap.GetProps()

	keeperWrap := table.NewSoRewardsKeeperWrap(e.db, &SINGLE_ID)
	if !keeperWrap.CheckExist() {
		return errors.New("Economist access rewards keeper error")
	}
	e.rewardsKeeper = keeperWrap.GetKeeper()
	return nil
}

func (e *Economist) Stop() error {
	return nil
}

func (e *Economist) updateRewardsKeeper() error {
	keeper := table.NewSoRewardsKeeperWrap(e.db, &SINGLE_ID)
	success := keeper.MdKeeper(e.rewardsKeeper)
	if !success {
		return errors.New("flush rewards keeper into db error")
	}
	return nil
}

func (e *Economist) Mint() error {
	blockCurrent := constants.PerBlockCurrent

	authorReward := blockCurrent * constants.RewardRateAuthor / constants.PERCENT
	replyReward := blockCurrent * constants.RewardRateAuthor / constants.PERCENT
	bpReward := blockCurrent * constants.RewardRateBP / constants.PERCENT

	e.globalProps.PostRewards.Value += uint64(authorReward)
	e.globalProps.ReplyRewards.Value += uint64(replyReward)

	_ = bpReward
	currentBp := e.globalProps.GetCurrentWitness().Value
	rewards := e.rewardsKeeper.GetRewards()
	if vest, ok := rewards[currentBp]; !ok {
		rewards[currentBp] = &prototype.Vest{Value: uint64(bpReward)}
	} else {
		vest.Value += uint64(bpReward)
	}

	return nil
}

//func (e *Economist) Do() error {
//	e.decayGlobalVotePower()
//	timestamp := e.globalProps.Time.UtcSeconds - uint32(constants.GenesisTime)
//	keyPrefix := fmt.Sprintf("cashout:%d_", common.GetBucket(timestamp))
//	postCashoutList := []string{}
//	replyCashoutList := []string{}
//	r := regexp.MustCompile(`cashout:(?P<bucket>\d+)_(?P<idx>\d+)`)
//	iter := e.db.NewIterator([]byte(keyPrefix), nil)
//	for iter.Next() {
//		if !iter.Valid() {
//			break
//		}
//		key, err := iter.Key()
//		if err != nil {
//			return err
//		}
//		value, err := iter.Value()
//		if err != nil {
//			return err
//		}
//		match := r.FindStringSubmatch(string(key))
//		if len(match) > 0 {
//			idx := match[2]
//			switch string(value) {
//			case "post":
//				postCashoutList = append(postCashoutList, idx)
//			case "reply":
//				replyCashoutList = append(replyCashoutList, idx)
//			}
//		}
//	}
//	if len(postCashoutList) > 0 {
//		e.postCashout(postCashoutList)
//	}
//
//	if len(replyCashoutList) > 0 {
//		e.replyCashout(replyCashoutList)
//	}
//
//	err := e.updateRewardsKeeper()
//	return err
//}

func (e *Economist) Do() error {
	e.decayGlobalVotePower()
	timestamp := e.globalProps.Time.UtcSeconds
	iterator := table.NewPostCashoutTimeWrap(e.db)
	var pids []*uint64
	err := iterator.ForEachByOrder(nil, &prototype.TimePointSec{UtcSeconds: timestamp}, func(mVal *uint64, sVal *prototype.TimePointSec, idx uint32) bool {
		pids = append(pids, mVal)
		return true
	})
	if err != nil {
		return err
	}
	var posts []*table.SoPostWrap
	var replies []*table.SoPostWrap
	for _, pid := range pids {
		post := table.NewSoPostWrap(e.db, pid)
		if post.GetParentId() == 0 {
			posts = append(posts, post)
		} else {
			replies = append(replies, post)
		}
	}
	if len(posts) > 0 {
		e.postCashout(posts)
	}

	if len(replies) > 0 {
		e.replyCashout(replies)
	}
	err = e.updateRewardsKeeper()
	return err
}

func (e *Economist) decayGlobalVotePower() {
	e.globalProps.WeightedVps -= e.globalProps.WeightedVps * constants.BlockInterval / constants.VpDecayTime
}

func (e *Economist) postCashout(posts []*table.SoPostWrap) {
	var vpAccumulator uint64 = 0
	for _, post := range posts {
		vpAccumulator += post.GetWeightedVp()
	}
	blockReward := vpAccumulator * e.globalProps.PostRewards.Value / e.globalProps.WeightedVps
	innerRewards := e.rewardsKeeper.Rewards
	for _, post := range posts {
		author := post.GetAuthor().Value
		reward := post.GetWeightedVp() * blockReward / vpAccumulator
		if vest, ok := innerRewards[author]; !ok {
			innerRewards[author] = &prototype.Vest{Value: reward}
		} else {
			vest.Value += reward
		}
		post.MdCashoutTime(&prototype.TimePointSec{UtcSeconds: math.MaxUint32})
	}
}

// use same algorithm to simplify
func (e *Economist) replyCashout(replies []*table.SoPostWrap) {
	var vpAccumulator uint64 = 0
	for _, reply := range replies {
		vpAccumulator += reply.GetWeightedVp()
	}
	blockReward := vpAccumulator * e.globalProps.ReplyRewards.Value / e.globalProps.WeightedVps
	innerRewards := e.rewardsKeeper.Rewards
	for _, reply := range replies {
		author := reply.GetAuthor().Value
		reward := reply.GetWeightedVp() * blockReward / vpAccumulator
		if vest, ok := e.rewardsKeeper.Rewards[author]; !ok {
			innerRewards[author] = &prototype.Vest{Value: reward}
		} else {
			vest.Value += reward
		}
		reply.MdCashoutTime(&prototype.TimePointSec{UtcSeconds: math.MaxUint32})
	}
}
