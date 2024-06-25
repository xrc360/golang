package gins_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/xrc360/golang/frame/gins"
	"github.com/xrc360/golang/internal/instance"
	"github.com/xrc360/golang/net/ghttp"
	"github.com/xrc360/golang/os/gcfg"
	"github.com/xrc360/golang/os/gctx"
	"github.com/xrc360/golang/os/gfile"
	"github.com/xrc360/golang/test/gtest"
)

func Test_Server(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			path                = gcfg.DefaultConfigFileName
			serverConfigContent = gtest.DataContent("server", "config.yaml")
			err                 = gfile.PutContents(path, serverConfigContent)
		)
		t.AssertNil(err)
		defer gfile.Remove(path)

		instance.Clear()
		defer instance.Clear()

		s := gins.Server("tempByInstanceName")
		s.BindHandler("/", func(r *ghttp.Request) {
			r.Response.Write("hello")
		})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		time.Sleep(100 * time.Millisecond)

		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
		client := gins.HttpClient()
		client.SetPrefix(prefix)
		t.Assert(client.GetContent(gctx.New(), "/"), "hello")
	})
}
