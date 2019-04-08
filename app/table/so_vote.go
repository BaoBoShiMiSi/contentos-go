package table

import (
	"errors"
	fmt "fmt"
	"reflect"
	"strings"

	"github.com/coschain/contentos-go/common/encoding/kope"
	"github.com/coschain/contentos-go/iservices"
	prototype "github.com/coschain/contentos-go/prototype"
	proto "github.com/golang/protobuf/proto"
)

////////////// SECTION Prefix Mark ///////////////
var (
	VoteVoterTable     uint32 = 2638131561
	VoteVoteTimeTable  uint32 = 277775897
	VotePostIdTable    uint32 = 3923440502
	VoteVoterUniTable  uint32 = 1965941220
	VotePostIdCell     uint32 = 3799570737
	VoteUpvoteCell     uint32 = 3751109702
	VoteVoteTimeCell   uint32 = 265476525
	VoteVoterCell      uint32 = 3085458610
	VoteWeightedVpCell uint32 = 334607886
)

////////////// SECTION Wrap Define ///////////////
type SoVoteWrap struct {
	dba      iservices.IDatabaseRW
	mainKey  *prototype.VoterId
	mKeyFlag int    //the flag of the main key exist state in db, -1:has not judged; 0:not exist; 1:already exist
	mKeyBuf  []byte //the buffer after the main key is encoded with prefix
	mBuf     []byte //the value after the main key is encoded
}

func NewSoVoteWrap(dba iservices.IDatabaseRW, key *prototype.VoterId) *SoVoteWrap {
	if dba == nil || key == nil {
		return nil
	}
	result := &SoVoteWrap{dba, key, -1, nil, nil}
	return result
}

func (s *SoVoteWrap) CheckExist() bool {
	if s.dba == nil {
		return false
	}
	if s.mKeyFlag != -1 {
		//if you have already obtained the existence status of the primary key, use it directly
		if s.mKeyFlag == 0 {
			return false
		}
		return true
	}
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return false
	}

	res, err := s.dba.Has(keyBuf)
	if err != nil {
		return false
	}
	if res == false {
		s.mKeyFlag = 0
	} else {
		s.mKeyFlag = 1
	}
	return res
}

func (s *SoVoteWrap) Create(f func(tInfo *SoVote)) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if s.mainKey == nil {
		return errors.New("the main key is nil")
	}
	val := &SoVote{}
	f(val)
	if val.Voter == nil {
		val.Voter = s.mainKey
	}
	if s.CheckExist() {
		return errors.New("the main key is already exist")
	}
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return err

	}
	err = s.saveAllMemKeys(val, true)
	if err != nil {
		s.delAllMemKeys(false, val)
		return err
	}

	// update srt list keys
	if err = s.insertAllSortKeys(val); err != nil {
		s.delAllSortKeys(false, val)
		s.dba.Delete(keyBuf)
		s.delAllMemKeys(false, val)
		return err
	}

	//update unique list
	if sucNames, err := s.insertAllUniKeys(val); err != nil {
		s.delAllSortKeys(false, val)
		s.delUniKeysWithNames(sucNames, val)
		s.dba.Delete(keyBuf)
		s.delAllMemKeys(false, val)
		return err
	}

	return nil
}

func (s *SoVoteWrap) getMainKeyBuf() ([]byte, error) {
	if s.mainKey == nil {
		return nil, errors.New("the main key is nil")
	}
	if s.mBuf == nil {
		var err error = nil
		s.mBuf, err = kope.Encode(s.mainKey)
		if err != nil {
			return nil, err
		}
	}
	return s.mBuf, nil
}

////////////// SECTION LKeys delete/insert ///////////////

func (s *SoVoteWrap) delSortKeyVoter(sa *SoVote) bool {
	if s.dba == nil || s.mainKey == nil {
		return false
	}
	val := SoListVoteByVoter{}
	if sa == nil {
		key, err := s.encodeMemKey("Voter")
		if err != nil {
			return false
		}
		buf, err := s.dba.Get(key)
		if err != nil {
			return false
		}
		ori := &SoMemVoteByVoter{}
		err = proto.Unmarshal(buf, ori)
		if err != nil {
			return false
		}
		val.Voter = ori.Voter
	} else {
		val.Voter = sa.Voter
	}

	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoVoteWrap) insertSortKeyVoter(sa *SoVote) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	val := SoListVoteByVoter{}
	val.Voter = sa.Voter
	buf, err := proto.Marshal(&val)
	if err != nil {
		return false
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Put(subBuf, buf)
	return ordErr == nil
}

