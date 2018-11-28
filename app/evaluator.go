package app

import (
	"github.com/coschain/contentos-go/app/table"
	"github.com/coschain/contentos-go/common"
	"github.com/coschain/contentos-go/common/constants"
	"github.com/coschain/contentos-go/prototype"
	"github.com/pkg/errors"
	"time"
)

func mustSuccess(b bool, val string) {
	if !b {
		panic(val)
	}
}

type AccountCreateEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.AccountCreateOperation
}

type TransferEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.TransferOperation
}

type PostEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.PostOperation
}
type ReplyEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.ReplyOperation
}
type VoteEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.VoteOperation
}
type BpRegisterEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.BpRegisterOperation
}
type BpUnregisterEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.BpUnregisterOperation
}

type BpVoteEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.BpVoteOperation
}

type FollowEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.FollowOperation
}

type TransferToVestingEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.TransferToVestingOperation
}

func (ev *AccountCreateEvaluator) Apply() {
	op := ev.op
	creatorWrap := table.NewSoAccountWrap(ev.ctx.db, op.Creator)

	opAssert(creatorWrap.CheckExist(), "creator not exist ")

	opAssert(creatorWrap.GetBalance().Value >= op.Fee.Value, "Insufficient balance to create account.")

	// check auth accounts
	for _, a := range op.Owner.AccountAuths {
		tmpAccountWrap := table.NewSoAccountWrap(ev.ctx.db, a.Name)
		opAssert(tmpAccountWrap.CheckExist(), "owner auth account not exist")
	}
	for _, a := range op.Active.AccountAuths {
		tmpAccountWrap := table.NewSoAccountWrap(ev.ctx.db, a.Name)
		opAssert(tmpAccountWrap.CheckExist(), "active auth account not exist")
	}
	for _, a := range op.Posting.AccountAuths {
		tmpAccountWrap := table.NewSoAccountWrap(ev.ctx.db, a.Name)
		opAssert(tmpAccountWrap.CheckExist(), "posting auth account not exist")
	}

	// sub creator's fee
	originBalance := creatorWrap.GetBalance()
	opAssertE(originBalance.Sub(op.Fee), "creator balance overflow")
	opAssert(creatorWrap.MdBalance(originBalance), "")

	// create account
	newAccountWrap := table.NewSoAccountWrap(ev.ctx.db, op.NewAccountName)
	opAssertE(newAccountWrap.Create(func(tInfo *table.SoAccount) {
		tInfo.Name = op.NewAccountName
		tInfo.Creator = op.Creator
		tInfo.CreatedTime = ev.ctx.control.HeadBlockTime()
		tInfo.Balance = prototype.NewCoin(0)
		tInfo.VestingShares = op.Fee.ToVest()
		tInfo.LastPostTime = &prototype.TimePointSec{UtcSeconds:uint32(time.Now().Second())}
	}), "duplicate create account object")

	// create account authority
	authorityWrap := table.NewSoAccountAuthorityObjectWrap(ev.ctx.db, op.NewAccountName)
	opAssertE(authorityWrap.Create(func(tInfo *table.SoAccountAuthorityObject) {
		tInfo.Account = op.NewAccountName
		tInfo.Posting = op.Posting
		tInfo.Active = op.Active
		tInfo.Owner = op.Owner
		tInfo.LastOwnerUpdate = prototype.NewTimePointSec(0)
	}), "duplicate create account authority object")

	// sub dynamic glaobal properties's total fee
	ev.ctx.control.TransferToVest(op.Fee)
}

func (ev *TransferEvaluator) Apply() {
	op := ev.op

	// @ active_challenged
	fromWrap := table.NewSoAccountWrap(ev.ctx.db, op.From)
	toWrap := table.NewSoAccountWrap(ev.ctx.db, op.To)

	opAssert(toWrap.CheckExist(), "To account do not exist ")

	fBalance := fromWrap.GetBalance()
	tBalance := toWrap.GetBalance()

	opAssertE(fBalance.Sub(op.Amount), "Insufficient balance to transfer.")
	opAssert(fromWrap.MdBalance(fBalance), "")

	opAssertE(tBalance.Add(op.Amount), "balance overflow")
	opAssert(toWrap.MdBalance(tBalance), "")
}

