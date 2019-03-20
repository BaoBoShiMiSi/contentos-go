package storage

import (
	"errors"
	"sync"
)

type dbDeque struct {
	db Database
	readFront bool
	sessions []*dbSession
	lock sync.RWMutex
}

func NewDBDeque(db Database, readFront bool) *dbDeque {
	return &dbDeque{ db: db, readFront: readFront }
}

func (dq *dbDeque) size() uint {
	return uint(1 + len(dq.sessions))
}

func (dq *dbDeque) Size() uint {
	dq.lock.RLock()
	defer dq.lock.RUnlock()

	return dq.size()
}

func (dq *dbDeque) front() Database {
	if len(dq.sessions) > 0 {
		return dq.sessions[len(dq.sessions) - 1]
	} else {
		return dq.db
	}
}

func (dq *dbDeque) back() Database {
	return dq.db
}

func (dq *dbDeque) pushFront() {
	dq.sessions = append(dq.sessions, &dbSession{
		db: dq.front(),
		mem: NewMemoryDatabase(),
		removals: make(map[string]bool),
	})
}

func (dq *dbDeque) PushFront() {
	dq.lock.Lock()
	defer dq.lock.Unlock()

	dq.pushFront()
}

func (dq *dbDeque) popFront(commit bool) (err error) {
	sessionCount := len(dq.sessions)
	if sessionCount == 0 {
		return errors.New("unexpected pop.")
	}
	if commit {
		err = dq.sessions[sessionCount - 1].commit()
	}
	dq.sessions = dq.sessions[:len(dq.sessions) - 1]
	return err
}

func (dq *dbDeque) PopFront(commit bool) error {
	dq.lock.Lock()
	defer dq.lock.Unlock()

	return dq.popFront(commit)
}

func (dq *dbDeque) pushBack() {
	dq.sessions = append([]*dbSession{ {
		db: dq.back(),
		mem: NewMemoryDatabase(),
		removals: make(map[string]bool),
	} }, dq.sessions...)

	if len(dq.sessions) > 1 {
		dq.sessions[1].db = dq.sessions[0]
	}
}

func (dq *dbDeque) PushBack() {
	dq.lock.Lock()
	defer dq.lock.Unlock()

	dq.pushBack()
}

func (dq *dbDeque) popBack(commit bool) (err error) {
	sessionCount := len(dq.sessions)
	if sessionCount == 0 {
		return errors.New("unexpected pop.")
	}
	if commit {
		err = dq.sessions[0].commit()
		if sessionCount > 1 {
			dq.sessions[1].db = dq.sessions[0].db
		}
	}
	dq.sessions = dq.sessions[1:]
	return err
}

func (dq *dbDeque) PopBack(commit bool) error {
	dq.lock.Lock()
	defer dq.lock.Unlock()

	return dq.popBack(commit)
}

func (dq *dbDeque) writerDB() Database {
	dq.lock.RLock()
	defer dq.lock.RUnlock()

	return dq.front()
}

func (dq *dbDeque) readerDB() Database {
	dq.lock.RLock()
	defer dq.lock.RUnlock()

	if dq.readFront {
		return dq.front()
	} else {
		return dq.back()
	}
}

func (dq *dbDeque) Close() {

}

func (dq *dbDeque) Has(key []byte) (bool, error) {
	return dq.readerDB().Has(key)
}

func (dq *dbDeque) Get(key []byte) ([]byte, error) {
	return dq.readerDB().Get(key)
}

func (dq *dbDeque) Put(key []byte, value []byte) error {
	return dq.writerDB().Put(key, value)
}

func (dq *dbDeque) Delete(key []byte) error {
	return dq.writerDB().Delete(key)
}

func (dq *dbDeque) NewIterator(start []byte, limit []byte) Iterator {
	return dq.readerDB().NewIterator(start, limit)
}

func (dq *dbDeque) NewReversedIterator(start []byte, limit []byte) Iterator {
	return dq.readerDB().NewReversedIterator(start, limit)
}

func (dq *dbDeque) DeleteIterator(it Iterator) {
	dq.readerDB().DeleteIterator(it)
}

func (dq *dbDeque) NewBatch() Batch {
	return dq.writerDB().NewBatch()
}

func (dq *dbDeque) DeleteBatch(b Batch) {
	dq.writerDB().DeleteBatch(b)
}
