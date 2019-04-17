package app

import (
	"github.com/coschain/contentos-go/app/table"
	"github.com/coschain/contentos-go/common/constants"
	"github.com/coschain/contentos-go/common/encoding/vme"
	"github.com/coschain/contentos-go/prototype"
	"github.com/coschain/contentos-go/vm"
	"github.com/coschain/contentos-go/vm/context"
	"github.com/coschain/contentos-go/vm/contract/abi"
	ct "github.com/coschain/contentos-go/vm/contract/table"
	"github.com/sirupsen/logrus"
	"math"
	"sort"
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

//type ClaimEvaluator struct {
//	BaseEvaluator
//	ctx *ApplyContext
//	op  *prototype.ClaimOperation
//}

type ReportEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.ReportOperation
}

type ConvertVestingEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.ConvertVestingOperation
}

// I can cat out this awkward claimall operation until I can get value from rpc resp
//type ClaimAllEvaluator struct {
//	BaseEvaluator
//	ctx *ApplyContext
//	op  *prototype.ClaimAllOperation
//}

type ContractDeployEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.ContractDeployOperation
}

type ContractApplyEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.ContractApplyOperation
}

type InternalContractApplyEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.InternalContractApplyOperation
}

func (ev *AccountCreateEvaluator) Apply() {
	op := ev.op
	creatorWrap := table.NewSoAccountWrap(ev.ctx.db, op.Creator)

	opAssert(creatorWrap.CheckExist(), "creator not exist ")

	opAssert(creatorWrap.GetBalance().Value >= op.Fee.Value, "Insufficient balance to create account.")

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
		tInfo.LastPostTime = ev.ctx.control.HeadBlockTime()
		tInfo.LastVoteTime = ev.ctx.control.HeadBlockTime()
		tInfo.NextPowerdownBlockNum = math.MaxUint32
		tInfo.EachPowerdownRate = &prototype.Vest{Value: 0}
		tInfo.ToPowerdown = &prototype.Vest{Value: 0}
		tInfo.HasPowerdown = &prototype.Vest{Value: 0}
		tInfo.Owner = op.Owner
		tInfo.LastOwnerUpdate = prototype.NewTimePointSec(0)
	}), "duplicate create account object")

	// create account authority
	//authorityWrap := table.NewSoAccountWrap(ev.ctx.db, op.NewAccountName)
	//opAssertE(authorityWrap.Create(func(tInfo *table.SoAccountAuthorityObject) {
	//	tInfo.Account = op.NewAccountName
	//	tInfo.Owner = op.Owner
	//	tInfo.LastOwnerUpdate = prototype.NewTimePointSec(0)
	//}), "duplicate create account authority object")

	// sub dynamic glaobal properties's total fee
	ev.ctx.control.TransferToVest(op.Fee)
	ev.ctx.control.ModifyProps(func(props *prototype.DynamicProperties) {
		props.TotalUserCnt++
	})
}

func (ev *TransferEvaluator) Apply() {
	op := ev.op

	// @ active_challenged
	fromWrap := table.NewSoAccountWrap(ev.ctx.db, op.From)
	toWrap := table.NewSoAccountWrap(ev.ctx.db, op.To)

	opAssert(toWrap.CheckExist(), "To account do not exist ")

	opAssert(op.From.Value != op.To.Value, "Transfer must between two different accounts")

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
	opAssert(elapsedSeconds > constants.MinPostInterval, "posting frequently")

	// default source is contentos
	opAssertE(idWrap.Create(func(t *table.SoPost) {
		t.PostId = op.Uuid
		t.Tags = op.Tags
		t.Title = op.Title
		t.Author = op.Owner
		t.Body = op.Content
		t.Created = ev.ctx.control.HeadBlockTime()
		//t.CashoutTime = &prototype.TimePointSec{UtcSeconds: ev.ctx.control.HeadBlockTime().UtcSeconds + uint32(constants.PostCashOutDelayTime)}
		t.CashoutBlockNum = ev.ctx.control.GetProps().HeadBlockNumber + constants.PostCashOutDelayBlock
		t.Depth = 0
		t.Children = 0
		t.RootId = t.PostId
		t.ParentId = 0
		t.RootId = 0
		t.Beneficiaries = op.Beneficiaries
		t.WeightedVp = 0
		t.VoteCnt = 0
		t.Rewards = &prototype.Vest{Value: 0}
		t.DappRewards = &prototype.Vest{Value: 0}
	}), "create post error")

	authorWrap.MdLastPostTime(ev.ctx.control.HeadBlockTime())

	ev.ctx.control.ModifyProps(func(props *prototype.DynamicProperties) {
		props.TotalPostCnt++
	})

	//timestamp := ev.ctx.control.HeadBlockTime().UtcSeconds + uint32(constants.PostCashOutDelayTime) - uint32(constants.GenesisTime)
	//key := fmt.Sprintf("cashout:%d_%d", common.GetBucket(timestamp), op.Uuid)
	//value := "post"
	//opAssertE(ev.ctx.db.Put([]byte(key), []byte(value)), "put post key into db error")

}

