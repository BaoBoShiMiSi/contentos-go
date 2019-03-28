// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_rewardsKeeper.proto

package table // import "github.com/coschain/contentos-go/app/table"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import prototype "github.com/coschain/contentos-go/prototype"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SoRewardsKeeper struct {
	Id                   int32                            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Keeper               *prototype.InternalRewardsKeeper `protobuf:"bytes,2,opt,name=keeper,proto3" json:"keeper,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *SoRewardsKeeper) Reset()         { *m = SoRewardsKeeper{} }
func (m *SoRewardsKeeper) String() string { return proto.CompactTextString(m) }
func (*SoRewardsKeeper) ProtoMessage()    {}
func (*SoRewardsKeeper) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_rewardsKeeper_679cac1776f8bc04, []int{0}
}
func (m *SoRewardsKeeper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoRewardsKeeper.Unmarshal(m, b)
}
func (m *SoRewardsKeeper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoRewardsKeeper.Marshal(b, m, deterministic)
}
func (dst *SoRewardsKeeper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoRewardsKeeper.Merge(dst, src)
}
func (m *SoRewardsKeeper) XXX_Size() int {
	return xxx_messageInfo_SoRewardsKeeper.Size(m)
}
func (m *SoRewardsKeeper) XXX_DiscardUnknown() {
	xxx_messageInfo_SoRewardsKeeper.DiscardUnknown(m)
}

var xxx_messageInfo_SoRewardsKeeper proto.InternalMessageInfo

func (m *SoRewardsKeeper) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SoRewardsKeeper) GetKeeper() *prototype.InternalRewardsKeeper {
	if m != nil {
		return m.Keeper
	}
	return nil
}

type SoMemRewardsKeeperById struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemRewardsKeeperById) Reset()         { *m = SoMemRewardsKeeperById{} }
func (m *SoMemRewardsKeeperById) String() string { return proto.CompactTextString(m) }
func (*SoMemRewardsKeeperById) ProtoMessage()    {}
func (*SoMemRewardsKeeperById) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_rewardsKeeper_679cac1776f8bc04, []int{1}
}
func (m *SoMemRewardsKeeperById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemRewardsKeeperById.Unmarshal(m, b)
}
func (m *SoMemRewardsKeeperById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemRewardsKeeperById.Marshal(b, m, deterministic)
}
func (dst *SoMemRewardsKeeperById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemRewardsKeeperById.Merge(dst, src)
}
func (m *SoMemRewardsKeeperById) XXX_Size() int {
	return xxx_messageInfo_SoMemRewardsKeeperById.Size(m)
}
func (m *SoMemRewardsKeeperById) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemRewardsKeeperById.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemRewardsKeeperById proto.InternalMessageInfo

func (m *SoMemRewardsKeeperById) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SoMemRewardsKeeperByKeeper struct {
	Keeper               *prototype.InternalRewardsKeeper `protobuf:"bytes,1,opt,name=keeper,proto3" json:"keeper,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *SoMemRewardsKeeperByKeeper) Reset()         { *m = SoMemRewardsKeeperByKeeper{} }
func (m *SoMemRewardsKeeperByKeeper) String() string { return proto.CompactTextString(m) }
func (*SoMemRewardsKeeperByKeeper) ProtoMessage()    {}
func (*SoMemRewardsKeeperByKeeper) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_rewardsKeeper_679cac1776f8bc04, []int{2}
}
func (m *SoMemRewardsKeeperByKeeper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemRewardsKeeperByKeeper.Unmarshal(m, b)
}
func (m *SoMemRewardsKeeperByKeeper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemRewardsKeeperByKeeper.Marshal(b, m, deterministic)
}
func (dst *SoMemRewardsKeeperByKeeper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemRewardsKeeperByKeeper.Merge(dst, src)
}
func (m *SoMemRewardsKeeperByKeeper) XXX_Size() int {
	return xxx_messageInfo_SoMemRewardsKeeperByKeeper.Size(m)
}
func (m *SoMemRewardsKeeperByKeeper) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemRewardsKeeperByKeeper.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemRewardsKeeperByKeeper proto.InternalMessageInfo

func (m *SoMemRewardsKeeperByKeeper) GetKeeper() *prototype.InternalRewardsKeeper {
	if m != nil {
		return m.Keeper
	}
	return nil
}

