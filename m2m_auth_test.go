package authparse

import (
	"testing"

	"github.com/waro163/auth-parse/utils"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticate(t *testing.T) {
	auth := M2mJwtAuthenticator{
		Audience: "test:resource",
		Issuer:   "auth-server",
		Scope:    "test:read",
		JwtParse: utils.NewJwtParse("http://localhost:8080/api/v1/jwks/get", nil, nil),
	}
	tokenStr := "eyJhbGciOiJSUzI1NiIsImtpZCI6ImM0MWFkZTBlLTUyZTgtNGIwMC1hNDAzLTY0MmVjMzY0M2JiYiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhdXRoLXNlcnZlciIsInN1YiI6InRlc3RfY2xpZW50X2lkIiwiYXVkIjpbInRlc3Q6cmVzb3VyY2UiXSwiZXhwIjoxNzIxNTcyMjkxLCJuYmYiOjE3MjE1Njg2OTEsImlhdCI6MTcyMTU2ODY5MSwic2NwIjoidGVzdDpyZWFkIn0.DSRFr4F3euS_ompq-bMDaKQ5t-Efk1Jbbh0IDHSKNq0j-WfSr4-FNinbc5XhGangI-fQkJWA-w-YUwuuIyeNSBujant1zxrPWd3aCn-v_q7fqAeCuAyxwfdfAA0W28mEBzqOmIbd6mby9emh1XlIv26VbYUcpgRXEf2yxNa6T5CUaGSQD_6nQfZSYPI7zFk7_oqH_hOsfxTqWRi0SnVyEaZqajvMz39YnLCzBi9RawAAcpL0PHNNWuDXMD-_29nOAafQPnhFZXXz-ENvaolcCn9AGiwJPZqQQAKhKw3dCpe_BxxbSyDBK0pTiKqiifs-tQhp_nvA3dbcFscIvAWamA"
	claims, err := auth.Authenticate(tokenStr)
	assert.Nil(t, err)
	t.Logf("claims: %#v", claims)
}
