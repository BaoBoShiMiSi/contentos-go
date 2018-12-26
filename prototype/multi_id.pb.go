// Code generated by protoc-gen-go. DO NOT EDIT.
// source: prototype/multi_id.proto

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

type FollowerRelation struct {
	Account              *AccountName `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Follower             *AccountName `protobuf:"bytes,2,opt,name=follower,proto3" json:"follower,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FollowerRelation) Reset()         { *m = FollowerRelation{} }
func (m *FollowerRelation) String() string { return proto.CompactTextString(m) }
func (*FollowerRelation) ProtoMessage()    {}
func (*FollowerRelation) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b47f83ece5cae8f, []int{0}
}

func (m *FollowerRelation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FollowerRelation.Unmarshal(m, b)
}
func (m *FollowerRelation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FollowerRelation.Marshal(b, m, deterministic)
}
func (m *FollowerRelation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FollowerRelation.Merge(m, src)
}
func (m *FollowerRelation) XXX_Size() int {
	return xxx_messageInfo_FollowerRelation.Size(m)
}
func (m *FollowerRelation) XXX_DiscardUnknown() {
	xxx_messageInfo_FollowerRelation.DiscardUnknown(m)
}

var xxx_messageInfo_FollowerRelation proto.InternalMessageInfo

func (m *FollowerRelation) GetAccount() *AccountName {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *FollowerRelation) GetFollower() *AccountName {
	if m != nil {
		return m.Follower
	}
	return nil
}

type FollowingRelation struct {
	Account              *AccountName `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Following            *AccountName `protobuf:"bytes,2,opt,name=following,proto3" json:"following,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FollowingRelation) Reset()         { *m = FollowingRelation{} }
func (m *FollowingRelation) String() string { return proto.CompactTextString(m) }
func (*FollowingRelation) ProtoMessage()    {}
func (*FollowingRelation) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b47f83ece5cae8f, []int{1}
}

func (m *FollowingRelation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FollowingRelation.Unmarshal(m, b)
}
func (m *FollowingRelation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FollowingRelation.Marshal(b, m, deterministic)
}
func (m *FollowingRelation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FollowingRelation.Merge(m, src)
}
func (m *FollowingRelation) XXX_Size() int {
	return xxx_messageInfo_FollowingRelation.Size(m)
}
func (m *FollowingRelation) XXX_DiscardUnknown() {
	xxx_messageInfo_FollowingRelation.DiscardUnknown(m)
}

var xxx_messageInfo_FollowingRelation proto.InternalMessageInfo

func (m *FollowingRelation) GetAccount() *AccountName {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *FollowingRelation) GetFollowing() *AccountName {
	if m != nil {
		return m.Following
	}
	return nil
}

type FollowerCreatedOrder struct {
	Account              *AccountName  `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	CreatedTime          *TimePointSec `protobuf:"bytes,2,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	Follower             *AccountName  `protobuf:"bytes,3,opt,name=follower,proto3" json:"follower,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FollowerCreatedOrder) Reset()         { *m = FollowerCreatedOrder{} }
func (m *FollowerCreatedOrder) String() string { return proto.CompactTextString(m) }
func (*FollowerCreatedOrder) ProtoMessage()    {}
func (*FollowerCreatedOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b47f83ece5cae8f, []int{2}
}

func (m *FollowerCreatedOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FollowerCreatedOrder.Unmarshal(m, b)
}
func (m *FollowerCreatedOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FollowerCreatedOrder.Marshal(b, m, deterministic)
}
func (m *FollowerCreatedOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FollowerCreatedOrder.Merge(m, src)
}
func (m *FollowerCreatedOrder) XXX_Size() int {
	return xxx_messageInfo_FollowerCreatedOrder.Size(m)
}
func (m *FollowerCreatedOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_FollowerCreatedOrder.DiscardUnknown(m)
}

var xxx_messageInfo_FollowerCreatedOrder proto.InternalMessageInfo

