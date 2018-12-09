// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_post.proto

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

type SoPost struct {
	PostId               uint64                            `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Category             string                            `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Author               *prototype.AccountName            `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Title                string                            `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string                            `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Tags                 []string                          `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	Created              *prototype.TimePointSec           `protobuf:"bytes,7,opt,name=created,proto3" json:"created,omitempty"`
	LastPayout           *prototype.TimePointSec           `protobuf:"bytes,8,opt,name=last_payout,json=lastPayout,proto3" json:"last_payout,omitempty"`
	Depth                uint32                            `protobuf:"varint,9,opt,name=depth,proto3" json:"depth,omitempty"`
	Children             uint32                            `protobuf:"varint,10,opt,name=children,proto3" json:"children,omitempty"`
	RootId               uint64                            `protobuf:"varint,11,opt,name=root_id,json=rootId,proto3" json:"root_id,omitempty"`
	ParentId             uint64                            `protobuf:"varint,12,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	VoteCnt              uint64                            `protobuf:"varint,13,opt,name=vote_cnt,json=voteCnt,proto3" json:"vote_cnt,omitempty"`
	Beneficiaries        []*prototype.BeneficiaryRouteType `protobuf:"bytes,14,rep,name=beneficiaries,proto3" json:"beneficiaries,omitempty"`
	CashoutTime          *prototype.TimePointSec           `protobuf:"bytes,15,opt,name=cashout_time,json=cashoutTime,proto3" json:"cashout_time,omitempty"`
	WeightedVp           uint64                            `protobuf:"varint,16,opt,name=weighted_vp,json=weightedVp,proto3" json:"weighted_vp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *SoPost) Reset()         { *m = SoPost{} }
func (m *SoPost) String() string { return proto.CompactTextString(m) }
func (*SoPost) ProtoMessage()    {}
func (*SoPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{0}
}

func (m *SoPost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoPost.Unmarshal(m, b)
}
func (m *SoPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoPost.Marshal(b, m, deterministic)
}
func (m *SoPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoPost.Merge(m, src)
}
func (m *SoPost) XXX_Size() int {
	return xxx_messageInfo_SoPost.Size(m)
}
func (m *SoPost) XXX_DiscardUnknown() {
	xxx_messageInfo_SoPost.DiscardUnknown(m)
}

var xxx_messageInfo_SoPost proto.InternalMessageInfo

func (m *SoPost) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func (m *SoPost) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *SoPost) GetAuthor() *prototype.AccountName {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *SoPost) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SoPost) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *SoPost) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *SoPost) GetCreated() *prototype.TimePointSec {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *SoPost) GetLastPayout() *prototype.TimePointSec {
	if m != nil {
		return m.LastPayout
	}
	return nil
}

func (m *SoPost) GetDepth() uint32 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func (m *SoPost) GetChildren() uint32 {
	if m != nil {
		return m.Children
	}
	return 0
}

func (m *SoPost) GetRootId() uint64 {
	if m != nil {
		return m.RootId
	}
	return 0
}

func (m *SoPost) GetParentId() uint64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *SoPost) GetVoteCnt() uint64 {
	if m != nil {
		return m.VoteCnt
	}
	return 0
}

func (m *SoPost) GetBeneficiaries() []*prototype.BeneficiaryRouteType {
	if m != nil {
		return m.Beneficiaries
	}
	return nil
}

func (m *SoPost) GetCashoutTime() *prototype.TimePointSec {
	if m != nil {
		return m.CashoutTime
	}
	return nil
}

func (m *SoPost) GetWeightedVp() uint64 {
	if m != nil {
		return m.WeightedVp
	}
	return 0
}

type SoMemPostByPostId struct {
	PostId               uint64   `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemPostByPostId) Reset()         { *m = SoMemPostByPostId{} }
func (m *SoMemPostByPostId) String() string { return proto.CompactTextString(m) }
func (*SoMemPostByPostId) ProtoMessage()    {}
func (*SoMemPostByPostId) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{1}
}

func (m *SoMemPostByPostId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemPostByPostId.Unmarshal(m, b)
}
func (m *SoMemPostByPostId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemPostByPostId.Marshal(b, m, deterministic)
}
func (m *SoMemPostByPostId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemPostByPostId.Merge(m, src)
}
func (m *SoMemPostByPostId) XXX_Size() int {
	return xxx_messageInfo_SoMemPostByPostId.Size(m)
}
func (m *SoMemPostByPostId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemPostByPostId.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemPostByPostId proto.InternalMessageInfo

func (m *SoMemPostByPostId) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

type SoMemPostByCategory struct {
	Category             string   `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemPostByCategory) Reset()         { *m = SoMemPostByCategory{} }
