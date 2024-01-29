package router

import (
	"net/http"

	"github.com/alvinfebriando/costumer-test/handler"
	"github.com/alvinfebriando/costumer-test/middleware"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Customer *handler.CustomerHandler
}

func New(h Handlers) http.Handler {
	r := gin.New()
	r.Use(middleware.Error())

	r.GET("/customers", h.Customer.ListCustomer)
	r.POST("/customers", h.Customer.AddCustomer)

	r.Use(gin.Recovery())

	return r
}