func (m *FollowerCreatedOrder) GetAccount() *AccountName {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *FollowerCreatedOrder) GetCreatedTime() *TimePointSec {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

func (m *FollowerCreatedOrder) GetFollower() *AccountName {
	if m != nil {
		return m.Follower
	}
	return nil
}

type FollowingCreatedOrder struct {
	Account              *AccountName  `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	CreatedTime          *TimePointSec `protobuf:"bytes,2,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	Following            *AccountName  `protobuf:"bytes,3,opt,name=following,proto3" json:"following,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FollowingCreatedOrder) Reset()         { *m = FollowingCreatedOrder{} }
func (m *FollowingCreatedOrder) String() string { return proto.CompactTextString(m) }
func (*FollowingCreatedOrder) ProtoMessage()    {}
func (*FollowingCreatedOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b47f83ece5cae8f, []int{3}
}

func (m *FollowingCreatedOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FollowingCreatedOrder.Unmarshal(m, b)
}
func (m *FollowingCreatedOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FollowingCreatedOrder.Marshal(b, m, deterministic)
}
func (m *FollowingCreatedOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FollowingCreatedOrder.Merge(m, src)
}
func (m *FollowingCreatedOrder) XXX_Size() int {
	return xxx_messageInfo_FollowingCreatedOrder.Size(m)
}
func (m *FollowingCreatedOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_FollowingCreatedOrder.DiscardUnknown(m)
}

var xxx_messageInfo_FollowingCreatedOrder proto.InternalMessageInfo

func (m *FollowingCreatedOrder) GetAccount() *AccountName {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *FollowingCreatedOrder) GetCreatedTime() *TimePointSec {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

func (m *FollowingCreatedOrder) GetFollowing() *AccountName {
	if m != nil {
		return m.Following
	}
	return nil
}

type PostCreatedOrder struct {
	Created              *TimePointSec `protobuf:"bytes,1,opt,name=created,proto3" json:"created,omitempty"`
	ParentId             uint64        `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PostCreatedOrder) Reset()         { *m = PostCreatedOrder{} }
func (m *PostCreatedOrder) String() string { return proto.CompactTextString(m) }
func (*PostCreatedOrder) ProtoMessage()    {}
func (*PostCreatedOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b47f83ece5cae8f, []int{4}
}

func (m *PostCreatedOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostCreatedOrder.Unmarshal(m, b)
}
func (m *PostCreatedOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostCreatedOrder.Marshal(b, m, deterministic)
}
func (m *PostCreatedOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostCreatedOrder.Merge(m, src)
}
func (m *PostCreatedOrder) XXX_Size() int {
	return xxx_messageInfo_PostCreatedOrder.Size(m)
}
func (m *PostCreatedOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_PostCreatedOrder.DiscardUnknown(m)
}

var xxx_messageInfo_PostCreatedOrder proto.InternalMessageInfo

func (m *PostCreatedOrder) GetCreated() *TimePointSec {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *PostCreatedOrder) GetParentId() uint64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

type VoterId struct {
	Voter                *AccountName `protobuf:"bytes,1,opt,name=voter,proto3" json:"voter,omitempty"`
	PostId               uint64       `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *VoterId) Reset()         { *m = VoterId{} }
func (m *VoterId) String() string { return proto.CompactTextString(m) }
func (*VoterId) ProtoMessage()    {}
func (*VoterId) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b47f83ece5cae8f, []int{5}
}

func (m *VoterId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoterId.Unmarshal(m, b)
}
func (m *VoterId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoterId.Marshal(b, m, deterministic)
}
func (m *VoterId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoterId.Merge(m, src)
}
func (m *VoterId) XXX_Size() int {
	return xxx_messageInfo_VoterId.Size(m)
}
func (m *VoterId) XXX_DiscardUnknown() {
	xxx_messageInfo_VoterId.DiscardUnknown(m)
}

var xxx_messageInfo_VoterId proto.InternalMessageInfo

func (m *VoterId) GetVoter() *AccountName {
	if m != nil {
		return m.Voter
	}
	return nil
}

func (m *VoterId) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

type BpVoterId struct {
	Voter                *AccountName `protobuf:"bytes,1,opt,name=voter,proto3" json:"voter,omitempty"`
	Witness              *AccountName `protobuf:"bytes,2,opt,name=witness,proto3" json:"witness,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BpVoterId) Reset()         { *m = BpVoterId{} }
func (m *BpVoterId) String() string { return proto.CompactTextString(m) }
func (*BpVoterId) ProtoMessage()    {}
func (*BpVoterId) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b47f83ece5cae8f, []int{6}
}

func (m *BpVoterId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BpVoterId.Unmarshal(m, b)
}
func (m *BpVoterId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BpVoterId.Marshal(b, m, deterministic)
}
func (m *BpVoterId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BpVoterId.Merge(m, src)
}
func (m *BpVoterId) XXX_Size() int {
	return xxx_messageInfo_BpVoterId.Size(m)
}
func (m *BpVoterId) XXX_DiscardUnknown() {
	xxx_messageInfo_BpVoterId.DiscardUnknown(m)
}

var xxx_messageInfo_BpVoterId proto.InternalMessageInfo

func (m *BpVoterId) GetVoter() *AccountName {
	if m != nil {
		return m.Voter
	}
	return nil
}

func (m *BpVoterId) GetWitness() *AccountName {
	if m != nil {
		return m.Witness
	}
	return nil
}

type BpWitnessId struct {
	Voter                *AccountName `protobuf:"bytes,1,opt,name=voter,proto3" json:"voter,omitempty"`
	Witness              *AccountName `protobuf:"bytes,2,opt,name=witness,proto3" json:"witness,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BpWitnessId) Reset()         { *m = BpWitnessId{} }
func (m *BpWitnessId) String() string { return proto.CompactTextString(m) }
func (*BpWitnessId) ProtoMessage()    {}
func (*BpWitnessId) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b47f83ece5cae8f, []int{7}
}

func (m *BpWitnessId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BpWitnessId.Unmarshal(m, b)
}
func (m *BpWitnessId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BpWitnessId.Marshal(b, m, deterministic)
}
func (m *BpWitnessId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BpWitnessId.Merge(m, src)
}
func (m *BpWitnessId) XXX_Size() int {
	return xxx_messageInfo_BpWitnessId.Size(m)
}
func (m *BpWitnessId) XXX_DiscardUnknown() {
	xxx_messageInfo_BpWitnessId.DiscardUnknown(m)
}

var xxx_messageInfo_BpWitnessId proto.InternalMessageInfo

func (m *BpWitnessId) GetVoter() *AccountName {
	if m != nil {
		return m.Voter
	}
	return nil
}

func (m *BpWitnessId) GetWitness() *AccountName {
	if m != nil {
		return m.Witness
	}
	return nil
}

type ContractId struct {
	Owner                *AccountName `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Cname                string       `protobuf:"bytes,2,opt,name=cname,proto3" json:"cname,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ContractId) Reset()         { *m = ContractId{} }
func (m *ContractId) String() string { return proto.CompactTextString(m) }
func (*ContractId) ProtoMessage()    {}
func (*ContractId) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b47f83ece5cae8f, []int{8}
}

func (m *ContractId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractId.Unmarshal(m, b)
}
func (m *ContractId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractId.Marshal(b, m, deterministic)
}
func (m *ContractId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractId.Merge(m, src)
}
func (m *ContractId) XXX_Size() int {
	return xxx_messageInfo_ContractId.Size(m)
}
func (m *ContractId) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractId.DiscardUnknown(m)
}

