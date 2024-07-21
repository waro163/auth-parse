package authparse

import (
	"auth-parse/utils"

	"github.com/golang-jwt/jwt/v5"
	ginmiddleware "github.com/waro163/gin-middleware"
)

type M2mJwtAuthenticator struct {
	Audience string
	Issuer   string
	Scope    string
	JwtParse utils.IJwtParser
}

func (auth *M2mJwtAuthenticator) Authenticate(token string) (interface{}, *ginmiddleware.Error) {
	if token == "" {
		return nil, &ginmiddleware.Error{
			Code:    -1,
			Message: "missing authorization token",
		}
	}
	parseToken, err := auth.JwtParse.ParseJwtToken(token, &M2mClaims{}, jwt.WithIssuer(auth.Issuer), jwt.WithIssuedAt(), jwt.WithAudience(auth.Audience), jwt.WithExpirationRequired())
	if err != nil {
		return nil, &ginmiddleware.Error{
			Code:    -2,
			Message: err.Error(),
		}
	}
	claims, ok := parseToken.Claims.(*M2mClaims)
	if !ok {
		return nil, &ginmiddleware.Error{
			Code:    -3,
			Message: "invalid claims",
		}
	}
	if auth.Scope != "" && claims.Scp != auth.Scope {
		return nil, &ginmiddleware.Error{
			Code:    -4,
			Message: "invalid scp",
		}
	}
	return claims, nil
}
