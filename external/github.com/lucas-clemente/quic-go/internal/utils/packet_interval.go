package utils

import "github.com/v2ray/v2ray-core/core/external/github.com/lucas-clemente/quic-go/internal/protocol"

// PacketInterval is an interval from one PacketNumber to the other
type PacketInterval struct {
	Start protocol.PacketNumber
	End   protocol.PacketNumber
}