func (ev *ReplyEvaluator) Apply() {
	op := ev.op
	cidWrap := table.NewSoPostWrap(ev.ctx.db, &op.Uuid)
	pidWrap := table.NewSoPostWrap(ev.ctx.db, &op.ParentUuid)

	opAssert(!cidWrap.CheckExist(), "post uuid exist")
	opAssert(pidWrap.CheckExist(), "parent uuid do not exist")

	opAssert(pidWrap.GetDepth()+1 < constants.PostMaxDepth, "reply depth error")

	authorWrap := table.NewSoAccountWrap(ev.ctx.db, op.Owner)
	elapsedSeconds := ev.ctx.control.HeadBlockTime().UtcSeconds - authorWrap.GetLastPostTime().UtcSeconds
	opAssert(elapsedSeconds > constants.MinPostInterval, "reply frequently")

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
		//t.CashoutTime = &prototype.TimePointSec{UtcSeconds: ev.ctx.control.HeadBlockTime().UtcSeconds + uint32(constants.PostCashOutDelayTime)}
		t.CashoutBlockNum = ev.ctx.control.GetProps().HeadBlockNumber + constants.PostCashOutDelayBlock
		t.Depth = pidWrap.GetDepth() + 1
		t.Children = 0
		t.RootId = rootId
		t.ParentId = op.ParentUuid
		t.WeightedVp = 0
		t.VoteCnt = 0
		t.Beneficiaries = op.Beneficiaries
		t.Rewards = &prototype.Vest{Value: 0}
		t.DappRewards = &prototype.Vest{Value: 0}
	}), "create reply error")

	authorWrap.MdLastPostTime(ev.ctx.control.HeadBlockTime())
	// Modify Parent Object
	opAssert(pidWrap.MdChildren(pidWrap.GetChildren()+1), "Modify Parent Children Error")

	//timestamp := ev.ctx.control.HeadBlockTime().UtcSeconds + uint32(constants.PostCashOutDelayTime) - uint32(constants.GenesisTime)
	//key := fmt.Sprintf("cashout:%d_%d", common.GetBucket(timestamp), op.Uuid)
	//value := "reply"
	//opAssertE(ev.ctx.db.Put([]byte(key), []byte(value)), "put reply key into db error")
}

