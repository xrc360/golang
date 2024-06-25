package gtime_test

import (
	"testing"

	"github.com/xrc360/golang/frame/g"
	"github.com/xrc360/golang/internal/json"
	"github.com/xrc360/golang/os/gtime"
	"github.com/xrc360/golang/test/gtest"
)

func Test_Json_Pointer(t *testing.T) {
	// Marshal
	gtest.C(t, func(t *gtest.T) {
		type MyTime struct {
			MyTime *gtime.Time
		}
		b, err := json.Marshal(MyTime{
			MyTime: gtime.NewFromStr("2006-01-02 15:04:05"),
		})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":"2006-01-02 15:04:05"}`)
	})
	gtest.C(t, func(t *gtest.T) {
		b, err := json.Marshal(g.Map{
			"MyTime": gtime.NewFromStr("2006-01-02 15:04:05"),
		})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":"2006-01-02 15:04:05"}`)
	})
	gtest.C(t, func(t *gtest.T) {
		b, err := json.Marshal(g.Map{
			"MyTime": *gtime.NewFromStr("2006-01-02 15:04:05"),
		})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":"2006-01-02 15:04:05"}`)
	})
	// Marshal nil
	gtest.C(t, func(t *gtest.T) {
		type MyTime struct {
			MyTime *gtime.Time
		}
		b, err := json.Marshal(&MyTime{})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":null}`)
	})
	// Marshal nil with json omitempty
	gtest.C(t, func(t *gtest.T) {
		type MyTime struct {
			MyTime *gtime.Time `json:"time,omitempty"`
		}
		b, err := json.Marshal(&MyTime{})
		t.AssertNil(err)
		t.Assert(b, `{}`)
	})
	// Unmarshal
	gtest.C(t, func(t *gtest.T) {
		var (
			myTime gtime.Time
			err    = json.UnmarshalUseNumber([]byte(`"2006-01-02 15:04:05"`), &myTime)
		)
		t.AssertNil(err)
		t.Assert(myTime.String(), "2006-01-02 15:04:05")
	})
}

func Test_Json_Struct(t *testing.T) {
	// Marshal struct.
	gtest.C(t, func(t *gtest.T) {
		type MyTime struct {
			MyTime gtime.Time
		}
		b, err := json.Marshal(MyTime{
			MyTime: *gtime.NewFromStr("2006-01-02 15:04:05"),
		})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":"2006-01-02 15:04:05"}`)
	})
	// Marshal pointer.
	gtest.C(t, func(t *gtest.T) {
		type MyTime struct {
			MyTime gtime.Time
		}
		b, err := json.Marshal(&MyTime{
			MyTime: *gtime.NewFromStr("2006-01-02 15:04:05"),
		})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":"2006-01-02 15:04:05"}`)
	})
	// Marshal nil
	gtest.C(t, func(t *gtest.T) {
		type MyTime struct {
			MyTime gtime.Time
		}
		b, err := json.Marshal(MyTime{})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":""}`)
	})
	// Marshal nil omitempty
	gtest.C(t, func(t *gtest.T) {
		type MyTime struct {
			MyTime gtime.Time `json:"time,omitempty"`
		}
		b, err := json.Marshal(MyTime{})
		t.AssertNil(err)
		t.Assert(b, `{"time":""}`)
	})

}
