package router

import (
	"net/http"

	"github.com/alvinfebriando/costumer-test/handler"
	"github.com/alvinfebriando/costumer-test/middleware"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Auth     *handler.AuthHandler
	Customer *handler.CustomerHandler
}

func New(h Handlers) http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Error())

	r.POST("/auth/register", h.Auth.Register)
	r.POST("/auth/login", h.Auth.Login)

	r.GET("/customers", middleware.Auth(), h.Customer.ListCustomer)
	r.PUT("/customers/:id", middleware.Auth(), h.Customer.UpdateCustomer)
	r.DELETE("/customers/:id", middleware.Auth(), h.Customer.DeleteCustomer)

	return r
}