var xxx_messageInfo_ContractId proto.InternalMessageInfo

func (m *ContractId) GetOwner() *AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *ContractId) GetCname() string {
	if m != nil {
		return m.Cname
	}
	return ""
}

type ContractDataId struct {
	Owner                *AccountName `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Cname                string       `protobuf:"bytes,2,opt,name=cname,proto3" json:"cname,omitempty"`
	Pos                  int32        `protobuf:"varint,3,opt,name=pos,proto3" json:"pos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ContractDataId) Reset()         { *m = ContractDataId{} }
func (m *ContractDataId) String() string { return proto.CompactTextString(m) }
func (*ContractDataId) ProtoMessage()    {}
func (*ContractDataId) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b47f83ece5cae8f, []int{9}
}

func (m *ContractDataId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractDataId.Unmarshal(m, b)
}
func (m *ContractDataId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractDataId.Marshal(b, m, deterministic)
}
func (m *ContractDataId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractDataId.Merge(m, src)
}
func (m *ContractDataId) XXX_Size() int {
	return xxx_messageInfo_ContractDataId.Size(m)
}
func (m *ContractDataId) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractDataId.DiscardUnknown(m)
}

var xxx_messageInfo_ContractDataId proto.InternalMessageInfo

func (m *ContractDataId) GetOwner() *AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *ContractDataId) GetCname() string {
	if m != nil {
		return m.Cname
	}
	return ""
}

