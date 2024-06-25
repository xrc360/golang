package gconv_test

import (
	"testing"
	"time"

	"github.com/xrc360/golang/frame/g"
	"github.com/xrc360/golang/test/gtest"
	"github.com/xrc360/golang/util/gconv"
)

type Duration time.Duration

// UnmarshalText unmarshal text to duration.
func (d *Duration) UnmarshalText(text []byte) error {
	tmp, err := time.ParseDuration(string(text))
	if err == nil {
		*d = Duration(tmp)
	}
	return err
}

func Test_Struct_CustomTimeDuration_Attribute(t *testing.T) {
	type A struct {
		Name    string
		Timeout Duration
	}
	gtest.C(t, func(t *gtest.T) {
		var a A
		err := gconv.Struct(g.Map{
			"name":    "john",
			"timeout": "1s",
		}, &a)
		t.AssertNil(err)
	})
}
