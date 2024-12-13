package authparse

import (
	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	Uid string `json:"uid"`
	jwt.RegisteredClaims
}
