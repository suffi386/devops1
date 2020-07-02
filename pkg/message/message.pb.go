// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package message

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

type ErrorDetail struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorDetail) Reset()         { *m = ErrorDetail{} }
func (m *ErrorDetail) String() string { return proto.CompactTextString(m) }
func (*ErrorDetail) ProtoMessage()    {}
func (*ErrorDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *ErrorDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorDetail.Unmarshal(m, b)
}
func (m *ErrorDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorDetail.Marshal(b, m, deterministic)
}
func (m *ErrorDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorDetail.Merge(m, src)
}
func (m *ErrorDetail) XXX_Size() int {
	return xxx_messageInfo_ErrorDetail.Size(m)
}
func (m *ErrorDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorDetail.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorDetail proto.InternalMessageInfo

func (m *ErrorDetail) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ErrorDetail) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type LocalizedMessage struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalizedMessage) Reset()         { *m = LocalizedMessage{} }
func (m *LocalizedMessage) String() string { return proto.CompactTextString(m) }
func (*LocalizedMessage) ProtoMessage()    {}
func (*LocalizedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *LocalizedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalizedMessage.Unmarshal(m, b)
}
func (m *LocalizedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalizedMessage.Marshal(b, m, deterministic)
}
func (m *LocalizedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalizedMessage.Merge(m, src)
}
func (m *LocalizedMessage) XXX_Size() int {
	return xxx_messageInfo_LocalizedMessage.Size(m)
}
func (m *LocalizedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalizedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LocalizedMessage proto.InternalMessageInfo

func (m *LocalizedMessage) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LocalizedMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ErrorDetail)(nil), "caos.zitadel.api.v1.ErrorDetail")
	proto.RegisterType((*LocalizedMessage)(nil), "caos.zitadel.api.v1.LocalizedMessage")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4e, 0x4e, 0xcc, 0x2f, 0xd6,
	0xab, 0xca, 0x2c, 0x49, 0x4c, 0x49, 0xcd, 0xd1, 0x4b, 0x2c, 0xc8, 0xd4, 0x2b, 0x33, 0x54, 0x32,
	0xe7, 0xe2, 0x76, 0x2d, 0x2a, 0xca, 0x2f, 0x72, 0x49, 0x2d, 0x49, 0xcc, 0xcc, 0x11, 0xe2, 0xe3,
	0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x92, 0xe0,
	0x62, 0x87, 0x1a, 0x22, 0xc1, 0x04, 0x16, 0x84, 0x71, 0x95, 0xec, 0xb8, 0x04, 0x7c, 0xf2, 0x93,
	0x13, 0x73, 0x32, 0xab, 0x52, 0x53, 0x7c, 0x21, 0x62, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95,
	0x50, 0xed, 0x20, 0x26, 0x6e, 0xfd, 0x4e, 0xaa, 0x51, 0xca, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49,
	0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x20, 0xa7, 0xe9, 0x43, 0x9d, 0xa6, 0x5f, 0x90, 0x9d, 0xae, 0x0f,
	0x55, 0x96, 0xc4, 0x06, 0x76, 0xbb, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x74, 0x88, 0xd8,
	0xcc, 0x00, 0x00, 0x00,
}
