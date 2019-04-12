package storage

//
// This file implements a simple database group and a dispatcher policy based on hashes of keys.
//

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/coschain/contentos-go/common"
	"strconv"
	"sync"
	"sync/atomic"
)

type KeyHashDispatcher struct {
	dbs []Database
}

func NewKeyHashDispatcher(databases []Database) *KeyHashDispatcher {
	dbs := make([]Database, 0, len(databases))
	for _, db := range databases {
		dbs = append(dbs, db)
	}
	return &KeyHashDispatcher{dbs: dbs}
}

func (dp *KeyHashDispatcher) MemberDatabases() []Database {
	return dp.dbs
}

func (dp *KeyHashDispatcher) DatabaseForKey(key []byte) int {
	if len(dp.dbs) > 0 {
		h := md5.Sum(key)
		n, _ := strconv.ParseUint(fmt.Sprintf("%x", h[:8]), 16, 64)
		return int(n % uint64(len(dp.dbs)))
	}
	return -1
}

func (dp *KeyHashDispatcher) DatabasesForKeyRange(start []byte, limit []byte) []int {
	n := len(dp.dbs)
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
	return idx
}

// the database group
type SimpleDatabaseGroup struct {
	dp      DatabaseDispatcher // key dispatching policy
	crashed int32              // non-zero if the group should stop service due to fatal errors
	lock    sync.RWMutex       // lock for db operations
	iterLock sync.RWMutex
	wal     WriteAheadLog      // write ahead log
}

func NewSimpleDatabaseGroup(dp DatabaseDispatcher, wal WriteAheadLog) (*SimpleDatabaseGroup, error) {
	if len(dp.MemberDatabases()) > 0 {
		g := &SimpleDatabaseGroup{dp: dp, wal: wal}

		// check for pending tasks in WAL, and try to finish them if there's any.
		// pending tasks are violation of atomicity. if we can't finish them all, the database is broken.
		if err := g.finishPendingTasks(); err != nil {
			return nil, err
		}
		return g, nil
	} else {
		return nil, errors.New("invalid member count or wal filepath.")
	}
}

func (g *SimpleDatabaseGroup) Crashed() bool {
	return atomic.LoadInt32(&g.crashed) != 0
}

func (g *SimpleDatabaseGroup) Close() {

}

func (g *SimpleDatabaseGroup) dbAt(idx int) Database {
	dbs := g.dp.MemberDatabases()
	if idx >= 0 && idx < len(dbs) {
		return dbs[idx]
	} else {
		return nil
	}
}

func (g *SimpleDatabaseGroup) dbForKey(key []byte) Database {
	if idx := g.dp.DatabaseForKey(key); idx >= 0 {
		return g.dbAt(idx)
	} else {
		return nil
	}
}

func (g *SimpleDatabaseGroup) Has(key []byte) (bool, error) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	if g.Crashed() {
		return false, errors.New("database group out of service due to fatal errors")
	}

	return g.dbForKey(key).Has(key)
}

func (g *SimpleDatabaseGroup) Get(key []byte) ([]byte, error) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	if g.Crashed() {
		return nil, errors.New("database group out of service due to fatal errors")
	}

	return g.dbForKey(key).Get(key)
}

func (g *SimpleDatabaseGroup) Put(key []byte, value []byte) error {
	g.lock.Lock()
	g.iterLock.Lock()
	defer g.iterLock.Unlock()
	defer g.lock.Unlock()

	if g.Crashed() {
		return errors.New("database group out of service due to fatal errors")
	}

	return g.dbForKey(key).Put(key, value)
}

func (g *SimpleDatabaseGroup) Delete(key []byte) error {
	g.lock.Lock()
	g.iterLock.Lock()
	defer g.iterLock.Unlock()
	defer g.lock.Unlock()

	if g.Crashed() {
		return errors.New("database group out of service due to fatal errors")
	}

	return g.dbForKey(key).Delete(key)
}

func (g *SimpleDatabaseGroup) Iterate(start, limit []byte, reverse bool, callback func(key, value []byte) bool) {
	g.lock.RLock()
	g.iterLock.RLock()
	defer g.iterLock.RUnlock()
	g.lock.RUnlock()

	if g.Crashed() {
		return
	}

	if it := NewMergedIterator(g.dp.MemberDatabases()); it != nil {
		it.Iterate(start, limit, reverse, callback)
	}
}

func (g *SimpleDatabaseGroup) NewBatch() Batch {
	if g.Crashed() {
		return nil
	}
	return &sdgBatch{
		g:   g,
		ops: make(map[int][]writeOp),
		rev: make(map[int][]writeOp),
	}
}

func (g *SimpleDatabaseGroup) DeleteBatch(b Batch) {

}

// db group batch
type sdgBatch struct {
	g    *SimpleDatabaseGroup // the db group
	ops  map[int][]writeOp    // operations of this batch, grouped by member db index
	rev  map[int][]writeOp    // reversed operations of this batch, grouped by member db index
	lock sync.RWMutex
}

func (b *sdgBatch) Put(key []byte, value []byte) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	dbIdx := b.g.dp.DatabaseForKey(key)
	// record the operation
	b.ops[dbIdx] = append(b.ops[dbIdx], writeOp{key, value, false})

	// record the reversed operation
	if oldval, err := b.g.dbAt(dbIdx).Get(key); err == nil {
		b.rev[dbIdx] = append(b.rev[dbIdx], writeOp{common.CopyBytes(key), common.CopyBytes(oldval), false})
	} else {
		b.rev[dbIdx] = append(b.rev[dbIdx], writeOp{common.CopyBytes(key), nil, true})
	}
	return nil
}

