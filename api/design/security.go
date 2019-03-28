package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

// BasicAuth defines a security scheme using basic authentication.
var BasicAuth = BasicAuthSecurity("basic_auth")

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint. Supports scopes "api:read" and "api:write".`)
	Scope("api:read", "Read-only access")
	Scope("api:write", "Read and write access")
})