func (s *SoVoteWrap) delSortKeyVoteTime(sa *SoVote) bool {
	if s.dba == nil || s.mainKey == nil {
		return false
	}
	val := SoListVoteByVoteTime{}
	if sa == nil {
		key, err := s.encodeMemKey("VoteTime")
		if err != nil {
			return false
		}
		buf, err := s.dba.Get(key)
		if err != nil {
			return false
		}
		ori := &SoMemVoteByVoteTime{}
		err = proto.Unmarshal(buf, ori)
		if err != nil {
			return false
		}
		val.VoteTime = ori.VoteTime
		val.Voter = s.mainKey

	} else {
		val.VoteTime = sa.VoteTime
		val.Voter = sa.Voter
	}

	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoVoteWrap) insertSortKeyVoteTime(sa *SoVote) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	val := SoListVoteByVoteTime{}
	val.Voter = sa.Voter
	val.VoteTime = sa.VoteTime
	buf, err := proto.Marshal(&val)
	if err != nil {
		return false
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Put(subBuf, buf)
	return ordErr == nil
}

func (s *SoVoteWrap) delSortKeyPostId(sa *SoVote) bool {
	if s.dba == nil || s.mainKey == nil {
		return false
	}
	val := SoListVoteByPostId{}
	if sa == nil {
		key, err := s.encodeMemKey("PostId")
		if err != nil {
			return false
		}
		buf, err := s.dba.Get(key)
		if err != nil {
			return false
		}
		ori := &SoMemVoteByPostId{}
		err = proto.Unmarshal(buf, ori)
		if err != nil {
			return false
		}
		val.PostId = ori.PostId
		val.Voter = s.mainKey

	} else {
		val.PostId = sa.PostId
		val.Voter = sa.Voter
	}

	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoVoteWrap) insertSortKeyPostId(sa *SoVote) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	val := SoListVoteByPostId{}
	val.Voter = sa.Voter
	val.PostId = sa.PostId
	buf, err := proto.Marshal(&val)
	if err != nil {
		return false
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Put(subBuf, buf)
	return ordErr == nil
}

func (s *SoVoteWrap) delAllSortKeys(br bool, val *SoVote) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if !s.delSortKeyVoter(val) {
		if br {
			return false
		} else {
			res = false
		}
	}
	if !s.delSortKeyVoteTime(val) {
		if br {
			return false
		} else {
			res = false
		}
	}
	if !s.delSortKeyPostId(val) {
		if br {
			return false
		} else {
			res = false
		}
	}

	return res
}

func (s *SoVoteWrap) insertAllSortKeys(val *SoVote) error {
	if s.dba == nil {
		return errors.New("insert sort Field fail,the db is nil ")
	}
	if val == nil {
		return errors.New("insert sort Field fail,get the SoVote fail ")
	}
	if !s.insertSortKeyVoter(val) {
		return errors.New("insert sort Field Voter fail while insert table ")
	}
	if !s.insertSortKeyVoteTime(val) {
		return errors.New("insert sort Field VoteTime fail while insert table ")
	}
	if !s.insertSortKeyPostId(val) {
		return errors.New("insert sort Field PostId fail while insert table ")
	}

	return nil
}

////////////// SECTION LKeys delete/insert //////////////

func (s *SoVoteWrap) RemoveVote() bool {
	if s.dba == nil {
		return false
	}
	val := &SoVote{}
	//delete sort list key
	if res := s.delAllSortKeys(true, nil); !res {
		return false
	}

	//delete unique list
	if res := s.delAllUniKeys(true, nil); !res {
		return false
	}

	err := s.delAllMemKeys(true, val)
	if err == nil {
		s.mKeyBuf = nil
		s.mKeyFlag = -1
		return true
	} else {
		return false
	}
}