func (ev *PostEvaluator) Apply() {
	op := ev.op
	idWrap := table.NewSoPostWrap(ev.ctx.db, &op.Uuid)
	opAssert(!idWrap.CheckExist(), "post uuid exist")

	authorWrap := table.NewSoAccountWrap(ev.ctx.db, op.Owner)
	elapsedSeconds := ev.ctx.control.HeadBlockTime().UtcSeconds - authorWrap.GetLastPostTime().UtcSeconds
	opAssert(elapsedSeconds > constants.MIN_POST_INTERVAL, "posting frequently")

	opAssertE(idWrap.Create(func(t *table.SoPost) {
		t.PostId = op.Uuid
		t.Tags = op.Tags
		t.Title = op.Title
		t.Author = op.Owner
		t.Body = op.Content
		t.Created = ev.ctx.control.HeadBlockTime()
		t.CashoutTime = &prototype.TimePointSec{UtcSeconds: ev.ctx.control.HeadBlockTime().UtcSeconds + uint32(constants.POST_CASHPUT_DELAY_TIME)}
		t.Depth = 0
		t.Children = 0
		t.RootId = t.PostId
		t.ParentId = 0
		t.RootId = 0
		t.Beneficiaries = op.Beneficiaries
		t.WeightedVp = 0
		t.VoteCnt = 0
	}), "create post error")

	authorWrap.MdLastPostTime(ev.ctx.control.HeadBlockTime())

	timestamp := ev.ctx.control.HeadBlockTime().UtcSeconds + uint32(constants.POST_CASHPUT_DELAY_TIME) - uint32(constants.GenesisTime)
	keyPrefix := "cashout:" + string(common.GetBucket(timestamp)) + "_"
	key := keyPrefix + string(op.Uuid)
	value := "post"
	opAssertE(ev.ctx.db.Put([]byte(key), []byte(value)), "put post key into db error")

}

func (ev *ReplyEvaluator) Apply() {
	op := ev.op
	cidWrap := table.NewSoPostWrap(ev.ctx.db, &op.Uuid)
	pidWrap := table.NewSoPostWrap(ev.ctx.db, &op.ParentUuid)

	opAssert(!cidWrap.CheckExist(), "post uuid exist")
	opAssert(pidWrap.CheckExist(), "parent uuid do not exist")

	opAssert(pidWrap.GetDepth()+1 < constants.POST_MAX_DEPTH, "reply depth error")

	authorWrap := table.NewSoAccountWrap(ev.ctx.db, op.Owner)
	elapsedSeconds := ev.ctx.control.HeadBlockTime().UtcSeconds - authorWrap.GetLastPostTime().UtcSeconds
	opAssert(elapsedSeconds > constants.MIN_POST_INTERVAL, "reply frequently")

	var rootId uint64
	if pidWrap.GetRootId() == 0 {
		rootId = pidWrap.GetPostId()
	} else {
		rootId = pidWrap.GetRootId()
	}

	opAssertE(cidWrap.Create(func(t *table.SoPost) {
		t.PostId = op.Uuid
		t.Tags = nil
		t.Title = ""
		t.Author = op.Owner
		t.Body = op.Content
		t.Created = ev.ctx.control.HeadBlockTime()
		t.CashoutTime = &prototype.TimePointSec{UtcSeconds: ev.ctx.control.HeadBlockTime().UtcSeconds + uint32(constants.POST_CASHPUT_DELAY_TIME)}
		t.Depth = pidWrap.GetDepth() + 1
		t.Children = 0
		t.RootId = rootId
		t.ParentId = op.ParentUuid
		t.VoteCnt = 0
		t.Beneficiaries = op.Beneficiaries
	}), "create reply error")

	authorWrap.MdLastPostTime(ev.ctx.control.HeadBlockTime())
	// Modify Parent Object
	opAssert(pidWrap.MdChildren(pidWrap.GetChildren()+1), "Modify Parent Children Error")

	timestamp := ev.ctx.control.HeadBlockTime().UtcSeconds + uint32(constants.POST_CASHPUT_DELAY_TIME) - uint32(constants.GenesisTime)
	keyPrefix := "cashout:" + string(common.GetBucket(timestamp)) + "_"
	key := keyPrefix + string(op.Uuid)
	value := "reply"
	opAssertE(ev.ctx.db.Put([]byte(key), []byte(value)), "put reply key into db error")
}

// upvote is true: upvote otherwise downvote
// no downvote has been supplied by command, so I ignore it
func (ev *VoteEvaluator) Apply() {
	op := ev.op

	voterWrap := table.NewSoAccountWrap(ev.ctx.db, op.Voter)
	elapsedSeconds := ev.ctx.control.HeadBlockTime().UtcSeconds - voterWrap.GetLastVoteTime().UtcSeconds
	opAssert(elapsedSeconds < constants.MIN_VOTE_INTERVAL, "voting frequently")

	voterId := prototype.VoterId{Voter: op.Voter, PostId: op.Idx}
	voteWrap := table.NewSoVoteWrap(ev.ctx.db, &voterId)
	postWrap := table.NewSoPostWrap(ev.ctx.db, &op.Idx)

	opAssert(!postWrap.CheckExist(), "post invalid")
	opAssert(!voteWrap.CheckExist(), "vote info exist")

	votePostWrap := table.NewVotePostIdWrap(ev.ctx.db)

	for voteIter := votePostWrap.QueryListByOrder(&op.Idx, nil); voteIter.Valid(); voteIter.Next() {
		voterId := votePostWrap.GetMainVal(voteIter)
		if voterId.Voter.Value == op.Voter.Value {
			opAssertE(errors.New("Vote Error"), "vote to a same post")
		}
	}

	regeneratedPower := constants.PERCENT * elapsedSeconds / constants.VOTE_REGENERATE_TIME
	var currentVp uint32
	votePower := voterWrap.GetVotePower() + regeneratedPower
	if votePower > constants.PERCENT {
		currentVp = constants.PERCENT
	} else {
		currentVp = votePower
	}
	usedVp := (currentVp + constants.VOTE_LIMITE_DURING_REGENERATE - 1) / constants.VOTE_LIMITE_DURING_REGENERATE

	voterWrap.MdVotePower(currentVp - usedVp)
	voterWrap.MdLastVoteTime(ev.ctx.control.HeadBlockTime())
	vesting := voterWrap.GetVestingShares().Value
	// todo: uint128
	weightedVp := vesting * uint64(usedVp)
	if postWrap.GetCashoutTime().UtcSeconds < ev.ctx.control.HeadBlockTime().UtcSeconds {
		lastVp := postWrap.GetWeightedVp()
		votePower := lastVp + weightedVp
		// add new vp into global
		ev.ctx.control.AddWeightedVP(weightedVp)
		// update post's weighted vp
		postWrap.MdWeightedVp(votePower)

		opAssertE(voteWrap.Create(func(t *table.SoVote) {
			t.Voter = &voterId
			t.PostId = op.Idx
			t.Upvote = true
			t.WeightedVp = weightedVp
			t.VoteTime = ev.ctx.control.HeadBlockTime()
		}), "create voter object error")

		opAssert(postWrap.MdVoteCnt(postWrap.GetVoteCnt()+1), "set vote count error")
	}
}

