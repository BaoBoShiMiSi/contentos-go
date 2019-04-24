package storage

// interface for insertion and updating
// methods must be thread safe
type DatabasePutter interface {
	// insert a new key-value pair, or update the value if the given key already exists
	Put(key []byte, value []byte) error
}

// interface for deletion
// methods must be thread safe
type DatabaseDeleter interface {
	// delete the given key and its value
	// if the given key does not exist, just return nil, indicating a successful deletion without doing anything.
	Delete(key []byte) error
}

type DatabaseWriter interface {
	DatabasePutter
	DatabaseDeleter
}

// interface for key & value query
// methods must be thread safe
type DatabaseGetter interface {
	// check existence of the given key
	Has(key []byte) (bool, error)

	// query the value of the given key
	Get(key []byte) ([]byte, error)
}

// interface for key-space range scan
// methods must be thread safe
type DatabaseScanner interface {
	// Iterate enumerates keys in range [start, limit) and calls callback with each enumerated key and its value.
	// If callback returns false, the enumeration stops immediately, otherwise all matched keys will be enumerated.
	// a nil start is the logical minimal key that is lesser than any existing keys
	// a nil limit is the logical maximum key that is greater than any existing keys
	Iterate(start, limit []byte, reverse bool, callback func(key, value []byte) bool)
}

// interface for transactional execution of multiple writes
// methods must be thread safe
type DatabaseBatcher interface {
	// create a batch which can pack DatabasePutter & DatabaseDeleter operations and execute them atomically
	NewBatch() Batch

	// release a Batch
	DeleteBatch(b Batch)
}

// interface for transaction executor
// methods must be thread safe
// write operations must be executed atomically
type Batch interface {
	DatabaseWriter

	// execute all batched operations
	Write() error

	// reset the batch to empty
	Reset()
}

// interface for full functional database
// methods must be thread safe
type Database interface {
	DatabaseGetter
	DatabasePutter
	DatabaseDeleter
	DatabaseScanner
	DatabaseBatcher
	Close()
}

// interface for transaction feature
// methods must be thread safe
// transaction sessions can be nested. BeginTransaction()/EndTransaction() must be paired.
type Transactional interface {
	// start a new transaction session
	BeginTransaction()

	// end current transaction session, commit or discard changes
	EndTransaction(commit bool) error

	// current transaction height
	TransactionHeight() uint
}

// interface for databases that support transactions
type TrxDatabase interface {
	Transactional
	Database
}

// interface for squashable transactional feature
type Squashable interface {
	Transactional

	// start a new transaction session named @tag
	BeginTransactionWithTag(tag string)

	// Squash() commits the transaction named @tag.
	// If the transaction is nested, its ancestor transactions will get committed too.
	Squash(tag string) error

	// RollbackTag() discards the transaction named @tag.
	// If any nested transactions were atop the target transaction, they will get discarded too.
	RollbackTag(tag string) error
}

type SquashDatabase interface {
	Squashable
	Database
}

// interface for revertible feature
// methods must be thread safe
type Revertible interface {
	// get current revision
	GetRevision() uint64

	// get current revision and base revision
	GetRevisionAndBase() (current uint64, base uint64)

	// revert to the given revision
	// you can only revert to a revision that is less than or equal to current revision.
	// after reverted to revision r, r will be the current revision.
	RevertToRevision(r uint64) error

	// rebase to the given revision
	// after rebased to revision r, r will be the minimal revision you can revert to.
	RebaseToRevision(r uint64) error

	EnableReversion(b bool) error
	ReversionEnabled() bool
}

// interface for databases that support reversion
type RevDatabase interface {
	Revertible
	Database
}

// interface for databases that support reversion and revision tagging
type TagRevertible interface {
	Revertible

	// tag a revision
	TagRevision(r uint64, tag string) error

	// get revision of a tag
	GetTagRevision(tag string) (uint64, error)

	// get tag of a revision
	GetRevisionTag(r uint64) string

	// revert to a revision by its tag
	RevertToTag(tag string) error

	// rebase to a revision by its tag
	RebaseToTag(tag string) error
}

// interface for databases that support reversion and revision tagging
type TagRevDatabase interface {
	TagRevertible
	Database
}

// interface for key->database mapping policy for a group of databases
type DatabaseDispatcher interface {
	// return members of database group.
	// members must be fixed once the DatabaseDispatcher object is created
	MemberDatabases() []Database

	// return the index number of the mapped member database
	DatabaseForKey(key []byte) int

	// return databases who possibly contains keys from given range
	DatabasesForKeyRange(start []byte, limit []byte) []int
}

// interface for a logical database consisting of a group of databases
type DatabaseGroup interface {
	DatabaseDispatcher
	Database
}