////////////// SECTION Members Get/Modify ///////////////
func (s *SoVoteWrap) getMemKeyPrefix(fName string) uint32 {
	if fName == "PostId" {
		return VotePostIdCell
	}
	if fName == "Upvote" {
		return VoteUpvoteCell
	}
	if fName == "VoteTime" {
		return VoteVoteTimeCell
	}
	if fName == "Voter" {
		return VoteVoterCell
	}
	if fName == "WeightedVp" {
		return VoteWeightedVpCell
	}

	return 0
}

func (s *SoVoteWrap) encodeMemKey(fName string) ([]byte, error) {
	if len(fName) < 1 || s.mainKey == nil {
		return nil, errors.New("field name or main key is empty")
	}
	pre := s.getMemKeyPrefix(fName)
	preBuf, err := kope.Encode(pre)
	if err != nil {
		return nil, err
	}
	mBuf, err := s.getMainKeyBuf()
	if err != nil {
		return nil, err
	}
	list := make([][]byte, 2)
	list[0] = preBuf
	list[1] = mBuf
	return kope.PackList(list), nil
}

func (s *SoVoteWrap) saveAllMemKeys(tInfo *SoVote, br bool) error {
	if s.dba == nil {
		return errors.New("save member Field fail , the db is nil")
	}

	if tInfo == nil {
		return errors.New("save member Field fail , the data is nil ")
	}
	var err error = nil
	errDes := ""
	if err = s.saveMemKeyPostId(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "PostId", err)
		}
	}
	if err = s.saveMemKeyUpvote(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "Upvote", err)
		}
	}
	if err = s.saveMemKeyVoteTime(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "VoteTime", err)
		}
	}
	if err = s.saveMemKeyVoter(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "Voter", err)
		}
	}
	if err = s.saveMemKeyWeightedVp(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "WeightedVp", err)
		}
	}

	if len(errDes) > 0 {
		return errors.New(errDes)
	}
	return err
}

func (s *SoVoteWrap) delAllMemKeys(br bool, tInfo *SoVote) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	t := reflect.TypeOf(*tInfo)
	errDesc := ""
	for k := 0; k < t.NumField(); k++ {
		name := t.Field(k).Name
		if len(name) > 0 && !strings.HasPrefix(name, "XXX_") {
			err := s.delMemKey(name)
			if err != nil {
				if br {
					return err
				}
				errDesc += fmt.Sprintf("delete the Field %s fail,error is %s;\n", name, err)
			}
		}
	}
	if len(errDesc) > 0 {
		return errors.New(errDesc)
	}
	return nil
}

func (s *SoVoteWrap) delMemKey(fName string) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if len(fName) <= 0 {
		return errors.New("the field name is empty ")
	}
	key, err := s.encodeMemKey(fName)
	if err != nil {
		return err
	}
	err = s.dba.Delete(key)
	return err
}

func (s *SoVoteWrap) saveMemKeyPostId(tInfo *SoVote) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemVoteByPostId{}
	val.PostId = tInfo.PostId
	key, err := s.encodeMemKey("PostId")
	if err != nil {
		return err
	}
	buf, err := proto.Marshal(&val)
	if err != nil {
		return err
	}
	err = s.dba.Put(key, buf)
	return err
}

func (s *SoVoteWrap) GetPostId() uint64 {
	res := true
	msg := &SoMemVoteByPostId{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("PostId")
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.PostId
			}
		}
	}
	if !res {
		var tmpValue uint64
		return tmpValue
	}
	return msg.PostId
}

func (s *SoVoteWrap) MdPostId(p uint64) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("PostId")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemVoteByPostId{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoVote{}
	sa.Voter = s.mainKey

	sa.PostId = ori.PostId

	if !s.delSortKeyPostId(sa) {
		return false
	}
	ori.PostId = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.PostId = p

	if !s.insertSortKeyPostId(sa) {
		return false
	}

	return true
}

func (s *SoVoteWrap) saveMemKeyUpvote(tInfo *SoVote) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemVoteByUpvote{}
	val.Upvote = tInfo.Upvote
	key, err := s.encodeMemKey("Upvote")
	if err != nil {
		return err
	}
	buf, err := proto.Marshal(&val)
	if err != nil {
		return err
	}
	err = s.dba.Put(key, buf)
	return err
}

