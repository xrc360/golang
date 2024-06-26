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

func Test_Router_Group_Hook1(t *testing.T) {
	s := g.Server(guid.S())
	group := s.Group("/api")
	group.GET("/handler", func(r *ghttp.Request) {
		r.Response.Write("1")
	})
	group.ALL("/handler", func(r *ghttp.Request) {
		r.Response.Write("0")
	}, ghttp.HookBeforeServe)
	group.ALL("/handler", func(r *ghttp.Request) {
		r.Response.Write("2")
	}, ghttp.HookAfterServe)

	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		t.Assert(client.GetContent(ctx, "/api/handler"), "012")
		t.Assert(client.PostContent(ctx, "/api/handler"), "02")
		t.Assert(client.GetContent(ctx, "/api/ThisDoesNotExist"), "Not Found")
	})
}

func Test_Router_Group_Hook2(t *testing.T) {
	s := g.Server(guid.S())
	group := s.Group("/api")
	group.GET("/handler", func(r *ghttp.Request) {
		r.Response.Write("1")
	})
	group.GET("/*", func(r *ghttp.Request) {
		r.Response.Write("0")
	}, ghttp.HookBeforeServe)
	group.GET("/*", func(r *ghttp.Request) {
		r.Response.Write("2")
	}, ghttp.HookAfterServe)

	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		t.Assert(client.GetContent(ctx, "/api/handler"), "012")
		t.Assert(client.PostContent(ctx, "/api/handler"), "Not Found")
		t.Assert(client.GetContent(ctx, "/api/ThisDoesNotExist"), "02")
		t.Assert(client.PostContent(ctx, "/api/ThisDoesNotExist"), "Not Found")
	})
}
