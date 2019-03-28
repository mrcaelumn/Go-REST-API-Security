package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// swagger
var _ = Resource("swagger", func() {
	Security(BasicAuth)
	Origin("*", func() {
		Methods("GET")
	})
	Files("swagger.json", "api/swagger/swagger.json")
})

// version
var VersionType = Type("version", func() {
	Attribute("version", String, "Application version", func() {
		Example("1.0")
	})
	Attribute("git", String, "Git commit hash", func() {
		Example("000000")
	})

	Required("version")
})

var _ = Resource("version", func() {
	Action("version", func() {
		Routing(GET("version"))
		Response(OK, VersionMedia)
		Metadata("swagger:summary", "Return application's version and commit hash")
	})
})

var _ = Resource("Action", func() {

	Security(JWTAuth, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})

	Action("getToken", func() {
		Description("Creates a valid JWT")
		Routing(GET("/getToken"))
		Security(BasicAuth)

		Response(NoContent, func() {
			Headers(func() {
				Header("Authorization", String, "Generated JWT")
			})
		})

		Response(Unauthorized)
		Response(Forbidden)
		Response(BadRequest, CustomeErrorMedia)
		Response(InternalServerError, CustomeErrorMedia)
		Metadata("swagger:summary", "Creates a valid JWT")
	})

	Action("request", func() {
		Routing(GET("/request"))

		Description("This action is secured with the jwt scheme")

		Response(OK, "application/json")
		Response(NoContent)
		Response(Unauthorized)
		Response(Forbidden)
		Response(BadRequest, CustomeErrorMedia)
		Response(InternalServerError, CustomeErrorMedia)
		Metadata("swagger:summary", "Get data")
	})
})
