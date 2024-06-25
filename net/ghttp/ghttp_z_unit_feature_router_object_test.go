package ghttp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/xrc360/golang/frame/g"
	"github.com/xrc360/golang/net/ghttp"
	"github.com/xrc360/golang/test/gtest"
	"github.com/xrc360/golang/util/guid"
)

type Object struct{}

func (o *Object) Init(r *ghttp.Request) {
	r.Response.Write("1")
}

func (o *Object) Shut(r *ghttp.Request) {
	r.Response.Write("2")
}

func (o *Object) Index(r *ghttp.Request) {
	r.Response.Write("Object Index")
}

func (o *Object) Show(r *ghttp.Request) {
	r.Response.Write("Object Show")
}

func (o *Object) Info(r *ghttp.Request) {
	r.Response.Write("Object Info")
}

func Test_Router_Object1(t *testing.T) {
	s := g.Server(guid.S())
	s.BindObject("/", new(Object))
	s.BindObject("/{.struct}/{.method}", new(Object))
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/"), "1Object Index2")
		t.Assert(client.GetContent(ctx, "/init"), "Not Found")
		t.Assert(client.GetContent(ctx, "/shut"), "Not Found")
		t.Assert(client.GetContent(ctx, "/index"), "1Object Index2")
		t.Assert(client.GetContent(ctx, "/show"), "1Object Show2")

		t.Assert(client.GetContent(ctx, "/object"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/init"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/shut"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/index"), "1Object Index2")
		t.Assert(client.GetContent(ctx, "/object/show"), "1Object Show2")

		t.Assert(client.GetContent(ctx, "/none-exist"), "Not Found")
	})
}

func Test_Router_Object2(t *testing.T) {
	s := g.Server(guid.S())
	s.BindObject("/object", new(Object), "Show, Info")
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/init"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/shut"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/index"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/show"), "1Object Show2")
		t.Assert(client.GetContent(ctx, "/object/info"), "1Object Info2")

		t.Assert(client.GetContent(ctx, "/none-exist"), "Not Found")
	})
}

func Test_Router_ObjectMethod(t *testing.T) {
	s := g.Server(guid.S())
	s.BindObjectMethod("/object-info", new(Object), "Info")
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/init"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/shut"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/index"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/show"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object/info"), "Not Found")
		t.Assert(client.GetContent(ctx, "/object-info"), "1Object Info2")

		t.Assert(client.GetContent(ctx, "/none-exist"), "Not Found")
	})
}
