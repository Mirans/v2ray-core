package congestion

import "github.com/v2ray/v2ray-core/core/external/github.com/lucas-clemente/quic-go/internal/protocol"

type connectionStats struct {
	slowstartPacketsLost protocol.PacketNumber
	slowstartBytesLost   protocol.ByteCount
}
