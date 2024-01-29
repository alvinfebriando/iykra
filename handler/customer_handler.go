package handler

import (
	"log"
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

func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	var customerId dto.RequestUri
	var request dto.AddCustomerRequest

	if err := c.ShouldBindUri(&customerId); err != nil {
		log.Println(err)
		_ = c.Error(err)
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(err)
		return
	}

	user, err := request.ToUser()
	if err != nil {
		_ = c.Error(err)
		return
	}

	user.Id = customerId.Id
	updatedCustomer, err := h.customerUsecase.UpdateCustomer(c.Request.Context(), user)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Data: dto.NewResponseFromUser(updatedCustomer),
	})

}

func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	var customerId dto.RequestUri
	if err := c.ShouldBindUri(&customerId); err != nil {
		log.Println(err)
		_ = c.Error(err)
		return
	}

	err := h.customerUsecase.DeleteCustomer(c.Request.Context(), customerId.Id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
