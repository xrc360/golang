// static service testing.

package ghttp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/xrcn/cg/frame/g"
	"github.com/xrcn/cg/net/ghttp"
	"github.com/xrcn/cg/test/gtest"
	"github.com/xrcn/cg/util/guid"
)

func Test_StatusHandler(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Server(guid.S())
		s.BindStatusHandlerByMap(map[int]ghttp.HandlerFunc{
			404: func(r *ghttp.Request) { r.Response.WriteOver("404") },
			502: func(r *ghttp.Request) { r.Response.WriteOver("502") },
		})
		s.BindHandler("/502", func(r *ghttp.Request) {
			r.Response.WriteStatusExit(502)
		})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/404"), "404")
		t.Assert(client.GetContent(ctx, "/502"), "502")
	})
}

func Test_StatusHandler_Multi(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Server(guid.S())
		s.BindStatusHandler(502, func(r *ghttp.Request) {
			r.Response.WriteOver("1")
		})
		s.BindStatusHandler(502, func(r *ghttp.Request) {
			r.Response.Write("2")
		})
		s.BindHandler("/502", func(r *ghttp.Request) {
			r.Response.WriteStatusExit(502)
		})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/502"), "12")
	})
}
