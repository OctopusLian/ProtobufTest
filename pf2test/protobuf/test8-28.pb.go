// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test8-28.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	test8-28.proto

It has these top-level messages:
	Test
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

type FOO int32

const (
	FOO_X FOO = 17
)

var FOO_name = map[int32]string{
	17: "X",
}
var FOO_value = map[string]int32{
	"X": 17,
}

func (x FOO) Enum() *FOO {
	p := new(FOO)
	*p = x
	return p
}
func (x FOO) String() string {
	return proto.EnumName(FOO_name, int32(x))
}
func (x *FOO) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FOO_value, data, "FOO")
	if err != nil {
		return err
	}
	*x = FOO(value)
	return nil
}
func (FOO) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type Test struct {
	Label            *string `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	Type             *int32  `protobuf:"varint,2,opt,name=type,def=77" json:"type,omitempty"`
	Reps             []int64 `protobuf:"varint,3,rep,name=reps" json:"reps,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Test) Reset()                    { *m = Test{} }
func (m *Test) String() string            { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()               {}
func (*Test) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

const Default_Test_Type int32 = 77

func (m *Test) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Test) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_Test_Type
}

func (m *Test) GetReps() []int64 {
	if m != nil {
		return m.Reps
	}
	return nil
}

func init() {
	proto.RegisterType((*Test)(nil), "protobuf.Test")
	proto.RegisterEnum("protobuf.FOO", FOO_name, FOO_value)
}

func init() { proto.RegisterFile("test8-28.proto", fileDescriptor0) }

var fileDescriptor1 = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x49, 0x2d, 0x2e,
	0xb1, 0xd0, 0x35, 0xb2, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5,
	0x69, 0x4a, 0x1e, 0x5c, 0x2c, 0x21, 0xa9, 0xc5, 0x25, 0x42, 0x22, 0x5c, 0xac, 0x39, 0x89, 0x49,
	0xa9, 0x39, 0x12, 0x8c, 0x0a, 0x4c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x18, 0x17, 0x4b, 0x49,
	0x65, 0x41, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0xab, 0x15, 0x93, 0xb9, 0x79, 0x10, 0x98, 0x2f,
	0x24, 0xc4, 0xc5, 0x52, 0x94, 0x5a, 0x50, 0x2c, 0xc1, 0xac, 0xc0, 0xac, 0xc1, 0x1c, 0x04, 0x66,
	0x6b, 0xf1, 0x70, 0x31, 0xbb, 0xf9, 0xfb, 0x0b, 0xb1, 0x72, 0x31, 0x46, 0x08, 0x08, 0x02, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xa3, 0x90, 0xe6, 0xb3, 0x72, 0x00, 0x00, 0x00,
}