// upvote is true: upvote otherwise downvote
// no downvote has been supplied by command, so I ignore it
func (ev *VoteEvaluator) Apply() {
	op := ev.op

	voterWrap := table.NewSoAccountWrap(ev.ctx.db, op.Voter)
	elapsedSeconds := ev.ctx.control.HeadBlockTime().UtcSeconds - voterWrap.GetLastVoteTime().UtcSeconds
	opAssert(elapsedSeconds > constants.MinVoteInterval, "voting frequently")

	voterId := prototype.VoterId{Voter: op.Voter, PostId: op.Idx}
	voteWrap := table.NewSoVoteWrap(ev.ctx.db, &voterId)
	postWrap := table.NewSoPostWrap(ev.ctx.db, &op.Idx)

	opAssert(postWrap.CheckExist(), "post invalid")
	opAssert(!voteWrap.CheckExist(), "vote info exist")

	//votePostWrap := table.NewVotePostIdWrap(ev.ctx.db)

	//for voteIter := votePostWrap.QueryListByOrder(&op.Idx, nil); voteIter.Valid(); voteIter.Next() {
	//	voterId := votePostWrap.GetMainVal(voteIter)
	//	if voterId.Voter.Value == op.Voter.Value {
	//		opAssertE(errors.New("Vote Error"), "vote to a same post")
	//	}
	//}

	regeneratedPower := constants.PERCENT * elapsedSeconds / constants.VoteRegenerateTime
	var currentVp uint32
	votePower := voterWrap.GetVotePower() + regeneratedPower
	if votePower > constants.PERCENT {
		currentVp = constants.PERCENT
	} else {
		currentVp = votePower
	}
	usedVp := (currentVp + constants.VoteLimitDuringRegenerate - 1) / constants.VoteLimitDuringRegenerate

	voterWrap.MdVotePower(currentVp - usedVp)
	voterWrap.MdLastVoteTime(ev.ctx.control.HeadBlockTime())
	vesting := voterWrap.GetVestingShares().Value
	// after constants.PERCENT replaced by 1000, max value is 10000000000 * 1000000 * 1000 / 30
	// 10000000000 * 1000000 * 1000 < 18446744073709552046 but 10000000000 * 1000000 > 9223372036854775807
	// so can not using int64 here
	weightedVp := vesting * uint64(usedVp)
	if postWrap.GetCashoutBlockNum() > ev.ctx.control.GetProps().HeadBlockNumber {
		lastVp := postWrap.GetWeightedVp()
		votePower := lastVp + weightedVp
		// add new vp into global
		//ev.ctx.control.AddWeightedVP(weightedVp)
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

func (ev *BpRegisterEvaluator) BpInWhiteList(bpName string) bool {
	switch bpName {
	case "initminer1":
		return true
	case "initminer2":
		return true
	case "initminer3":
		return true
	case "initminer4":
		return true
	}
	return false
}

func (ev *BpRegisterEvaluator) Apply() {
	op := ev.op

	opAssert(ev.BpInWhiteList(op.Owner.Value), "bp name not in white list")

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
		opAssert(voteCnt < constants.MaxBpVoteCount, "vote count exceeding")

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

func (ev *ConvertVestingEvaluator) Apply() {
	op := ev.op
	accWrap := table.NewSoAccountWrap(ev.ctx.db, op.From)
	opAssert(accWrap.CheckExist(), "account do not exist")
	opAssert(op.Amount.Value >= uint64(1e6), "At least 1 vesting should be converted")
	opAssert(accWrap.GetVestingShares().Value >= op.Amount.Value, "vesting balance not enough")
	globalProps := ev.ctx.control.GetProps()
	//timestamp := globalProps.Time.UtcSeconds
	currentBlock := globalProps.HeadBlockNumber
	eachRate := op.Amount.Value / constants.ConvertWeeks
	//accWrap.MdNextPowerdownTime(&prototype.TimePointSec{UtcSeconds: timestamp + constants.POWER_DOWN_INTERVAL})
	accWrap.MdNextPowerdownBlockNum(currentBlock + constants.PowerDownBlockInterval)
	accWrap.MdEachPowerdownRate(&prototype.Vest{Value: eachRate})
	accWrap.MdHasPowerdown(&prototype.Vest{Value: 0})
	accWrap.MdToPowerdown(op.Amount)
}

//func (ev *ClaimEvaluator) Apply() {
//	op := ev.op
//
//	account := op.Account
//	accWrap := table.NewSoAccountWrap(ev.ctx.db, account)
//
//	opAssert(accWrap.CheckExist(), "claim account do not exist")
//
//	var i int32 = 1
//	keeperWrap := table.NewSoRewardsKeeperWrap(ev.ctx.db, &i)
//	opAssert(keeperWrap.CheckExist(), "reward keeper do not exist")
//
//	keeper := keeperWrap.GetKeeper()
//	innerRewards := keeper.Rewards
//
//	amount := op.Amount
//
//	if val, ok := innerRewards[account.Value]; ok {
//		rewardBalance := val.Value
//		var reward uint64
//		if rewardBalance >= amount && rewardBalance-amount <= rewardBalance {
//			reward = amount
//		} else {
//			reward = rewardBalance
//		}
//		if reward > 0 {
//			vestingBalance := accWrap.GetVestingShares()
//			accWrap.MdVestingShares(&prototype.Vest{Value: vestingBalance.Value + reward})
//			val.Value -= reward
//			keeperWrap.MdKeeper(keeper)
//		} else {
//			// do nothing
//		}
//	} else {
//		opAssert(ok, "No remains reward on chain")
//	}
//
//}

type byTag []int32

func (c byTag) Len() int {
	return len(c)
}
func (c byTag) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c byTag) Less(i, j int) bool {
	return c[i] < c[j]
}

func mergeTags(existed []int32, new []prototype.ReportOperationTag) []int32 {
	len1 := len(existed)
	len2 := len(new)
	tmp := make([]int32, 0, len2)
	for i := 0; i < len2; i++ {
		tmp[i] = int32(new[i])
	}
	sort.Sort(byTag(existed))
	sort.Sort(byTag(tmp))

	res := make([]int32, 0, len1+len2)
	i := 0
	j := 0
	for {
		if i == len1 || j == len2 {
			break
		}
		if existed[i] <= tmp[j] {
			res = append(res, existed[i])
			if existed[i] == tmp[j] {
				j++
			}
			i++
		} else if existed[i] > tmp[j] {
			res = append(res, tmp[j])
			j++
		}
	}
	if i < len1 {
		res = append(res, existed[i:]...)
	}
	if j < len2 {
		res = append(res, tmp[i:]...)
	}

	return res
}

func (ev *ReportEvaluator) Apply() {
	op := ev.op
	post := table.NewSoPostWrap(ev.ctx.db, &op.Reported)
	opAssert(post.CheckExist(), "the reported post doesn't exist")
	report := table.NewSoReportListWrap(ev.ctx.db, &op.Reported)
	if op.IsArbitration {
		opAssert(report.CheckExist(), "cannot arbitrate a non-existed post")
		if op.IsApproved {
			post.RemovePost()
			report.RemoveReportList()
			return
		}

		report.MdIsArbitrated(true)
	} else {
		if report.CheckExist() {
			if report.GetIsArbitrated() {
				opAssert(false, "cannot report a legal post")
			}
			report.MdReportedTimes(report.GetReportedTimes() + 1)
			existedTags := report.GetTags()
			newTags := op.ReportTag
			report.MdTags(mergeTags(existedTags, newTags))
			return
		}

		report.Create(func(tInfo *table.SoReportList) {
			tInfo.Uuid = op.Reported
			tInfo.ReportedTimes = 1
			tags := make([]int32, len(op.ReportTag))
			for i := range op.ReportTag {
				tags[i] = int32(op.ReportTag[i])
			}
			tInfo.Tags = tags
			tInfo.IsArbitrated = false
		})
	}
}
//
//func (ev *ClaimAllEvaluator) Apply() {
//	op := ev.op
//
//	account := op.Account
//	accWrap := table.NewSoAccountWrap(ev.ctx.db, account)
//
//	opAssert(accWrap.CheckExist(), "claim account do not exist")
//
//	var i int32 = 1
//	keeperWrap := table.NewSoRewardsKeeperWrap(ev.ctx.db, &i)
//	opAssert(keeperWrap.CheckExist(), "reward keeper do not exist")
//
//	keeper := keeperWrap.GetKeeper()
//	innerRewards := keeper.Rewards
//
//	if val, ok := innerRewards[account.Value]; ok {
//		reward := val.Value
//		if reward > 0 {
//			vestingBalance := accWrap.GetVestingShares()
//			accWrap.MdVestingShares(&prototype.Vest{Value: vestingBalance.Value + reward})
//			val.Value -= reward
//			keeperWrap.MdKeeper(keeper)
//		} else {
//			// do nothing
//		}
//	} else {
//		opAssert(ok, "No remains reward on chain")
//	}
//
//}

func (ev *ContractDeployEvaluator) Apply() {
	op := ev.op

	cid := prototype.ContractId{Owner: op.Owner, Cname: op.Contract}
	scid := table.NewSoContractWrap(ev.ctx.db, &cid)

	opAssert(!scid.CheckExist(), "contract name exist")

	_, err := abi.UnmarshalABI([]byte(op.GetAbi()))
	if err != nil {
		opAssertE(err, "invalid contract abi")
	}

	vmCtx := vmcontext.NewContextFromDeployOp(op, nil)

	cosVM := vm.NewCosVM(vmCtx, nil, nil, nil)

	opAssertE(cosVM.Validate(), "validate code failed")

	opAssertE(scid.Create(func(t *table.SoContract) {
		t.Code = op.Code
		t.Id = &cid
		t.CreatedTime = ev.ctx.control.HeadBlockTime()
		t.Abi = op.Abi
		t.Balance = prototype.NewCoin(0)
	}), "create contract data error")
}

func (ev *ContractApplyEvaluator) Apply() {
	op := ev.op

	cid := prototype.ContractId{Owner: op.Owner, Cname: op.Contract}
	scid := table.NewSoContractWrap(ev.ctx.db, &cid)
	opAssert(scid.CheckExist(), "contract name doesn't exist")

	acc := table.NewSoAccountWrap(ev.ctx.db, op.Caller)
	opAssert(acc.CheckExist(), "account doesn't exist")

	balance := acc.GetBalance().Value
	// fixme, should base on minicos
	balanceExchange := balance * constants.BaseRate

	opAssert(balanceExchange >= op.Gas.Value, "balance can not pay gas fee")

	// the amount is also minicos or cos ?
	// here I assert it is minicos
	// also, I think balance base on minicos is far more reliable.
	opAssert(balanceExchange-op.Gas.Value > op.Amount.Value, "balance does not have enough fund to transfer after paid gas fee")

	code := scid.GetCode()

	var err error
	var abiInterface abi.IContractABI
	var paramsData []byte
	var tables *ct.ContractTables

	if abiInterface, err = abi.UnmarshalABI([]byte(scid.GetAbi())); err != nil {
		opAssertE(err, "invalid contract abi")
	}
	if m := abiInterface.MethodByName(op.Method); m != nil {
		paramsData, err = vme.EncodeFromJson([]byte(op.Params), m.Args().Type())
		if err != nil {
			opAssertE(err, "invalid contract parameters")
		}
	} else {
		opAssert(false, "unknown contract method: "+op.Method)
	}

	if abiInterface != nil {
		tables = ct.NewContractTables(op.Owner.Value, op.Contract, abiInterface, ev.ctx.db)
	}

	vmCtx := vmcontext.NewContextFromApplyOp(op, paramsData, code, abiInterface, tables, ev.ctx.vmInjector)
	// should be active ?
	//defer func() {
	//	_ := recover()
	//}()

	cosVM := vm.NewCosVM(vmCtx, ev.ctx.db, ev.ctx.control.GetProps(), logrus.New())

	ret, err := cosVM.Run()
	spentGas := cosVM.SpentGas()
	// need extra query db, is it a good way or should I pass account object as parameter?
	// deductgasfee and usertranfer could be panic (rarely, I can't image how it happens)
	// the panic should catch then return or bubble it ?

	// TODO merge, temp fix
	opAssertE(err, "execute vm error")

	vmCtx.Injector.DeductGasFee(op.Caller.Value, spentGas)
	if err != nil {
		vmCtx.Injector.Error(ret, err.Error())
	} else {
		if op.Amount.Value > 0 {
			vmCtx.Injector.TransferFromUserToContract(op.Caller.Value, op.Contract, op.Owner.Value, op.Amount.Value)
		}
	}
}

func (ev *InternalContractApplyEvaluator) Apply() {
	op := ev.op

	fromContract := table.NewSoContractWrap(ev.ctx.db, &prototype.ContractId{Owner: op.FromOwner, Cname: op.FromContract})
	opAssert(fromContract.CheckExist(), "fromContract contract doesn't exist")

	toContract := table.NewSoContractWrap(ev.ctx.db, &prototype.ContractId{Owner: op.ToOwner, Cname: op.ToContract})
	opAssert(toContract.CheckExist(), "toContract contract doesn't exist")

	caller := table.NewSoAccountWrap(ev.ctx.db, op.FromCaller)
	opAssert(caller.CheckExist(), "caller account doesn't exist")

	opAssert(caller.GetBalance().Value*constants.BaseRate >= op.Gas.Value, "caller balance less than gas")
	opAssert(fromContract.GetBalance().Value >= op.Amount.Value, "fromContract balance less than transfer amount")

	code := toContract.GetCode()

	var err error
	var abiInterface abi.IContractABI
	var tables *ct.ContractTables

	if abiInterface, err = abi.UnmarshalABI([]byte(toContract.GetAbi())); err != nil {
		opAssertE(err, "invalid toContract abi")
	}
	if m := abiInterface.MethodByName(op.ToMethod); m != nil {
		_, err = vme.DecodeToJson(op.Params, m.Args().Type(), false)
		if err != nil {
			opAssertE(err, "invalid contract parameters")
		}
	} else {
		opAssert(false, "unknown contract method: "+op.ToMethod)
	}

	if abiInterface != nil {
		tables = ct.NewContractTables(op.ToOwner.Value, op.ToContract, abiInterface, ev.ctx.db)
	}

	vmCtx := vmcontext.NewContextFromInternalApplyOp(op, code, abiInterface, tables, ev.ctx.vmInjector)
	cosVM := vm.NewCosVM(vmCtx, ev.ctx.db, ev.ctx.control.GetProps(), logrus.New())
	ret, err := cosVM.Run()

	vmCtx.Injector.DeductGasFee(op.FromCaller.Value, cosVM.SpentGas())
	if err != nil {
		vmCtx.Injector.Error(ret, err.Error())
	} else {
		if op.Amount.Value > 0 {
			vmCtx.Injector.TransferFromContractToContract(op.FromContract, op.FromOwner.Value, op.ToContract, op.ToOwner.Value, op.Amount.Value)
		}
	}
}
