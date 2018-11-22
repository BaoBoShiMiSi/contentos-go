package table

import (
	"bytes"
	"errors"

	"github.com/coschain/contentos-go/common/encoding"
	"github.com/coschain/contentos-go/iservices"
	prototype "github.com/coschain/contentos-go/prototype"
	proto "github.com/golang/protobuf/proto"
)

////////////// SECTION Prefix Mark ///////////////
var (
	VoteTable         = []byte("VoteTable")
	VoteVoteTimeTable = []byte("VoteVoteTimeTable")
	VotePostIdTable   = []byte("VotePostIdTable")
	VoteVoterUniTable = []byte("VoteVoterUniTable")
)

////////////// SECTION Wrap Define ///////////////
type SoVoteWrap struct {
	dba     iservices.IDatabaseService
	mainKey *prototype.VoterId
}

func NewSoVoteWrap(dba iservices.IDatabaseService, key *prototype.VoterId) *SoVoteWrap {
	result := &SoVoteWrap{dba, key}
	return result
}

func (s *SoVoteWrap) CheckExist() bool {
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return false
	}

	res, err := s.dba.Has(keyBuf)
	if err != nil {
		return false
	}

	return res
}

func (s *SoVoteWrap) Create(f func(tInfo *SoVote)) error {
	val := &SoVote{}
	f(val)
	if val.Voter == nil {
		return errors.New("the mainkey is nil")
	}
	if s.CheckExist() {
		return errors.New("the mainkey is already exist")
	}
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return err

	}
	resBuf, err := proto.Marshal(val)
	if err != nil {
		return err
	}
	err = s.dba.Put(keyBuf, resBuf)
	if err != nil {
		return err
	}

	// update sort list keys

	if !s.insertSortKeyVoteTime(val) {
		return errors.New("insert sort Field VoteTime while insert table ")
	}

	if !s.insertSortKeyPostId(val) {
		return errors.New("insert sort Field PostId while insert table ")
	}

	//update unique list
	if !s.insertUniKeyVoter(val) {
		return errors.New("insert unique Field prototype.VoterId				 while insert table ")
	}

	return nil
}

////////////// SECTION LKeys delete/insert ///////////////

func (s *SoVoteWrap) delSortKeyVoteTime(sa *SoVote) bool {
	val := SoListVoteByVoteTime{}
	val.VoteTime = sa.VoteTime
	val.Voter = sa.Voter
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoVoteWrap) insertSortKeyVoteTime(sa *SoVote) bool {
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
	val := SoListVoteByPostId{}
	val.PostId = sa.PostId
	val.Voter = sa.Voter
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoVoteWrap) insertSortKeyPostId(sa *SoVote) bool {
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

////////////// SECTION LKeys delete/insert //////////////

func (s *SoVoteWrap) RemoveVote() bool {
	sa := s.getVote()
	if sa == nil {
		return false
	}
	//delete sort list key
	if !s.delSortKeyVoteTime(sa) {
		return false
	}
	if !s.delSortKeyPostId(sa) {
		return false
	}

	//delete unique list
	if !s.delUniKeyVoter(sa) {
		return false
	}

	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return false
	}
	return s.dba.Delete(keyBuf) == nil
}

////////////// SECTION Members Get/Modify ///////////////
func (s *SoVoteWrap) GetPostId() uint64 {
	res := s.getVote()

	if res == nil {
		var tmpValue uint64
		return tmpValue
	}
	return res.PostId
}

func (s *SoVoteWrap) MdPostId(p uint64) bool {
	sa := s.getVote()
	if sa == nil {
		return false
	}

	if !s.delSortKeyPostId(sa) {
		return false
	}
	sa.PostId = p
	if !s.update(sa) {
		return false
	}

	if !s.insertSortKeyPostId(sa) {
		return false
	}

	return true
}

func (s *SoVoteWrap) GetVoteTime() *prototype.TimePointSec {
	res := s.getVote()

	if res == nil {
		return nil

	}
	return res.VoteTime
}

func (s *SoVoteWrap) MdVoteTime(p *prototype.TimePointSec) bool {
	sa := s.getVote()
	if sa == nil {
		return false
	}

	if !s.delSortKeyVoteTime(sa) {
		return false
	}
	sa.VoteTime = p
	if !s.update(sa) {
		return false
	}

	if !s.insertSortKeyVoteTime(sa) {
		return false
	}

	return true
}

func (s *SoVoteWrap) GetVoter() *prototype.VoterId {
	res := s.getVote()

	if res == nil {
		return nil

	}
	return res.Voter
}

////////////// SECTION List Keys ///////////////
type SVoteVoteTimeWrap struct {
	Dba iservices.IDatabaseService
}

func (s *SVoteVoteTimeWrap) DelIterater(iterator iservices.IDatabaseIterator) {
	if iterator == nil || !iterator.Valid() {
		return
	}
	s.Dba.DeleteIterator(iterator)
}

func (s *SVoteVoteTimeWrap) GetMainVal(iterator iservices.IDatabaseIterator) *prototype.VoterId {
	if iterator == nil || !iterator.Valid() {
		return nil
	}
	val, err := iterator.Value()

	if err != nil {
		return nil
	}

	res := &SoListVoteByVoteTime{}
	err = proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.Voter

}

