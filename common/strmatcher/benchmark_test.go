package strmatcher_test

import (
	"strconv"
	"testing"

	"github.com/v2ray/v2ray-core/common"
	. "github.com/v2ray/v2ray-core/common/strmatcher"
)

func BenchmarkDomainMatcherGroup(b *testing.B) {
	g := new(DomainMatcherGroup)

	for i := 1; i <= 1024; i++ {
		g.Add(strconv.Itoa(i)+".github.com/v2ray/v2ray-core", uint32(i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = g.Match("0.github.com/v2ray/v2ray-core")
	}
}

func BenchmarkFullMatcherGroup(b *testing.B) {
	g := new(FullMatcherGroup)

	for i := 1; i <= 1024; i++ {
		g.Add(strconv.Itoa(i)+".github.com/v2ray/v2ray-core", uint32(i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = g.Match("0.github.com/v2ray/v2ray-core")
	}
}

func BenchmarkMarchGroup(b *testing.B) {
	g := new(MatcherGroup)
	for i := 1; i <= 1024; i++ {
		m, err := Domain.New(strconv.Itoa(i) + ".github.com/v2ray/v2ray-core")
		common.Must(err)
		g.Add(m)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = g.Match("0.github.com/v2ray/v2ray-core")
	}
}
