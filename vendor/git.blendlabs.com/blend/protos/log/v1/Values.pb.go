// Code generated by protoc-gen-go. DO NOT EDIT.
// source: log/v1/Values.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// [START messages]
type Values struct {
	// body of the message
	Values map[string]string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Values) Reset()                    { *m = Values{} }
func (m *Values) String() string            { return proto.CompactTextString(m) }
func (*Values) ProtoMessage()               {}
func (*Values) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *Values) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*Values)(nil), "logv1.Values")
}

func init() { proto.RegisterFile("log/v1/Values.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0xc9, 0x4f, 0xd7,
	0x2f, 0x33, 0xd4, 0x0f, 0x4b, 0xcc, 0x29, 0x4d, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0xcd, 0xc9, 0x4f, 0x2f, 0x33, 0x54, 0x2a, 0xe3, 0x62, 0x83, 0x08, 0x0b, 0x19, 0x72, 0xb1,
	0x95, 0x81, 0x59, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x92, 0x7a, 0x60, 0x15, 0x7a, 0x50,
	0x5d, 0x10, 0xca, 0x35, 0xaf, 0xa4, 0xa8, 0x32, 0x08, 0xaa, 0x50, 0xca, 0x92, 0x8b, 0x1b, 0x49,
	0x58, 0x48, 0x80, 0x8b, 0x39, 0x3b, 0xb5, 0x52, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc4,
	0x14, 0x12, 0xe1, 0x62, 0x05, 0x2b, 0x95, 0x60, 0x02, 0x8b, 0x41, 0x38, 0x56, 0x4c, 0x16, 0x8c,
	0x4e, 0xea, 0x51, 0xaa, 0xe9, 0x99, 0x25, 0x7a, 0x49, 0x39, 0xa9, 0x79, 0x29, 0x39, 0x89, 0x49,
	0xc5, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x60, 0x9e, 0x3e, 0xd8, 0x7d, 0xc5, 0xfa, 0x10, 0x47, 0x27,
	0xb1, 0x81, 0xb9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xa9, 0xe5, 0xad, 0xc5, 0x00,
	0x00, 0x00,
}