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

type NamesObject struct{}

func (o *NamesObject) ShowName(r *ghttp.Request) {
	r.Response.Write("Object Show Name")
}

func Test_NameToUri_FullName(t *testing.T) {
	s := g.Server(guid.S())
	s.SetNameToUriType(ghttp.UriTypeFullName)
	s.BindObject("/{.struct}/{.method}", new(NamesObject))
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetBrowserMode(true)
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/NamesObject"), "Not Found")
		t.Assert(client.GetContent(ctx, "/NamesObject/ShowName"), "Object Show Name")
	})
}

func Test_NameToUri_AllLower(t *testing.T) {
	s := g.Server(guid.S())
	s.SetNameToUriType(ghttp.UriTypeAllLower)
	s.BindObject("/{.struct}/{.method}", new(NamesObject))
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetBrowserMode(true)
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/NamesObject"), "Not Found")
		t.Assert(client.GetContent(ctx, "/namesobject/showname"), "Object Show Name")
	})
}

func Test_NameToUri_Camel(t *testing.T) {
	s := g.Server(guid.S())
	s.SetNameToUriType(ghttp.UriTypeCamel)
	s.BindObject("/{.struct}/{.method}", new(NamesObject))
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetBrowserMode(true)
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/NamesObject"), "Not Found")
		t.Assert(client.GetContent(ctx, "/namesObject/showName"), "Object Show Name")
	})
}

func Test_NameToUri_Default(t *testing.T) {
	s := g.Server(guid.S())
	s.SetNameToUriType(ghttp.UriTypeDefault)
	s.BindObject("/{.struct}/{.method}", new(NamesObject))
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetBrowserMode(true)
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/NamesObject"), "Not Found")
		t.Assert(client.GetContent(ctx, "/names-object/show-name"), "Object Show Name")
	})
}
