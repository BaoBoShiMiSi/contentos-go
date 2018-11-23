package prototype

type SkipFlag uint32

const (
	Skip_nothing                SkipFlag = 0
	Skip_transaction_signatures SkipFlag = 1 << 0
	Skip_apply_transaction      SkipFlag = 1 << 1
)
