package storage

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"
)

type SquashableDatabase struct {
	TransactionalDatabase

	tags map[string]uint
	tagsByIdx map[uint]string
	lock sync.RWMutex
}

func NewSquashableDatabase(db Database, dirtyRead bool) *SquashableDatabase {
	return &SquashableDatabase{
		TransactionalDatabase: TransactionalDatabase{
			dbDeque: dbDeque{
				db:        db,
				readFront: dirtyRead,
			},
		},
		tags: make(map[string]uint),
		tagsByIdx: make(map[uint]string),
	}
}

func (db *SquashableDatabase) BeginTransaction() {
	db.lock.Lock()
	defer db.lock.Unlock()

	db.TransactionalDatabase.BeginTransaction()
}

func (db *SquashableDatabase) EndTransaction(commit bool) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	err := db.TransactionalDatabase.EndTransaction(commit)

	frontIdx := db.Size() - 1
	if poppedTag, ok := db.tagsByIdx[frontIdx]; ok {
		delete(db.tagsByIdx, frontIdx)
		delete(db.tags, poppedTag)
	}

	return err
}

func (db *SquashableDatabase) BeginTransactionWithTag(tag string) {
	db.lock.Lock()
	defer db.lock.Unlock()

	db.TransactionalDatabase.BeginTransaction()
	frontIdx := db.Size() - 2
	db.tags[tag] = frontIdx
	db.tagsByIdx[frontIdx] = tag
}

func (db *SquashableDatabase) Squash(tag string) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	if idx, ok := db.tags[tag]; ok {
		count := int(idx) + 1
		if err := db.PopBackN(count, true); err != nil {
			fmt.Printf("pop fail,the error is %s",err)
			return err
		}
		newTags := make(map[string]uint)
		newTagsByIdx := make(map[uint]string)
		for i, t := range db.tagsByIdx {
			if i > idx {
				newTagsByIdx[i - idx - 1] = t
				newTags[t] = i - idx - 1
			}
		}
		db.tags, db.tagsByIdx = newTags, newTagsByIdx
		return nil
	}
	return errors.New("unknown tag: " + tag)
}

func (db *SquashableDatabase) RollbackTag(tag string) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	if idx, ok := db.tags[tag]; ok {
		count := int(db.Size()-1) - int(idx)
		for i := 0; i < count; i++ {
			if err := db.PopFront(false); err != nil {
				return err
			}
		}
		newTags := make(map[string]uint)
		newTagsByIdx := make(map[uint]string)
		for i, t := range db.tagsByIdx {
			if i < idx {
				newTagsByIdx[i] = t
				newTags[t] = i
			}
		}
		db.tags, db.tagsByIdx = newTags, newTagsByIdx
		return nil
	}
	return errors.New("unknown tag: " + tag)
}
