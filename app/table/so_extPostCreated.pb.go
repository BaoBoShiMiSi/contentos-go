// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_extPostCreated.proto

package table

import (
	fmt "fmt"
	prototype "github.com/coschain/contentos-go/prototype"
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

type SoExtPostCreated struct {
	PostId               uint64                      `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	CreatedOrder         *prototype.PostCreatedOrder `protobuf:"bytes,2,opt,name=created_order,json=createdOrder,proto3" json:"created_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SoExtPostCreated) Reset()         { *m = SoExtPostCreated{} }
func (m *SoExtPostCreated) String() string { return proto.CompactTextString(m) }
func (*SoExtPostCreated) ProtoMessage()    {}
func (*SoExtPostCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_5aa22736cbf04dfd, []int{0}
}

func (m *SoExtPostCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoExtPostCreated.Unmarshal(m, b)
}
func (m *SoExtPostCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoExtPostCreated.Marshal(b, m, deterministic)
}
func (m *SoExtPostCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoExtPostCreated.Merge(m, src)
}
func (m *SoExtPostCreated) XXX_Size() int {
	return xxx_messageInfo_SoExtPostCreated.Size(m)
}
func (m *SoExtPostCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_SoExtPostCreated.DiscardUnknown(m)
}

var xxx_messageInfo_SoExtPostCreated proto.InternalMessageInfo

func (m *SoExtPostCreated) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func (m *SoExtPostCreated) GetCreatedOrder() *prototype.PostCreatedOrder {
	if m != nil {
		return m.CreatedOrder
	}
	return nil
}

type SoMemExtPostCreatedByPostId struct {
	PostId               uint64   `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemExtPostCreatedByPostId) Reset()         { *m = SoMemExtPostCreatedByPostId{} }
func (m *SoMemExtPostCreatedByPostId) String() string { return proto.CompactTextString(m) }
func (*SoMemExtPostCreatedByPostId) ProtoMessage()    {}
func (*SoMemExtPostCreatedByPostId) Descriptor() ([]byte, []int) {
	return fileDescriptor_5aa22736cbf04dfd, []int{1}
}

func (m *SoMemExtPostCreatedByPostId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtPostCreatedByPostId.Unmarshal(m, b)
}
func (m *SoMemExtPostCreatedByPostId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtPostCreatedByPostId.Marshal(b, m, deterministic)
}
func (m *SoMemExtPostCreatedByPostId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtPostCreatedByPostId.Merge(m, src)
}
func (m *SoMemExtPostCreatedByPostId) XXX_Size() int {
	return xxx_messageInfo_SoMemExtPostCreatedByPostId.Size(m)
}
func (m *SoMemExtPostCreatedByPostId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtPostCreatedByPostId.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtPostCreatedByPostId proto.InternalMessageInfo

func (m *SoMemExtPostCreatedByPostId) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

type SoMemExtPostCreatedByCreatedOrder struct {
	CreatedOrder         *prototype.PostCreatedOrder `protobuf:"bytes,1,opt,name=created_order,json=createdOrder,proto3" json:"created_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SoMemExtPostCreatedByCreatedOrder) Reset()         { *m = SoMemExtPostCreatedByCreatedOrder{} }
func (m *SoMemExtPostCreatedByCreatedOrder) String() string { return proto.CompactTextString(m) }
func (*SoMemExtPostCreatedByCreatedOrder) ProtoMessage()    {}
func (*SoMemExtPostCreatedByCreatedOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_5aa22736cbf04dfd, []int{2}
}

func (m *SoMemExtPostCreatedByCreatedOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtPostCreatedByCreatedOrder.Unmarshal(m, b)
}
func (m *SoMemExtPostCreatedByCreatedOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtPostCreatedByCreatedOrder.Marshal(b, m, deterministic)
}
func (m *SoMemExtPostCreatedByCreatedOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtPostCreatedByCreatedOrder.Merge(m, src)
}
func (m *SoMemExtPostCreatedByCreatedOrder) XXX_Size() int {
	return xxx_messageInfo_SoMemExtPostCreatedByCreatedOrder.Size(m)
}
func (m *SoMemExtPostCreatedByCreatedOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtPostCreatedByCreatedOrder.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtPostCreatedByCreatedOrder proto.InternalMessageInfo

func (m *SoMemExtPostCreatedByCreatedOrder) GetCreatedOrder() *prototype.PostCreatedOrder {
	if m != nil {
		return m.CreatedOrder
	}
	return nil
}

type SoListExtPostCreatedByCreatedOrder struct {
	CreatedOrder         *prototype.PostCreatedOrder `protobuf:"bytes,1,opt,name=created_order,json=createdOrder,proto3" json:"created_order,omitempty"`
	PostId               uint64                      `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SoListExtPostCreatedByCreatedOrder) Reset()         { *m = SoListExtPostCreatedByCreatedOrder{} }
func (m *SoListExtPostCreatedByCreatedOrder) String() string { return proto.CompactTextString(m) }
func (*SoListExtPostCreatedByCreatedOrder) ProtoMessage()    {}
func (*SoListExtPostCreatedByCreatedOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_5aa22736cbf04dfd, []int{3}
}

func (m *SoListExtPostCreatedByCreatedOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtPostCreatedByCreatedOrder.Unmarshal(m, b)
}
func (m *SoListExtPostCreatedByCreatedOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtPostCreatedByCreatedOrder.Marshal(b, m, deterministic)
}
func (m *SoListExtPostCreatedByCreatedOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtPostCreatedByCreatedOrder.Merge(m, src)
}
func (m *SoListExtPostCreatedByCreatedOrder) XXX_Size() int {
	return xxx_messageInfo_SoListExtPostCreatedByCreatedOrder.Size(m)
}
func (m *SoListExtPostCreatedByCreatedOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtPostCreatedByCreatedOrder.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtPostCreatedByCreatedOrder proto.InternalMessageInfo

func (m *SoListExtPostCreatedByCreatedOrder) GetCreatedOrder() *prototype.PostCreatedOrder {
	if m != nil {
		return m.CreatedOrder
	}
	return nil
}

func (m *SoListExtPostCreatedByCreatedOrder) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

type SoUniqueExtPostCreatedByPostId struct {
	PostId               uint64   `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoUniqueExtPostCreatedByPostId) Reset()         { *m = SoUniqueExtPostCreatedByPostId{} }
func (m *SoUniqueExtPostCreatedByPostId) String() string { return proto.CompactTextString(m) }
func (*SoUniqueExtPostCreatedByPostId) ProtoMessage()    {}
func (*SoUniqueExtPostCreatedByPostId) Descriptor() ([]byte, []int) {
	return fileDescriptor_5aa22736cbf04dfd, []int{4}
}

func (m *SoUniqueExtPostCreatedByPostId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueExtPostCreatedByPostId.Unmarshal(m, b)
}
func (m *SoUniqueExtPostCreatedByPostId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueExtPostCreatedByPostId.Marshal(b, m, deterministic)
}
func (m *SoUniqueExtPostCreatedByPostId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueExtPostCreatedByPostId.Merge(m, src)
}
func (m *SoUniqueExtPostCreatedByPostId) XXX_Size() int {
	return xxx_messageInfo_SoUniqueExtPostCreatedByPostId.Size(m)
}
func (m *SoUniqueExtPostCreatedByPostId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueExtPostCreatedByPostId.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueExtPostCreatedByPostId proto.InternalMessageInfo

func (m *SoUniqueExtPostCreatedByPostId) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func init() {
	proto.RegisterType((*SoExtPostCreated)(nil), "table.so_extPostCreated")
	proto.RegisterType((*SoMemExtPostCreatedByPostId)(nil), "table.so_mem_extPostCreated_by_post_id")
	proto.RegisterType((*SoMemExtPostCreatedByCreatedOrder)(nil), "table.so_mem_extPostCreated_by_created_order")
	proto.RegisterType((*SoListExtPostCreatedByCreatedOrder)(nil), "table.so_list_extPostCreated_by_created_order")
	proto.RegisterType((*SoUniqueExtPostCreatedByPostId)(nil), "table.so_unique_extPostCreated_by_post_id")
}

func init() { proto.RegisterFile("app/table/so_extPostCreated.proto", fileDescriptor_5aa22736cbf04dfd) }

var fileDescriptor_5aa22736cbf04dfd = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x2c, 0x28, 0xd0,
	0x2f, 0x49, 0x4c, 0xca, 0x49, 0xd5, 0x2f, 0xce, 0x8f, 0x4f, 0xad, 0x28, 0x09, 0xc8, 0x2f, 0x2e,
	0x71, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05,
	0x4b, 0x4b, 0x89, 0x80, 0x79, 0x25, 0x95, 0x05, 0xa9, 0xfa, 0x20, 0x02, 0x22, 0xa9, 0x54, 0xc0,
	0x25, 0x88, 0xa1, 0x4f, 0x48, 0x9c, 0x8b, 0xbd, 0x20, 0xbf, 0xb8, 0x24, 0x3e, 0x33, 0x45, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x25, 0x88, 0x0d, 0xc4, 0xf5, 0x4c, 0x11, 0x72, 0xe2, 0xe2, 0x4d, 0x86,
	0xa8, 0x89, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0x92, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd5,
	0x83, 0x9b, 0xad, 0x07, 0xd6, 0x88, 0xa2, 0x28, 0x88, 0x07, 0xca, 0xf5, 0x07, 0xf1, 0x94, 0xac,
	0xb9, 0x14, 0x8a, 0xf3, 0xe3, 0x73, 0x53, 0x73, 0xd1, 0x6c, 0x8d, 0x4f, 0xaa, 0x8c, 0x87, 0xda,
	0x8a, 0xd3, 0x01, 0x4a, 0x39, 0x5c, 0x6a, 0x38, 0x35, 0xa3, 0x58, 0x8a, 0xe9, 0x54, 0x46, 0xd2,
	0x9d, 0xda, 0xc6, 0xc8, 0xa5, 0x5e, 0x9c, 0x1f, 0x9f, 0x93, 0x59, 0x5c, 0x42, 0x0f, 0xfb, 0x90,
	0xbd, 0xcd, 0x84, 0xe2, 0x6d, 0x3b, 0x2e, 0xe5, 0xe2, 0xfc, 0xf8, 0xd2, 0xbc, 0xcc, 0xc2, 0xd2,
	0x54, 0x32, 0x82, 0xcd, 0x49, 0x23, 0x4a, 0x2d, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39,
	0x3f, 0x57, 0x3f, 0x39, 0xbf, 0x38, 0x39, 0x23, 0x31, 0x33, 0x4f, 0x3f, 0x39, 0x3f, 0xaf, 0x24,
	0x35, 0xaf, 0x24, 0xbf, 0x58, 0x37, 0x3d, 0x1f, 0x92, 0x88, 0x92, 0xd8, 0xc0, 0xce, 0x35, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x3b, 0x45, 0x0d, 0x58, 0x02, 0x00, 0x00,
}
