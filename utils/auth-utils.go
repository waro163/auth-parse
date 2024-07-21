package utils

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type IJwtParser interface {
	ParseJwtToken(string, jwt.Claims, ...jwt.ParserOption) (*jwt.Token, error)
}

type JwtParse struct {
	jwksUrl string
	client  IReqeustClient
	cache   ICache
}

func NewJwtParse(jwksUrl string, client IReqeustClient, cache ICache) IJwtParser {
	parse := &JwtParse{
		jwksUrl: jwksUrl,
		client:  client,
		cache:   cache,
	}
	if client == nil {
		parse.client = DefaultClient
	}
	if cache == nil {
		parse.cache = DefaultMemoryCache
	}
	return parse
}

func (p *JwtParse) ParseJwtToken(token string, claims jwt.Claims, options ...jwt.ParserOption) (*jwt.Token, error) {
	parseToken, err := jwt.ParseWithClaims(token, claims, p.getPublicSecret, options...)
	return parseToken, err
}

func (p *JwtParse) getPublicSecret(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("unexpected signed method: %s", token.Header["alg"])
	}
	kid, ok := token.Header["kid"].(string)
	if !ok || kid == "" {
		return nil, fmt.Errorf("invalid kid %s", kid)
	}
	publicKey, err := p.cache.Get(kid)
	if err == nil && publicKey != nil {
		return publicKey, nil
	}
	req, err := http.NewRequest(http.MethodGet, p.jwksUrl, nil)
	if err != nil {
		return nil, err
	}
	response, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var pubCerts []PublicKey
	if err := json.NewDecoder(response.Body).Decode(&pubCerts); err != nil {
		return nil, err
	}
	var currentPubCert string
	for _, pubCert := range pubCerts {
		if pubCert.Kid == kid {
			currentPubCert = pubCert.Public
			break
		}
	}
	if currentPubCert == "" {
		return nil, fmt.Errorf("not found kid %s in jwks", kid)
	}
	rsaPublic, err := base64.StdEncoding.DecodeString(currentPubCert)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(rsaPublic)
	rsaPublicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	p.cache.Set(kid, rsaPublicKey, 0)
	return rsaPublicKey, nil
}