func (m *ContractDataId) GetPos() int32 {
	if m != nil {
		return m.Pos
	}
	return 0
}

type ReplyCreatedOrder struct {
	ParentId             uint64        `protobuf:"varint,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Created              *TimePointSec `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReplyCreatedOrder) Reset()         { *m = ReplyCreatedOrder{} }
func (m *ReplyCreatedOrder) String() string { return proto.CompactTextString(m) }
func (*ReplyCreatedOrder) ProtoMessage()    {}
func (*ReplyCreatedOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b47f83ece5cae8f, []int{10}
}

func (m *ReplyCreatedOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyCreatedOrder.Unmarshal(m, b)
}
func (m *ReplyCreatedOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyCreatedOrder.Marshal(b, m, deterministic)
}
func (m *ReplyCreatedOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyCreatedOrder.Merge(m, src)
}
func (m *ReplyCreatedOrder) XXX_Size() int {
	return xxx_messageInfo_ReplyCreatedOrder.Size(m)
}
func (m *ReplyCreatedOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyCreatedOrder.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyCreatedOrder proto.InternalMessageInfo

func (m *ReplyCreatedOrder) GetParentId() uint64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *ReplyCreatedOrder) GetCreated() *TimePointSec {
	if m != nil {
		return m.Created
	}
	return nil
}

func init() {
	proto.RegisterType((*FollowerRelation)(nil), "prototype.follower_relation")
	proto.RegisterType((*FollowingRelation)(nil), "prototype.following_relation")
	proto.RegisterType((*FollowerCreatedOrder)(nil), "prototype.follower_created_order")
	proto.RegisterType((*FollowingCreatedOrder)(nil), "prototype.following_created_order")
	proto.RegisterType((*PostCreatedOrder)(nil), "prototype.post_created_order")
	proto.RegisterType((*VoterId)(nil), "prototype.voter_id")
	proto.RegisterType((*BpVoterId)(nil), "prototype.bp_voter_id")
	proto.RegisterType((*BpWitnessId)(nil), "prototype.bp_witness_id")
	proto.RegisterType((*ContractId)(nil), "prototype.contract_id")
	proto.RegisterType((*ContractDataId)(nil), "prototype.contract_data_id")
	proto.RegisterType((*ReplyCreatedOrder)(nil), "prototype.reply_created_order")
}

func init() { proto.RegisterFile("prototype/multi_id.proto", fileDescriptor_7b47f83ece5cae8f) }

