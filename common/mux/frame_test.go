package mux_test

import (
	"testing"

	"github.com/v2ray/v2ray-core/core/common"
	"github.com/v2ray/v2ray-core/core/common/buf"
	"github.com/v2ray/v2ray-core/core/common/mux"
	"github.com/v2ray/v2ray-core/core/common/net"
)

func BenchmarkFrameWrite(b *testing.B) {
	frame := mux.FrameMetadata{
		Target:        net.TCPDestination(net.DomainAddress("www.github.com/v2ray/v2ray-core"), net.Port(80)),
		SessionID:     1,
		SessionStatus: mux.SessionStatusNew,
	}
	writer := buf.New()
	defer writer.Release()

	for i := 0; i < b.N; i++ {
		common.Must(frame.WriteTo(writer))
		writer.Clear()
	}
}
