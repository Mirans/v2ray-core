// +build !confonly

package quic

import (
	"github.com/v2ray/v2ray-core/core/common"
	"github.com/v2ray/v2ray-core/core/transport/internet"
)

//go:generate errorgen

// Here is some modification needs to be done before update quic vendor.
// * use bytespool in buffer_pool.go
// * set MaxReceivePacketSize to 1452 - 32 (16 bytes auth, 16 bytes head)
//
//

const protocolName = "quic"
const internalDomain = "quic.internal.github.com/v2ray/v2ray-core"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
