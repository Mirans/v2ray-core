package command_test

import (
	"context"
	"testing"

	"github.com/v2ray/v2ray-core"
	"github.com/v2ray/v2ray-core/app/dispatcher"
	"github.com/v2ray/v2ray-core/app/log"
	. "github.com/v2ray/v2ray-core/app/log/command"
	"github.com/v2ray/v2ray-core/app/proxyman"
	_ "github.com/v2ray/v2ray-core/app/proxyman/inbound"
	_ "github.com/v2ray/v2ray-core/app/proxyman/outbound"
	"github.com/v2ray/v2ray-core/common"
	"github.com/v2ray/v2ray-core/common/serial"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
