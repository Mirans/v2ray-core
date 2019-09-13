package strmatcher_test

import (
	"testing"

	"github.com/v2ray/v2ray-core/common"
	. "github.com/v2ray/v2ray-core/common/strmatcher"
)

func TestMatcher(t *testing.T) {
	cases := []struct {
		pattern string
		mType   Type
		input   string
		output  bool
	}{
		{
			pattern: "github.com/v2ray/v2ray-core",
			mType:   Domain,
			input:   "www.github.com/v2ray/v2ray-core",
			output:  true,
		},
		{
			pattern: "github.com/v2ray/v2ray-core",
			mType:   Domain,
			input:   "github.com/v2ray/v2ray-core",
			output:  true,
		},
		{
			pattern: "github.com/v2ray/v2ray-core",
			mType:   Domain,
			input:   "www.v3ray.com",
			output:  false,
		},
		{
			pattern: "github.com/v2ray/v2ray-core",
			mType:   Domain,
			input:   "2ray.com",
			output:  false,
		},
		{
			pattern: "github.com/v2ray/v2ray-core",
			mType:   Domain,
			input:   "xgithub.com/v2ray/v2ray-core",
			output:  false,
		},
		{
			pattern: "github.com/v2ray/v2ray-core",
			mType:   Full,
			input:   "github.com/v2ray/v2ray-core",
			output:  true,
		},
		{
			pattern: "github.com/v2ray/v2ray-core",
			mType:   Full,
			input:   "xgithub.com/v2ray/v2ray-core",
			output:  false,
		},
		{
			pattern: "github.com/v2ray/v2ray-core",
			mType:   Regex,
			input:   "v2rayxcom",
			output:  true,
		},
	}
	for _, test := range cases {
		matcher, err := test.mType.New(test.pattern)
		common.Must(err)
		if m := matcher.Match(test.input); m != test.output {
			t.Error("unexpected output: ", m, " for test case ", test)
		}
	}
}
