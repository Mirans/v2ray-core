package strmatcher_test

import (
	"testing"

	. "github.com/v2ray/v2ray-core/core/common/strmatcher"
)

func TestFullMatcherGroup(t *testing.T) {
	g := new(FullMatcherGroup)
	g.Add("github.com/v2ray/v2ray-core", 1)
	g.Add("google.com", 2)
	g.Add("x.a.com", 3)

	testCases := []struct {
		Domain string
		Result uint32
	}{
		{
			Domain: "github.com/v2ray/v2ray-core",
			Result: 1,
		},
		{
			Domain: "y.com",
			Result: 0,
		},
	}

	for _, testCase := range testCases {
		r := g.Match(testCase.Domain)
		if r != testCase.Result {
			t.Error("Failed to match domain: ", testCase.Domain, ", expect ", testCase.Result, ", but got ", r)
		}
	}
}

func TestEmptyFullMatcherGroup(t *testing.T) {
	g := new(FullMatcherGroup)
	r := g.Match("github.com/v2ray/v2ray-core")
	if r != 0 {
		t.Error("Expect 0, but ", r)
	}
}
