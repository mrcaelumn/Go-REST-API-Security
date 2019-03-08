package design

import (
	"net/http"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("go-rest-security", func() {
	Version("1.0")
	Title("go-rest-security APIs")
	Description("API for go-rest-security")
	Scheme("https")
	BasePath("/gorestsecurity")
	Consumes("application/json")
	Produces("application/json")

	ResponseTemplate(BadRequest, func() {
		Status(http.StatusBadRequest)
		Media(CustomeErrorMedia)
		Description("BadRequest is returned if input object is missing " +
			"required attributes or their values are out of range.")
	})

	ResponseTemplate(Unauthorized, func() {
		Status(http.StatusUnauthorized)
		Media(CustomeErrorMedia)
		Description("Unauthorized is returned when user request does not " +
			"contain authentication token or authentication is invalid. " +
			"The response must include a valid \"WWW-Authenticate\" header.")
		Headers(func() {
			Header("WWW-Authenticate", func() {
				Description(`https://tools.ietf.org/html/rfc7235`)
				Default("Bearer")
			})
		})
	})

	ResponseTemplate(Forbidden, func() {
		Status(http.StatusForbidden)
		Media(CustomeErrorMedia)
		Description("Forbidden is returned when user is not authorized " +
			"to perform an action.")
	})
})
