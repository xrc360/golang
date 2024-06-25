// static service testing.

package ghttp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/xrc360/golang/errors/gcode"
	"github.com/xrc360/golang/errors/gerror"
	"github.com/xrc360/golang/frame/g"
	"github.com/xrc360/golang/net/ghttp"
	"github.com/xrc360/golang/test/gtest"
	"github.com/xrc360/golang/util/guid"
)

func Test_Error_Code(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Server(guid.S())
		s.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(func(r *ghttp.Request) {
				r.Middleware.Next()
				r.Response.ClearBuffer()
				r.Response.Write(gerror.Code(r.GetError()))
			})
			group.ALL("/", func(r *ghttp.Request) {
				panic(gerror.NewCode(gcode.New(10000, "", nil), "test error"))
			})
		})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(c.GetContent(ctx, "/"), "10000")
	})
}
