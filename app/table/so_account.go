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
	AccountCreatedTimeTable        uint32 = 2128286283
	AccountBalanceTable            uint32 = 4012029019
	AccountVestingSharesTable      uint32 = 3830877790
	AccountBpVoteCountTable        uint32 = 2264397557
	AccountNameUniTable            uint32 = 2528390520
	AccountBalanceCell             uint32 = 2894785396
	AccountBpVoteCountCell         uint32 = 2131409895
	AccountCreatedTimeCell         uint32 = 826305594
	AccountCreatorCell             uint32 = 1804791917
	AccountLastPostTimeCell        uint32 = 3226532373
	AccountLastStakeTimeCell       uint32 = 3774075190
	AccountLastVoteTimeCell        uint32 = 1980371646
	AccountNameCell                uint32 = 1725869739
	AccountStakeVestingCell        uint32 = 1603133992
	AccountStaminaCell             uint32 = 674022235
	AccountStaminaFreeCell         uint32 = 676517039
	AccountStaminaFreeUseBlockCell uint32 = 985510361
	AccountStaminaUseBlockCell     uint32 = 3536676248
	AccountVestingSharesCell       uint32 = 57659323
	AccountVotePowerCell           uint32 = 2246508735
)

////////////// SECTION Wrap Define ///////////////
type SoAccountWrap struct {
	dba      iservices.IDatabaseService
	mainKey  *prototype.AccountName
	mKeyFlag int    //the flag of the main key exist state in db, -1:has not judged; 0:not exist; 1:already exist
	mKeyBuf  []byte //the buffer after the main key is encoded with prefix
	mBuf     []byte //the value after the main key is encoded
}

func NewSoAccountWrap(dba iservices.IDatabaseService, key *prototype.AccountName) *SoAccountWrap {
	if dba == nil || key == nil {
		return nil
	}
	result := &SoAccountWrap{dba, key, -1, nil, nil}
	return result
}

