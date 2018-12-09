package table

import (
	"errors"
	fmt "fmt"
	"reflect"
	"strings"

	"github.com/coschain/contentos-go/common/encoding/kope"
	"github.com/coschain/contentos-go/iservices"
	proto "github.com/golang/protobuf/proto"
)

////////////// SECTION Prefix Mark ///////////////
var (
	WitnessScheduleObjectTable      = []byte("WitnessScheduleObjectTable")
	WitnessScheduleObjectIdUniTable = []byte("WitnessScheduleObjectIdUniTable")
)

////////////// SECTION Wrap Define ///////////////
type SoWitnessScheduleObjectWrap struct {
	dba     iservices.IDatabaseService
	mainKey *int32
}

func NewSoWitnessScheduleObjectWrap(dba iservices.IDatabaseService, key *int32) *SoWitnessScheduleObjectWrap {
	if dba == nil || key == nil {
		return nil
	}
	result := &SoWitnessScheduleObjectWrap{dba, key}
	return result
}

func (s *SoWitnessScheduleObjectWrap) CheckExist() bool {
	if s.dba == nil {
		return false
	}
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

func (s *SoWitnessScheduleObjectWrap) Create(f func(tInfo *SoWitnessScheduleObject)) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if s.mainKey == nil {
		return errors.New("the main key is nil")
	}
	val := &SoWitnessScheduleObject{}
	f(val)
	if s.CheckExist() {
		return errors.New("the main key is already exist")
	}
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return err

	}
	err = s.saveAllMemKeys(val, true)
	if err != nil {
		return err
	}

	// update sort list keys
	if err = s.insertAllSortKeys(val); err != nil {
		s.delAllSortKeys(false, val)
		s.dba.Delete(keyBuf)
		s.delAllMemKeys(false, val)
		return err
	}

	//update unique list
	if err = s.insertAllUniKeys(val); err != nil {
		s.delAllSortKeys(false, val)
		s.delAllUniKeys(false, val)
		s.dba.Delete(keyBuf)
		s.delAllMemKeys(false, val)
		return err
	}

	return nil
}

func (s *SoWitnessScheduleObjectWrap) encodeMemKey(fName string) ([]byte, error) {
	if len(fName) < 1 || s.mainKey == nil {
		return nil, errors.New("field name or main key is empty")
	}
	pre := "WitnessScheduleObject" + fName + "cell"
	kList := []interface{}{pre, s.mainKey}
	key, err := kope.EncodeSlice(kList)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func (so *SoWitnessScheduleObjectWrap) saveAllMemKeys(tInfo *SoWitnessScheduleObject, br bool) error {
	if so.dba == nil {
		return errors.New("save member Field fail , the db is nil")
	}

	if tInfo == nil {
		return errors.New("save member Field fail , the data is nil ")
	}
	var err error = nil
	errDes := ""
	if err = so.saveMemKeyCurrentShuffledWitness(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "CurrentShuffledWitness", err)
		}
	}
	if err = so.saveMemKeyId(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "Id", err)
		}
	}

	if len(errDes) > 0 {
		return errors.New(errDes)
	}
	return err
}

func (so *SoWitnessScheduleObjectWrap) delAllMemKeys(br bool, tInfo *SoWitnessScheduleObject) error {
	if so.dba == nil {
		return errors.New("the db is nil")
	}
	t := reflect.TypeOf(*tInfo)
	errDesc := ""
	for k := 0; k < t.NumField(); k++ {
		name := t.Field(k).Name
		if len(name) > 0 && !strings.HasPrefix(name, "XXX_") {
			err := so.delMemKey(name)
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

func (so *SoWitnessScheduleObjectWrap) delMemKey(fName string) error {
	if so.dba == nil {
		return errors.New("the db is nil")
	}
	if len(fName) <= 0 {
		return errors.New("the field name is empty ")
	}
	key, err := so.encodeMemKey(fName)
	if err != nil {
		return err
	}
	err = so.dba.Delete(key)
	return err
}

////////////// SECTION LKeys delete/insert ///////////////

func (s *SoWitnessScheduleObjectWrap) delAllSortKeys(br bool, val *SoWitnessScheduleObject) bool {
	if s.dba == nil {
		return false
	}
	res := true

	return res
}

func (s *SoWitnessScheduleObjectWrap) insertAllSortKeys(val *SoWitnessScheduleObject) error {
	if s.dba == nil {
		return errors.New("insert sort Field fail,the db is nil ")
	}
	if val == nil {
		return errors.New("insert sort Field fail,get the SoWitnessScheduleObject fail ")
	}

	return nil
}

////////////// SECTION LKeys delete/insert //////////////

func (s *SoWitnessScheduleObjectWrap) RemoveWitnessScheduleObject() bool {
	if s.dba == nil {
		return false
	}
	val := &SoWitnessScheduleObject{}
	//delete sort list key
	if res := s.delAllSortKeys(true, nil); !res {
		return false
	}

	//delete unique list
	if res := s.delAllUniKeys(true, nil); !res {
		return false
	}

	err := s.delAllMemKeys(true, val)
	return err == nil
}

////////////// SECTION Members Get/Modify ///////////////
func (s *SoWitnessScheduleObjectWrap) saveMemKeyCurrentShuffledWitness(tInfo *SoWitnessScheduleObject) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemWitnessScheduleObjectByCurrentShuffledWitness{}
	val.CurrentShuffledWitness = tInfo.CurrentShuffledWitness
	key, err := s.encodeMemKey("CurrentShuffledWitness")
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

func (s *SoWitnessScheduleObjectWrap) GetCurrentShuffledWitness() []string {
	res := true
	msg := &SoMemWitnessScheduleObjectByCurrentShuffledWitness{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("CurrentShuffledWitness")
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
				return msg.CurrentShuffledWitness
			}
		}
	}
	if !res {
		var tmpValue []string
		return tmpValue
	}
	return msg.CurrentShuffledWitness
}

func (s *SoWitnessScheduleObjectWrap) MdCurrentShuffledWitness(p []string) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("CurrentShuffledWitness")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemWitnessScheduleObjectByCurrentShuffledWitness{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoWitnessScheduleObject{}
	sa.Id = *s.mainKey
	sa.CurrentShuffledWitness = ori.CurrentShuffledWitness

	ori.CurrentShuffledWitness = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.CurrentShuffledWitness = p

	return true
}

func (s *SoWitnessScheduleObjectWrap) saveMemKeyId(tInfo *SoWitnessScheduleObject) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemWitnessScheduleObjectById{}
	val.Id = tInfo.Id
	key, err := s.encodeMemKey("Id")
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

func (s *SoWitnessScheduleObjectWrap) GetId() int32 {
	res := true
	msg := &SoMemWitnessScheduleObjectById{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("Id")
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
				return msg.Id
			}
		}
	}
	if !res {
		var tmpValue int32
		return tmpValue
	}
	return msg.Id
}