func (s *SoVoteWrap) GetUpvote() bool {
	res := true
	msg := &SoMemVoteByUpvote{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("Upvote")
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.Upvote
			}
		}
	}
	if !res {
		var tmpValue bool
		return tmpValue
	}
	return msg.Upvote
}

func (s *SoVoteWrap) MdUpvote(p bool) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("Upvote")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemVoteByUpvote{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoVote{}
	sa.Voter = s.mainKey

	sa.Upvote = ori.Upvote

	ori.Upvote = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.Upvote = p

	return true
}

func (s *SoVoteWrap) saveMemKeyVoteTime(tInfo *SoVote) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemVoteByVoteTime{}
	val.VoteTime = tInfo.VoteTime
	key, err := s.encodeMemKey("VoteTime")
	if err != nil {
		return err
	}
	buf, err := proto.Marshal(&val)
	if err != nil {
		return err
	}
	err = s.dba.Put(key, buf)
	return err
}

func (s *SoVoteWrap) GetVoteTime() *prototype.TimePointSec {
	res := true
	msg := &SoMemVoteByVoteTime{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("VoteTime")
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.VoteTime
			}
		}
	}
	if !res {
		return nil

	}
	return msg.VoteTime
}

func (s *SoVoteWrap) MdVoteTime(p *prototype.TimePointSec) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("VoteTime")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemVoteByVoteTime{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoVote{}
	sa.Voter = s.mainKey

	sa.VoteTime = ori.VoteTime

	if !s.delSortKeyVoteTime(sa) {
		return false
	}
	ori.VoteTime = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.VoteTime = p

	if !s.insertSortKeyVoteTime(sa) {
		return false
	}

	return true
}

func (s *SoVoteWrap) saveMemKeyVoter(tInfo *SoVote) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemVoteByVoter{}
	val.Voter = tInfo.Voter
	key, err := s.encodeMemKey("Voter")
	if err != nil {
		return err
	}
	buf, err := proto.Marshal(&val)
	if err != nil {
		return err
	}
	err = s.dba.Put(key, buf)
	return err
}

func (s *SoVoteWrap) GetVoter() *prototype.VoterId {
	res := true
	msg := &SoMemVoteByVoter{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("Voter")
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.Voter
			}
		}
	}
	if !res {
		return nil

	}
	return msg.Voter
}

func (s *SoVoteWrap) saveMemKeyWeightedVp(tInfo *SoVote) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemVoteByWeightedVp{}
	val.WeightedVp = tInfo.WeightedVp
	key, err := s.encodeMemKey("WeightedVp")
	if err != nil {
		return err
	}
	buf, err := proto.Marshal(&val)
	if err != nil {
		return err
	}
	err = s.dba.Put(key, buf)
	return err
}

func (s *SoVoteWrap) GetWeightedVp() uint64 {
	res := true
	msg := &SoMemVoteByWeightedVp{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("WeightedVp")
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.WeightedVp
			}
		}
	}
	if !res {
		var tmpValue uint64
		return tmpValue
	}
	return msg.WeightedVp
}

func (s *SoVoteWrap) MdWeightedVp(p uint64) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("WeightedVp")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemVoteByWeightedVp{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoVote{}
	sa.Voter = s.mainKey

	sa.WeightedVp = ori.WeightedVp

	ori.WeightedVp = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.WeightedVp = p

	return true
}

////////////// SECTION List Keys ///////////////
type SVoteVoterWrap struct {
	Dba iservices.IDatabaseRW
}

func NewVoteVoterWrap(db iservices.IDatabaseRW) *SVoteVoterWrap {
	if db == nil {
		return nil
	}
	wrap := SVoteVoterWrap{Dba: db}
	return &wrap
}

func (s *SVoteVoterWrap) GetMainVal(val []byte) *prototype.VoterId {
	res := &SoListVoteByVoter{}
	err := proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.Voter

}

func (s *SVoteVoterWrap) GetSubVal(val []byte) *prototype.VoterId {
	res := &SoListVoteByVoter{}
	err := proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
	return res.Voter

}

