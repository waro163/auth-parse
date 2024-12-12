package authparse

import (
	"github.com/waro163/auth-parse/utils"

	"github.com/golang-jwt/jwt/v5"
	gmw "github.com/waro163/gin-middleware"
)

var (
	_ gmw.IAuthenticator = (*M2mJwtAuthenticator)(nil)
)

type M2mJwtAuthenticator struct {
	Audience string
	Issuer   string
	Scope    string
	JwtParse utils.IJwtParser
}

func (auth *M2mJwtAuthenticator) Authenticate(token string) (interface{}, *gmw.Error) {
	if token == "" {
		return nil, &gmw.Error{
			Code:    -1,
			Message: "missing authorization token",
		}
	}
	parseToken, err := auth.JwtParse.ParseJwtToken(token, &M2mClaims{}, jwt.WithIssuer(auth.Issuer), jwt.WithIssuedAt(), jwt.WithAudience(auth.Audience), jwt.WithExpirationRequired())
	if err != nil {
		return nil, &gmw.Error{
			Code:    -2,
			Message: err.Error(),
		}
	}
	claims, ok := parseToken.Claims.(*M2mClaims)
	if !ok {
		return nil, &gmw.Error{
			Code:    -3,
			Message: "invalid claims",
		}
	}
	if auth.Scope != "" && claims.Scp != auth.Scope {
		return nil, &gmw.Error{
			Code:    -4,
			Message: "invalid scp",
		}
	}
	return claims, nil
}
