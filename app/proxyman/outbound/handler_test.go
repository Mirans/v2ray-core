package outbound_test

import (
	"testing"

	. "github.com/v2ray/v2ray-core/core/app/proxyman/outbound"
	"github.com/v2ray/v2ray-core/core/features/outbound"
)

func TestInterfaces(t *testing.T) {
	_ = (outbound.Handler)(new(Handler))
	_ = (outbound.Manager)(new(Manager))
}