func (m *SoListVoteByVoter) OpeEncode() ([]byte, error) {
	pre := VoteVoterTable
	sub := m.Voter
	if sub == nil {
		return nil, errors.New("the pro Voter is nil")
	}
	sub1 := m.Voter
	if sub1 == nil {
		return nil, errors.New("the mainkey Voter is nil")
	}
	kList := []interface{}{pre, sub, sub1}
	kBuf, cErr := kope.EncodeSlice(kList)
	return kBuf, cErr
}

//Query srt by order
//
//start = nil  end = nil (query the db from start to end)
//start = nil (query from start the db)
//end = nil (query to the end of db)
//
//f: callback for each traversal , primary 、sub key、idx(the number of times it has been iterated)
//as arguments to the callback function
//if the return value of f is true,continue iterating until the end iteration;
//otherwise stop iteration immediately
//
//lastMainKey: the main key of the last one of last page
//lastSubVal: the value  of the last one of last page
//
func (s *SVoteVoterWrap) ForEachByOrder(start *prototype.VoterId, end *prototype.VoterId, lastMainKey *prototype.VoterId,
	lastSubVal *prototype.VoterId, f func(mVal *prototype.VoterId, sVal *prototype.VoterId, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := VoteVoterTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
		if lastMainKey != nil {
			skeyList = append(skeyList, lastMainKey, kope.MinimalKey)
		}
	} else {
		if lastMainKey != nil && lastSubVal != nil {
			skeyList = append(skeyList, lastSubVal, lastMainKey, kope.MinimalKey)
		}
		skeyList = append(skeyList, kope.MinimalKey)
	}
	sBuf, cErr := kope.EncodeSlice(skeyList)
	if cErr != nil {
		return cErr
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	} else {
		eKeyList = append(eKeyList, kope.MaximumKey)
	}
	eBuf, cErr := kope.EncodeSlice(eKeyList)
	if cErr != nil {
		return cErr
	}
	var idx uint32 = 0
	s.Dba.Iterate(sBuf, eBuf, false, func(key, value []byte) bool {
		idx++
		return f(s.GetMainVal(value), s.GetSubVal(value), idx)
	})
	return nil
}

////////////// SECTION List Keys ///////////////
type SVoteVoteTimeWrap struct {
	Dba iservices.IDatabaseRW
}

func NewVoteVoteTimeWrap(db iservices.IDatabaseRW) *SVoteVoteTimeWrap {
	if db == nil {
		return nil
	}
	wrap := SVoteVoteTimeWrap{Dba: db}
	return &wrap
}

func (s *SVoteVoteTimeWrap) GetMainVal(val []byte) *prototype.VoterId {
	res := &SoListVoteByVoteTime{}
	err := proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.Voter

}

func (s *SVoteVoteTimeWrap) GetSubVal(val []byte) *prototype.TimePointSec {
	res := &SoListVoteByVoteTime{}
	err := proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
	return res.VoteTime

}

func (m *SoListVoteByVoteTime) OpeEncode() ([]byte, error) {
	pre := VoteVoteTimeTable
	sub := m.VoteTime
	if sub == nil {
		return nil, errors.New("the pro VoteTime is nil")
	}
	sub1 := m.Voter
	if sub1 == nil {
		return nil, errors.New("the mainkey Voter is nil")
	}
	kList := []interface{}{pre, sub, sub1}
	kBuf, cErr := kope.EncodeSlice(kList)
	return kBuf, cErr
}

//Query srt by order
//
//start = nil  end = nil (query the db from start to end)
//start = nil (query from start the db)
//end = nil (query to the end of db)
//
//f: callback for each traversal , primary 、sub key、idx(the number of times it has been iterated)
//as arguments to the callback function
//if the return value of f is true,continue iterating until the end iteration;
//otherwise stop iteration immediately
//
//lastMainKey: the main key of the last one of last page
//lastSubVal: the value  of the last one of last page
//
func (s *SVoteVoteTimeWrap) ForEachByOrder(start *prototype.TimePointSec, end *prototype.TimePointSec, lastMainKey *prototype.VoterId,
	lastSubVal *prototype.TimePointSec, f func(mVal *prototype.VoterId, sVal *prototype.TimePointSec, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := VoteVoteTimeTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
		if lastMainKey != nil {
			skeyList = append(skeyList, lastMainKey, kope.MinimalKey)
		}
	} else {
		if lastMainKey != nil && lastSubVal != nil {
			skeyList = append(skeyList, lastSubVal, lastMainKey, kope.MinimalKey)
		}
		skeyList = append(skeyList, kope.MinimalKey)
	}
	sBuf, cErr := kope.EncodeSlice(skeyList)
	if cErr != nil {
		return cErr
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	} else {
		eKeyList = append(eKeyList, kope.MaximumKey)
	}
	eBuf, cErr := kope.EncodeSlice(eKeyList)
	if cErr != nil {
		return cErr
	}
	var idx uint32 = 0
	s.Dba.Iterate(sBuf, eBuf, false, func(key, value []byte) bool {
		idx++
		return f(s.GetMainVal(value), s.GetSubVal(value), idx)
	})
	return nil
}

