package json

//go:generate errorgen

import (
	"io"

	"github.com/v2ray/v2ray-core"
	"github.com/v2ray/v2ray-core/common"
	"github.com/v2ray/v2ray-core/common/buf"
	"github.com/v2ray/v2ray-core/common/platform/ctlcmd"
)

func init() {
	common.Must(core.RegisterConfigLoader(&core.ConfigFormat{
		Name:      "JSON",
		Extension: []string{"json"},
		Loader: func(input io.Reader) (*core.Config, error) {
			jsonContent, err := ctlcmd.Run([]string{"config"}, input)
			if err != nil {
				return nil, newError("failed to execute v2ctl to convert config file.").Base(err).AtWarning()
			}
			return core.LoadConfig("protobuf", "", &buf.MultiBufferContainer{
				MultiBuffer: jsonContent,
			})
		},
	}))
}