func (ev *BpRegisterEvaluator) Apply() {
	op := ev.op
	witnessWrap := table.NewSoWitnessWrap(ev.ctx.db, op.Owner)

	opAssert(!witnessWrap.CheckExist(), "witness already exist")

	opAssertE(witnessWrap.Create(func(t *table.SoWitness) {
		t.Owner = op.Owner
		t.CreatedTime = ev.ctx.control.HeadBlockTime()
		t.Url = op.Url
		t.SigningKey = op.BlockSigningKey

		// TODO add others
	}), "add witness record error")
}

func (ev *BpUnregisterEvaluator) Apply() {
	// unregister op cost too much cpu time
	panic("not yet implement")

}

func (ev *BpVoteEvaluator) Apply() {
	op := ev.op

	voterAccount := table.NewSoAccountWrap(ev.ctx.db, op.Voter)
	voteCnt := voterAccount.GetBpVoteCount()

	voterId := &prototype.BpVoterId{Voter: op.Voter, Witness: op.Witness}
	witnessId := &prototype.BpWitnessId{Voter: op.Voter, Witness: op.Witness}
	vidWrap := table.NewSoWitnessVoteWrap(ev.ctx.db, voterId)

	witAccWrap := table.NewSoAccountWrap(ev.ctx.db, op.Voter)
	opAssert(witAccWrap.CheckExist(), "witness account do not exist ")

	witnessWrap := table.NewSoWitnessWrap(ev.ctx.db, op.Witness)

	if op.Cancel {
		opAssert(voteCnt > 0, "vote count must not be 0")
		opAssert(vidWrap.CheckExist(), "vote record not exist")
		opAssert(vidWrap.RemoveWitnessVote(), "remove vote record error")
		opAssert(witnessWrap.GetVoteCount() > 0, "witness data error")
		opAssert(witnessWrap.MdVoteCount(witnessWrap.GetVoteCount()-1), "set witness data error")
		opAssert(voterAccount.MdBpVoteCount(voteCnt-1), "set voter data error")
	} else {
		opAssert(voteCnt < constants.MAX_BP_VOTE_COUNT, "vote count exceeding")

		opAssertE(vidWrap.Create(func(t *table.SoWitnessVote) {
			t.VoteTime = ev.ctx.control.HeadBlockTime()
			t.VoterId = voterId
			t.WitnessId = witnessId
		}), "add vote record error")

		opAssert(voterAccount.MdBpVoteCount(voteCnt+1), "set voter data error")
		opAssert(witnessWrap.MdVoteCount(witnessWrap.GetVoteCount()+1), "set witness data error")
	}

}

func (ev *FollowEvaluator) Apply() {
	op := ev.op

	acctWrap := table.NewSoAccountWrap(ev.ctx.db, op.Account)
	opAssert(acctWrap.CheckExist(), "follow account do not exist ")

	acctWrap = table.NewSoAccountWrap(ev.ctx.db, op.FAccount)
	opAssert(acctWrap.CheckExist(), "follow f_account do not exist ")
}

func (ev *TransferToVestingEvaluator) Apply() {
	op := ev.op

	fidWrap := table.NewSoAccountWrap(ev.ctx.db, op.From)
	tidWrap := table.NewSoAccountWrap(ev.ctx.db, op.To)

	opAssert(tidWrap.CheckExist(), "to account do not exist")

	fBalance := fidWrap.GetBalance()
	tVests := tidWrap.GetVestingShares()
	addVests := prototype.NewVest(op.Amount.Value)

	opAssertE(fBalance.Sub(op.Amount), "balance not enough")
	opAssert(fidWrap.MdBalance(fBalance), "set from new balance error")

	opAssertE(tVests.Add(addVests), "vests error")
	opAssert(tidWrap.MdVestingShares(tVests), "set to new vests error")

	ev.ctx.control.TransferToVest(op.Amount)
}
