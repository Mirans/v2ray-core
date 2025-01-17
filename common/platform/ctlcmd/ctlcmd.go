package ctlcmd

import (
	"io"
	"os"
	"os/exec"

	"github.com/v2ray/v2ray-core/common/buf"
	"github.com/v2ray/v2ray-core/common/platform"
)

//go:generate errorgen

func Run(args []string, input io.Reader) (buf.MultiBuffer, error) {
	v2ctl := platform.GetToolLocation("v2ctl")
	if _, err := os.Stat(v2ctl); err != nil {
		return nil, newError("v2ctl doesn't exist").Base(err)
	}

	var errBuffer buf.MultiBufferContainer
	var outBuffer buf.MultiBufferContainer

	cmd := exec.Command(v2ctl, args...)
	cmd.Stderr = &errBuffer
	cmd.Stdout = &outBuffer
	cmd.SysProcAttr = getSysProcAttr()
	if input != nil {
		cmd.Stdin = input
	}

	if err := cmd.Start(); err != nil {
		return nil, newError("failed to start v2ctl").Base(err)
	}

	if err := cmd.Wait(); err != nil {
		msg := "failed to execute v2ctl"
		if errBuffer.Len() > 0 {
			msg += ": " + errBuffer.MultiBuffer.String()
		}
		return nil, newError(msg).Base(err)
	}

	return outBuffer.MultiBuffer, nil
}