var fileDescriptor_7b47f83ece5cae8f = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x52, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0xd5, 0x36, 0xa4, 0x49, 0x26, 0x20, 0x15, 0x53, 0x91, 0x00, 0x17, 0xe4, 0x13, 0x42, 0x34,
	0x51, 0x89, 0xb8, 0x71, 0xe2, 0xd6, 0xeb, 0x8a, 0x13, 0x17, 0xcb, 0xb1, 0xdd, 0xd4, 0xd2, 0xae,
	0xc7, 0x78, 0x27, 0x44, 0x15, 0x12, 0x9f, 0x86, 0xc4, 0x9f, 0xa1, 0xf5, 0xee, 0x3a, 0x4d, 0x0f,
	0x4d, 0x14, 0x10, 0x5c, 0x56, 0x33, 0xb3, 0xcf, 0xef, 0x3d, 0x8f, 0x1f, 0x4c, 0x7d, 0x40, 0x42,
	0xba, 0xf5, 0x66, 0x5e, 0xae, 0x0b, 0xb2, 0xc2, 0xea, 0x59, 0x1c, 0xb1, 0x51, 0xfa, 0xf3, 0xf2,
	0x7c, 0x0b, 0xaa, 0x3f, 0x0d, 0x80, 0x7f, 0x87, 0xa7, 0xd7, 0x58, 0x14, 0xb8, 0x31, 0x41, 0x04,
	0x53, 0x48, 0xb2, 0xe8, 0xd8, 0x25, 0x0c, 0xa4, 0x52, 0xb8, 0x76, 0x34, 0xcd, 0x5e, 0x67, 0x6f,
	0xc6, 0xef, 0x27, 0xb3, 0x74, 0x78, 0xd6, 0xfe, 0x11, 0x4e, 0x96, 0x26, 0xef, 0x70, 0x6c, 0x01,
	0xc3, 0x8e, 0x67, 0x7a, 0xf2, 0xf0, 0x99, 0x04, 0xe4, 0x3f, 0x80, 0x35, 0xb5, 0x75, 0xab, 0x3f,
	0x52, 0xff, 0x00, 0xa3, 0x44, 0xb4, 0x4f, 0x7e, 0x8b, 0xe4, 0x3f, 0x33, 0x78, 0x9e, 0x6e, 0xaf,
	0x82, 0x91, 0x64, 0xb4, 0xc0, 0xa0, 0x4d, 0x38, 0xc6, 0xc4, 0x47, 0x78, 0xdc, 0x71, 0x90, 0x2d,
	0x4d, 0xeb, 0xe3, 0xc5, 0x9d, 0x73, 0xf5, 0x58, 0x78, 0xb4, 0x8e, 0x44, 0x65, 0x54, 0x3e, 0x6e,
	0xe1, 0x9f, 0x6d, 0x69, 0x76, 0x16, 0xd8, 0x3b, 0x74, 0x81, 0xbf, 0x32, 0x98, 0x6c, 0x37, 0xf8,
	0x9f, 0x6f, 0xb0, 0xf3, 0x08, 0xbd, 0x83, 0x1f, 0xe1, 0x1a, 0x98, 0xc7, 0x8a, 0xee, 0xb9, 0x5f,
	0xc0, 0xa0, 0x1d, 0xb4, 0xee, 0x1f, 0x70, 0xd1, 0x21, 0xd9, 0x2b, 0x18, 0x79, 0x19, 0x8c, 0x23,
	0x61, 0x75, 0x34, 0xff, 0x28, 0x1f, 0x36, 0x83, 0x2b, 0xcd, 0x73, 0x18, 0x7e, 0x43, 0x32, 0x41,
	0x58, 0xcd, 0x2e, 0xa0, 0x1f, 0xeb, 0x7d, 0x9b, 0x69, 0x50, 0x6c, 0x02, 0x83, 0x68, 0x31, 0xb1,
	0x9e, 0xd6, 0xed, 0x95, 0xe6, 0x08, 0xe3, 0xa5, 0x17, 0xc7, 0xd2, 0x5e, 0xc2, 0x60, 0x63, 0xc9,
	0x99, 0xaa, 0xda, 0x97, 0xd9, 0x0e, 0xc7, 0xbf, 0xc2, 0x93, 0xa5, 0x17, 0x6d, 0xf7, 0x6f, 0x24,
	0x73, 0x18, 0x2b, 0x74, 0x14, 0xa4, 0xa2, 0x56, 0x10, 0x37, 0xee, 0x00, 0xc1, 0x88, 0x62, 0xe7,
	0xd0, 0x57, 0x75, 0x1f, 0xe5, 0x46, 0x79, 0xd3, 0x70, 0x0b, 0x67, 0x89, 0x53, 0x4b, 0x92, 0x7f,
	0x8b, 0x98, 0x9d, 0x41, 0xcf, 0x63, 0x15, 0xd3, 0xd7, 0xcf, 0xeb, 0x92, 0xaf, 0xe0, 0x59, 0x30,
	0xbe, 0xb8, 0xbd, 0x97, 0xaf, 0x9d, 0xa8, 0x64, 0xbb, 0x51, 0xb9, 0x1b, 0xbe, 0x93, 0x43, 0xc3,
	0xf7, 0xe9, 0xdd, 0x97, 0xb7, 0x2b, 0x4b, 0x37, 0xeb, 0xe5, 0x4c, 0x61, 0x39, 0x57, 0x58, 0xa9,
	0x1b, 0x69, 0xdd, 0xbc, 0xbe, 0xa7, 0x71, 0x84, 0xd5, 0xc5, 0x0a, 0xe7, 0x89, 0x65, 0x79, 0x1a,
	0xcb, 0xc5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xcf, 0x37, 0xf4, 0xbb, 0x05, 0x00, 0x00,
}