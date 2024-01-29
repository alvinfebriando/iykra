package router

import (
	"net/http"

	"github.com/alvinfebriando/costumer-test/handler"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Customer *handler.CustomerHandler
}

func New(h Handlers) http.Handler {
	r := gin.New()

	r.GET("/customers", h.Customer.ListCustomer)

	r.Use(gin.Recovery())

	return r
}
