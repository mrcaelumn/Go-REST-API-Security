package custommiddleware

import (
	"context"
	"crypto/rsa"
	"io/ioutil"
	"net/http"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/gofrs/uuid"
	"github.com/mrcaelumn/Go-REST-API-Security/api/app"
)

// NewJWTMiddleware creates a middleware that checks for the presence of a JWT Authorization header
// and validates its content. A real app would probably use goa's JWT security middleware instead.
//
// Note: the code below assumes the example is compiled against the master branch of goa.
// If compiling against goa v1 the call to jwt.New needs to be:
//
//    middleware := jwt.New(keys, ForceFail(), app.NewJWTSecurity())
func NewJWTMiddleware() (goa.Middleware, error) {
	keys, err := LoadJWTPublicKeys()
	if err != nil {
		return nil, err
	}
	return jwt.New(keys, ForceFail(), app.NewJWTSecurity()), nil
}

// LoadJWTPublicKeys loads PEM encoded RSA public keys used to validata and decrypt the JWT.
func LoadJWTPublicKeys() (*rsa.PublicKey, error) {
	keyFiles, err := ioutil.ReadFile("./file/app.key.pub")
	if err != nil {
		return nil, err
	}

	key, err := jwtgo.ParseRSAPublicKeyFromPEM(keyFiles)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// ForceFail is a middleware illustrating the use of validation middleware with JWT auth.  It checks
// for the presence of a "fail" query string and fails validation if set to the value "true".
func ForceFail() goa.Middleware {
	errValidationFailed := goa.NewErrorClass("validation_failed", 401)
	forceFail := func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			if f, ok := req.URL.Query()["fail"]; ok {
				if f[0] == "true" {
					return errValidationFailed("forcing failure to illustrate Validation middleware")
				}
			}
			return h(ctx, rw, req)
		}
	}
	fm, _ := goa.NewMiddleware(forceFail)
	return fm
}

// Generate Json Web Token
func GenerateJWT(user string) *jwtgo.Token {

	token := jwtgo.New(jwtgo.SigningMethodRS512)
	in10m := time.Now().Add(time.Duration(10) * time.Minute).Unix()
	token.Claims = jwtgo.MapClaims{
		"iss":    "Issuer",                         // who creates the token and signs it
		"aud":    "Audience",                       // to whom the token is intended to be sent
		"exp":    in10m,                            // time when the token will expire (10 minutes from now)
		"jti":    uuid.Must(uuid.NewV4()).String(), // a unique identifier for the token
		"iat":    time.Now().Unix(),                // when the token was issued/created (now)
		"nbf":    2,                                // time before which the token is not yet valid (2 minutes ago)
		"sub":    "subject",                        // the subject/principal is whom the token is about
		"scopes": "api:access",                     // token scope - not a standard claim
		"user":   user,                             // username
	}
	return token
}
