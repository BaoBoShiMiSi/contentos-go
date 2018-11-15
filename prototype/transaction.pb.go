// Code generated by protoc-gen-go. DO NOT EDIT.
// source: prototype/transaction.proto

package prototype

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Operation struct {
	// Types that are valid to be assigned to Op:
	//	*Operation_Op1
	//	*Operation_Op2
	//	*Operation_Op3
	//	*Operation_Op4
	//	*Operation_Op5
	//	*Operation_Op6
	//	*Operation_Op7
	//	*Operation_Op8
	//	*Operation_Op9
	//	*Operation_Op10
	Op                   isOperation_Op `protobuf_oneof:"op"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{0}
}

func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (m *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(m, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

type isOperation_Op interface {
	isOperation_Op()
}

type Operation_Op1 struct {
	Op1 *AccountCreateOperation `protobuf:"bytes,1,opt,name=op1,proto3,oneof"`
}

type Operation_Op2 struct {
	Op2 *TransferOperation `protobuf:"bytes,2,opt,name=op2,proto3,oneof"`
}

type Operation_Op3 struct {
	Op3 *BpRegisterOperation `protobuf:"bytes,3,opt,name=op3,proto3,oneof"`
}

type Operation_Op4 struct {
	Op4 *BpUnregisterOperation `protobuf:"bytes,4,opt,name=op4,proto3,oneof"`
}

type Operation_Op5 struct {
	Op5 *BpVoteOperation `protobuf:"bytes,5,opt,name=op5,proto3,oneof"`
}

type Operation_Op6 struct {
	Op6 *PostOperation `protobuf:"bytes,6,opt,name=op6,proto3,oneof"`
}

type Operation_Op7 struct {
	Op7 *ReplayOperation `protobuf:"bytes,7,opt,name=op7,proto3,oneof"`
}

type Operation_Op8 struct {
	Op8 *FollowOperation `protobuf:"bytes,8,opt,name=op8,proto3,oneof"`
}

type Operation_Op9 struct {
	Op9 *VoteOperation `protobuf:"bytes,9,opt,name=op9,proto3,oneof"`
}

type Operation_Op10 struct {
	Op10 *TransferToVestingOperation `protobuf:"bytes,10,opt,name=op10,proto3,oneof"`
}

func (*Operation_Op1) isOperation_Op() {}

func (*Operation_Op2) isOperation_Op() {}

func (*Operation_Op3) isOperation_Op() {}

func (*Operation_Op4) isOperation_Op() {}

func (*Operation_Op5) isOperation_Op() {}

func (*Operation_Op6) isOperation_Op() {}

func (*Operation_Op7) isOperation_Op() {}

func (*Operation_Op8) isOperation_Op() {}

func (*Operation_Op9) isOperation_Op() {}

func (*Operation_Op10) isOperation_Op() {}

func (m *Operation) GetOp() isOperation_Op {
	if m != nil {
		return m.Op
	}
	return nil
}

func (m *Operation) GetOp1() *AccountCreateOperation {
	if x, ok := m.GetOp().(*Operation_Op1); ok {
		return x.Op1
	}
	return nil
}

func (m *Operation) GetOp2() *TransferOperation {
	if x, ok := m.GetOp().(*Operation_Op2); ok {
		return x.Op2
	}
	return nil
}

func (m *Operation) GetOp3() *BpRegisterOperation {
	if x, ok := m.GetOp().(*Operation_Op3); ok {
		return x.Op3
	}
	return nil
}

func (m *Operation) GetOp4() *BpUnregisterOperation {
	if x, ok := m.GetOp().(*Operation_Op4); ok {
		return x.Op4
	}
	return nil
}

func (m *Operation) GetOp5() *BpVoteOperation {
	if x, ok := m.GetOp().(*Operation_Op5); ok {
		return x.Op5
	}
	return nil
}

func (m *Operation) GetOp6() *PostOperation {
	if x, ok := m.GetOp().(*Operation_Op6); ok {
		return x.Op6
	}
	return nil
}

func (m *Operation) GetOp7() *ReplayOperation {
	if x, ok := m.GetOp().(*Operation_Op7); ok {
		return x.Op7
	}
	return nil
}

func (m *Operation) GetOp8() *FollowOperation {
	if x, ok := m.GetOp().(*Operation_Op8); ok {
		return x.Op8
	}
	return nil
}

func (m *Operation) GetOp9() *VoteOperation {
	if x, ok := m.GetOp().(*Operation_Op9); ok {
		return x.Op9
	}
	return nil
}

func (m *Operation) GetOp10() *TransferToVestingOperation {
	if x, ok := m.GetOp().(*Operation_Op10); ok {
		return x.Op10
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Operation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Operation_OneofMarshaler, _Operation_OneofUnmarshaler, _Operation_OneofSizer, []interface{}{
		(*Operation_Op1)(nil),
		(*Operation_Op2)(nil),
		(*Operation_Op3)(nil),
		(*Operation_Op4)(nil),
		(*Operation_Op5)(nil),
		(*Operation_Op6)(nil),
		(*Operation_Op7)(nil),
		(*Operation_Op8)(nil),
		(*Operation_Op9)(nil),
		(*Operation_Op10)(nil),
	}
}

func _Operation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Operation)
	// op
	switch x := m.Op.(type) {
	case *Operation_Op1:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Op1); err != nil {
			return err
		}
	case *Operation_Op2:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Op2); err != nil {
			return err
		}
	case *Operation_Op3:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Op3); err != nil {
			return err
		}
	case *Operation_Op4:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Op4); err != nil {
			return err
		}
	case *Operation_Op5:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Op5); err != nil {
			return err
		}
	case *Operation_Op6:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Op6); err != nil {
			return err
		}
	case *Operation_Op7:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Op7); err != nil {
			return err
		}
	case *Operation_Op8:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Op8); err != nil {
			return err
		}
	case *Operation_Op9:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Op9); err != nil {
			return err
		}
	case *Operation_Op10:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Op10); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Operation.Op has unexpected type %T", x)
	}
	return nil
}

func _Operation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Operation)
	switch tag {
	case 1: // op.op1
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AccountCreateOperation)
		err := b.DecodeMessage(msg)
		m.Op = &Operation_Op1{msg}
		return true, err
	case 2: // op.op2
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransferOperation)
		err := b.DecodeMessage(msg)
		m.Op = &Operation_Op2{msg}
		return true, err
	case 3: // op.op3
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BpRegisterOperation)
		err := b.DecodeMessage(msg)
		m.Op = &Operation_Op3{msg}
		return true, err
	case 4: // op.op4
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BpUnregisterOperation)
		err := b.DecodeMessage(msg)
		m.Op = &Operation_Op4{msg}
		return true, err
	case 5: // op.op5
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BpVoteOperation)
		err := b.DecodeMessage(msg)
		m.Op = &Operation_Op5{msg}
		return true, err
	case 6: // op.op6
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PostOperation)
		err := b.DecodeMessage(msg)
		m.Op = &Operation_Op6{msg}
		return true, err
	case 7: // op.op7
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ReplayOperation)
		err := b.DecodeMessage(msg)
		m.Op = &Operation_Op7{msg}
		return true, err
	case 8: // op.op8
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FollowOperation)
		err := b.DecodeMessage(msg)
		m.Op = &Operation_Op8{msg}
		return true, err
	case 9: // op.op9
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(VoteOperation)
		err := b.DecodeMessage(msg)
		m.Op = &Operation_Op9{msg}
		return true, err
	case 10: // op.op10
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransferToVestingOperation)
		err := b.DecodeMessage(msg)
		m.Op = &Operation_Op10{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Operation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Operation)
	// op
	switch x := m.Op.(type) {
	case *Operation_Op1:
		s := proto.Size(x.Op1)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_Op2:
		s := proto.Size(x.Op2)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_Op3:
		s := proto.Size(x.Op3)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_Op4:
		s := proto.Size(x.Op4)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_Op5:
		s := proto.Size(x.Op5)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_Op6:
		s := proto.Size(x.Op6)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_Op7:
		s := proto.Size(x.Op7)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_Op8:
		s := proto.Size(x.Op8)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_Op9:
		s := proto.Size(x.Op9)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_Op10:
		s := proto.Size(x.Op10)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// transaction
type Transaction struct {
	RefBlockNum          uint32        `protobuf:"varint,1,opt,name=ref_block_num,json=refBlockNum,proto3" json:"ref_block_num,omitempty"`
	RefBlockPrefix       uint32        `protobuf:"varint,2,opt,name=ref_block_prefix,json=refBlockPrefix,proto3" json:"ref_block_prefix,omitempty"`
	Expiration           *TimePointSec `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	Operations           []*Operation  `protobuf:"bytes,4,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{1}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetRefBlockNum() uint32 {
	if m != nil {
		return m.RefBlockNum
	}
	return 0
}

func (m *Transaction) GetRefBlockPrefix() uint32 {
	if m != nil {
		return m.RefBlockPrefix
	}
	return 0
}

func (m *Transaction) GetExpiration() *TimePointSec {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (m *Transaction) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

type SignedTransaction struct {
	Trx                  *Transaction     `protobuf:"bytes,1,opt,name=trx,proto3" json:"trx,omitempty"`
	Signatures           []*SignatureType `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SignedTransaction) Reset()         { *m = SignedTransaction{} }