func (s *SoAccountWrap) CheckExist() bool {
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

func (s *SoAccountWrap) Create(f func(tInfo *SoAccount)) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if s.mainKey == nil {
		return errors.New("the main key is nil")
	}
	val := &SoAccount{}
	f(val)
	if val.Name == nil {
		val.Name = s.mainKey
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

func (s *SoAccountWrap) getMainKeyBuf() ([]byte, error) {
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

func (s *SoAccountWrap) delSortKeyCreatedTime(sa *SoAccount) bool {
	if s.dba == nil || s.mainKey == nil {
		return false
	}
	val := SoListAccountByCreatedTime{}
	if sa == nil {
		key, err := s.encodeMemKey("CreatedTime")
		if err != nil {
			return false
		}
		buf, err := s.dba.Get(key)
		if err != nil {
			return false
		}
		ori := &SoMemAccountByCreatedTime{}
		err = proto.Unmarshal(buf, ori)
		if err != nil {
			return false
		}
		val.CreatedTime = ori.CreatedTime
		val.Name = s.mainKey

	} else {
		val.CreatedTime = sa.CreatedTime
		val.Name = sa.Name
	}

	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoAccountWrap) insertSortKeyCreatedTime(sa *SoAccount) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	val := SoListAccountByCreatedTime{}
	val.Name = sa.Name
	val.CreatedTime = sa.CreatedTime
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

func (s *SoAccountWrap) delSortKeyBalance(sa *SoAccount) bool {
	if s.dba == nil || s.mainKey == nil {
		return false
	}
	val := SoListAccountByBalance{}
	if sa == nil {
		key, err := s.encodeMemKey("Balance")
		if err != nil {
			return false
		}
		buf, err := s.dba.Get(key)
		if err != nil {
			return false
		}
		ori := &SoMemAccountByBalance{}
		err = proto.Unmarshal(buf, ori)
		if err != nil {
			return false
		}
		val.Balance = ori.Balance
		val.Name = s.mainKey

	} else {
		val.Balance = sa.Balance
		val.Name = sa.Name
	}

	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoAccountWrap) insertSortKeyBalance(sa *SoAccount) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	val := SoListAccountByBalance{}
	val.Name = sa.Name
	val.Balance = sa.Balance
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

func (s *SoAccountWrap) delSortKeyVestingShares(sa *SoAccount) bool {
	if s.dba == nil || s.mainKey == nil {
		return false
	}
	val := SoListAccountByVestingShares{}
	if sa == nil {
		key, err := s.encodeMemKey("VestingShares")
		if err != nil {
			return false
		}
		buf, err := s.dba.Get(key)
		if err != nil {
			return false
		}
		ori := &SoMemAccountByVestingShares{}
		err = proto.Unmarshal(buf, ori)
		if err != nil {
			return false
		}
		val.VestingShares = ori.VestingShares
		val.Name = s.mainKey

	} else {
		val.VestingShares = sa.VestingShares
		val.Name = sa.Name
	}

	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoAccountWrap) insertSortKeyVestingShares(sa *SoAccount) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	val := SoListAccountByVestingShares{}
	val.Name = sa.Name
	val.VestingShares = sa.VestingShares
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

func (s *SoAccountWrap) delSortKeyBpVoteCount(sa *SoAccount) bool {
	if s.dba == nil || s.mainKey == nil {
		return false
	}
	val := SoListAccountByBpVoteCount{}
	if sa == nil {
		key, err := s.encodeMemKey("BpVoteCount")
		if err != nil {
			return false
		}
		buf, err := s.dba.Get(key)
		if err != nil {
			return false
		}
		ori := &SoMemAccountByBpVoteCount{}
		err = proto.Unmarshal(buf, ori)
		if err != nil {
			return false
		}
		val.BpVoteCount = ori.BpVoteCount
		val.Name = s.mainKey

	} else {
		val.BpVoteCount = sa.BpVoteCount
		val.Name = sa.Name
	}

	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoAccountWrap) insertSortKeyBpVoteCount(sa *SoAccount) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	val := SoListAccountByBpVoteCount{}
	val.Name = sa.Name
	val.BpVoteCount = sa.BpVoteCount
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

func (s *SoAccountWrap) delAllSortKeys(br bool, val *SoAccount) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if !s.delSortKeyCreatedTime(val) {
		if br {
			return false
		} else {
			res = false
		}
	}
	if !s.delSortKeyBalance(val) {
		if br {
			return false
		} else {
			res = false
		}
	}
	if !s.delSortKeyVestingShares(val) {
		if br {
			return false
		} else {
			res = false
		}
	}
	if !s.delSortKeyBpVoteCount(val) {
		if br {
			return false
		} else {
			res = false
		}
	}

	return res
}

func (s *SoAccountWrap) insertAllSortKeys(val *SoAccount) error {
	if s.dba == nil {
		return errors.New("insert sort Field fail,the db is nil ")
	}
	if val == nil {
		return errors.New("insert sort Field fail,get the SoAccount fail ")
	}
	if !s.insertSortKeyCreatedTime(val) {
		return errors.New("insert sort Field CreatedTime fail while insert table ")
	}
	if !s.insertSortKeyBalance(val) {
		return errors.New("insert sort Field Balance fail while insert table ")
	}
	if !s.insertSortKeyVestingShares(val) {
		return errors.New("insert sort Field VestingShares fail while insert table ")
	}
	if !s.insertSortKeyBpVoteCount(val) {
		return errors.New("insert sort Field BpVoteCount fail while insert table ")
	}

	return nil
}

////////////// SECTION LKeys delete/insert //////////////

func (s *SoAccountWrap) RemoveAccount() bool {
	if s.dba == nil {
		return false
	}
	val := &SoAccount{}
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
func (s *SoAccountWrap) getMemKeyPrefix(fName string) uint32 {
	if fName == "Balance" {
		return AccountBalanceCell
	}
	if fName == "BpVoteCount" {
		return AccountBpVoteCountCell
	}
	if fName == "CreatedTime" {
		return AccountCreatedTimeCell
	}
	if fName == "Creator" {
		return AccountCreatorCell
	}
	if fName == "LastPostTime" {
		return AccountLastPostTimeCell
	}
	if fName == "LastStakeTime" {
		return AccountLastStakeTimeCell
	}
	if fName == "LastVoteTime" {
		return AccountLastVoteTimeCell
	}
	if fName == "Name" {
		return AccountNameCell
	}
	if fName == "StakeVesting" {
		return AccountStakeVestingCell
	}
	if fName == "Stamina" {
		return AccountStaminaCell
	}
	if fName == "StaminaFree" {
		return AccountStaminaFreeCell
	}
	if fName == "StaminaFreeUseBlock" {
		return AccountStaminaFreeUseBlockCell
	}
	if fName == "StaminaUseBlock" {
		return AccountStaminaUseBlockCell
	}
	if fName == "VestingShares" {
		return AccountVestingSharesCell
	}
	if fName == "VotePower" {
		return AccountVotePowerCell
	}

	return 0
}

func (s *SoAccountWrap) encodeMemKey(fName string) ([]byte, error) {
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

func (s *SoAccountWrap) saveAllMemKeys(tInfo *SoAccount, br bool) error {
	if s.dba == nil {
		return errors.New("save member Field fail , the db is nil")
	}

	if tInfo == nil {
		return errors.New("save member Field fail , the data is nil ")
	}
	var err error = nil
	errDes := ""
	if err = s.saveMemKeyBalance(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "Balance", err)
		}
	}
	if err = s.saveMemKeyBpVoteCount(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "BpVoteCount", err)
		}
	}
	if err = s.saveMemKeyCreatedTime(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "CreatedTime", err)
		}
	}
	if err = s.saveMemKeyCreator(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "Creator", err)
		}
	}
	if err = s.saveMemKeyLastPostTime(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "LastPostTime", err)
		}
	}
	if err = s.saveMemKeyLastStakeTime(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "LastStakeTime", err)
		}
	}
	if err = s.saveMemKeyLastVoteTime(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "LastVoteTime", err)
		}
	}
	if err = s.saveMemKeyName(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "Name", err)
		}
	}
	if err = s.saveMemKeyStakeVesting(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "StakeVesting", err)
		}
	}
	if err = s.saveMemKeyStamina(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "Stamina", err)
		}
	}
	if err = s.saveMemKeyStaminaFree(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "StaminaFree", err)
		}
	}
	if err = s.saveMemKeyStaminaFreeUseBlock(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "StaminaFreeUseBlock", err)
		}
	}
	if err = s.saveMemKeyStaminaUseBlock(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "StaminaUseBlock", err)
		}
	}
	if err = s.saveMemKeyVestingShares(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "VestingShares", err)
		}
	}
	if err = s.saveMemKeyVotePower(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "VotePower", err)
		}
	}

	if len(errDes) > 0 {
		return errors.New(errDes)
	}
	return err
}