type SoUniqueRewardsKeeperById struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoUniqueRewardsKeeperById) Reset()         { *m = SoUniqueRewardsKeeperById{} }
func (m *SoUniqueRewardsKeeperById) String() string { return proto.CompactTextString(m) }
func (*SoUniqueRewardsKeeperById) ProtoMessage()    {}
func (*SoUniqueRewardsKeeperById) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_rewardsKeeper_679cac1776f8bc04, []int{3}
}
func (m *SoUniqueRewardsKeeperById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueRewardsKeeperById.Unmarshal(m, b)
}
func (m *SoUniqueRewardsKeeperById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueRewardsKeeperById.Marshal(b, m, deterministic)
}
func (dst *SoUniqueRewardsKeeperById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueRewardsKeeperById.Merge(dst, src)
}
func (m *SoUniqueRewardsKeeperById) XXX_Size() int {
	return xxx_messageInfo_SoUniqueRewardsKeeperById.Size(m)
}
func (m *SoUniqueRewardsKeeperById) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueRewardsKeeperById.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueRewardsKeeperById proto.InternalMessageInfo

func (m *SoUniqueRewardsKeeperById) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*SoRewardsKeeper)(nil), "table.so_rewardsKeeper")
	proto.RegisterType((*SoMemRewardsKeeperById)(nil), "table.so_mem_rewardsKeeper_by_id")
	proto.RegisterType((*SoMemRewardsKeeperByKeeper)(nil), "table.so_mem_rewardsKeeper_by_keeper")
	proto.RegisterType((*SoUniqueRewardsKeeperById)(nil), "table.so_unique_rewardsKeeper_by_id")
}

func init() {
	proto.RegisterFile("app/table/so_rewardsKeeper.proto", fileDescriptor_so_rewardsKeeper_679cac1776f8bc04)
}

var fileDescriptor_so_rewardsKeeper_679cac1776f8bc04 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xbf, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0x49, 0xe1, 0x6e, 0x88, 0x20, 0x52, 0x1c, 0x8e, 0x03, 0xa5, 0x64, 0x3a, 0xe4, 0x4c,
	0x40, 0x37, 0x47, 0x57, 0xb7, 0x1b, 0x45, 0x0c, 0xf9, 0xf1, 0xb8, 0x06, 0xdb, 0xbc, 0x98, 0xa4,
	0x48, 0xff, 0x7b, 0x31, 0xd5, 0xa2, 0x05, 0x41, 0x97, 0x40, 0xf8, 0xfe, 0xfa, 0xf0, 0x68, 0xa3,
	0x42, 0x10, 0x59, 0xe9, 0x0e, 0x44, 0x42, 0x19, 0xe1, 0x4d, 0x45, 0x9b, 0x1e, 0x00, 0x02, 0x44,
	0x1e, 0x22, 0x66, 0xac, 0x57, 0x45, 0xdd, 0x9e, 0x97, 0x5f, 0x1e, 0x03, 0x88, 0x8f, 0x67, 0x12,
	0xd9, 0x33, 0x3d, 0x5b, 0xc6, 0xea, 0x53, 0x5a, 0x39, 0xbb, 0x21, 0x0d, 0xd9, 0xad, 0x0e, 0x95,
	0xb3, 0xf5, 0x1d, 0x5d, 0xbf, 0x14, 0x65, 0x53, 0x35, 0x64, 0x77, 0x72, 0xc3, 0xf8, 0x5c, 0xc5,
	0x9d, 0xcf, 0x10, 0xbd, 0xea, 0xbe, 0x2a, 0xe4, 0xe4, 0x3c, 0x7c, 0x26, 0xd8, 0x9e, 0x6e, 0x13,
	0xca, 0x1e, 0xfa, 0x9f, 0x1b, 0x52, 0x8f, 0xd2, 0xd9, 0xe5, 0x12, 0x7b, 0xa2, 0x97, 0xbf, 0xb9,
	0xa7, 0xbe, 0x6f, 0x2c, 0xe4, 0xdf, 0x2c, 0x82, 0x5e, 0x24, 0x94, 0x83, 0x77, 0xaf, 0x03, 0xfc,
	0x05, 0xe7, 0x7e, 0xff, 0x78, 0x75, 0x74, 0xb9, 0x1d, 0x34, 0x37, 0xd8, 0x0b, 0x83, 0xc9, 0xb4,
	0xca, 0x79, 0x61, 0xd0, 0x67, 0xf0, 0x19, 0xd3, 0xf5, 0x11, 0xc5, 0x7c, 0x7e, 0xbd, 0x2e, 0x24,
	0xb7, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x06, 0x8c, 0x62, 0xdb, 0x92, 0x01, 0x00, 0x00,
}
