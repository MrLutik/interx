// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kira/slashing/query.proto

package slashing

import (
	context "context"
	fmt "fmt"
	base "github.com/KiraCore/interx/proto-gen/cosmos/base"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// ValidatorSigningInfo defines a validator's signing info for monitoring their
// liveness activity.
type ValidatorSigningInfo struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// height at which validator was first a candidate OR was activated
	StartHeight int64 `protobuf:"varint,2,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty"`
	// timestamp validator cannot be activated until
	InactiveUntil *timestamp.Timestamp `protobuf:"bytes,3,opt,name=inactive_until,json=inactiveUntil,proto3" json:"inactive_until,omitempty"`
	// mischance confidence counter - count when it's in active status
	MischanceConfidence int64 `protobuf:"varint,4,opt,name=mischance_confidence,json=mischanceConfidence,proto3" json:"mischance_confidence,omitempty"`
	// missed blocks counter after mischance confidence - count when it's in active status
	Mischance int64 `protobuf:"varint,5,opt,name=mischance,proto3" json:"mischance,omitempty"`
	// last signed block height by the validator
	LastPresentBlock int64 `protobuf:"varint,6,opt,name=last_present_block,json=lastPresentBlock,proto3" json:"last_present_block,omitempty"`
	// missed blocks counter (to avoid scanning the array every time)
	MissedBlocksCounter int64 `protobuf:"varint,7,opt,name=missed_blocks_counter,json=missedBlocksCounter,proto3" json:"missed_blocks_counter,omitempty"`
	// count produced blocks so far by a validator
	ProducedBlocksCounter int64    `protobuf:"varint,8,opt,name=produced_blocks_counter,json=producedBlocksCounter,proto3" json:"produced_blocks_counter,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *ValidatorSigningInfo) Reset()         { *m = ValidatorSigningInfo{} }
func (m *ValidatorSigningInfo) String() string { return proto.CompactTextString(m) }
func (*ValidatorSigningInfo) ProtoMessage()    {}
func (*ValidatorSigningInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5858f9240a85cd9f, []int{0}
}

func (m *ValidatorSigningInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorSigningInfo.Unmarshal(m, b)
}
func (m *ValidatorSigningInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorSigningInfo.Marshal(b, m, deterministic)
}
func (m *ValidatorSigningInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSigningInfo.Merge(m, src)
}
func (m *ValidatorSigningInfo) XXX_Size() int {
	return xxx_messageInfo_ValidatorSigningInfo.Size(m)
}
func (m *ValidatorSigningInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSigningInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSigningInfo proto.InternalMessageInfo

func (m *ValidatorSigningInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ValidatorSigningInfo) GetStartHeight() int64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *ValidatorSigningInfo) GetInactiveUntil() *timestamp.Timestamp {
	if m != nil {
		return m.InactiveUntil
	}
	return nil
}

func (m *ValidatorSigningInfo) GetMischanceConfidence() int64 {
	if m != nil {
		return m.MischanceConfidence
	}
	return 0
}

func (m *ValidatorSigningInfo) GetMischance() int64 {
	if m != nil {
		return m.Mischance
	}
	return 0
}

func (m *ValidatorSigningInfo) GetLastPresentBlock() int64 {
	if m != nil {
		return m.LastPresentBlock
	}
	return 0
}

func (m *ValidatorSigningInfo) GetMissedBlocksCounter() int64 {
	if m != nil {
		return m.MissedBlocksCounter
	}
	return 0
}

func (m *ValidatorSigningInfo) GetProducedBlocksCounter() int64 {
	if m != nil {
		return m.ProducedBlocksCounter
	}
	return 0
}

// QuerySigningInfoRequest is the request type for the Query/SigningInfo RPC
// method
type QuerySigningInfoRequest struct {
	// cons_address is the address to query signing info of
	ConsAddress          string   `protobuf:"bytes,1,opt,name=cons_address,json=consAddress,proto3" json:"cons_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuerySigningInfoRequest) Reset()         { *m = QuerySigningInfoRequest{} }
func (m *QuerySigningInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySigningInfoRequest) ProtoMessage()    {}
func (*QuerySigningInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5858f9240a85cd9f, []int{1}
}

func (m *QuerySigningInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuerySigningInfoRequest.Unmarshal(m, b)
}
func (m *QuerySigningInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuerySigningInfoRequest.Marshal(b, m, deterministic)
}
func (m *QuerySigningInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySigningInfoRequest.Merge(m, src)
}
func (m *QuerySigningInfoRequest) XXX_Size() int {
	return xxx_messageInfo_QuerySigningInfoRequest.Size(m)
}
func (m *QuerySigningInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySigningInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySigningInfoRequest proto.InternalMessageInfo

func (m *QuerySigningInfoRequest) GetConsAddress() string {
	if m != nil {
		return m.ConsAddress
	}
	return ""
}

// QuerySigningInfoResponse is the response type for the Query/SigningInfo RPC
// method
type QuerySigningInfoResponse struct {
	// val_signing_info is the signing info of requested val cons address
	ValSigningInfo       *ValidatorSigningInfo `protobuf:"bytes,1,opt,name=val_signing_info,json=valSigningInfo,proto3" json:"val_signing_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *QuerySigningInfoResponse) Reset()         { *m = QuerySigningInfoResponse{} }
func (m *QuerySigningInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySigningInfoResponse) ProtoMessage()    {}
func (*QuerySigningInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5858f9240a85cd9f, []int{2}
}

func (m *QuerySigningInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuerySigningInfoResponse.Unmarshal(m, b)
}
func (m *QuerySigningInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuerySigningInfoResponse.Marshal(b, m, deterministic)
}
func (m *QuerySigningInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySigningInfoResponse.Merge(m, src)
}
func (m *QuerySigningInfoResponse) XXX_Size() int {
	return xxx_messageInfo_QuerySigningInfoResponse.Size(m)
}
func (m *QuerySigningInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySigningInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySigningInfoResponse proto.InternalMessageInfo

func (m *QuerySigningInfoResponse) GetValSigningInfo() *ValidatorSigningInfo {
	if m != nil {
		return m.ValSigningInfo
	}
	return nil
}

// QuerySigningInfosRequest is the request type for the Query/SigningInfos RPC
// method
type QuerySigningInfosRequest struct {
	Pagination           *base.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	All                  bool              `protobuf:"varint,2,opt,name=all,proto3" json:"all,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *QuerySigningInfosRequest) Reset()         { *m = QuerySigningInfosRequest{} }
func (m *QuerySigningInfosRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySigningInfosRequest) ProtoMessage()    {}
func (*QuerySigningInfosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5858f9240a85cd9f, []int{3}
}

func (m *QuerySigningInfosRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuerySigningInfosRequest.Unmarshal(m, b)
}
func (m *QuerySigningInfosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuerySigningInfosRequest.Marshal(b, m, deterministic)
}
func (m *QuerySigningInfosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySigningInfosRequest.Merge(m, src)
}
func (m *QuerySigningInfosRequest) XXX_Size() int {
	return xxx_messageInfo_QuerySigningInfosRequest.Size(m)
}
func (m *QuerySigningInfosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySigningInfosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySigningInfosRequest proto.InternalMessageInfo

func (m *QuerySigningInfosRequest) GetPagination() *base.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *QuerySigningInfosRequest) GetAll() bool {
	if m != nil {
		return m.All
	}
	return false
}

// QuerySigningInfosResponse is the response type for the Query/SigningInfos RPC
// method
type QuerySigningInfosResponse struct {
	// info is the signing info of all validators
	Info                 []*ValidatorSigningInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
	Pagination           *base.PageResponse      `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *QuerySigningInfosResponse) Reset()         { *m = QuerySigningInfosResponse{} }
func (m *QuerySigningInfosResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySigningInfosResponse) ProtoMessage()    {}
func (*QuerySigningInfosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5858f9240a85cd9f, []int{4}
}

func (m *QuerySigningInfosResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuerySigningInfosResponse.Unmarshal(m, b)
}
func (m *QuerySigningInfosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuerySigningInfosResponse.Marshal(b, m, deterministic)
}
func (m *QuerySigningInfosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySigningInfosResponse.Merge(m, src)
}
func (m *QuerySigningInfosResponse) XXX_Size() int {
	return xxx_messageInfo_QuerySigningInfosResponse.Size(m)
}
func (m *QuerySigningInfosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySigningInfosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySigningInfosResponse proto.InternalMessageInfo

func (m *QuerySigningInfosResponse) GetInfo() []*ValidatorSigningInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *QuerySigningInfosResponse) GetPagination() *base.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*ValidatorSigningInfo)(nil), "kira.slashing.ValidatorSigningInfo")
	proto.RegisterType((*QuerySigningInfoRequest)(nil), "kira.slashing.QuerySigningInfoRequest")
	proto.RegisterType((*QuerySigningInfoResponse)(nil), "kira.slashing.QuerySigningInfoResponse")
	proto.RegisterType((*QuerySigningInfosRequest)(nil), "kira.slashing.QuerySigningInfosRequest")
	proto.RegisterType((*QuerySigningInfosResponse)(nil), "kira.slashing.QuerySigningInfosResponse")
}

func init() {
	proto.RegisterFile("kira/slashing/query.proto", fileDescriptor_5858f9240a85cd9f)
}

var fileDescriptor_5858f9240a85cd9f = []byte{
	// 790 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0x5f, 0x37, 0xcd, 0x6e, 0x77, 0x92, 0x5d, 0x95, 0x69, 0xab, 0x7a, 0xa3, 0x95, 0xe2, 0x1a,
	0x69, 0x37, 0x5a, 0x51, 0x0f, 0x31, 0x12, 0x87, 0x08, 0x90, 0x9a, 0x48, 0xfc, 0x11, 0x42, 0x2a,
	0x2e, 0x70, 0xe0, 0x62, 0x4d, 0xec, 0x89, 0x33, 0xaa, 0x33, 0xe3, 0x7a, 0xc6, 0xa1, 0x15, 0xe2,
	0xc2, 0x27, 0x40, 0xe5, 0xc2, 0xb1, 0x12, 0x9f, 0x80, 0x4f, 0x01, 0x07, 0x6e, 0x5c, 0x38, 0x95,
	0x0b, 0x07, 0xce, 0xfd, 0x04, 0x68, 0x66, 0xec, 0xd4, 0x21, 0x2d, 0xd1, 0x9e, 0xe2, 0x79, 0xbf,
	0xdf, 0x9b, 0xf7, 0x7b, 0xf3, 0x7e, 0x93, 0x01, 0xcf, 0x4e, 0x69, 0x8e, 0x91, 0x48, 0xb1, 0x98,
	0x52, 0x96, 0xa0, 0xb3, 0x82, 0xe4, 0x17, 0x5e, 0x96, 0x73, 0xc9, 0xe1, 0x13, 0x05, 0x79, 0x15,
	0xd4, 0xd9, 0x4d, 0x78, 0xc2, 0x35, 0x82, 0xd4, 0x97, 0x21, 0x75, 0x9e, 0x27, 0x9c, 0x27, 0x29,
	0x41, 0x38, 0xa3, 0x08, 0x33, 0xc6, 0x25, 0x96, 0x94, 0x33, 0x51, 0xa2, 0xdd, 0x12, 0xd5, 0xab,
	0x71, 0x31, 0x41, 0x92, 0xce, 0x88, 0x90, 0x78, 0x96, 0x95, 0x04, 0xf3, 0x13, 0x1d, 0x26, 0x84,
	0x1d, 0xf2, 0x8c, 0x30, 0x9c, 0xd1, 0xb9, 0x8f, 0x78, 0xa6, 0x37, 0xb9, 0x63, 0xc3, 0x57, 0x11,
	0x17, 0x33, 0x2e, 0xd0, 0x18, 0x0b, 0x62, 0xc4, 0xa2, 0x79, 0x7f, 0x4c, 0x24, 0xee, 0xa3, 0x0c,
	0x27, 0x94, 0x69, 0xb2, 0xe1, 0xba, 0xbf, 0x36, 0xc0, 0xee, 0x57, 0x38, 0xa5, 0x31, 0x96, 0x3c,
	0x3f, 0xa1, 0x09, 0xa3, 0x2c, 0xf9, 0x84, 0x4d, 0x38, 0xb4, 0xc1, 0x23, 0x1c, 0xc7, 0x39, 0x11,
	0xc2, 0xb6, 0x1c, 0xab, 0xf7, 0x38, 0xa8, 0x96, 0x70, 0x00, 0xda, 0x42, 0xe2, 0x5c, 0x86, 0x53,
	0x42, 0x93, 0xa9, 0xb4, 0x37, 0x1c, 0xab, 0xd7, 0x18, 0xee, 0xdf, 0x5c, 0x77, 0x77, 0x2e, 0xf0,
	0x2c, 0x1d, 0xb8, 0x75, 0xd4, 0x0d, 0x5a, 0x7a, 0xf9, 0xb1, 0x5e, 0xc1, 0x18, 0x3c, 0xa5, 0x0c,
	0x47, 0x92, 0xce, 0x49, 0x58, 0x30, 0x49, 0x53, 0xbb, 0xe1, 0x58, 0xbd, 0x96, 0xdf, 0xf1, 0xcc,
	0x21, 0x78, 0xd5, 0x21, 0x78, 0x5f, 0x54, 0x87, 0x30, 0x3c, 0xf8, 0xed, 0xba, 0xfb, 0xe0, 0xe6,
	0xba, 0xbb, 0x67, 0x76, 0x5f, 0xce, 0x77, 0x7f, 0xf8, 0xab, 0x6b, 0x05, 0x4f, 0xaa, 0xe0, 0x97,
	0x2a, 0x06, 0xfb, 0x60, 0x77, 0x46, 0x45, 0x34, 0xc5, 0x2c, 0x22, 0x61, 0xc4, 0xd9, 0x84, 0xc6,
	0x84, 0x45, 0xc4, 0xde, 0x54, 0x4a, 0x83, 0x9d, 0x05, 0x36, 0x5a, 0x40, 0xf0, 0x39, 0x78, 0xbc,
	0x08, 0xdb, 0x4d, 0xcd, 0xbb, 0x0d, 0xc0, 0xb7, 0x00, 0x4c, 0xb1, 0x90, 0x61, 0x96, 0x13, 0x41,
	0x98, 0x0c, 0xc7, 0x29, 0x8f, 0x4e, 0xed, 0x87, 0x9a, 0xb6, 0xad, 0x90, 0x63, 0x03, 0x0c, 0x55,
	0x1c, 0xfa, 0x60, 0x6f, 0x46, 0x85, 0x20, 0xb1, 0xe1, 0x89, 0x30, 0xe2, 0x05, 0x93, 0x24, 0xb7,
	0x1f, 0x2d, 0xea, 0x0b, 0x12, 0x6b, 0xae, 0x18, 0x19, 0x08, 0xbe, 0x0b, 0xf6, 0xb3, 0x9c, 0xc7,
	0x45, 0xb4, 0x9a, 0xb5, 0xa5, 0xb3, 0xf6, 0x2a, 0x78, 0x29, 0x6f, 0xb0, 0xf5, 0xd3, 0x55, 0xf7,
	0xc1, 0x3f, 0x57, 0x5d, 0xcb, 0x7d, 0x0f, 0xec, 0x7f, 0xae, 0x66, 0x5d, 0x1b, 0x62, 0x40, 0xce,
	0x0a, 0x22, 0x24, 0x3c, 0x00, 0xed, 0x88, 0x33, 0x11, 0x2e, 0x0f, 0xb4, 0xa5, 0x62, 0x47, 0x26,
	0xe4, 0x72, 0x60, 0xaf, 0x66, 0x8b, 0x8c, 0x33, 0x41, 0xe0, 0x09, 0xd8, 0x9e, 0xe3, 0x34, 0x14,
	0x06, 0x0a, 0x29, 0x9b, 0x70, 0xbd, 0x45, 0xcb, 0x7f, 0xd3, 0x5b, 0xb2, 0xbf, 0x77, 0x97, 0x93,
	0x86, 0x9b, 0x6a, 0x7e, 0xc1, 0xd3, 0x39, 0x4e, 0x6b, 0x51, 0x57, 0xae, 0x16, 0x14, 0x95, 0xde,
	0x0f, 0x01, 0xb8, 0x35, 0x6a, 0x59, 0xea, 0x85, 0x67, 0x5c, 0xed, 0x29, 0x57, 0x7b, 0xe6, 0x0a,
	0x96, 0xae, 0xf6, 0x8e, 0x71, 0x42, 0xca, 0xdc, 0xa0, 0x96, 0x09, 0xb7, 0x41, 0x03, 0xa7, 0xa9,
	0x36, 0xe8, 0x56, 0xa0, 0x3e, 0xdd, 0x9f, 0x2d, 0xf0, 0xec, 0x8e, 0xb2, 0x65, 0xa3, 0xef, 0x83,
	0xcd, 0xb2, 0xb9, 0xc6, 0xeb, 0x35, 0xa7, 0xd3, 0xe0, 0x47, 0x4b, 0xb2, 0x37, 0xb4, 0xec, 0x97,
	0x6b, 0x65, 0x9b, 0xda, 0x75, 0xdd, 0xfe, 0x2f, 0x0d, 0xd0, 0xd4, 0x2a, 0xe1, 0x9f, 0x16, 0x68,
	0xd5, 0x6f, 0xe5, 0x8b, 0xff, 0x68, 0xba, 0x67, 0xe2, 0x9d, 0x97, 0x6b, 0x79, 0xa6, 0xac, 0x7b,
	0x7e, 0x79, 0xf4, 0x19, 0x68, 0x6a, 0x5d, 0xf0, 0x40, 0xb3, 0x9c, 0x92, 0xe6, 0x28, 0x9e, 0xc3,
	0x27, 0x0e, 0x76, 0x16, 0x9d, 0x77, 0xd6, 0x53, 0xbe, 0xff, 0xe3, 0xef, 0x1f, 0x37, 0x5c, 0xe8,
	0xe8, 0x3f, 0xbb, 0x39, 0x4e, 0x79, 0x46, 0x72, 0x75, 0x3a, 0x02, 0x7d, 0x5b, 0xf7, 0xe0, 0x77,
	0xf0, 0x77, 0x0b, 0xb4, 0xeb, 0x53, 0x80, 0xeb, 0x34, 0x57, 0xf6, 0xe8, 0xf4, 0xd6, 0x13, 0xcb,
	0xee, 0xe8, 0xff, 0x77, 0x27, 0x94, 0xf6, 0x85, 0x72, 0xd1, 0x59, 0x4f, 0xd1, 0xdd, 0xed, 0xc0,
	0x37, 0x56, 0xba, 0x1b, 0x16, 0x97, 0x47, 0x1f, 0xc0, 0xa6, 0xdf, 0xe8, 0x7b, 0x6f, 0xbf, 0xb2,
	0xac, 0xdc, 0x07, 0xed, 0x24, 0x38, 0x1e, 0x1d, 0x26, 0x58, 0x92, 0x6f, 0xf0, 0x05, 0x74, 0xa7,
	0x52, 0x66, 0x62, 0x80, 0x50, 0x42, 0xe5, 0xb4, 0x18, 0x7b, 0x11, 0x9f, 0xa1, 0x4f, 0x69, 0x8e,
	0x47, 0x3c, 0x27, 0x88, 0xaa, 0xab, 0x7c, 0xfe, 0xb5, 0x7f, 0x3f, 0x66, 0x1e, 0x07, 0xf5, 0x02,
	0xa0, 0xa5, 0xb7, 0x68, 0xfc, 0x50, 0x03, 0xef, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xae, 0xd2,
	0xad, 0x44, 0xa3, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// SigningInfo queries the signing info of given cons address
	SigningInfo(ctx context.Context, in *QuerySigningInfoRequest, opts ...grpc.CallOption) (*QuerySigningInfoResponse, error)
	// SigningInfos queries signing info of all validators
	SigningInfos(ctx context.Context, in *QuerySigningInfosRequest, opts ...grpc.CallOption) (*QuerySigningInfosResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) SigningInfo(ctx context.Context, in *QuerySigningInfoRequest, opts ...grpc.CallOption) (*QuerySigningInfoResponse, error) {
	out := new(QuerySigningInfoResponse)
	err := c.cc.Invoke(ctx, "/kira.slashing.Query/SigningInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SigningInfos(ctx context.Context, in *QuerySigningInfosRequest, opts ...grpc.CallOption) (*QuerySigningInfosResponse, error) {
	out := new(QuerySigningInfosResponse)
	err := c.cc.Invoke(ctx, "/kira.slashing.Query/SigningInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// SigningInfo queries the signing info of given cons address
	SigningInfo(context.Context, *QuerySigningInfoRequest) (*QuerySigningInfoResponse, error)
	// SigningInfos queries signing info of all validators
	SigningInfos(context.Context, *QuerySigningInfosRequest) (*QuerySigningInfosResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) SigningInfo(ctx context.Context, req *QuerySigningInfoRequest) (*QuerySigningInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SigningInfo not implemented")
}
func (*UnimplementedQueryServer) SigningInfos(ctx context.Context, req *QuerySigningInfosRequest) (*QuerySigningInfosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SigningInfos not implemented")
}

func RegisterQueryServer(s *grpc.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_SigningInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySigningInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SigningInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.slashing.Query/SigningInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SigningInfo(ctx, req.(*QuerySigningInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SigningInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySigningInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SigningInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.slashing.Query/SigningInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SigningInfos(ctx, req.(*QuerySigningInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kira.slashing.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SigningInfo",
			Handler:    _Query_SigningInfo_Handler,
		},
		{
			MethodName: "SigningInfos",
			Handler:    _Query_SigningInfos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kira/slashing/query.proto",
}
