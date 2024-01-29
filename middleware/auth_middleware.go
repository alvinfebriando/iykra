package middleware

import (
	"context"
	"strings"

	"github.com/alvinfebriando/costumer-test/apperror"
	"github.com/alvinfebriando/costumer-test/appjwt"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		token, err := extractBearerToken(bearerToken)
		if err != nil {
			c.Abort()
			_ = c.Error(err)
			return
		}

		newJwt := appjwt.NewJwt()
		claims, err := newJwt.ValidateToken(token)
		if err != nil {
			c.Abort()
			_ = c.Error(apperror.NewInvalidTokenError())
			return
		}

		ctx := context.WithValue(c.Request.Context(), "user_id", claims.Id)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func extractBearerToken(bearerToken string) (string, error) {
	if bearerToken == "" {
		return "", apperror.NewMissingTokenError()
	}
	token := strings.Split(bearerToken, " ")
	if len(token) != 2 {
		return "", apperror.NewInvalidTokenError()
	}
	return token[1], nil
}