func (m *SignedTransaction) String() string { return proto.CompactTextString(m) }
func (*SignedTransaction) ProtoMessage()    {}
func (*SignedTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{2}
}

func (m *SignedTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedTransaction.Unmarshal(m, b)
}
func (m *SignedTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedTransaction.Marshal(b, m, deterministic)
}
func (m *SignedTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedTransaction.Merge(m, src)
}
func (m *SignedTransaction) XXX_Size() int {
	return xxx_messageInfo_SignedTransaction.Size(m)
}
func (m *SignedTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_SignedTransaction proto.InternalMessageInfo

func (m *SignedTransaction) GetTrx() *Transaction {
	if m != nil {
		return m.Trx
	}
	return nil
}

func (m *SignedTransaction) GetSignatures() []*SignatureType {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type TransactionInvoice struct {
	Status               uint32   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	VmError              bool     `protobuf:"varint,2,opt,name=vm_error,json=vmError,proto3" json:"vm_error,omitempty"`
	VmErrorCode          uint32   `protobuf:"varint,3,opt,name=vm_error_code,json=vmErrorCode,proto3" json:"vm_error_code,omitempty"`
	VmErrorMsg           string   `protobuf:"bytes,4,opt,name=vm_error_msg,json=vmErrorMsg,proto3" json:"vm_error_msg,omitempty"`
	GasUsage             uint64   `protobuf:"varint,5,opt,name=gas_usage,json=gasUsage,proto3" json:"gas_usage,omitempty"`
	VmConsole            string   `protobuf:"bytes,6,opt,name=vm_console,json=vmConsole,proto3" json:"vm_console,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionInvoice) Reset()         { *m = TransactionInvoice{} }
func (m *TransactionInvoice) String() string { return proto.CompactTextString(m) }
func (*TransactionInvoice) ProtoMessage()    {}
func (*TransactionInvoice) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{3}
}

func (m *TransactionInvoice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionInvoice.Unmarshal(m, b)
}
func (m *TransactionInvoice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionInvoice.Marshal(b, m, deterministic)
}
func (m *TransactionInvoice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionInvoice.Merge(m, src)
}
func (m *TransactionInvoice) XXX_Size() int {
	return xxx_messageInfo_TransactionInvoice.Size(m)
}
func (m *TransactionInvoice) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionInvoice.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionInvoice proto.InternalMessageInfo

func (m *TransactionInvoice) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *TransactionInvoice) GetVmError() bool {
	if m != nil {
		return m.VmError
	}
	return false
}

func (m *TransactionInvoice) GetVmErrorCode() uint32 {
	if m != nil {
		return m.VmErrorCode
	}
	return 0
}

func (m *TransactionInvoice) GetVmErrorMsg() string {
	if m != nil {
		return m.VmErrorMsg
	}
	return ""
}

func (m *TransactionInvoice) GetGasUsage() uint64 {
	if m != nil {
		return m.GasUsage
	}
	return 0
}

func (m *TransactionInvoice) GetVmConsole() string {
	if m != nil {
		return m.VmConsole
	}
	return ""
}

type TransactionWrapper struct {
	SigTrx               *SignedTransaction  `protobuf:"bytes,1,opt,name=sig_trx,json=sigTrx,proto3" json:"sig_trx,omitempty"`
	Invoice              *TransactionInvoice `protobuf:"bytes,2,opt,name=invoice,proto3" json:"invoice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TransactionWrapper) Reset()         { *m = TransactionWrapper{} }
func (m *TransactionWrapper) String() string { return proto.CompactTextString(m) }
func (*TransactionWrapper) ProtoMessage()    {}
func (*TransactionWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{4}
}

func (m *TransactionWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionWrapper.Unmarshal(m, b)
}
func (m *TransactionWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionWrapper.Marshal(b, m, deterministic)
}
func (m *TransactionWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionWrapper.Merge(m, src)
}
func (m *TransactionWrapper) XXX_Size() int {
	return xxx_messageInfo_TransactionWrapper.Size(m)
}
func (m *TransactionWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionWrapper proto.InternalMessageInfo

func (m *TransactionWrapper) GetSigTrx() *SignedTransaction {
	if m != nil {
		return m.SigTrx
	}
	return nil
}

func (m *TransactionWrapper) GetInvoice() *TransactionInvoice {
	if m != nil {
		return m.Invoice
	}
	return nil
}

// block
type BlockHeader struct {
	Previous              *Sha256       `protobuf:"bytes,1,opt,name=previous,proto3" json:"previous,omitempty"`
	Timestamp             *TimePointSec `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Witness               string        `protobuf:"bytes,3,opt,name=witness,proto3" json:"witness,omitempty"`
	TransactionMerkleRoot *Sha256       `protobuf:"bytes,4,opt,name=transaction_merkle_root,json=transactionMerkleRoot,proto3" json:"transaction_merkle_root,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}      `json:"-"`
	XXX_unrecognized      []byte        `json:"-"`
	XXX_sizecache         int32         `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{5}
}

func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetPrevious() *Sha256 {
	if m != nil {
		return m.Previous
	}
	return nil
}

func (m *BlockHeader) GetTimestamp() *TimePointSec {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *BlockHeader) GetWitness() string {
	if m != nil {
		return m.Witness
	}
	return ""
}

func (m *BlockHeader) GetTransactionMerkleRoot() *Sha256 {
	if m != nil {
		return m.TransactionMerkleRoot
	}
	return nil
}

type SignedBlockHeader struct {
	Header               *BlockHeader   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	WitnessSignature     *SignatureType `protobuf:"bytes,2,opt,name=witness_signature,json=witnessSignature,proto3" json:"witness_signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SignedBlockHeader) Reset()         { *m = SignedBlockHeader{} }
func (m *SignedBlockHeader) String() string { return proto.CompactTextString(m) }
func (*SignedBlockHeader) ProtoMessage()    {}
func (*SignedBlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{6}
}

func (m *SignedBlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedBlockHeader.Unmarshal(m, b)
}
func (m *SignedBlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedBlockHeader.Marshal(b, m, deterministic)
}
func (m *SignedBlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedBlockHeader.Merge(m, src)
}
func (m *SignedBlockHeader) XXX_Size() int {
	return xxx_messageInfo_SignedBlockHeader.Size(m)
}
func (m *SignedBlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedBlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_SignedBlockHeader proto.InternalMessageInfo

func (m *SignedBlockHeader) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SignedBlockHeader) GetWitnessSignature() *SignatureType {
	if m != nil {
		return m.WitnessSignature
	}
	return nil
}

type SignedBlock struct {
	SignedHeader         *SignedBlockHeader    `protobuf:"bytes,1,opt,name=signed_header,json=signedHeader,proto3" json:"signed_header,omitempty"`
	Transactions         []*TransactionWrapper `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SignedBlock) Reset()         { *m = SignedBlock{} }
func (m *SignedBlock) String() string { return proto.CompactTextString(m) }
func (*SignedBlock) ProtoMessage()    {}
func (*SignedBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{7}
}

func (m *SignedBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedBlock.Unmarshal(m, b)
}
func (m *SignedBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedBlock.Marshal(b, m, deterministic)
}
func (m *SignedBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedBlock.Merge(m, src)
}
func (m *SignedBlock) XXX_Size() int {
	return xxx_messageInfo_SignedBlock.Size(m)
}
func (m *SignedBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedBlock.DiscardUnknown(m)
}

var xxx_messageInfo_SignedBlock proto.InternalMessageInfo

func (m *SignedBlock) GetSignedHeader() *SignedBlockHeader {
	if m != nil {
		return m.SignedHeader
	}
	return nil
}

func (m *SignedBlock) GetTransactions() []*TransactionWrapper {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func init() {
	proto.RegisterType((*Operation)(nil), "prototype.operation")
	proto.RegisterType((*Transaction)(nil), "prototype.transaction")
	proto.RegisterType((*SignedTransaction)(nil), "prototype.signed_transaction")
	proto.RegisterType((*TransactionInvoice)(nil), "prototype.transaction_invoice")
	proto.RegisterType((*TransactionWrapper)(nil), "prototype.transaction_wrapper")
	proto.RegisterType((*BlockHeader)(nil), "prototype.block_header")
	proto.RegisterType((*SignedBlockHeader)(nil), "prototype.signed_block_header")
	proto.RegisterType((*SignedBlock)(nil), "prototype.signed_block")
}

func init() { proto.RegisterFile("prototype/transaction.proto", fileDescriptor_f3aa2bc02ae1e20c) }

var fileDescriptor_f3aa2bc02ae1e20c = []byte{
	// 807 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xc1, 0x6e, 0xdb, 0x36,
	0x18, 0xc7, 0xe7, 0xc4, 0xb3, 0xad, 0xcf, 0xf1, 0xd0, 0xb2, 0x5d, 0xab, 0x34, 0x4b, 0x11, 0x68,
	0x17, 0x63, 0x58, 0xe2, 0xc6, 0x49, 0x9c, 0xe4, 0xb2, 0x43, 0x82, 0x0d, 0xdd, 0xa1, 0xc3, 0xc0,
	0x6d, 0x97, 0x5d, 0x08, 0x46, 0xa1, 0x15, 0xa2, 0x96, 0x48, 0x90, 0xb4, 0x93, 0x3c, 0xc1, 0x6e,
	0xbb, 0xee, 0xb4, 0x17, 0x1a, 0xf6, 0x06, 0x7b, 0x99, 0x81, 0x14, 0x2d, 0x53, 0xb2, 0xd3, 0x9b,
	0xf8, 0xf1, 0xff, 0x93, 0xfe, 0xdf, 0xc7, 0x3f, 0x05, 0x7b, 0x52, 0x09, 0x23, 0xcc, 0xa3, 0x64,
	0x23, 0xa3, 0x68, 0xa1, 0x69, 0x6a, 0xb8, 0x28, 0x8e, 0x5c, 0x15, 0x45, 0xd5, 0xe6, 0x9b, 0x97,
	0x81, 0xee, 0x51, 0xb2, 0x52, 0xf0, 0x66, 0x77, 0x55, 0x15, 0x92, 0x29, 0xba, 0x62, 0x93, 0xbf,
	0xdb, 0x10, 0x55, 0x35, 0x74, 0x0e, 0xdb, 0x42, 0x1e, 0xc7, 0xad, 0x83, 0xd6, 0xb0, 0x3f, 0xfe,
	0xfa, 0xa8, 0xc2, 0x8e, 0x68, 0x9a, 0x8a, 0x79, 0x61, 0x48, 0xaa, 0x18, 0x35, 0x8c, 0x54, 0xc4,
	0xfb, 0xcf, 0xb0, 0x25, 0xd0, 0xb1, 0x05, 0xc7, 0xf1, 0x96, 0x03, 0xf7, 0x03, 0xd0, 0xb9, 0x9d,
	0x32, 0xd5, 0x44, 0xc6, 0xe8, 0xd4, 0x22, 0x27, 0xf1, 0xb6, 0x43, 0x0e, 0x02, 0xe4, 0x46, 0x12,
	0xc5, 0x32, 0xae, 0xcd, 0x3a, 0x75, 0x82, 0x26, 0x96, 0x3a, 0x8d, 0xdb, 0x8e, 0x4a, 0xea, 0xd4,
	0xbc, 0x78, 0x8a, 0x3b, 0x45, 0xef, 0x2c, 0x77, 0x16, 0x7f, 0xee, 0xb8, 0xaf, 0xea, 0xdc, 0x42,
	0xac, 0xb7, 0x74, 0x86, 0x0e, 0x2d, 0x31, 0x89, 0x3b, 0x8e, 0xd8, 0x0d, 0x08, 0x29, 0xb4, 0x69,
	0xca, 0x27, 0x68, 0x64, 0xe5, 0xe7, 0x71, 0xd7, 0xc9, 0xf7, 0x02, 0xb9, 0x62, 0x72, 0x46, 0x1f,
	0x9b, 0xc0, 0x79, 0x09, 0x5c, 0xc4, 0xbd, 0x35, 0x60, 0x2a, 0x66, 0x33, 0x71, 0xdf, 0x04, 0x2e,
	0x4a, 0x43, 0x97, 0x71, 0xb4, 0x66, 0x68, 0x93, 0xff, 0x4b, 0xf4, 0x1d, 0xb4, 0x85, 0x3c, 0x7e,
	0x17, 0x83, 0xd3, 0x0f, 0x37, 0x9d, 0x89, 0x11, 0x64, 0xc1, 0xb4, 0xe1, 0x45, 0x56, 0xc3, 0x1d,
	0x77, 0xd5, 0x86, 0x2d, 0x21, 0x93, 0x7f, 0x5b, 0xd0, 0x0f, 0x12, 0x87, 0x12, 0x18, 0x28, 0x36,
	0x25, 0x37, 0x33, 0x91, 0x7e, 0x24, 0xc5, 0x3c, 0x77, 0x59, 0x19, 0xe0, 0xbe, 0x62, 0xd3, 0x2b,
	0x5b, 0xfb, 0x69, 0x9e, 0xa3, 0x21, 0x3c, 0x5b, 0x69, 0xa4, 0x62, 0x53, 0xfe, 0xe0, 0x92, 0x31,
	0xc0, 0x5f, 0x2c, 0x65, 0x3f, 0xbb, 0x2a, 0xba, 0x04, 0x60, 0x0f, 0x92, 0x97, 0x5f, 0xf6, 0x51,
	0x08, 0x3b, 0x33, 0x3c, 0x67, 0x44, 0x0a, 0x5e, 0x18, 0xa2, 0x59, 0x8a, 0x03, 0x31, 0x3a, 0x05,
	0xa8, 0x3c, 0xeb, 0xb8, 0x7d, 0xb0, 0x3d, 0xec, 0x8f, 0x5f, 0x06, 0x68, 0xb5, 0x89, 0x03, 0x5d,
	0xf2, 0x08, 0x48, 0xf3, 0xac, 0x60, 0xb7, 0x24, 0x6c, 0x6a, 0x08, 0xdb, 0x46, 0x3d, 0xf8, 0xd8,
	0xbf, 0x6a, 0x4e, 0xaa, 0x14, 0x61, 0x2b, 0xb1, 0x86, 0x2d, 0x4f, 0xcd, 0x5c, 0x31, 0x1d, 0x6f,
	0xb9, 0xaf, 0x86, 0x86, 0xab, 0x4d, 0x62, 0x97, 0x38, 0x10, 0x27, 0xff, 0xb4, 0xe0, 0x45, 0xf0,
	0x3e, 0xc2, 0x8b, 0x85, 0xe0, 0x29, 0x43, 0xaf, 0xa0, 0xa3, 0x0d, 0x35, 0x73, 0xed, 0x47, 0xe9,
	0x57, 0x68, 0x17, 0x7a, 0x8b, 0x9c, 0x30, 0xa5, 0x84, 0x72, 0xd3, 0xeb, 0xe1, 0xee, 0x22, 0xff,
	0xde, 0x2e, 0xed, 0x21, 0x2c, 0xb7, 0x48, 0x2a, 0x6e, 0x99, 0x9b, 0xdc, 0x00, 0xf7, 0xfd, 0xfe,
	0xb5, 0xb8, 0x65, 0xe8, 0x00, 0x76, 0x2a, 0x4d, 0xae, 0x33, 0x77, 0x63, 0x22, 0x0c, 0x5e, 0xf2,
	0x41, 0x67, 0x68, 0x0f, 0xa2, 0x8c, 0x6a, 0x32, 0xd7, 0x34, 0x63, 0xee, 0x62, 0xb4, 0x71, 0x2f,
	0xa3, 0xfa, 0x37, 0xbb, 0x46, 0xfb, 0x00, 0x8b, 0x9c, 0xa4, 0xa2, 0xd0, 0x62, 0xc6, 0xdc, 0x25,
	0x88, 0x70, 0xb4, 0xc8, 0xaf, 0xcb, 0x42, 0xf2, 0x47, 0xa3, 0x99, 0x7b, 0x45, 0xa5, 0x64, 0x0a,
	0x4d, 0xa0, 0xab, 0x79, 0x46, 0x56, 0xd3, 0xdc, 0x6f, 0x0c, 0xa7, 0x3e, 0x79, 0xdc, 0xd1, 0x3c,
	0xfb, 0x55, 0x3d, 0xa0, 0x0b, 0xe8, 0xfa, 0x79, 0xf8, 0x7f, 0xc8, 0xdb, 0xcd, 0xa7, 0xb0, 0x9c,
	0x1a, 0x5e, 0xca, 0x93, 0xff, 0x5a, 0xb0, 0x53, 0x26, 0xed, 0x8e, 0xd1, 0x5b, 0xa6, 0xd0, 0x21,
	0xf4, 0xa4, 0x62, 0x0b, 0x2e, 0xfc, 0x44, 0xfb, 0xe3, 0xe7, 0xa1, 0x87, 0x3b, 0x3a, 0x3e, 0x9b,
	0xe0, 0x4a, 0x82, 0xce, 0x21, 0xb2, 0x29, 0xd3, 0x86, 0xe6, 0xd2, 0x7f, 0xfb, 0x13, 0x09, 0x5c,
	0x69, 0x51, 0x0c, 0xdd, 0x7b, 0x6e, 0x0a, 0xa6, 0xb5, 0x1b, 0x7f, 0x84, 0x97, 0x4b, 0xf4, 0x23,
	0xbc, 0x0e, 0x2d, 0xe7, 0x4c, 0x7d, 0x9c, 0x31, 0xa2, 0x84, 0x30, 0xfe, 0xbf, 0xb5, 0xc1, 0xd0,
	0x97, 0x01, 0xf1, 0xc1, 0x01, 0x58, 0x08, 0x93, 0xfc, 0xd9, 0x82, 0x17, 0x7e, 0x6c, 0xb5, 0x26,
	0x47, 0xd0, 0x29, 0x9f, 0x7c, 0x8b, 0xaf, 0xc3, 0x3f, 0x5a, 0x20, 0xc4, 0x5e, 0x86, 0x7e, 0x80,
	0xe7, 0xde, 0x1e, 0xa9, 0x32, 0xb9, 0xa1, 0xdd, 0x46, 0x7e, 0x9f, 0x79, 0xe6, 0x97, 0x65, 0x39,
	0xf9, 0xab, 0x05, 0x3b, 0xa1, 0x21, 0x74, 0x0d, 0x03, 0xbf, 0xae, 0x19, 0x7a, 0xbb, 0x7e, 0xee,
	0x35, 0x5f, 0xfe, 0x25, 0xef, 0x4b, 0x77, 0x57, 0xb0, 0x13, 0xf4, 0xbf, 0xbc, 0x58, 0x4f, 0x65,
	0xc0, 0x87, 0x0d, 0xd7, 0x98, 0xab, 0x6f, 0x7f, 0xff, 0x26, 0xe3, 0xe6, 0x6e, 0x7e, 0x73, 0x94,
	0x8a, 0x7c, 0x94, 0x0a, 0x9d, 0xde, 0x51, 0x5e, 0x8c, 0x52, 0x51, 0x18, 0x56, 0x18, 0xa1, 0x0f,
	0x33, 0x31, 0xaa, 0xde, 0x77, 0xd3, 0x71, 0x8f, 0x27, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x9f,
	0x74, 0xaf, 0xb4, 0x59, 0x07, 0x00, 0x00,
}
