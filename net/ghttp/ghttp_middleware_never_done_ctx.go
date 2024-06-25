package ghttp

// MiddlewareNeverDoneCtx sets the context never done for current process.
func MiddlewareNeverDoneCtx(r *Request) {
	r.SetCtx(r.GetNeverDoneCtx())
	r.Middleware.Next()
}
