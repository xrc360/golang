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

func TestRequest_GetRemoteIp(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Server(guid.S())
		s.BindHandler("/", func(r *ghttp.Request) {
			r.Response.Write(r.GetRemoteIp())
		})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		time.Sleep(100 * time.Millisecond)

		clientV4 := g.Client()
		clientV4.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		clientV6 := g.Client()
		clientV6.SetPrefix(fmt.Sprintf("http://[::1]:%d", s.GetListenedPort()))

		t.Assert(clientV4.GetContent(ctx, "/"), "127.0.0.1")
		t.Assert(clientV6.GetContent(ctx, "/"), "::1")
	})
}
