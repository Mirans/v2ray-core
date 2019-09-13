package log

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

type Severity int32

const (
	Severity_Unknown Severity = 0
	Severity_Error   Severity = 1
	Severity_Warning Severity = 2
	Severity_Info    Severity = 3
	Severity_Debug   Severity = 4
)

var Severity_name = map[int32]string{
	0: "Unknown",
	1: "Error",
	2: "Warning",
	3: "Info",
	4: "Debug",
}

var Severity_value = map[string]int32{
	"Unknown": 0,
	"Error":   1,
	"Warning": 2,
	"Info":    3,
	"Debug":   4,
}

func (x Severity) String() string {
	return proto.EnumName(Severity_name, int32(x))
}

func (Severity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_903ab33e40dced38, []int{0}
}

func init() {
	proto.RegisterEnum("v2ray.core.common.log.Severity", Severity_name, Severity_value)
}

func init() {
	proto.RegisterFile("github.com/v2ray/v2ray-core/core/common/log/log.proto", fileDescriptor_903ab33e40dced38)
}

var fileDescriptor_903ab33e40dced38 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0xce, 0xcf, 0xcd, 0xcd,
	0xcf, 0xd3, 0xcf, 0xc9, 0x4f, 0x07, 0x61, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x51, 0x98,
	0xa2, 0xa2, 0x54, 0x3d, 0x88, 0x02, 0xbd, 0x9c, 0xfc, 0x74, 0x2d, 0x17, 0x2e, 0x8e, 0xe0, 0xd4,
	0xb2, 0xd4, 0xa2, 0xcc, 0x92, 0x4a, 0x21, 0x6e, 0x2e, 0xf6, 0xd0, 0xbc, 0xec, 0xbc, 0xfc, 0xf2,
	0x3c, 0x01, 0x06, 0x21, 0x4e, 0x2e, 0x56, 0xd7, 0xa2, 0xa2, 0xfc, 0x22, 0x01, 0x46, 0x90, 0x78,
	0x78, 0x62, 0x51, 0x5e, 0x66, 0x5e, 0xba, 0x00, 0x93, 0x10, 0x07, 0x17, 0x8b, 0x67, 0x5e, 0x5a,
	0xbe, 0x00, 0x33, 0x48, 0x85, 0x4b, 0x6a, 0x52, 0x69, 0xba, 0x00, 0x8b, 0x93, 0x15, 0x97, 0x64,
	0x72, 0x7e, 0xae, 0x1e, 0x56, 0x2b, 0x02, 0x18, 0xa3, 0x98, 0x73, 0xf2, 0xd3, 0x57, 0x31, 0x89,
	0x86, 0x19, 0x05, 0x25, 0x56, 0xea, 0x39, 0x83, 0xa4, 0x9d, 0x21, 0xd2, 0x3e, 0xf9, 0xe9, 0x49,
	0x6c, 0x60, 0xf7, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x69, 0xef, 0x0e, 0xc6, 0x00,
	0x00, 0x00,
}
