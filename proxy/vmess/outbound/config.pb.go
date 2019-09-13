package outbound

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protocol "github.com/v2ray/v2ray-core/core/common/protocol"
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

type Config struct {
	Receiver             []*protocol.ServerEndpoint `protobuf:"bytes,1,rep,name=Receiver,proto3" json:"Receiver,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc22c8b653a4f7ef, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetReceiver() []*protocol.ServerEndpoint {
	if m != nil {
		return m.Receiver
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "v2ray.core.proxy.vmess.outbound.Config")
}

func init() {
	proto.RegisterFile("github.com/v2ray/v2ray-core/core/proxy/vmess/outbound/config.proto", fileDescriptor_bc22c8b653a4f7ef)
}

var fileDescriptor_bc22c8b653a4f7ef = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8e, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0x39, 0x85, 0xe3, 0x58, 0xbb, 0xab, 0xc4, 0xe6, 0x44, 0x1b, 0xb1, 0x98, 0x95, 0xf8,
	0x06, 0x06, 0x6d, 0x0d, 0x09, 0xa4, 0xb0, 0x91, 0x64, 0x32, 0x4a, 0xc0, 0xdd, 0x59, 0x66, 0x93,
	0xc5, 0xbc, 0x92, 0x4f, 0x29, 0x4e, 0x5c, 0x11, 0x1b, 0xbb, 0x29, 0xbe, 0xff, 0xfb, 0xc6, 0xdc,
	0xa4, 0x42, 0xba, 0x05, 0x90, 0x9d, 0x45, 0x16, 0xb2, 0x41, 0xf8, 0x7d, 0xb1, 0xc9, 0x51, 0x8c,
	0x96, 0xe7, 0xa9, 0xe7, 0xd9, 0x0f, 0x16, 0xd9, 0xbf, 0x8c, 0xaf, 0x10, 0x84, 0x27, 0xde, 0x1f,
	0xf2, 0x42, 0x08, 0x94, 0x06, 0xa5, 0x21, 0xd3, 0x67, 0x7f, 0x95, 0xc8, 0xce, 0xb1, 0xb7, 0xba,
	0x46, 0x7e, 0xb3, 0x91, 0x24, 0x91, 0x3c, 0xc7, 0x40, 0xb8, 0x2a, 0x2f, 0x2a, 0xb3, 0x2d, 0x35,
	0xb1, 0x7f, 0x30, 0xbb, 0x9a, 0x90, 0xc6, 0x44, 0x72, 0xba, 0x39, 0x3f, 0xbe, 0x3a, 0x29, 0xae,
	0xe1, 0x57, 0x6f, 0x55, 0x41, 0x56, 0x41, 0xa3, 0xaa, 0x7b, 0x3f, 0x04, 0x1e, 0xfd, 0x54, 0xff,
	0x6c, 0xef, 0x1a, 0x73, 0x89, 0xec, 0xe0, 0x9f, 0x57, 0xab, 0xcd, 0xd3, 0x2e, 0xdf, 0x1f, 0x47,
	0x87, 0xb6, 0xa8, 0xbb, 0x05, 0xca, 0x2f, 0xba, 0x52, 0xba, 0x55, 0xfa, 0xf1, 0x9b, 0xe8, 0xb7,
	0xda, 0xbd, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x80, 0xb5, 0x19, 0x90, 0x34, 0x01, 0x00, 0x00,
}
