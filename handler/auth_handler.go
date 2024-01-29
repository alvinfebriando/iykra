package handler

import (
	"net/http"

	"github.com/alvinfebriando/costumer-test/dto"
	"github.com/alvinfebriando/costumer-test/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	usecase usecase.AuthUsecase
}

func NewAuthHandler(u usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		usecase: u,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var request dto.AddCustomerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(err)
		return
	}
	user, err := request.ToUser()
	if err != nil {
		_ = c.Error(err)
		return
	}

	createdUser, err := h.usecase.Register(c.Request.Context(), user)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Data: dto.NewResponseFromUser(createdUser),
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request dto.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(err)
		return
	}
	user := request.ToUser()
	token, err := h.usecase.Login(c.Request.Context(), user)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Data: token,
	})
}
