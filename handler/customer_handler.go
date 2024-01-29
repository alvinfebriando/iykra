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

	c.JSON(http.StatusOK, fetchedCustomers)
}