func (s *SoAccountWrap) delAllMemKeys(br bool, tInfo *SoAccount) error {
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

func (s *SoAccountWrap) delMemKey(fName string) error {
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

func (s *SoAccountWrap) saveMemKeyBalance(tInfo *SoAccount) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemAccountByBalance{}
	val.Balance = tInfo.Balance
	key, err := s.encodeMemKey("Balance")
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

func (s *SoAccountWrap) GetBalance() *prototype.Coin {
	res := true
	msg := &SoMemAccountByBalance{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("Balance")
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
				return msg.Balance
			}
		}
	}
	if !res {
		return nil

	}
	return msg.Balance
}

func (s *SoAccountWrap) MdBalance(p *prototype.Coin) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("Balance")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemAccountByBalance{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoAccount{}
	sa.Name = s.mainKey

	sa.Balance = ori.Balance

	if !s.delSortKeyBalance(sa) {
		return false
	}
	ori.Balance = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.Balance = p

	if !s.insertSortKeyBalance(sa) {
		return false
	}

	return true
}

func (s *SoAccountWrap) saveMemKeyBpVoteCount(tInfo *SoAccount) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemAccountByBpVoteCount{}
	val.BpVoteCount = tInfo.BpVoteCount
	key, err := s.encodeMemKey("BpVoteCount")
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

func (s *SoAccountWrap) GetBpVoteCount() uint32 {
	res := true
	msg := &SoMemAccountByBpVoteCount{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("BpVoteCount")
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
				return msg.BpVoteCount
			}
		}
	}
	if !res {
		var tmpValue uint32
		return tmpValue
	}
	return msg.BpVoteCount
}

func (s *SoAccountWrap) MdBpVoteCount(p uint32) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("BpVoteCount")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemAccountByBpVoteCount{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoAccount{}
	sa.Name = s.mainKey

	sa.BpVoteCount = ori.BpVoteCount

	if !s.delSortKeyBpVoteCount(sa) {
		return false
	}
	ori.BpVoteCount = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.BpVoteCount = p

	if !s.insertSortKeyBpVoteCount(sa) {
		return false
	}

	return true
}

func (s *SoAccountWrap) saveMemKeyCreatedTime(tInfo *SoAccount) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemAccountByCreatedTime{}
	val.CreatedTime = tInfo.CreatedTime
	key, err := s.encodeMemKey("CreatedTime")
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

func (s *SoAccountWrap) GetCreatedTime() *prototype.TimePointSec {
	res := true
	msg := &SoMemAccountByCreatedTime{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("CreatedTime")
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
				return msg.CreatedTime
			}
		}
	}
	if !res {
		return nil

	}
	return msg.CreatedTime
}

func (s *SoAccountWrap) MdCreatedTime(p *prototype.TimePointSec) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("CreatedTime")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemAccountByCreatedTime{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoAccount{}
	sa.Name = s.mainKey

	sa.CreatedTime = ori.CreatedTime

	if !s.delSortKeyCreatedTime(sa) {
		return false
	}
	ori.CreatedTime = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.CreatedTime = p

	if !s.insertSortKeyCreatedTime(sa) {
		return false
	}

	return true
}

func (s *SoAccountWrap) saveMemKeyCreator(tInfo *SoAccount) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemAccountByCreator{}
	val.Creator = tInfo.Creator
	key, err := s.encodeMemKey("Creator")
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

func (s *SoAccountWrap) GetCreator() *prototype.AccountName {
	res := true
	msg := &SoMemAccountByCreator{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("Creator")
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
				return msg.Creator
			}
		}
	}
	if !res {
		return nil

	}
	return msg.Creator
}

func (s *SoAccountWrap) MdCreator(p *prototype.AccountName) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("Creator")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemAccountByCreator{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoAccount{}
	sa.Name = s.mainKey

	sa.Creator = ori.Creator

	ori.Creator = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.Creator = p

	return true
}

