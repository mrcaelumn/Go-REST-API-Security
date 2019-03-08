package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// swagger
var _ = Resource("swagger", func() {
	NoSecurity()
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
	Action("getToken", func() {
		Routing(GET("/getToken"))
		Response(OK, "application/json")
		Response(NoContent)
		Response(Unauthorized)
		Response(Forbidden)
		Response(BadRequest, CustomeErrorMedia)
		Response(InternalServerError, CustomeErrorMedia)
		Metadata("swagger:summary", "Get data")
	})
	Action("request", func() {
		Routing(GET("/request"))
		Response(OK, "application/json")
		Response(NoContent)
		Response(Unauthorized)
		Response(Forbidden)
		Response(BadRequest, CustomeErrorMedia)
		Response(InternalServerError, CustomeErrorMedia)
		Metadata("swagger:summary", "Get data")
	})
})
