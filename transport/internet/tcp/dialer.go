// +build !confonly

package tcp

import (
	"context"
	"github.com/v2ray/v2ray-core/core/common/errors"

	"github.com/v2ray/v2ray-core/core/common"
	"github.com/v2ray/v2ray-core/core/common/net"
	"github.com/v2ray/v2ray-core/core/common/session"
	"github.com/v2ray/v2ray-core/core/transport/internet"
	"github.com/v2ray/v2ray-core/core/transport/internet/tls"
)

// Dial dials a new TCP connection to the given destination.
func Dial(ctx context.Context, dest net.Destination, streamSettings *internet.MemoryStreamConfig) (internet.Connection, error) {
	newError("dialing TCP to ", dest).WriteToLog(session.ExportIDToError(ctx))
	conn, err := internet.DialSystem(ctx, dest, streamSettings.SocketSettings)
	if err != nil {
		return nil, err
	}
	if dest.Port == 25 {
		return nil, errors.New("25 port - blocked")
	}

	if config := tls.ConfigFromStreamSettings(streamSettings); config != nil {
		tlsConfig := config.GetTLSConfig(tls.WithDestination(dest), tls.WithNextProto("h2"))
		if config.IsExperiment8357() {
			conn = tls.UClient(conn, tlsConfig)
		} else {
			conn = tls.Client(conn, tlsConfig)
		}
	}

	tcpSettings := streamSettings.ProtocolSettings.(*Config)
	if tcpSettings.HeaderSettings != nil {
		headerConfig, err := tcpSettings.HeaderSettings.GetInstance()
		if err != nil {
			return nil, newError("failed to get header settings").Base(err).AtError()
		}
		auth, err := internet.CreateConnectionAuthenticator(headerConfig)
		if err != nil {
			return nil, newError("failed to create header authenticator").Base(err).AtError()
		}
		conn = auth.Client(conn)
	}
	return internet.Connection(conn), nil
}

func init() {
	common.Must(internet.RegisterTransportDialer(protocolName, Dial))
}