func (s *SoAccountWrap) saveMemKeyLastPostTime(tInfo *SoAccount) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemAccountByLastPostTime{}
	val.LastPostTime = tInfo.LastPostTime
	key, err := s.encodeMemKey("LastPostTime")
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

func (s *SoAccountWrap) GetLastPostTime() *prototype.TimePointSec {
	res := true
	msg := &SoMemAccountByLastPostTime{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("LastPostTime")
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
				return msg.LastPostTime
			}
		}
	}
	if !res {
		return nil

	}
	return msg.LastPostTime
}

func (s *SoAccountWrap) MdLastPostTime(p *prototype.TimePointSec) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("LastPostTime")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemAccountByLastPostTime{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoAccount{}
	sa.Name = s.mainKey

	sa.LastPostTime = ori.LastPostTime

	ori.LastPostTime = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.LastPostTime = p

	return true
}

func (s *SoAccountWrap) saveMemKeyLastStakeTime(tInfo *SoAccount) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemAccountByLastStakeTime{}
	val.LastStakeTime = tInfo.LastStakeTime
	key, err := s.encodeMemKey("LastStakeTime")
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

func (s *SoAccountWrap) GetLastStakeTime() *prototype.TimePointSec {
	res := true
	msg := &SoMemAccountByLastStakeTime{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("LastStakeTime")
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
				return msg.LastStakeTime
			}
		}
	}
	if !res {
		return nil

	}
	return msg.LastStakeTime
}

func (s *SoAccountWrap) MdLastStakeTime(p *prototype.TimePointSec) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("LastStakeTime")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemAccountByLastStakeTime{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoAccount{}
	sa.Name = s.mainKey

	sa.LastStakeTime = ori.LastStakeTime

	ori.LastStakeTime = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.LastStakeTime = p

	return true
}

func (s *SoAccountWrap) saveMemKeyLastVoteTime(tInfo *SoAccount) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemAccountByLastVoteTime{}
	val.LastVoteTime = tInfo.LastVoteTime
	key, err := s.encodeMemKey("LastVoteTime")
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

func (s *SoAccountWrap) GetLastVoteTime() *prototype.TimePointSec {
	res := true
	msg := &SoMemAccountByLastVoteTime{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("LastVoteTime")
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
				return msg.LastVoteTime
			}
		}
	}
	if !res {
		return nil

	}
	return msg.LastVoteTime
}

func (s *SoAccountWrap) MdLastVoteTime(p *prototype.TimePointSec) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("LastVoteTime")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemAccountByLastVoteTime{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoAccount{}
	sa.Name = s.mainKey

	sa.LastVoteTime = ori.LastVoteTime

	ori.LastVoteTime = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.LastVoteTime = p

	return true
}

func (s *SoAccountWrap) saveMemKeyName(tInfo *SoAccount) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemAccountByName{}
	val.Name = tInfo.Name
	key, err := s.encodeMemKey("Name")
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

func (s *SoAccountWrap) GetName() *prototype.AccountName {
	res := true
	msg := &SoMemAccountByName{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("Name")
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
				return msg.Name
			}
		}
	}
	if !res {
		return nil

	}
	return msg.Name
}

func (s *SoAccountWrap) saveMemKeyStakeVesting(tInfo *SoAccount) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemAccountByStakeVesting{}
	val.StakeVesting = tInfo.StakeVesting
	key, err := s.encodeMemKey("StakeVesting")
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

func (s *SoAccountWrap) GetStakeVesting() *prototype.Vest {
	res := true
	msg := &SoMemAccountByStakeVesting{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("StakeVesting")
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
				return msg.StakeVesting
			}
		}
	}
	if !res {
		return nil

	}
	return msg.StakeVesting
}

func (s *SoAccountWrap) MdStakeVesting(p *prototype.Vest) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("StakeVesting")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemAccountByStakeVesting{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoAccount{}
	sa.Name = s.mainKey

	sa.StakeVesting = ori.StakeVesting

	ori.StakeVesting = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.StakeVesting = p

	return true
}

func (s *SoAccountWrap) saveMemKeyStamina(tInfo *SoAccount) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemAccountByStamina{}
	val.Stamina = tInfo.Stamina
	key, err := s.encodeMemKey("Stamina")
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

func (s *SoAccountWrap) GetStamina() uint64 {
	res := true
	msg := &SoMemAccountByStamina{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("Stamina")
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
				return msg.Stamina
			}
		}
	}
	if !res {
		var tmpValue uint64
		return tmpValue
	}
	return msg.Stamina
}

func (s *SoAccountWrap) MdStamina(p uint64) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("Stamina")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemAccountByStamina{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoAccount{}
	sa.Name = s.mainKey

	sa.Stamina = ori.Stamina

	ori.Stamina = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.Stamina = p

	return true
}

