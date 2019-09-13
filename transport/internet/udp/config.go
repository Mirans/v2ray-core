package udp

import (
	"github.com/v2ray/v2ray-core/core/common"
	"github.com/v2ray/v2ray-core/core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
