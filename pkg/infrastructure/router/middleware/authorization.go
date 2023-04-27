package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/i-akbarshoh/task-manager/pkg/config"
	"log"
	"net/http"
	"strings"
)

func Authorizer() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		enforcer, err := casbin.NewEnforcer(config.C.Casbin.Model, config.C.Casbin.Policy)
		if err != nil {
			log.Fatal("enforcer not initialized, ", err)
			return
		}

		claims, err := extractClaims(accessToken, []byte(config.C.JWT.SigningKey))
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		path := c.Request.URL.Path
		method := c.Request.Method
		role := claims["role"]
		fmt.Println(role, path, method)
		ok, err := enforcer.Enforce(role, path, method)
		if err != nil {
			log.Println("could not enforce:", err)
			c.Abort()
			return
		}

		if !ok {
			c.JSON(http.StatusForbidden, map[string]string{
				"error": "user not allowed, there is a problem with authorization",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func extractClaims(tokenString string, signingKey []byte) (jwtgo.MapClaims, error) {
	claims := jwtgo.MapClaims{}
	if tokenString == "" {
		claims["role"] = "unauthorized"
		return claims, nil
	}
	if strings.Contains(tokenString, "Basic") {
		claims["role"] = "unauthorized"
		return claims, nil
	}
	token, err := jwtgo.ParseWithClaims(tokenString, claims, func(token *jwtgo.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwtgo.MapClaims)
	if !(ok && token.Valid) {
		err = fmt.Errorf("invalid jwt token")
		return nil, err
	}

	return claims, nil
}
