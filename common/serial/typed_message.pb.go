package serial

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

// TypedMessage is a serialized proto message along with its type name.
type TypedMessage struct {
	// The name of the message type, retrieved from protobuf API.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Serialized proto message.
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypedMessage) Reset()         { *m = TypedMessage{} }
func (m *TypedMessage) String() string { return proto.CompactTextString(m) }
func (*TypedMessage) ProtoMessage()    {}
func (*TypedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb3cdb51e9fc84d, []int{0}
}

func (m *TypedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypedMessage.Unmarshal(m, b)
}
func (m *TypedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypedMessage.Marshal(b, m, deterministic)
}
func (m *TypedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypedMessage.Merge(m, src)
}
func (m *TypedMessage) XXX_Size() int {
	return xxx_messageInfo_TypedMessage.Size(m)
}
func (m *TypedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TypedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TypedMessage proto.InternalMessageInfo

func (m *TypedMessage) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TypedMessage) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*TypedMessage)(nil), "v2ray.core.common.serial.TypedMessage")
}

func init() {
	proto.RegisterFile("github.com/v2ray/v2ray-core/common/serial/typed_message.proto", fileDescriptor_0bb3cdb51e9fc84d)
}

var fileDescriptor_0bb3cdb51e9fc84d = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x28, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0xce, 0xcf, 0xcd, 0xcd,
	0xcf, 0xd3, 0x2f, 0x4e, 0x2d, 0xca, 0x4c, 0xcc, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x4d, 0x89, 0xcf,
	0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x80, 0xe9,
	0x28, 0x4a, 0xd5, 0x83, 0xa8, 0xd6, 0x83, 0xa8, 0x56, 0xb2, 0xe0, 0xe2, 0x09, 0x01, 0x69, 0xf0,
	0x85, 0xa8, 0x17, 0x12, 0xe2, 0x62, 0x01, 0x19, 0x20, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04,
	0x66, 0x0b, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0xf0,
	0x04, 0x41, 0x38, 0x4e, 0xce, 0x5c, 0x32, 0xc9, 0xf9, 0xb9, 0x7a, 0xb8, 0x4c, 0x0e, 0x60, 0x8c,
	0x62, 0x83, 0xb0, 0x56, 0x31, 0x49, 0x84, 0x19, 0x05, 0x25, 0x56, 0xea, 0x39, 0x83, 0x14, 0x39,
	0x43, 0x14, 0x05, 0x83, 0xa5, 0x92, 0xd8, 0xc0, 0xee, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x2f, 0x28, 0xe9, 0x5b, 0xd3, 0x00, 0x00, 0x00,
}