func (s *SoAccountWrap) saveMemKeyStaminaFree(tInfo *SoAccount) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemAccountByStaminaFree{}
	val.StaminaFree = tInfo.StaminaFree
	key, err := s.encodeMemKey("StaminaFree")
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

func (s *SoAccountWrap) GetStaminaFree() uint64 {
	res := true
	msg := &SoMemAccountByStaminaFree{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("StaminaFree")
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
				return msg.StaminaFree
			}
		}
	}
	if !res {
		var tmpValue uint64
		return tmpValue
	}
	return msg.StaminaFree
}

func (s *SoAccountWrap) MdStaminaFree(p uint64) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("StaminaFree")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemAccountByStaminaFree{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoAccount{}
	sa.Name = s.mainKey

	sa.StaminaFree = ori.StaminaFree

	ori.StaminaFree = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.StaminaFree = p

	return true
}

func (s *SoAccountWrap) saveMemKeyStaminaFreeUseBlock(tInfo *SoAccount) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemAccountByStaminaFreeUseBlock{}
	val.StaminaFreeUseBlock = tInfo.StaminaFreeUseBlock
	key, err := s.encodeMemKey("StaminaFreeUseBlock")
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

func (s *SoAccountWrap) GetStaminaFreeUseBlock() uint64 {
	res := true
	msg := &SoMemAccountByStaminaFreeUseBlock{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("StaminaFreeUseBlock")
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
				return msg.StaminaFreeUseBlock
			}
		}
	}
	if !res {
		var tmpValue uint64
		return tmpValue
	}
	return msg.StaminaFreeUseBlock
}

func (s *SoAccountWrap) MdStaminaFreeUseBlock(p uint64) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("StaminaFreeUseBlock")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemAccountByStaminaFreeUseBlock{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoAccount{}
	sa.Name = s.mainKey

	sa.StaminaFreeUseBlock = ori.StaminaFreeUseBlock

	ori.StaminaFreeUseBlock = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.StaminaFreeUseBlock = p

	return true
}

func (s *SoAccountWrap) saveMemKeyStaminaUseBlock(tInfo *SoAccount) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemAccountByStaminaUseBlock{}
	val.StaminaUseBlock = tInfo.StaminaUseBlock
	key, err := s.encodeMemKey("StaminaUseBlock")
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

func (s *SoAccountWrap) GetStaminaUseBlock() uint64 {
	res := true
	msg := &SoMemAccountByStaminaUseBlock{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("StaminaUseBlock")
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
				return msg.StaminaUseBlock
			}
		}
	}
	if !res {
		var tmpValue uint64
		return tmpValue
	}
	return msg.StaminaUseBlock
}

func (s *SoAccountWrap) MdStaminaUseBlock(p uint64) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("StaminaUseBlock")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemAccountByStaminaUseBlock{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoAccount{}
	sa.Name = s.mainKey

	sa.StaminaUseBlock = ori.StaminaUseBlock

	ori.StaminaUseBlock = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.StaminaUseBlock = p

	return true
}

func (s *SoAccountWrap) saveMemKeyVestingShares(tInfo *SoAccount) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemAccountByVestingShares{}
	val.VestingShares = tInfo.VestingShares
	key, err := s.encodeMemKey("VestingShares")
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

func (s *SoAccountWrap) GetVestingShares() *prototype.Vest {
	res := true
	msg := &SoMemAccountByVestingShares{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("VestingShares")
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
				return msg.VestingShares
			}
		}
	}
	if !res {
		return nil

	}
	return msg.VestingShares
}

func (s *SoAccountWrap) MdVestingShares(p *prototype.Vest) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("VestingShares")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemAccountByVestingShares{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoAccount{}
	sa.Name = s.mainKey

	sa.VestingShares = ori.VestingShares

	if !s.delSortKeyVestingShares(sa) {
		return false
	}
	ori.VestingShares = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.VestingShares = p

	if !s.insertSortKeyVestingShares(sa) {
		return false
	}

	return true
}

func (s *SoAccountWrap) saveMemKeyVotePower(tInfo *SoAccount) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemAccountByVotePower{}
	val.VotePower = tInfo.VotePower
	key, err := s.encodeMemKey("VotePower")
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

func (s *SoAccountWrap) GetVotePower() uint32 {
	res := true
	msg := &SoMemAccountByVotePower{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("VotePower")
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
				return msg.VotePower
			}
		}
	}
	if !res {
		var tmpValue uint32
		return tmpValue
	}
	return msg.VotePower
}

func (s *SoAccountWrap) MdVotePower(p uint32) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("VotePower")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemAccountByVotePower{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoAccount{}
	sa.Name = s.mainKey

	sa.VotePower = ori.VotePower

	ori.VotePower = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.VotePower = p

	return true
}

////////////// SECTION List Keys ///////////////
type SAccountCreatedTimeWrap struct {
	Dba iservices.IDatabaseService
}

