// +build !confonly

package tcp

import (
	"github.com/v2ray/v2ray-core/core/common"
	"github.com/v2ray/v2ray-core/core/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