/////////////// SECTION Private function ////////////////

func (s *SoWitnessScheduleObjectWrap) update(sa *SoWitnessScheduleObject) bool {
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

func (s *SoWitnessScheduleObjectWrap) getWitnessScheduleObject() *SoWitnessScheduleObject {
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

	res := &SoWitnessScheduleObject{}
	if proto.Unmarshal(resBuf, res) != nil {
		return nil
	}
	return res
}

func (s *SoWitnessScheduleObjectWrap) encodeMainKey() ([]byte, error) {
	pre := "WitnessScheduleObject" + "Id" + "cell"
	sub := s.mainKey
	if sub == nil {
		return nil, errors.New("the mainKey is nil")
	}
	kList := []interface{}{pre, sub}
	kBuf, cErr := kope.EncodeSlice(kList)
	return kBuf, cErr
}

////////////// Unique Query delete/insert/query ///////////////

func (s *SoWitnessScheduleObjectWrap) delAllUniKeys(br bool, val *SoWitnessScheduleObject) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if !s.delUniKeyId(val) {
		if br {
			return false
		} else {
			res = false
		}
	}

	return res
}

func (s *SoWitnessScheduleObjectWrap) insertAllUniKeys(val *SoWitnessScheduleObject) error {
	if s.dba == nil {
		return errors.New("insert uniuqe Field fail,the db is nil ")
	}
	if val == nil {
		return errors.New("insert uniuqe Field fail,get the SoWitnessScheduleObject fail ")
	}
	if !s.insertUniKeyId(val) {
		return errors.New("insert unique Field Id fail while insert table ")
	}

	return nil
}

func (s *SoWitnessScheduleObjectWrap) delUniKeyId(sa *SoWitnessScheduleObject) bool {
	if s.dba == nil {
		return false
	}
	pre := WitnessScheduleObjectIdUniTable
	kList := []interface{}{pre}
	if sa != nil {

		sub := sa.Id
		kList = append(kList, sub)
	} else {
		key, err := s.encodeMemKey("Id")
		if err != nil {
			return false
		}
		buf, err := s.dba.Get(key)
		if err != nil {
			return false
		}
		ori := &SoMemWitnessScheduleObjectById{}
		err = proto.Unmarshal(buf, ori)
		if err != nil {
			return false
		}
		sub := ori.Id
		kList = append(kList, sub)

	}
	kBuf, err := kope.EncodeSlice(kList)
	if err != nil {
		return false
	}
	return s.dba.Delete(kBuf) == nil
}

func (s *SoWitnessScheduleObjectWrap) insertUniKeyId(sa *SoWitnessScheduleObject) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	uniWrap := UniWitnessScheduleObjectIdWrap{}
	uniWrap.Dba = s.dba
	res := uniWrap.UniQueryId(&sa.Id)

	if res != nil {
		//the unique key is already exist
		return false
	}
	val := SoUniqueWitnessScheduleObjectById{}
	val.Id = sa.Id

	buf, err := proto.Marshal(&val)

	if err != nil {
		return false
	}

	pre := WitnessScheduleObjectIdUniTable
	sub := sa.Id
	kList := []interface{}{pre, sub}
	kBuf, err := kope.EncodeSlice(kList)
	if err != nil {
		return false
	}
	return s.dba.Put(kBuf, buf) == nil

}

type UniWitnessScheduleObjectIdWrap struct {
	Dba iservices.IDatabaseService
}

func NewUniWitnessScheduleObjectIdWrap(db iservices.IDatabaseService) *UniWitnessScheduleObjectIdWrap {
	if db == nil {
		return nil
	}
	wrap := UniWitnessScheduleObjectIdWrap{Dba: db}
	return &wrap
}

func (s *UniWitnessScheduleObjectIdWrap) UniQueryId(start *int32) *SoWitnessScheduleObjectWrap {
	if start == nil || s.Dba == nil {
		return nil
	}
	pre := WitnessScheduleObjectIdUniTable
	kList := []interface{}{pre, start}
	bufStartkey, err := kope.EncodeSlice(kList)
	val, err := s.Dba.Get(bufStartkey)
	if err == nil {
		res := &SoUniqueWitnessScheduleObjectById{}
		rErr := proto.Unmarshal(val, res)
		if rErr == nil {
			wrap := NewSoWitnessScheduleObjectWrap(s.Dba, &res.Id)
			return wrap
		}
	}
	return nil
}