func (s *SVoteVoteTimeWrap) GetSubVal(iterator iservices.IDatabaseIterator) *prototype.TimePointSec {
	if iterator == nil || !iterator.Valid() {
		return nil
	}

	val, err := iterator.Value()

	if err != nil {
		return nil
	}
	res := &SoListVoteByVoteTime{}
	err = proto.Unmarshal(val, res)
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
	kBuf, cErr := encoding.EncodeSlice(kList, false)
	return kBuf, cErr
}

//Query sort by order
//start = nil  end = nil (query the db from start to end)
//start = nil (query from start the db)
//end = nil (query to the end of db)
func (s *SVoteVoteTimeWrap) QueryListByOrder(start *prototype.TimePointSec, end *prototype.TimePointSec) iservices.IDatabaseIterator {
	pre := VoteVoteTimeTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
	}
	sBuf, cErr := encoding.EncodeSlice(skeyList, false)
	if cErr != nil {
		return nil
	}
	if start != nil && end == nil {
		iter := s.Dba.NewIterator(sBuf, nil)
		return iter
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	}
	eBuf, cErr := encoding.EncodeSlice(eKeyList, false)
	if cErr != nil {
		return nil
	}

	res := bytes.Compare(sBuf, eBuf)
	if res == 0 {
		eBuf = nil
	} else if res == 1 {
		//reverse order
		return nil
	}
	iter := s.Dba.NewIterator(sBuf, eBuf)

	return iter
}

////////////// SECTION List Keys ///////////////
type SVotePostIdWrap struct {
	Dba iservices.IDatabaseService
}

func (s *SVotePostIdWrap) DelIterater(iterator iservices.IDatabaseIterator) {
	if iterator == nil || !iterator.Valid() {
		return
	}
	s.Dba.DeleteIterator(iterator)
}

func (s *SVotePostIdWrap) GetMainVal(iterator iservices.IDatabaseIterator) *prototype.VoterId {
	if iterator == nil || !iterator.Valid() {
		return nil
	}
	val, err := iterator.Value()

	if err != nil {
		return nil
	}

	res := &SoListVoteByPostId{}
	err = proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.Voter

}

func (s *SVotePostIdWrap) GetSubVal(iterator iservices.IDatabaseIterator) *uint64 {
	if iterator == nil || !iterator.Valid() {
		return nil
	}

	val, err := iterator.Value()

	if err != nil {
		return nil
	}
	res := &SoListVoteByPostId{}
	err = proto.Unmarshal(val, res)
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
	kBuf, cErr := encoding.EncodeSlice(kList, false)
	return kBuf, cErr
}

//Query sort by order
//start = nil  end = nil (query the db from start to end)
//start = nil (query from start the db)
//end = nil (query to the end of db)
func (s *SVotePostIdWrap) QueryListByOrder(start *uint64, end *uint64) iservices.IDatabaseIterator {
	pre := VotePostIdTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
	}
	sBuf, cErr := encoding.EncodeSlice(skeyList, false)
	if cErr != nil {
		return nil
	}
	if start != nil && end == nil {
		iter := s.Dba.NewIterator(sBuf, nil)
		return iter
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	}
	eBuf, cErr := encoding.EncodeSlice(eKeyList, false)
	if cErr != nil {
		return nil
	}

	res := bytes.Compare(sBuf, eBuf)
	if res == 0 {
		eBuf = nil
	} else if res == 1 {
		//reverse order
		return nil
	}
	iter := s.Dba.NewIterator(sBuf, eBuf)

	return iter
}

/////////////// SECTION Private function ////////////////

func (s *SoVoteWrap) update(sa *SoVote) bool {
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
	pre := VoteTable
	sub := s.mainKey
	if sub == nil {
		return nil, errors.New("the mainKey is nil")
	}
	kList := []interface{}{pre, sub}
	kBuf, cErr := encoding.EncodeSlice(kList, false)
	return kBuf, cErr
}

////////////// Unique Query delete/insert/query ///////////////

func (s *SoVoteWrap) delUniKeyVoter(sa *SoVote) bool {
	pre := VoteVoterUniTable
	sub := sa.Voter
	kList := []interface{}{pre, sub}
	kBuf, err := encoding.EncodeSlice(kList, false)
	if err != nil {
		return false
	}
	return s.dba.Delete(kBuf) == nil
}

func (s *SoVoteWrap) insertUniKeyVoter(sa *SoVote) bool {
	uniWrap := UniVoteVoterWrap{}
	uniWrap.Dba = s.dba

	res := uniWrap.UniQueryVoter(sa.Voter)
	if res != nil {
		//the unique key is already exist
		return false
	}
	val := SoUniqueVoteByVoter{}
	val.Voter = sa.Voter

	buf, err := proto.Marshal(&val)

	if err != nil {
		return false
	}

	pre := VoteVoterUniTable
	sub := sa.Voter
	kList := []interface{}{pre, sub}
	kBuf, err := encoding.EncodeSlice(kList, false)
	if err != nil {
		return false
	}
	return s.dba.Put(kBuf, buf) == nil

}

type UniVoteVoterWrap struct {
	Dba iservices.IDatabaseService
}

func (s *UniVoteVoterWrap) UniQueryVoter(start *prototype.VoterId) *SoVoteWrap {
	pre := VoteVoterUniTable
	kList := []interface{}{pre, start}
	bufStartkey, err := encoding.EncodeSlice(kList, false)
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