// Code generated by protoc-gen-go. DO NOT EDIT.
// source: arith.proto

package arith

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ArithRequest struct {
	A                    *int32   `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B                    *int32   `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArithRequest) Reset()         { *m = ArithRequest{} }
func (m *ArithRequest) String() string { return proto.CompactTextString(m) }
func (*ArithRequest) ProtoMessage()    {}
func (*ArithRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_211289c5d02710c5, []int{0}
}

func (m *ArithRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArithRequest.Unmarshal(m, b)
}
func (m *ArithRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArithRequest.Marshal(b, m, deterministic)
}
func (m *ArithRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArithRequest.Merge(m, src)
}
func (m *ArithRequest) XXX_Size() int {
	return xxx_messageInfo_ArithRequest.Size(m)
}
func (m *ArithRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ArithRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ArithRequest proto.InternalMessageInfo

func (m *ArithRequest) GetA() int32 {
	if m != nil && m.A != nil {
		return *m.A
	}
	return 0
}

func (m *ArithRequest) GetB() int32 {
	if m != nil && m.B != nil {
		return *m.B
	}
	return 0
}

type ArithResponse struct {
	Val                  *int32   `protobuf:"varint,1,opt,name=val" json:"val,omitempty"`
	Quo                  *int32   `protobuf:"varint,2,opt,name=quo" json:"quo,omitempty"`
	Rem                  *int32   `protobuf:"varint,3,opt,name=rem" json:"rem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArithResponse) Reset()         { *m = ArithResponse{} }
func (m *ArithResponse) String() string { return proto.CompactTextString(m) }
func (*ArithResponse) ProtoMessage()    {}
func (*ArithResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_211289c5d02710c5, []int{1}
}

func (m *ArithResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArithResponse.Unmarshal(m, b)
}
func (m *ArithResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArithResponse.Marshal(b, m, deterministic)
}
func (m *ArithResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArithResponse.Merge(m, src)
}
func (m *ArithResponse) XXX_Size() int {
	return xxx_messageInfo_ArithResponse.Size(m)
}
func (m *ArithResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArithResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArithResponse proto.InternalMessageInfo

func (m *ArithResponse) GetVal() int32 {
	if m != nil && m.Val != nil {
		return *m.Val
	}
	return 0
}

func (m *ArithResponse) GetQuo() int32 {
	if m != nil && m.Quo != nil {
		return *m.Quo
	}
	return 0
}

func (m *ArithResponse) GetRem() int32 {
	if m != nil && m.Rem != nil {
		return *m.Rem
	}
	return 0
}

func init() {
	proto.RegisterType((*ArithRequest)(nil), "arith.ArithRequest")
	proto.RegisterType((*ArithResponse)(nil), "arith.ArithResponse")
}

func init() {
	proto.RegisterFile("arith.proto", fileDescriptor_211289c5d02710c5)
}

var fileDescriptor_211289c5d02710c5 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2c, 0xca, 0x2c,
	0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xb4, 0xb8, 0x78, 0x1c,
	0x41, 0x8c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x1e, 0x2e, 0xc6, 0x44, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xd6, 0x20, 0xc6, 0x44, 0x10, 0x2f, 0x49, 0x82, 0x09, 0xc2, 0x4b, 0x52, 0x72,
	0xe5, 0xe2, 0x85, 0xaa, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe0, 0x62, 0x2e, 0x4b,
	0xcc, 0x81, 0x2a, 0x07, 0x31, 0x41, 0x22, 0x85, 0xa5, 0xf9, 0x50, 0x2d, 0x20, 0x26, 0x48, 0xa4,
	0x28, 0x35, 0x57, 0x82, 0x19, 0x22, 0x52, 0x94, 0x9a, 0x6b, 0x54, 0x05, 0xb5, 0x32, 0x38, 0xb5,
	0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8, 0x94, 0x8b, 0x23, 0xb7, 0x34, 0xa7, 0x24, 0xb3, 0x20, 0xa7,
	0x52, 0x48, 0x58, 0x0f, 0xe2, 0x46, 0x64, 0x37, 0x49, 0x89, 0xa0, 0x0a, 0x42, 0x2d, 0x37, 0xe6,
	0x62, 0x4b, 0xc9, 0x2c, 0xcb, 0x4c, 0x49, 0x25, 0x41, 0x93, 0x13, 0x73, 0x03, 0x23, 0x23, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x7b, 0xaf, 0xde, 0x98, 0x08, 0x01, 0x00, 0x00,
}
