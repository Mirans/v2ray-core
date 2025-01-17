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

// Address of a network host. It may be either an IP address or a domain address.
type IPOrDomain struct {
	// Types that are valid to be assigned to Address:
	//	*IPOrDomain_Ip
	//	*IPOrDomain_Domain
	Address              isIPOrDomain_Address `protobuf_oneof:"address"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *IPOrDomain) Reset()         { *m = IPOrDomain{} }
func (m *IPOrDomain) String() string { return proto.CompactTextString(m) }
func (*IPOrDomain) ProtoMessage()    {}
func (*IPOrDomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ba9b4a4304e7c1f, []int{0}
}

func (m *IPOrDomain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPOrDomain.Unmarshal(m, b)
}
func (m *IPOrDomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPOrDomain.Marshal(b, m, deterministic)
}
func (m *IPOrDomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPOrDomain.Merge(m, src)
}
func (m *IPOrDomain) XXX_Size() int {
	return xxx_messageInfo_IPOrDomain.Size(m)
}
func (m *IPOrDomain) XXX_DiscardUnknown() {
	xxx_messageInfo_IPOrDomain.DiscardUnknown(m)
}

var xxx_messageInfo_IPOrDomain proto.InternalMessageInfo

type isIPOrDomain_Address interface {
	isIPOrDomain_Address()
}

type IPOrDomain_Ip struct {
	Ip []byte `protobuf:"bytes,1,opt,name=ip,proto3,oneof"`
}

type IPOrDomain_Domain struct {
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3,oneof"`
}

func (*IPOrDomain_Ip) isIPOrDomain_Address() {}

func (*IPOrDomain_Domain) isIPOrDomain_Address() {}

func (m *IPOrDomain) GetAddress() isIPOrDomain_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *IPOrDomain) GetIp() []byte {
	if x, ok := m.GetAddress().(*IPOrDomain_Ip); ok {
		return x.Ip
	}
	return nil
}

func (m *IPOrDomain) GetDomain() string {
	if x, ok := m.GetAddress().(*IPOrDomain_Domain); ok {
		return x.Domain
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*IPOrDomain) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*IPOrDomain_Ip)(nil),
		(*IPOrDomain_Domain)(nil),
	}
}

func init() {
	proto.RegisterType((*IPOrDomain)(nil), "v2ray.core.common.net.IPOrDomain")
}

func init() {
	proto.RegisterFile("github.com/v2ray/v2ray-core/common/net/address.proto", fileDescriptor_4ba9b4a4304e7c1f)
}

var fileDescriptor_4ba9b4a4304e7c1f = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0xce, 0xcf, 0xcd, 0xcd,
	0xcf, 0xd3, 0xcf, 0x4b, 0x2d, 0xd1, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0x85, 0x29, 0x2c, 0x4a, 0xd5, 0x83, 0x28, 0xd2, 0xcb, 0x4b, 0x2d,
	0x51, 0x72, 0xe6, 0xe2, 0xf2, 0x0c, 0xf0, 0x2f, 0x72, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0x13, 0x12,
	0xe0, 0x62, 0xca, 0x2c, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0xf1, 0x60, 0x08, 0x62, 0xca, 0x2c,
	0x10, 0x92, 0xe0, 0x62, 0x4b, 0x01, 0xcb, 0x49, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x7a, 0x30, 0x04,
	0x41, 0xf9, 0x4e, 0x9c, 0x5c, 0xec, 0x50, 0x1b, 0x9c, 0xac, 0xb8, 0x24, 0x93, 0xf3, 0x73, 0xf5,
	0xb0, 0xda, 0x10, 0xc0, 0x18, 0xc5, 0x9c, 0x97, 0x5a, 0xb2, 0x8a, 0x49, 0x34, 0xcc, 0x28, 0x28,
	0xb1, 0x52, 0xcf, 0x19, 0x24, 0xed, 0x0c, 0x91, 0xf6, 0x4b, 0x2d, 0x49, 0x62, 0x03, 0x3b, 0xcf,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xb4, 0xfa, 0x61, 0xc9, 0x00, 0x00, 0x00,
}
