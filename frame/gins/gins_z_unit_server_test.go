package gins_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/xrcn/cg/frame/gins"
	"github.com/xrcn/cg/internal/instance"
	"github.com/xrcn/cg/net/ghttp"
	"github.com/xrcn/cg/os/gcfg"
	"github.com/xrcn/cg/os/gctx"
	"github.com/xrcn/cg/os/gfile"
	"github.com/xrcn/cg/test/gtest"
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
