package main

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	parse "github.com/waro163/auth-parse"
	"github.com/waro163/auth-parse/utils"

	"github.com/gin-gonic/gin"
	gmw "github.com/waro163/gin-middleware"
)

func main() {
	r := gin.New()
	writer := os.Stdout
	log := &gmw.RequestLog{Output: writer}

	r.Use(gmw.AddRequestID(), log.AddRequestLog(), gmw.Recovery(writer))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	auth := gmw.AuthMiddleware{
		GetToken: &gmw.JWTHeaderAuthToken,
		Authenticator: &parse.M2mJwtAuthenticator{
			Audience: "test:resource",
			Issuer:   "auth-server",
			Scope:    "test:read",
			JwtParse: utils.NewJwtParse("http://localhost:8080/api/v1/jwks/get", nil, nil),
		},
	}
	r.GET("/auth", auth.ValidateAuth(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "auth",
		})
	})

	commonAuth := gmw.AuthMiddleware{
		GetToken: &gmw.JWTHeaderAuthToken,
		Authenticator: &parse.CommonJwtAuthenticator{
			Options: []jwt.ParserOption{
				jwt.WithIssuedAt(),
				jwt.WithExpirationRequired(),
			},
			JwtParse: utils.NewJwtParse("http://localhost:8080/api/v1/jwks/get", nil, nil),
		},
	}
	r.GET("/common_auth", commonAuth.ValidateAuth(), func(c *gin.Context) {
		claims, ok := c.Get(gmw.AuthPayload)
		if !ok {
			c.JSON(401, gin.H{
				"message": "not found",
			})
			return
		}
		myClaims, ok := claims.(jwt.MapClaims)
		if !ok {
			c.JSON(500, gin.H{
				"message": "some thing is wrong",
			})
			return
		}
		uid := myClaims["uid"]
		c.JSON(200, gin.H{
			"message": "common auth",
			"uid":     uid,
		})
	})

	r.Run(":8081")
}
