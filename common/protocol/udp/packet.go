package udp

import (
	"github.com/v2ray/v2ray-core/common/buf"
	"github.com/v2ray/v2ray-core/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
