package ghttp

import (
	"github.com/xrcn/cg/internal/json"
)

// MiddlewareJsonBody validates and returns request body whether JSON format.
func MiddlewareJsonBody(r *Request) {
	requestBody := r.GetBody()
	if len(requestBody) > 0 {
		if !json.Valid(requestBody) {
			r.SetError(ErrNeedJsonBody)
			return
		}
	}
	r.Middleware.Next()
}
