package custommiddleware

import (
	"context"
	"net/http"

	"github.com/goadesign/goa"
)

var (
	// ErrUnauthorized is the error returned for unauthorized requests.
	ErrUnauthorized = goa.NewErrorClass("unauthorized", 401)
)

// NewBasicAuthMiddleware creates a middleware that checks for the presence of a basic auth header
// and validates its content.
func NewBasicAuthMiddleware() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log basic auth info
			// fmt.Println(req.BasicAuth())
			user, pass, ok := req.BasicAuth()
			// A real app would do something more interesting here
			if !ok {
				goa.LogInfo(ctx, "failed basic auth")
				return ErrUnauthorized("missing auth")
			}
			if user != "youruser" || pass != "yourpass" {
				goa.LogInfo(ctx, "failed basic auth")
				return ErrUnauthorized("username or password wrong!")
			}
			// Proceed
			goa.LogInfo(ctx, "basic", "user", user, "pass", pass)
			return h(ctx, rw, req)
		}
	}
}
