package authparse

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/waro163/auth-parse/utils"
	gmw "github.com/waro163/gin-middleware"
)

var (
	_ gmw.IAuthenticator = (*CommonJwtAuthenticator)(nil)
)

type CommonJwtAuthenticator struct {
	Options  []jwt.ParserOption
	JwtParse utils.IJwtParser
}

func (auth *CommonJwtAuthenticator) Authenticate(token string) (interface{}, *gmw.Error) {
	if token == "" {
		return nil, &gmw.Error{
			Code:    -1,
			Message: "missing authorization token",
		}
	}
	claims := jwt.MapClaims{}
	_, err := auth.JwtParse.ParseJwtToken(token, &claims, auth.Options...)
	if err != nil {
		return nil, &gmw.Error{
			Code:    -2,
			Message: err.Error(),
		}
	}
	return claims, nil
}