////////////// SECTION List Keys ///////////////
type SVotePostIdWrap struct {
	Dba iservices.IDatabaseRW
}

func NewVotePostIdWrap(db iservices.IDatabaseRW) *SVotePostIdWrap {
	if db == nil {
		return nil
	}
	wrap := SVotePostIdWrap{Dba: db}
	return &wrap
}

func (s *SVotePostIdWrap) GetMainVal(val []byte) *prototype.VoterId {
	res := &SoListVoteByPostId{}
	err := proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.Voter

}

func (s *SVotePostIdWrap) GetSubVal(val []byte) *uint64 {
	res := &SoListVoteByPostId{}
	err := proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
	return &res.PostId

}

func (m *SoListVoteByPostId) OpeEncode() ([]byte, error) {
	pre := VotePostIdTable
	sub := m.PostId

	sub1 := m.Voter
	if sub1 == nil {
		return nil, errors.New("the mainkey Voter is nil")
	}
	kList := []interface{}{pre, sub, sub1}
	kBuf, cErr := kope.EncodeSlice(kList)
	return kBuf, cErr
}

//Query srt by order
//
//start = nil  end = nil (query the db from start to end)
//start = nil (query from start the db)
//end = nil (query to the end of db)
//
//f: callback for each traversal , primary 、sub key、idx(the number of times it has been iterated)
//as arguments to the callback function
//if the return value of f is true,continue iterating until the end iteration;
//otherwise stop iteration immediately
//
//lastMainKey: the main key of the last one of last page
//lastSubVal: the value  of the last one of last page
//
func (s *SVotePostIdWrap) ForEachByOrder(start *uint64, end *uint64, lastMainKey *prototype.VoterId,
	lastSubVal *uint64, f func(mVal *prototype.VoterId, sVal *uint64, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := VotePostIdTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
		if lastMainKey != nil {
			skeyList = append(skeyList, lastMainKey, kope.MinimalKey)
		}
	} else {
		if lastMainKey != nil && lastSubVal != nil {
			skeyList = append(skeyList, lastSubVal, lastMainKey, kope.MinimalKey)
		}
		skeyList = append(skeyList, kope.MinimalKey)
	}
	sBuf, cErr := kope.EncodeSlice(skeyList)
	if cErr != nil {
		return cErr
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	} else {
		eKeyList = append(eKeyList, kope.MaximumKey)
	}
	eBuf, cErr := kope.EncodeSlice(eKeyList)
	if cErr != nil {
		return cErr
	}
	var idx uint32 = 0
	s.Dba.Iterate(sBuf, eBuf, false, func(key, value []byte) bool {
		idx++
		return f(s.GetMainVal(value), s.GetSubVal(value), idx)
	})
	return nil
}

/////////////// SECTION Private function ////////////////

func (s *SoVoteWrap) update(sa *SoVote) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	buf, err := proto.Marshal(sa)
	if err != nil {
		return false
	}

	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return false
	}

	return s.dba.Put(keyBuf, buf) == nil
}

func (s *SoVoteWrap) getVote() *SoVote {
	if s.dba == nil {
		return nil
	}
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return nil
	}
	resBuf, err := s.dba.Get(keyBuf)

	if err != nil {
		return nil
	}

	res := &SoVote{}
	if proto.Unmarshal(resBuf, res) != nil {
		return nil
	}
	return res
}

