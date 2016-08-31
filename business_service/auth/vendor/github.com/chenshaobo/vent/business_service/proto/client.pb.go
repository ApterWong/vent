// Code generated by protoc-gen-go.
// source: client.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	client.proto
	microRpc.proto

It has these top-level messages:
	RegisterC2S
	RegisterS2C
	RegisterReq
	RegisterRes
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type RegisterC2S struct {
	PhoneNumber uint64 `protobuf:"varint,1,opt,name=phone_number,json=phoneNumber" json:"phone_number,omitempty"`
	Password    string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *RegisterC2S) Reset()                    { *m = RegisterC2S{} }
func (m *RegisterC2S) String() string            { return proto1.CompactTextString(m) }
func (*RegisterC2S) ProtoMessage()               {}
func (*RegisterC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RegisterS2C struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=err_code,json=errCode" json:"err_code,omitempty"`
}

func (m *RegisterS2C) Reset()                    { *m = RegisterS2C{} }
func (m *RegisterS2C) String() string            { return proto1.CompactTextString(m) }
func (*RegisterS2C) ProtoMessage()               {}
func (*RegisterS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto1.RegisterType((*RegisterC2S)(nil), "proto.RegisterC2s")
	proto1.RegisterType((*RegisterS2C)(nil), "proto.RegisterS2c")
}

func init() { proto1.RegisterFile("client.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x3e, 0x5c, 0xdc,
	0x41, 0xa9, 0xe9, 0x99, 0xc5, 0x25, 0xa9, 0x45, 0xce, 0x46, 0xc5, 0x42, 0x8a, 0x5c, 0x3c, 0x05,
	0x19, 0xf9, 0x79, 0xa9, 0xf1, 0x79, 0xa5, 0xb9, 0x49, 0xa9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x2c, 0x41, 0xdc, 0x60, 0x31, 0x3f, 0xb0, 0x90, 0x90, 0x14, 0x17, 0x47, 0x41, 0x62, 0x71, 0x71,
	0x79, 0x7e, 0x51, 0x8a, 0x04, 0x13, 0x50, 0x9a, 0x33, 0x08, 0xce, 0x57, 0xd2, 0x40, 0x98, 0x16,
	0x6c, 0x94, 0x2c, 0x24, 0xc9, 0xc5, 0x91, 0x5a, 0x54, 0x14, 0x9f, 0x9c, 0x9f, 0x92, 0x0a, 0x36,
	0x89, 0x37, 0x88, 0x1d, 0xc8, 0x77, 0x06, 0x72, 0x93, 0xd8, 0xc0, 0xd6, 0x1b, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xcc, 0x39, 0x13, 0xb0, 0x95, 0x00, 0x00, 0x00,
}