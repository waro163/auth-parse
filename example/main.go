package main

import (
	"os"

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
			"message": "pong",
		})
	})

	r.Run(":8081")
}