func (m *SoMemPostByCategory) String() string { return proto.CompactTextString(m) }
func (*SoMemPostByCategory) ProtoMessage()    {}
func (*SoMemPostByCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{2}
}

func (m *SoMemPostByCategory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemPostByCategory.Unmarshal(m, b)
}
func (m *SoMemPostByCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemPostByCategory.Marshal(b, m, deterministic)
}
func (m *SoMemPostByCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemPostByCategory.Merge(m, src)
}
func (m *SoMemPostByCategory) XXX_Size() int {
	return xxx_messageInfo_SoMemPostByCategory.Size(m)
}
func (m *SoMemPostByCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemPostByCategory.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemPostByCategory proto.InternalMessageInfo

func (m *SoMemPostByCategory) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

type SoMemPostByAuthor struct {
	Author               *prototype.AccountName `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoMemPostByAuthor) Reset()         { *m = SoMemPostByAuthor{} }
func (m *SoMemPostByAuthor) String() string { return proto.CompactTextString(m) }
func (*SoMemPostByAuthor) ProtoMessage()    {}
func (*SoMemPostByAuthor) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{3}
}

func (m *SoMemPostByAuthor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemPostByAuthor.Unmarshal(m, b)
}
func (m *SoMemPostByAuthor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemPostByAuthor.Marshal(b, m, deterministic)
}
func (m *SoMemPostByAuthor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemPostByAuthor.Merge(m, src)
}
func (m *SoMemPostByAuthor) XXX_Size() int {
	return xxx_messageInfo_SoMemPostByAuthor.Size(m)
}
func (m *SoMemPostByAuthor) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemPostByAuthor.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemPostByAuthor proto.InternalMessageInfo

func (m *SoMemPostByAuthor) GetAuthor() *prototype.AccountName {
	if m != nil {
		return m.Author
	}
	return nil
}

type SoMemPostByTitle struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemPostByTitle) Reset()         { *m = SoMemPostByTitle{} }
func (m *SoMemPostByTitle) String() string { return proto.CompactTextString(m) }
func (*SoMemPostByTitle) ProtoMessage()    {}
func (*SoMemPostByTitle) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{4}
}

func (m *SoMemPostByTitle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemPostByTitle.Unmarshal(m, b)
}
func (m *SoMemPostByTitle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemPostByTitle.Marshal(b, m, deterministic)
}
func (m *SoMemPostByTitle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemPostByTitle.Merge(m, src)
}
func (m *SoMemPostByTitle) XXX_Size() int {
	return xxx_messageInfo_SoMemPostByTitle.Size(m)
}
func (m *SoMemPostByTitle) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemPostByTitle.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemPostByTitle proto.InternalMessageInfo

func (m *SoMemPostByTitle) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type SoMemPostByBody struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemPostByBody) Reset()         { *m = SoMemPostByBody{} }
func (m *SoMemPostByBody) String() string { return proto.CompactTextString(m) }
func (*SoMemPostByBody) ProtoMessage()    {}
func (*SoMemPostByBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{5}
}

func (m *SoMemPostByBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemPostByBody.Unmarshal(m, b)
}
func (m *SoMemPostByBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemPostByBody.Marshal(b, m, deterministic)
}
func (m *SoMemPostByBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemPostByBody.Merge(m, src)
}
func (m *SoMemPostByBody) XXX_Size() int {
	return xxx_messageInfo_SoMemPostByBody.Size(m)
}
func (m *SoMemPostByBody) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemPostByBody.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemPostByBody proto.InternalMessageInfo

func (m *SoMemPostByBody) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type SoMemPostByTags struct {
	Tags                 []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemPostByTags) Reset()         { *m = SoMemPostByTags{} }
func (m *SoMemPostByTags) String() string { return proto.CompactTextString(m) }
func (*SoMemPostByTags) ProtoMessage()    {}
func (*SoMemPostByTags) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{6}
}

func (m *SoMemPostByTags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemPostByTags.Unmarshal(m, b)
}
func (m *SoMemPostByTags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemPostByTags.Marshal(b, m, deterministic)
}
func (m *SoMemPostByTags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemPostByTags.Merge(m, src)
}
func (m *SoMemPostByTags) XXX_Size() int {
	return xxx_messageInfo_SoMemPostByTags.Size(m)
}
func (m *SoMemPostByTags) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemPostByTags.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemPostByTags proto.InternalMessageInfo

func (m *SoMemPostByTags) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type SoMemPostByCreated struct {
	Created              *prototype.TimePointSec `protobuf:"bytes,1,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoMemPostByCreated) Reset()         { *m = SoMemPostByCreated{} }
func (m *SoMemPostByCreated) String() string { return proto.CompactTextString(m) }
func (*SoMemPostByCreated) ProtoMessage()    {}
func (*SoMemPostByCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{7}
}

func (m *SoMemPostByCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemPostByCreated.Unmarshal(m, b)
}
func (m *SoMemPostByCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemPostByCreated.Marshal(b, m, deterministic)
}
func (m *SoMemPostByCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemPostByCreated.Merge(m, src)
}
func (m *SoMemPostByCreated) XXX_Size() int {
	return xxx_messageInfo_SoMemPostByCreated.Size(m)
}
func (m *SoMemPostByCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemPostByCreated.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemPostByCreated proto.InternalMessageInfo

func (m *SoMemPostByCreated) GetCreated() *prototype.TimePointSec {
	if m != nil {
		return m.Created
	}
	return nil
}

type SoMemPostByLastPayout struct {
	LastPayout           *prototype.TimePointSec `protobuf:"bytes,1,opt,name=last_payout,json=lastPayout,proto3" json:"last_payout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoMemPostByLastPayout) Reset()         { *m = SoMemPostByLastPayout{} }
func (m *SoMemPostByLastPayout) String() string { return proto.CompactTextString(m) }
func (*SoMemPostByLastPayout) ProtoMessage()    {}
func (*SoMemPostByLastPayout) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{8}
}

func (m *SoMemPostByLastPayout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemPostByLastPayout.Unmarshal(m, b)
}
func (m *SoMemPostByLastPayout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemPostByLastPayout.Marshal(b, m, deterministic)
}
func (m *SoMemPostByLastPayout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemPostByLastPayout.Merge(m, src)
}
func (m *SoMemPostByLastPayout) XXX_Size() int {
	return xxx_messageInfo_SoMemPostByLastPayout.Size(m)
}
func (m *SoMemPostByLastPayout) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemPostByLastPayout.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemPostByLastPayout proto.InternalMessageInfo

func (m *SoMemPostByLastPayout) GetLastPayout() *prototype.TimePointSec {
	if m != nil {
		return m.LastPayout
	}
	return nil
}

type SoMemPostByDepth struct {
	Depth                uint32   `protobuf:"varint,1,opt,name=depth,proto3" json:"depth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemPostByDepth) Reset()         { *m = SoMemPostByDepth{} }
func (m *SoMemPostByDepth) String() string { return proto.CompactTextString(m) }
func (*SoMemPostByDepth) ProtoMessage()    {}
func (*SoMemPostByDepth) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{9}
}

func (m *SoMemPostByDepth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemPostByDepth.Unmarshal(m, b)
}
func (m *SoMemPostByDepth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemPostByDepth.Marshal(b, m, deterministic)
}
func (m *SoMemPostByDepth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemPostByDepth.Merge(m, src)
}
func (m *SoMemPostByDepth) XXX_Size() int {
	return xxx_messageInfo_SoMemPostByDepth.Size(m)
}
func (m *SoMemPostByDepth) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemPostByDepth.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemPostByDepth proto.InternalMessageInfo

func (m *SoMemPostByDepth) GetDepth() uint32 {
	if m != nil {
		return m.Depth
	}
	return 0
}

type SoMemPostByChildren struct {
	Children             uint32   `protobuf:"varint,1,opt,name=children,proto3" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemPostByChildren) Reset()         { *m = SoMemPostByChildren{} }
func (m *SoMemPostByChildren) String() string { return proto.CompactTextString(m) }
func (*SoMemPostByChildren) ProtoMessage()    {}
func (*SoMemPostByChildren) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{10}
}

func (m *SoMemPostByChildren) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemPostByChildren.Unmarshal(m, b)
}
func (m *SoMemPostByChildren) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemPostByChildren.Marshal(b, m, deterministic)
}
func (m *SoMemPostByChildren) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemPostByChildren.Merge(m, src)
}
func (m *SoMemPostByChildren) XXX_Size() int {
	return xxx_messageInfo_SoMemPostByChildren.Size(m)
}
func (m *SoMemPostByChildren) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemPostByChildren.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemPostByChildren proto.InternalMessageInfo

func (m *SoMemPostByChildren) GetChildren() uint32 {
	if m != nil {
		return m.Children
	}
	return 0
}

type SoMemPostByRootId struct {
	RootId               uint64   `protobuf:"varint,1,opt,name=root_id,json=rootId,proto3" json:"root_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemPostByRootId) Reset()         { *m = SoMemPostByRootId{} }
func (m *SoMemPostByRootId) String() string { return proto.CompactTextString(m) }
func (*SoMemPostByRootId) ProtoMessage()    {}
func (*SoMemPostByRootId) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{11}
}

func (m *SoMemPostByRootId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemPostByRootId.Unmarshal(m, b)
}
func (m *SoMemPostByRootId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemPostByRootId.Marshal(b, m, deterministic)
}
func (m *SoMemPostByRootId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemPostByRootId.Merge(m, src)
}
func (m *SoMemPostByRootId) XXX_Size() int {
	return xxx_messageInfo_SoMemPostByRootId.Size(m)
}
func (m *SoMemPostByRootId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemPostByRootId.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemPostByRootId proto.InternalMessageInfo

func (m *SoMemPostByRootId) GetRootId() uint64 {
	if m != nil {
		return m.RootId
	}
	return 0
}

type SoMemPostByParentId struct {
	ParentId             uint64   `protobuf:"varint,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemPostByParentId) Reset()         { *m = SoMemPostByParentId{} }
func (m *SoMemPostByParentId) String() string { return proto.CompactTextString(m) }
func (*SoMemPostByParentId) ProtoMessage()    {}
func (*SoMemPostByParentId) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{12}
}

func (m *SoMemPostByParentId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemPostByParentId.Unmarshal(m, b)
}
func (m *SoMemPostByParentId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemPostByParentId.Marshal(b, m, deterministic)
}
func (m *SoMemPostByParentId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemPostByParentId.Merge(m, src)
}
func (m *SoMemPostByParentId) XXX_Size() int {
	return xxx_messageInfo_SoMemPostByParentId.Size(m)
}
func (m *SoMemPostByParentId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemPostByParentId.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemPostByParentId proto.InternalMessageInfo

func (m *SoMemPostByParentId) GetParentId() uint64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

type SoMemPostByVoteCnt struct {
	VoteCnt              uint64   `protobuf:"varint,1,opt,name=vote_cnt,json=voteCnt,proto3" json:"vote_cnt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemPostByVoteCnt) Reset()         { *m = SoMemPostByVoteCnt{} }
func (m *SoMemPostByVoteCnt) String() string { return proto.CompactTextString(m) }
func (*SoMemPostByVoteCnt) ProtoMessage()    {}
func (*SoMemPostByVoteCnt) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{13}
}

func (m *SoMemPostByVoteCnt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemPostByVoteCnt.Unmarshal(m, b)
}
func (m *SoMemPostByVoteCnt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemPostByVoteCnt.Marshal(b, m, deterministic)
}
func (m *SoMemPostByVoteCnt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemPostByVoteCnt.Merge(m, src)
}
func (m *SoMemPostByVoteCnt) XXX_Size() int {
	return xxx_messageInfo_SoMemPostByVoteCnt.Size(m)
}
func (m *SoMemPostByVoteCnt) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemPostByVoteCnt.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemPostByVoteCnt proto.InternalMessageInfo

func (m *SoMemPostByVoteCnt) GetVoteCnt() uint64 {
	if m != nil {
		return m.VoteCnt
	}
	return 0
}

type SoMemPostByBeneficiaries struct {
	Beneficiaries        []*prototype.BeneficiaryRouteType `protobuf:"bytes,1,rep,name=beneficiaries,proto3" json:"beneficiaries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *SoMemPostByBeneficiaries) Reset()         { *m = SoMemPostByBeneficiaries{} }
func (m *SoMemPostByBeneficiaries) String() string { return proto.CompactTextString(m) }
func (*SoMemPostByBeneficiaries) ProtoMessage()    {}
func (*SoMemPostByBeneficiaries) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{14}
}

func (m *SoMemPostByBeneficiaries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemPostByBeneficiaries.Unmarshal(m, b)
}
func (m *SoMemPostByBeneficiaries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemPostByBeneficiaries.Marshal(b, m, deterministic)
}
func (m *SoMemPostByBeneficiaries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemPostByBeneficiaries.Merge(m, src)
}
func (m *SoMemPostByBeneficiaries) XXX_Size() int {
	return xxx_messageInfo_SoMemPostByBeneficiaries.Size(m)
}
func (m *SoMemPostByBeneficiaries) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemPostByBeneficiaries.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemPostByBeneficiaries proto.InternalMessageInfo

func (m *SoMemPostByBeneficiaries) GetBeneficiaries() []*prototype.BeneficiaryRouteType {
	if m != nil {
		return m.Beneficiaries
	}
	return nil
}

type SoMemPostByCashoutTime struct {
	CashoutTime          *prototype.TimePointSec `protobuf:"bytes,1,opt,name=cashout_time,json=cashoutTime,proto3" json:"cashout_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoMemPostByCashoutTime) Reset()         { *m = SoMemPostByCashoutTime{} }
func (m *SoMemPostByCashoutTime) String() string { return proto.CompactTextString(m) }
func (*SoMemPostByCashoutTime) ProtoMessage()    {}
func (*SoMemPostByCashoutTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{15}
}

func (m *SoMemPostByCashoutTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemPostByCashoutTime.Unmarshal(m, b)
}
func (m *SoMemPostByCashoutTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemPostByCashoutTime.Marshal(b, m, deterministic)
}
func (m *SoMemPostByCashoutTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemPostByCashoutTime.Merge(m, src)
}
func (m *SoMemPostByCashoutTime) XXX_Size() int {
	return xxx_messageInfo_SoMemPostByCashoutTime.Size(m)
}
func (m *SoMemPostByCashoutTime) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemPostByCashoutTime.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemPostByCashoutTime proto.InternalMessageInfo

func (m *SoMemPostByCashoutTime) GetCashoutTime() *prototype.TimePointSec {
	if m != nil {
		return m.CashoutTime
	}
	return nil
}

type SoMemPostByWeightedVp struct {
	WeightedVp           uint64   `protobuf:"varint,1,opt,name=weighted_vp,json=weightedVp,proto3" json:"weighted_vp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemPostByWeightedVp) Reset()         { *m = SoMemPostByWeightedVp{} }
func (m *SoMemPostByWeightedVp) String() string { return proto.CompactTextString(m) }
func (*SoMemPostByWeightedVp) ProtoMessage()    {}
func (*SoMemPostByWeightedVp) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{16}
}

func (m *SoMemPostByWeightedVp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemPostByWeightedVp.Unmarshal(m, b)
}
func (m *SoMemPostByWeightedVp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemPostByWeightedVp.Marshal(b, m, deterministic)
}
func (m *SoMemPostByWeightedVp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemPostByWeightedVp.Merge(m, src)
}
func (m *SoMemPostByWeightedVp) XXX_Size() int {
	return xxx_messageInfo_SoMemPostByWeightedVp.Size(m)
}
func (m *SoMemPostByWeightedVp) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemPostByWeightedVp.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemPostByWeightedVp proto.InternalMessageInfo

func (m *SoMemPostByWeightedVp) GetWeightedVp() uint64 {
	if m != nil {
		return m.WeightedVp
	}
	return 0
}

type SoListPostByCreated struct {
	Created              *prototype.TimePointSec `protobuf:"bytes,1,opt,name=created,proto3" json:"created,omitempty"`
	PostId               uint64                  `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoListPostByCreated) Reset()         { *m = SoListPostByCreated{} }
func (m *SoListPostByCreated) String() string { return proto.CompactTextString(m) }
func (*SoListPostByCreated) ProtoMessage()    {}
func (*SoListPostByCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{17}
}

func (m *SoListPostByCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListPostByCreated.Unmarshal(m, b)
}
func (m *SoListPostByCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListPostByCreated.Marshal(b, m, deterministic)
}
func (m *SoListPostByCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListPostByCreated.Merge(m, src)
}
func (m *SoListPostByCreated) XXX_Size() int {
	return xxx_messageInfo_SoListPostByCreated.Size(m)
}
func (m *SoListPostByCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListPostByCreated.DiscardUnknown(m)
}

var xxx_messageInfo_SoListPostByCreated proto.InternalMessageInfo

func (m *SoListPostByCreated) GetCreated() *prototype.TimePointSec {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *SoListPostByCreated) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

type SoUniquePostByPostId struct {
	PostId               uint64   `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoUniquePostByPostId) Reset()         { *m = SoUniquePostByPostId{} }
func (m *SoUniquePostByPostId) String() string { return proto.CompactTextString(m) }
func (*SoUniquePostByPostId) ProtoMessage()    {}
func (*SoUniquePostByPostId) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{18}
}

func (m *SoUniquePostByPostId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniquePostByPostId.Unmarshal(m, b)
}
func (m *SoUniquePostByPostId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniquePostByPostId.Marshal(b, m, deterministic)
}
func (m *SoUniquePostByPostId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniquePostByPostId.Merge(m, src)
}
func (m *SoUniquePostByPostId) XXX_Size() int {
	return xxx_messageInfo_SoUniquePostByPostId.Size(m)
}
func (m *SoUniquePostByPostId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniquePostByPostId.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniquePostByPostId proto.InternalMessageInfo

func (m *SoUniquePostByPostId) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func init() {
	proto.RegisterType((*SoPost)(nil), "table.so_post")
	proto.RegisterType((*SoMemPostByPostId)(nil), "table.so_mem_post_by_post_id")
	proto.RegisterType((*SoMemPostByCategory)(nil), "table.so_mem_post_by_category")
	proto.RegisterType((*SoMemPostByAuthor)(nil), "table.so_mem_post_by_author")
	proto.RegisterType((*SoMemPostByTitle)(nil), "table.so_mem_post_by_title")
	proto.RegisterType((*SoMemPostByBody)(nil), "table.so_mem_post_by_body")
	proto.RegisterType((*SoMemPostByTags)(nil), "table.so_mem_post_by_tags")
	proto.RegisterType((*SoMemPostByCreated)(nil), "table.so_mem_post_by_created")
	proto.RegisterType((*SoMemPostByLastPayout)(nil), "table.so_mem_post_by_last_payout")
	proto.RegisterType((*SoMemPostByDepth)(nil), "table.so_mem_post_by_depth")
	proto.RegisterType((*SoMemPostByChildren)(nil), "table.so_mem_post_by_children")
	proto.RegisterType((*SoMemPostByRootId)(nil), "table.so_mem_post_by_root_id")
	proto.RegisterType((*SoMemPostByParentId)(nil), "table.so_mem_post_by_parent_id")
	proto.RegisterType((*SoMemPostByVoteCnt)(nil), "table.so_mem_post_by_vote_cnt")
	proto.RegisterType((*SoMemPostByBeneficiaries)(nil), "table.so_mem_post_by_beneficiaries")
	proto.RegisterType((*SoMemPostByCashoutTime)(nil), "table.so_mem_post_by_cashout_time")
	proto.RegisterType((*SoMemPostByWeightedVp)(nil), "table.so_mem_post_by_weighted_vp")
	proto.RegisterType((*SoListPostByCreated)(nil), "table.so_list_post_by_created")
	proto.RegisterType((*SoUniquePostByPostId)(nil), "table.so_unique_post_by_post_id")
}

func init() { proto.RegisterFile("app/table/so_post.proto", fileDescriptor_aaea493edfbbad11) }

var fileDescriptor_aaea493edfbbad11 = []byte{
	// 659 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0xc0, 0xe5, 0xad, 0x6b, 0xbb, 0xd7, 0x0d, 0x90, 0x19, 0xd4, 0xdb, 0x90, 0x28, 0x3e, 0xa0,
	0x22, 0x41, 0x2b, 0xb6, 0x21, 0x24, 0x04, 0x17, 0x38, 0xc0, 0x0e, 0x48, 0xa8, 0x42, 0x08, 0xc1,
	0xc1, 0x72, 0x13, 0x93, 0x58, 0x6a, 0xe3, 0x90, 0xbc, 0x0c, 0xf5, 0x23, 0xf2, 0xad, 0x90, 0x9d,
	0xa6, 0xcd, 0x9f, 0xa2, 0x75, 0x88, 0x4b, 0x9b, 0xf7, 0xc7, 0x89, 0xf3, 0xfa, 0xfb, 0xb9, 0xd0,
	0x97, 0x71, 0x3c, 0x46, 0x39, 0x9d, 0xa9, 0x71, 0x6a, 0x44, 0x6c, 0x52, 0x1c, 0xc5, 0x89, 0x41,
	0x43, 0xf7, 0x5c, 0xf2, 0xe4, 0xc8, 0x45, 0xb8, 0x88, 0xd5, 0xd8, 0x7e, 0xe4, 0x45, 0xfe, 0xbb,
	0x05, 0x9d, 0x65, 0x3b, 0xed, 0x43, 0xc7, 0x7e, 0x0b, 0xed, 0x33, 0x32, 0x20, 0xc3, 0xd6, 0xa4,
	0x6d, 0xc3, 0x4b, 0x9f, 0x9e, 0x40, 0xd7, 0x93, 0xa8, 0x02, 0x93, 0x2c, 0xd8, 0xce, 0x80, 0x0c,
	0xf7, 0x27, 0xab, 0x98, 0x8e, 0xa1, 0x2d, 0x33, 0x0c, 0x4d, 0xc2, 0x76, 0x07, 0x64, 0xd8, 0x3b,
	0xeb, 0x8f, 0x56, 0xcf, 0x19, 0x49, 0xcf, 0x33, 0x59, 0x84, 0x22, 0x92, 0x73, 0x35, 0x59, 0xb6,
	0xd1, 0x23, 0xd8, 0x43, 0x8d, 0x33, 0xc5, 0x5a, 0xee, 0x4e, 0x79, 0x40, 0x29, 0xb4, 0xa6, 0xc6,
	0x5f, 0xb0, 0x3d, 0x97, 0x74, 0xd7, 0x36, 0x87, 0x32, 0x48, 0x59, 0x7b, 0xb0, 0x6b, 0x73, 0xf6,
	0x9a, 0x9e, 0x43, 0xc7, 0x4b, 0x94, 0x44, 0xe5, 0xb3, 0x8e, 0x7b, 0xde, 0x71, 0xe9, 0x79, 0xa8,
	0xe7, 0x4a, 0xc4, 0x46, 0x47, 0x28, 0x52, 0xe5, 0x4d, 0x8a, 0x4e, 0xfa, 0x0a, 0x7a, 0x33, 0x99,
	0xa2, 0x88, 0xe5, 0xc2, 0x64, 0xc8, 0xba, 0xd7, 0x2d, 0x04, 0xdb, 0xfd, 0xc9, 0x35, 0xdb, 0xed,
	0xfa, 0x2a, 0xc6, 0x90, 0xed, 0x0f, 0xc8, 0xf0, 0x70, 0x92, 0x07, 0x6e, 0x22, 0xa1, 0x9e, 0xf9,
	0x89, 0x8a, 0x18, 0xb8, 0xc2, 0x2a, 0xb6, 0x63, 0x4c, 0x8c, 0x71, 0x63, 0xec, 0xe5, 0x63, 0xb4,
	0xe1, 0xa5, 0x4f, 0x4f, 0x61, 0x3f, 0x96, 0x89, 0x8a, 0x5c, 0xe9, 0xc0, 0x95, 0xba, 0x79, 0xe2,
	0xd2, 0xa7, 0xc7, 0xd0, 0xbd, 0x32, 0xa8, 0x84, 0x17, 0x21, 0x3b, 0x74, 0xb5, 0x8e, 0x8d, 0xdf,
	0x45, 0x48, 0xdf, 0xc3, 0xe1, 0x54, 0x45, 0xea, 0x87, 0xf6, 0xb4, 0x4c, 0xb4, 0x4a, 0xd9, 0xad,
	0xc1, 0xee, 0xb0, 0x77, 0xf6, 0xa8, 0xf4, 0x02, 0xeb, 0xfa, 0x42, 0x24, 0x26, 0x43, 0x25, 0x6c,
	0x7a, 0x52, 0x5d, 0x47, 0x5f, 0xc3, 0x81, 0x27, 0xd3, 0xd0, 0x64, 0x28, 0xec, 0x1b, 0xb3, 0xdb,
	0xd7, 0x0d, 0xa2, 0xb7, 0x6c, 0xff, 0xac, 0xe7, 0x8a, 0x3e, 0x84, 0xde, 0x2f, 0xa5, 0x83, 0x10,
	0x95, 0x2f, 0xae, 0x62, 0x76, 0xc7, 0x6d, 0x12, 0x8a, 0xd4, 0x97, 0x98, 0x3f, 0x87, 0xfb, 0xa9,
	0x11, 0x73, 0x35, 0x77, 0x38, 0x89, 0xe9, 0x42, 0x2c, 0x71, 0xfa, 0x2b, 0x59, 0xfc, 0x05, 0xf4,
	0x6b, 0x4b, 0x56, 0x60, 0x95, 0xa1, 0x23, 0x55, 0xe8, 0xf8, 0x07, 0xb8, 0x57, 0x5b, 0xb6, 0x84,
	0x6b, 0x4d, 0x23, 0xd9, 0x8a, 0x46, 0xfe, 0x14, 0x8e, 0x6a, 0x77, 0xca, 0x79, 0x5c, 0x51, 0x4a,
	0x4a, 0x94, 0xf2, 0x27, 0x70, 0xb7, 0xd6, 0x5d, 0x80, 0xea, 0xe0, 0x25, 0x6b, 0x78, 0x37, 0xb4,
	0x3a, 0x7e, 0x0b, 0xa6, 0xc9, 0x9a, 0x69, 0xfe, 0xb1, 0x31, 0xb7, 0x02, 0xdc, 0x12, 0xed, 0x64,
	0x5b, 0xda, 0xf9, 0x57, 0x38, 0xa9, 0xdd, 0xae, 0x04, 0x7f, 0xdd, 0x05, 0x72, 0x03, 0x17, 0x36,
	0x0c, 0x2b, 0xb7, 0x61, 0xe5, 0x08, 0x29, 0x39, 0xb2, 0xe9, 0xb7, 0x2d, 0x14, 0x29, 0xeb, 0x43,
	0xaa, 0xfa, 0x6c, 0xa0, 0x68, 0x69, 0x53, 0x59, 0x2c, 0x52, 0x16, 0x8b, 0xbf, 0x04, 0x56, 0x07,
	0xaf, 0xf0, 0xac, 0x2a, 0x1d, 0xa9, 0x4a, 0xc7, 0x2f, 0x1a, 0x5b, 0x2c, 0x1c, 0xac, 0xf8, 0x48,
	0x2a, 0x3e, 0xf2, 0x00, 0x1e, 0xd4, 0x29, 0xa8, 0x68, 0xd6, 0xf0, 0x95, 0xfc, 0x9b, 0xaf, 0xfc,
	0x3b, 0x9c, 0x36, 0xec, 0x58, 0xeb, 0xdb, 0xd0, 0x99, 0xdc, 0x44, 0x67, 0xfe, 0xa6, 0x81, 0x49,
	0xc9, 0xee, 0xba, 0xec, 0xa4, 0x21, 0x7b, 0xe0, 0x46, 0x37, 0xd3, 0x16, 0xa5, 0xff, 0x40, 0x6d,
	0xf9, 0x88, 0xd8, 0xa9, 0x1c, 0x11, 0x17, 0x70, 0x9c, 0x1a, 0x91, 0x45, 0xfa, 0x67, 0xa6, 0xb6,
	0x3e, 0x58, 0xde, 0x0e, 0xbf, 0x3d, 0x0e, 0x34, 0x86, 0xd9, 0x74, 0xe4, 0x99, 0xf9, 0xd8, 0x33,
	0xa9, 0x17, 0x4a, 0x1d, 0x8d, 0x3d, 0x13, 0xa1, 0x8a, 0xd0, 0xa4, 0xcf, 0x02, 0x93, 0xff, 0x59,
	0x4e, 0xdb, 0x6e, 0x6f, 0xe7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xf2, 0x44, 0x44, 0x40,
	0x07, 0x00, 0x00,
}