func (b *sdgBatch) Delete(key []byte) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	dbIdx := b.g.dp.DatabaseForKey(key)
	// record the operation
	b.ops[dbIdx] = append(b.ops[dbIdx], writeOp{common.CopyBytes(key), nil, true})
	// record the reversed operation
	if oldval, err := b.g.dbAt(dbIdx).Get(key); err == nil {
		b.rev[dbIdx] = append(b.rev[dbIdx], writeOp{common.CopyBytes(key), common.CopyBytes(oldval), false})
	}
	return nil
}

func (b *sdgBatch) Reset() {
	b.lock.Lock()
	defer b.lock.Unlock()

	b.ops = make(map[int][]writeOp)
	b.rev = make(map[int][]writeOp)
}

func (b *sdgBatch) Write() error {
	b.lock.Lock()
	defer b.lock.Unlock()

	// prepare member batches
	// dbBatches[member_db_idx] = { batch, batch_for_reversion }
	dbBatches := make(map[int][]Batch)
	dbTasks := make([]*WriteTask, 0, len(b.ops))
	dbTaskIds := make([]uint64, 0, len(b.ops))
	for idx, w := range b.ops {
		batch := b.g.dbAt(idx).NewBatch()
		rbatch := b.g.dbAt(idx).NewBatch()
		r := b.rev[idx]
		for i, wop := range w {
			if wop.Del {
				batch.Delete(wop.Key)
			} else {
				batch.Put(wop.Key, wop.Value)
			}
			rop := r[i]
			if rop.Del {
				rbatch.Delete(rop.Key)
			} else {
				rbatch.Put(rop.Key, rop.Value)
			}
		}
		dbBatches[idx] = append(dbBatches[idx], batch, rbatch)

		taskId := b.g.wal.NewTaskID()
		dbTasks = append(dbTasks, &WriteTask{taskId, strconv.Itoa(idx), w})
		dbTaskIds = append(dbTaskIds, taskId)
	}

	// @result will hold the execution result of each member batch
	result := make(map[int]bool)
	for idx := range dbBatches {
		result[idx] = false
	}

	// write ahead logging
	if err := b.g.wal.PutTasks(dbTasks); err != nil {
		return errors.New("failed updating wal: " + err.Error())
	}

	// run all batches in parallel
	var wg sync.WaitGroup
	wg.Add(len(result))
	for idx, batches := range dbBatches {
		go func(idx int, batch Batch) {
			defer wg.Done()
			if err := batch.Write(); err == nil {
				result[idx] = true
			}
		}(idx, batches[0])
	}
	wg.Wait()

	// check if all member batches succeeded
	ok := true
	for _, r := range result {
		if !r {
			ok = false
			break
		}
	}
	var err error
	if !ok {
		// if some member batches failed,
		// we have to revert succeeded ones so that atomicity keeps
		wg.Add(len(result))
		for idx, batches := range dbBatches {
			go func(idx int, rbatch Batch) {
				defer wg.Done()
				if result[idx] {
					if err := rbatch.Write(); err == nil {
						result[idx] = false
					}
				}
			}(idx, batches[1])
		}
		wg.Wait()

		// check if all reversions successfully done
		var fataldb []int
		for idx, r := range result {
			if r {
				fataldb = append(fataldb, idx)
			}
		}
		if len(fataldb) == 0 {
			err = errors.New("some of databases failed batch writing")
		} else {
			// this is really really bad. some member databases are out of service.
			// unrecoverable atomicity violation makes database group totally crashed.
			err = errors.New(fmt.Sprintf("FATAL: Atomicity violation due to failed recoveries on databases %v", fataldb))
			atomic.StoreInt32(&b.g.crashed, 1)
		}
	}

	// release member batches
	for idx, batches := range dbBatches {
		b.g.dbAt(idx).DeleteBatch(batches[0])
		b.g.dbAt(idx).DeleteBatch(batches[1])
	}

	// delete write ahead logs after db writing
	b.g.wal.DeleteTasks(dbTaskIds)
	return err
}

func (g *SimpleDatabaseGroup) finishPendingTasks() error {
	tasks, err := g.wal.GetTasks()
	if err != nil {
		return err
	}
	for _, task := range tasks {
		taskId := task.TaskID
		dbIdx, err := strconv.ParseInt(task.DatabaseID, 10, 32)
		if err != nil {
			return errors.New("cannot convert wal data to integeral database id:" + err.Error())
		}
		db := g.dbAt(int(dbIdx))
		if db == nil {
			return errors.New(fmt.Sprintf("failed accessing member database #%d. invalid wal data or unavailable member database.", dbIdx))
		}

		b := db.NewBatch()
		defer db.DeleteBatch(b)

		for _, op := range task.Operations {
			if op.Del {
				err = b.Delete(op.Key)
			} else {
				err = b.Put(op.Key, op.Value)
			}
			if err != nil {
				return errors.New(fmt.Sprintf("failed creating batch on member database #%d: %s", dbIdx, err.Error()))
			}
		}
		if err = b.Write(); err != nil {
			return errors.New(fmt.Sprintf("failed writing batch to member database #%d: %s", dbIdx, err.Error()))
		}
		if err = g.wal.DeleteTask(taskId); err != nil {
			return errors.New(fmt.Sprintf("cannot delete task %d from wal: %s", taskId, err.Error()))
		}
	}
	return nil
}
