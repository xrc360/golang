package ghttp

import "net/http"

// WrapF is a helper function for wrapping http.HandlerFunc and returns a ghttp.HandlerFunc.
func WrapF(f http.HandlerFunc) HandlerFunc {
	return func(r *Request) {
		f(r.Response.Writer, r.Request)
	}
}

// WrapH is a helper function for wrapping http.Handler and returns a ghttp.HandlerFunc.
func WrapH(h http.Handler) HandlerFunc {
	return func(r *Request) {
		h.ServeHTTP(r.Response.Writer, r.Request)
	}
}