func (s *SoVoteWrap) encodeMainKey() ([]byte, error) {
	if s.mKeyBuf != nil {
		return s.mKeyBuf, nil
	}
	pre := s.getMemKeyPrefix("Voter")
	sub := s.mainKey
	if sub == nil {
		return nil, errors.New("the mainKey is nil")
	}
	preBuf, err := kope.Encode(pre)
	if err != nil {
		return nil, err
	}
	mBuf, err := s.getMainKeyBuf()
	if err != nil {
		return nil, err
	}
	list := make([][]byte, 2)
	list[0] = preBuf
	list[1] = mBuf
	s.mKeyBuf = kope.PackList(list)
	return s.mKeyBuf, nil
}

////////////// Unique Query delete/insert/query ///////////////

func (s *SoVoteWrap) delAllUniKeys(br bool, val *SoVote) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if !s.delUniKeyVoter(val) {
		if br {
			return false
		} else {
			res = false
		}
	}

	return res
}

func (s *SoVoteWrap) delUniKeysWithNames(names map[string]string, val *SoVote) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if len(names["Voter"]) > 0 {
		if !s.delUniKeyVoter(val) {
			res = false
		}
	}

	return res
}

func (s *SoVoteWrap) insertAllUniKeys(val *SoVote) (map[string]string, error) {
	if s.dba == nil {
		return nil, errors.New("insert uniuqe Field fail,the db is nil ")
	}
	if val == nil {
		return nil, errors.New("insert uniuqe Field fail,get the SoVote fail ")
	}
	sucFields := map[string]string{}
	if !s.insertUniKeyVoter(val) {
		return sucFields, errors.New("insert unique Field Voter fail while insert table ")
	}
	sucFields["Voter"] = "Voter"

	return sucFields, nil
}

func (s *SoVoteWrap) delUniKeyVoter(sa *SoVote) bool {
	if s.dba == nil {
		return false
	}
	pre := VoteVoterUniTable
	kList := []interface{}{pre}
	if sa != nil {

		if sa.Voter == nil {
			return false
		}

		sub := sa.Voter
		kList = append(kList, sub)
	} else {
		key, err := s.encodeMemKey("Voter")
		if err != nil {
			return false
		}
		buf, err := s.dba.Get(key)
		if err != nil {
			return false
		}
		ori := &SoMemVoteByVoter{}
		err = proto.Unmarshal(buf, ori)
		if err != nil {
			return false
		}
		sub := ori.Voter
		kList = append(kList, sub)

	}
	kBuf, err := kope.EncodeSlice(kList)
	if err != nil {
		return false
	}
	return s.dba.Delete(kBuf) == nil
}

func (s *SoVoteWrap) insertUniKeyVoter(sa *SoVote) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	pre := VoteVoterUniTable
	sub := sa.Voter
	kList := []interface{}{pre, sub}
	kBuf, err := kope.EncodeSlice(kList)
	if err != nil {
		return false
	}
	res, err := s.dba.Has(kBuf)
	if err == nil && res == true {
		//the unique key is already exist
		return false
	}
	val := SoUniqueVoteByVoter{}
	val.Voter = sa.Voter

	buf, err := proto.Marshal(&val)

	if err != nil {
		return false
	}

	return s.dba.Put(kBuf, buf) == nil

}

type UniVoteVoterWrap struct {
	Dba iservices.IDatabaseRW
}

func NewUniVoteVoterWrap(db iservices.IDatabaseRW) *UniVoteVoterWrap {
	if db == nil {
		return nil
	}
	wrap := UniVoteVoterWrap{Dba: db}
	return &wrap
}

func (s *UniVoteVoterWrap) UniQueryVoter(start *prototype.VoterId) *SoVoteWrap {
	if start == nil || s.Dba == nil {
		return nil
	}
	pre := VoteVoterUniTable
	kList := []interface{}{pre, start}
	bufStartkey, err := kope.EncodeSlice(kList)
	val, err := s.Dba.Get(bufStartkey)
	if err == nil {
		res := &SoUniqueVoteByVoter{}
		rErr := proto.Unmarshal(val, res)
		if rErr == nil {
			wrap := NewSoVoteWrap(s.Dba, res.Voter)

			return wrap
		}
	}
	return nil
}
