package handler

import (
	"net/http"

	"github.com/alvinfebriando/costumer-test/dto"
	"github.com/alvinfebriando/costumer-test/usecase"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerUsecase usecase.CustomerUsecase
}

func NewCustomerHandler(customerUsecase usecase.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{customerUsecase: customerUsecase}
}

func (h *CustomerHandler) ListCustomer(c *gin.Context) {
	var request dto.ListCustomerQueryParam
	if err := c.ShouldBindQuery(&request); err != nil {
		_ = c.Error(err)
		return
	}

	query := request.ToQuery()
	fetchedCustomers, err := h.customerUsecase.GetAllCustomers(c.Request.Context(), query)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Data: dto.NewResponsesFromUsers(fetchedCustomers),
	})
}

func (h *CustomerHandler) AddCustomer(c *gin.Context) {
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

	createdCustomer, err := h.customerUsecase.AddCustomer(c.Request.Context(), user)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Data: dto.NewResponseFromUser(createdCustomer),
	})
}
