package g_test

import (
	"github.com/xrc360/golang/frame/g"
	"github.com/xrc360/golang/net/ghttp"
)

func ExampleServer() {
	// A hello world example.
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("hello world")
	})
	s.SetPort(8999)
	s.Run()
}
