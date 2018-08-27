// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto2test.proto

/*
Package packet is a generated protocol buffer package.

It is generated from these files:
	proto2test.proto

It has these top-level messages:
	Header
	ByteMessage
	StringMessage
*/
package protobuf

import proto "ProtobufTest/pf2test/github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Header struct {
	MessageId        *string `protobuf:"bytes,1,req,name=messageId" json:"messageId,omitempty"`
	Topic            *string `protobuf:"bytes,2,req,name=topic" json:"topic,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Header) GetMessageId() string {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return ""
}

func (m *Header) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

type ByteMessage struct {
	Header           *Header `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Body             []byte  `protobuf:"bytes,2,req,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ByteMessage) Reset()                    { *m = ByteMessage{} }
func (m *ByteMessage) String() string            { return proto.CompactTextString(m) }
func (*ByteMessage) ProtoMessage()               {}
func (*ByteMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ByteMessage) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ByteMessage) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type StringMessage struct {
	Header           *Header `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Body             *string `protobuf:"bytes,2,req,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StringMessage) Reset()                    { *m = StringMessage{} }
func (m *StringMessage) String() string            { return proto.CompactTextString(m) }
func (*StringMessage) ProtoMessage()               {}
func (*StringMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StringMessage) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *StringMessage) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Header)(nil), "packet.Header")
	proto.RegisterType((*ByteMessage)(nil), "packet.ByteMessage")
	proto.RegisterType((*StringMessage)(nil), "packet.StringMessage")
}

func init() { proto.RegisterFile("proto2test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0x37, 0x2a, 0x49, 0x2d, 0x2e, 0xd1, 0x03, 0x33, 0x85, 0xd8, 0x0a, 0x12, 0x93, 0xb3, 0x53,
	0x4b, 0x94, 0x6c, 0xb8, 0xd8, 0x3c, 0x52, 0x13, 0x53, 0x52, 0x8b, 0x84, 0x64, 0xb8, 0x38, 0x73,
	0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x3d, 0x53, 0x24, 0x18, 0x15, 0x98, 0x34, 0x38, 0x83, 0x10,
	0x02, 0x42, 0x22, 0x5c, 0xac, 0x25, 0xf9, 0x05, 0x99, 0xc9, 0x12, 0x4c, 0x60, 0x19, 0x08, 0x47,
	0xc9, 0x93, 0x8b, 0xdb, 0xa9, 0xb2, 0x24, 0xd5, 0x17, 0xa2, 0x4c, 0x48, 0x8d, 0x8b, 0x2d, 0x03,
	0x6c, 0x18, 0x58, 0x3f, 0xb7, 0x11, 0x9f, 0x1e, 0xc4, 0x16, 0x3d, 0x88, 0x15, 0x41, 0x50, 0x59,
	0x21, 0x21, 0x2e, 0x96, 0xa4, 0xfc, 0x94, 0x4a, 0xb0, 0x59, 0x3c, 0x41, 0x60, 0xb6, 0x92, 0x37,
	0x17, 0x6f, 0x70, 0x49, 0x51, 0x66, 0x5e, 0x3a, 0x25, 0x86, 0x71, 0x42, 0x0c, 0x03, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x43, 0x6b, 0xc5, 0x6a, 0xf0, 0x00, 0x00, 0x00,
}
