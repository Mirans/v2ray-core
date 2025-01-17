package udp

import (
	"github.com/v2ray/v2ray-core/common"
	"github.com/v2ray/v2ray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
