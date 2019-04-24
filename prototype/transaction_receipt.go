package prototype

const StatusSuccess  = 200
const StatusDeductGas  = 201
const StatusError  = 500

func (m *TransactionReceipt) Validate() error {
	return nil
}

func (m *TransactionReceipt) IsSuccess() bool {
	return m.Status == StatusSuccess || m.Status == StatusDeductGas
}