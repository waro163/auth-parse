package authparse

import "github.com/golang-jwt/jwt/v5"

type M2mClaims struct {
	jwt.RegisteredClaims
	Scp string `json:"scp,omitempty"`
}
