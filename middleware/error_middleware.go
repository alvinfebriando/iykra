package middleware

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/alvinfebriando/costumer-test/apperror"
	"github.com/alvinfebriando/costumer-test/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) < 1 {
			return
		}

		err := c.Errors.Last().Err
		message := strings.Split(err.Error(), "\n")

		var sErr *json.SyntaxError
		var uErr *json.UnmarshalTypeError
		var vErr validator.ValidationErrors
		var cErr *apperror.ClientError

		isClientError := false
		if errors.As(err, &cErr) {
			isClientError = true
			err = cErr.UnWrap()
		}

		switch {
		case errors.Is(err, io.EOF):
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: message,
			})
		case errors.As(err, &sErr):
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: message,
			})
		case errors.As(err, &uErr):
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: message,
			})
		case errors.As(err, &vErr):
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
				Error: message,
			})
		case isClientError:
			c.AbortWithStatusJSON(cErr.HttpStatusCode(), dto.Response{
				Error: message,
			})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
				Error: message,
			})
		}
	}
}
