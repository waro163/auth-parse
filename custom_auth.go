package authparse

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/waro163/auth-parse/utils"
	gmw "github.com/waro163/gin-middleware"
)

const (
	ACCESS_TOKEN_SUBJECT  = "access_token"
	REFRESH_TOKEN_SUBJECT = "refresh_token"
)

type CustomJwtAuthenticator struct {
	Audience string
	Issuer   string
	Subject  string
	JwtParse utils.IJwtParser
}

func (auth *CustomJwtAuthenticator) Authenticate(token string) (interface{}, *gmw.Error) {
	if token == "" {
		return nil, &gmw.Error{
			Code:    -1,
			Message: "missing authorization token",
		}
	}
	claims := MyCustomClaims{}
	_, err := auth.JwtParse.ParseJwtToken(token, &claims, jwt.WithAudience(auth.Audience), jwt.WithIssuer(auth.Issuer), jwt.WithIssuedAt(), jwt.WithExpirationRequired(), jwt.WithSubject(auth.Subject))
	if err != nil {
		return nil, &gmw.Error{
			Code:    -2,
			Message: err.Error(),
		}
	}
	return claims, nil
}
