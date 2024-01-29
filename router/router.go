package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct{}

func New(handlers Handlers) http.Handler {
	r := gin.New()

	r.Use(gin.Recovery())
	
	return r
}
