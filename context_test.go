package core_test

import (
	"context"
	"testing"

	. "github.com/v2ray/v2ray-core"
)

func TestContextPanic(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("expect panic, but nil")
		}
	}()

	MustFromContext(context.Background())
}
