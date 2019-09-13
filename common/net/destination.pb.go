package net

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

// Endpoint of a network connection.
type Endpoint struct {
	Network              Network     `protobuf:"varint,1,opt,name=network,proto3,enum=v2ray.core.common.net.Network" json:"network,omitempty"`
	Address              *IPOrDomain `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Port                 uint32      `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_77acfe1424029862, []int{0}
}

func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetNetwork() Network {
	if m != nil {
		return m.Network
	}
	return Network_Unknown
}

func (m *Endpoint) GetAddress() *IPOrDomain {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Endpoint) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "v2ray.core.common.net.Endpoint")
}

func init() {
	proto.RegisterFile("github.com/v2ray/v2ray-core/common/net/destination.proto", fileDescriptor_77acfe1424029862)
}

var fileDescriptor_77acfe1424029862 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2e, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0xce, 0xcf, 0xcd, 0xcd,
	0xcf, 0xd3, 0xcf, 0x4b, 0x2d, 0xd1, 0x4f, 0x49, 0x2d, 0x2e, 0xc9, 0xcc, 0x4b, 0x2c, 0xc9, 0xcc,
	0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x85, 0x29, 0x2e, 0x4a, 0xd5, 0x83, 0x28,
	0xd4, 0xcb, 0x4b, 0x2d, 0x91, 0x52, 0xc7, 0x6d, 0x46, 0x5e, 0x6a, 0x49, 0x79, 0x7e, 0x51, 0x36,
	0x44, 0x3f, 0x3e, 0x85, 0x89, 0x29, 0x29, 0x45, 0xa9, 0xc5, 0xc5, 0x10, 0x85, 0x4a, 0x53, 0x19,
	0xb9, 0x38, 0x5c, 0xf3, 0x52, 0x0a, 0xf2, 0x33, 0xf3, 0x4a, 0x84, 0x2c, 0xb8, 0xd8, 0xa1, 0xc6,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19, 0xc9, 0xe9, 0x61, 0x75, 0x87, 0x9e, 0x1f, 0x44, 0x55,
	0x10, 0x4c, 0xb9, 0x90, 0x35, 0x17, 0x3b, 0xd4, 0x5c, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23,
	0x45, 0x1c, 0x3a, 0x3d, 0x03, 0xfc, 0x8b, 0x5c, 0xf2, 0x73, 0x13, 0x33, 0xf3, 0x82, 0x60, 0x3a,
	0x84, 0x84, 0xb8, 0x58, 0x0a, 0xf2, 0x8b, 0x4a, 0x24, 0x98, 0x15, 0x18, 0x35, 0x78, 0x83, 0xc0,
	0x6c, 0x27, 0x2b, 0x2e, 0xc9, 0xe4, 0xfc, 0x5c, 0xec, 0x86, 0x04, 0x30, 0x46, 0x31, 0xe7, 0xa5,
	0x96, 0xac, 0x62, 0x12, 0x0d, 0x33, 0x0a, 0x4a, 0xac, 0xd4, 0x73, 0x06, 0x49, 0x3b, 0x43, 0xa4,
	0xfd, 0x52, 0x4b, 0x92, 0xd8, 0xc0, 0x5e, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x75, 0xe4,
	0xce, 0x58, 0x72, 0x01, 0x00, 0x00,
}