func NewAccountCreatedTimeWrap(db iservices.IDatabaseService) *SAccountCreatedTimeWrap {
	if db == nil {
		return nil
	}
	wrap := SAccountCreatedTimeWrap{Dba: db}
	return &wrap
}

func (s *SAccountCreatedTimeWrap) DelIterator(iterator iservices.IDatabaseIterator) {
	if iterator == nil || !iterator.Valid() {
		return
	}
	s.Dba.DeleteIterator(iterator)
}

func (s *SAccountCreatedTimeWrap) GetMainVal(iterator iservices.IDatabaseIterator) *prototype.AccountName {
	if iterator == nil || !iterator.Valid() {
		return nil
	}
	val, err := iterator.Value()

	if err != nil {
		return nil
	}

	res := &SoListAccountByCreatedTime{}
	err = proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.Name

}

func (s *SAccountCreatedTimeWrap) GetSubVal(iterator iservices.IDatabaseIterator) *prototype.TimePointSec {
	if iterator == nil || !iterator.Valid() {
		return nil
	}

	val, err := iterator.Value()

	if err != nil {
		return nil
	}
	res := &SoListAccountByCreatedTime{}
	err = proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
	return res.CreatedTime

}

func (m *SoListAccountByCreatedTime) OpeEncode() ([]byte, error) {
	pre := AccountCreatedTimeTable
	sub := m.CreatedTime
	if sub == nil {
		return nil, errors.New("the pro CreatedTime is nil")
	}
	sub1 := m.Name
	if sub1 == nil {
		return nil, errors.New("the mainkey Name is nil")
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
func (s *SAccountCreatedTimeWrap) ForEachByOrder(start *prototype.TimePointSec, end *prototype.TimePointSec, lastMainKey *prototype.AccountName,
	lastSubVal *prototype.TimePointSec, f func(mVal *prototype.AccountName, sVal *prototype.TimePointSec, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := AccountCreatedTimeTable
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
	iterator := s.Dba.NewIterator(sBuf, eBuf)
	if iterator == nil {
		return errors.New("there is no data in range")
	}
	var idx uint32 = 0
	for iterator.Next() {
		idx++
		if isContinue := f(s.GetMainVal(iterator), s.GetSubVal(iterator), idx); !isContinue {
			break
		}
	}
	s.DelIterator(iterator)
	return nil
}

////////////// SECTION List Keys ///////////////
type SAccountBalanceWrap struct {
	Dba iservices.IDatabaseService
}

func NewAccountBalanceWrap(db iservices.IDatabaseService) *SAccountBalanceWrap {
	if db == nil {
		return nil
	}
	wrap := SAccountBalanceWrap{Dba: db}
	return &wrap
}

func (s *SAccountBalanceWrap) DelIterator(iterator iservices.IDatabaseIterator) {
	if iterator == nil || !iterator.Valid() {
		return
	}
	s.Dba.DeleteIterator(iterator)
}

func (s *SAccountBalanceWrap) GetMainVal(iterator iservices.IDatabaseIterator) *prototype.AccountName {
	if iterator == nil || !iterator.Valid() {
		return nil
	}
	val, err := iterator.Value()

	if err != nil {
		return nil
	}

	res := &SoListAccountByBalance{}
	err = proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.Name

}

func (s *SAccountBalanceWrap) GetSubVal(iterator iservices.IDatabaseIterator) *prototype.Coin {
	if iterator == nil || !iterator.Valid() {
		return nil
	}

	val, err := iterator.Value()

	if err != nil {
		return nil
	}
	res := &SoListAccountByBalance{}
	err = proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
	return res.Balance

}

func (m *SoListAccountByBalance) OpeEncode() ([]byte, error) {
	pre := AccountBalanceTable
	sub := m.Balance
	if sub == nil {
		return nil, errors.New("the pro Balance is nil")
	}
	sub1 := m.Name
	if sub1 == nil {
		return nil, errors.New("the mainkey Name is nil")
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
func (s *SAccountBalanceWrap) ForEachByOrder(start *prototype.Coin, end *prototype.Coin, lastMainKey *prototype.AccountName,
	lastSubVal *prototype.Coin, f func(mVal *prototype.AccountName, sVal *prototype.Coin, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := AccountBalanceTable
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
	iterator := s.Dba.NewIterator(sBuf, eBuf)
	if iterator == nil {
		return errors.New("there is no data in range")
	}
	var idx uint32 = 0
	for iterator.Next() {
		idx++
		if isContinue := f(s.GetMainVal(iterator), s.GetSubVal(iterator), idx); !isContinue {
			break
		}
	}
	s.DelIterator(iterator)
	return nil
}

//Query srt by reverse order
//
//f: callback for each traversal , primary 、sub key、idx(the number of times it has been iterated)
//as arguments to the callback function
//if the return value of f is true,continue iterating until the end iteration;
//otherwise stop iteration immediately
//
//lastMainKey: the main key of the last one of last page
//lastSubVal: the value  of the last one of last page
//
func (s *SAccountBalanceWrap) ForEachByRevOrder(start *prototype.Coin, end *prototype.Coin, lastMainKey *prototype.AccountName,
	lastSubVal *prototype.Coin, f func(mVal *prototype.AccountName, sVal *prototype.Coin, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := AccountBalanceTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
		if lastMainKey != nil {
			skeyList = append(skeyList, lastMainKey)
		}
	} else {
		if lastMainKey != nil && lastSubVal != nil {
			skeyList = append(skeyList, lastSubVal, lastMainKey)
		}
		skeyList = append(skeyList, kope.MaximumKey)
	}
	sBuf, cErr := kope.EncodeSlice(skeyList)
	if cErr != nil {
		return cErr
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	}
	eBuf, cErr := kope.EncodeSlice(eKeyList)
	if cErr != nil {
		return cErr
	}
	//reverse the start and end when create ReversedIterator to query by reverse order
	iterator := s.Dba.NewReversedIterator(eBuf, sBuf)
	if iterator == nil {
		return errors.New("there is no data in range")
	}
	var idx uint32 = 0
	for iterator.Next() {
		idx++
		if isContinue := f(s.GetMainVal(iterator), s.GetSubVal(iterator), idx); !isContinue {
			break
		}
	}
	s.DelIterator(iterator)
	return nil
}

////////////// SECTION List Keys ///////////////
type SAccountVestingSharesWrap struct {
	Dba iservices.IDatabaseService
}

func NewAccountVestingSharesWrap(db iservices.IDatabaseService) *SAccountVestingSharesWrap {
	if db == nil {
		return nil
	}
	wrap := SAccountVestingSharesWrap{Dba: db}
	return &wrap
}

func (s *SAccountVestingSharesWrap) DelIterator(iterator iservices.IDatabaseIterator) {
	if iterator == nil || !iterator.Valid() {
		return
	}
	s.Dba.DeleteIterator(iterator)
}

func (s *SAccountVestingSharesWrap) GetMainVal(iterator iservices.IDatabaseIterator) *prototype.AccountName {
	if iterator == nil || !iterator.Valid() {
		return nil
	}
	val, err := iterator.Value()

	if err != nil {
		return nil
	}

	res := &SoListAccountByVestingShares{}
	err = proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.Name

}

func (s *SAccountVestingSharesWrap) GetSubVal(iterator iservices.IDatabaseIterator) *prototype.Vest {
	if iterator == nil || !iterator.Valid() {
		return nil
	}

	val, err := iterator.Value()

	if err != nil {
		return nil
	}
	res := &SoListAccountByVestingShares{}
	err = proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
	return res.VestingShares

}

func (m *SoListAccountByVestingShares) OpeEncode() ([]byte, error) {
	pre := AccountVestingSharesTable
	sub := m.VestingShares
	if sub == nil {
		return nil, errors.New("the pro VestingShares is nil")
	}
	sub1 := m.Name
	if sub1 == nil {
		return nil, errors.New("the mainkey Name is nil")
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
func (s *SAccountVestingSharesWrap) ForEachByOrder(start *prototype.Vest, end *prototype.Vest, lastMainKey *prototype.AccountName,
	lastSubVal *prototype.Vest, f func(mVal *prototype.AccountName, sVal *prototype.Vest, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := AccountVestingSharesTable
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
	iterator := s.Dba.NewIterator(sBuf, eBuf)
	if iterator == nil {
		return errors.New("there is no data in range")
	}
	var idx uint32 = 0
	for iterator.Next() {
		idx++
		if isContinue := f(s.GetMainVal(iterator), s.GetSubVal(iterator), idx); !isContinue {
			break
		}
	}
	s.DelIterator(iterator)
	return nil
}

////////////// SECTION List Keys ///////////////
type SAccountBpVoteCountWrap struct {
	Dba iservices.IDatabaseService
}

func NewAccountBpVoteCountWrap(db iservices.IDatabaseService) *SAccountBpVoteCountWrap {
	if db == nil {
		return nil
	}
	wrap := SAccountBpVoteCountWrap{Dba: db}
	return &wrap
}

func (s *SAccountBpVoteCountWrap) DelIterator(iterator iservices.IDatabaseIterator) {
	if iterator == nil || !iterator.Valid() {
		return
	}
	s.Dba.DeleteIterator(iterator)
}

func (s *SAccountBpVoteCountWrap) GetMainVal(iterator iservices.IDatabaseIterator) *prototype.AccountName {
	if iterator == nil || !iterator.Valid() {
		return nil
	}
	val, err := iterator.Value()

	if err != nil {
		return nil
	}

	res := &SoListAccountByBpVoteCount{}
	err = proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.Name

}

func (s *SAccountBpVoteCountWrap) GetSubVal(iterator iservices.IDatabaseIterator) *uint32 {
	if iterator == nil || !iterator.Valid() {
		return nil
	}

	val, err := iterator.Value()

	if err != nil {
		return nil
	}
	res := &SoListAccountByBpVoteCount{}
	err = proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
	return &res.BpVoteCount

}

func (m *SoListAccountByBpVoteCount) OpeEncode() ([]byte, error) {
	pre := AccountBpVoteCountTable
	sub := m.BpVoteCount

	sub1 := m.Name
	if sub1 == nil {
		return nil, errors.New("the mainkey Name is nil")
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
func (s *SAccountBpVoteCountWrap) ForEachByOrder(start *uint32, end *uint32, lastMainKey *prototype.AccountName,
	lastSubVal *uint32, f func(mVal *prototype.AccountName, sVal *uint32, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := AccountBpVoteCountTable
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
	iterator := s.Dba.NewIterator(sBuf, eBuf)
	if iterator == nil {
		return errors.New("there is no data in range")
	}
	var idx uint32 = 0
	for iterator.Next() {
		idx++
		if isContinue := f(s.GetMainVal(iterator), s.GetSubVal(iterator), idx); !isContinue {
			break
		}
	}
	s.DelIterator(iterator)
	return nil
}

/////////////// SECTION Private function ////////////////

func (s *SoAccountWrap) update(sa *SoAccount) bool {
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

func (s *SoAccountWrap) getAccount() *SoAccount {
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

	res := &SoAccount{}
	if proto.Unmarshal(resBuf, res) != nil {
		return nil
	}
	return res
}

func (s *SoAccountWrap) encodeMainKey() ([]byte, error) {
	if s.mKeyBuf != nil {
		return s.mKeyBuf, nil
	}
	pre := s.getMemKeyPrefix("Name")
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

func (s *SoAccountWrap) delAllUniKeys(br bool, val *SoAccount) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if !s.delUniKeyName(val) {
		if br {
			return false
		} else {
			res = false
		}
	}

	return res
}

func (s *SoAccountWrap) delUniKeysWithNames(names map[string]string, val *SoAccount) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if len(names["Name"]) > 0 {
		if !s.delUniKeyName(val) {
			res = false
		}
	}

	return res
}

func (s *SoAccountWrap) insertAllUniKeys(val *SoAccount) (map[string]string, error) {
	if s.dba == nil {
		return nil, errors.New("insert uniuqe Field fail,the db is nil ")
	}
	if val == nil {
		return nil, errors.New("insert uniuqe Field fail,get the SoAccount fail ")
	}
	sucFields := map[string]string{}
	if !s.insertUniKeyName(val) {
		return sucFields, errors.New("insert unique Field Name fail while insert table ")
	}
	sucFields["Name"] = "Name"

	return sucFields, nil
}

func (s *SoAccountWrap) delUniKeyName(sa *SoAccount) bool {
	if s.dba == nil {
		return false
	}
	pre := AccountNameUniTable
	kList := []interface{}{pre}
	if sa != nil {

		if sa.Name == nil {
			return false
		}

		sub := sa.Name
		kList = append(kList, sub)
	} else {
		key, err := s.encodeMemKey("Name")
		if err != nil {
			return false
		}
		buf, err := s.dba.Get(key)
		if err != nil {
			return false
		}
		ori := &SoMemAccountByName{}
		err = proto.Unmarshal(buf, ori)
		if err != nil {
			return false
		}
		sub := ori.Name
		kList = append(kList, sub)

	}
	kBuf, err := kope.EncodeSlice(kList)
	if err != nil {
		return false
	}
	return s.dba.Delete(kBuf) == nil
}

func (s *SoAccountWrap) insertUniKeyName(sa *SoAccount) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	pre := AccountNameUniTable
	sub := sa.Name
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
	val := SoUniqueAccountByName{}
	val.Name = sa.Name

	buf, err := proto.Marshal(&val)

	if err != nil {
		return false
	}

	return s.dba.Put(kBuf, buf) == nil

}

type UniAccountNameWrap struct {
	Dba iservices.IDatabaseService
}

func NewUniAccountNameWrap(db iservices.IDatabaseService) *UniAccountNameWrap {
	if db == nil {
		return nil
	}
	wrap := UniAccountNameWrap{Dba: db}
	return &wrap
}

func (s *UniAccountNameWrap) UniQueryName(start *prototype.AccountName) *SoAccountWrap {
	if start == nil || s.Dba == nil {
		return nil
	}
	pre := AccountNameUniTable
	kList := []interface{}{pre, start}
	bufStartkey, err := kope.EncodeSlice(kList)
	val, err := s.Dba.Get(bufStartkey)
	if err == nil {
		res := &SoUniqueAccountByName{}
		rErr := proto.Unmarshal(val, res)
		if rErr == nil {
			wrap := NewSoAccountWrap(s.Dba, res.Name)

			return wrap
		}
	}
	return nil
}
