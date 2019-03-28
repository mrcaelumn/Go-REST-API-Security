package controller

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/mrcaelumn/go-rest-api-security/api/app"
	"github.com/mrcaelumn/go-rest-api-security/api/custommiddleware"
)

// ActionController implements the Action resource.
type ActionController struct {
	*goa.Controller
	privateKey *rsa.PrivateKey
}

// NewActionController creates a Action controller.
func NewActionController(service *goa.Service) (*ActionController, error) {
	b, err := ioutil.ReadFile("./file/app.key")
	if err != nil {
		return nil, err
	}
	privKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("jwt: failed to load private key: %s", err) // bug
	}
	return &ActionController{
		Controller: service.NewController("ActionController"),
		privateKey: privKey,
	}, nil
}

// GetToken runs the getToken action.
func (c *ActionController) GetToken(ctx *app.GetTokenActionContext) error {
	// ActionController_GetToken: start_implement

	// Put your logic here
	user, _, ok := ctx.Request.BasicAuth()
	if !ok {
		goa.LogInfo(ctx, "failed basic auth")
		return ctx.BadRequest(&app.GorestsecurityError{Msg: "failed basic auth", Code: "001"})
	}
	// fmt.Println("User: ", user)
	token := custommiddleware.GenerateJWT(user)
	signedToken, err := token.SignedString(c.privateKey)
	if err != nil {
		return fmt.Errorf("failed to sign token: %s", err) // internal error
	}
	// Set auth header for client retrieval
	ctx.ResponseData.Header().Set("Authorization", "Bearer "+signedToken)

	// Send response
	return ctx.NoContent()
	// ActionController_GetToken: end_implement
}

// Request runs the request action.
func (c *ActionController) Request(ctx *app.RequestActionContext) error {
	// ActionController_Request: start_implement

	// Put your logic here

	// Retrieve the token claims
	token := jwt.ContextJWT(ctx)
	if token == nil {
		return fmt.Errorf("JWT token is missing from context") // internal error
	}
	claims := token.Claims.(jwtgo.MapClaims)
	fmt.Println(claims)
	// Use the claims to authorize
	subject := claims["user"]
	if subject != "subject" {
		// A real app would probably use an "Unauthorized" response here

		return ctx.OK([]byte("OK Unauthorized"))
	}

	return ctx.OK([]byte("OK"))
	// ActionController_Request: end_implement
}
