package ghttp

// MiddlewareCORS is a middleware handler for CORS with default options.
func MiddlewareCORS(r *Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
