package dto

import (
	"time"

	"github.com/alvinfebriando/costumer-test/entity"
	"github.com/alvinfebriando/costumer-test/valueobject"
)

type ListCustomerQueryParam struct {
	SortBy *string `form:"sort_by" binding:"omitempty,oneof=name date_of_birth"`
	Order  *string `form:"order" binding:"omitempty,oneof=asc desc"`
	Limit  *int    `form:"limit" binding:"omitempty,numeric,min=1"`
	Page   *int    `form:"page" binding:"omitempty,numeric,min=1"`
}

func (qp *ListCustomerQueryParam) ToQuery() *valueobject.Query {
	query := valueobject.NewQuery()

	if qp.Page != nil {
		query.WithPage(*qp.Page)
	}

	if qp.Limit != nil {
		query.WithLimit(*qp.Limit)
	}

	if qp.SortBy != nil {
		query.WithSortBy(*qp.SortBy)
	}

	if qp.Order != nil {
		query.WithOrder(valueobject.Order(*qp.Order))
	}

	return query
}

type AddCustomerRequest struct {
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
}

func (r *AddCustomerRequest) ToUser() (*entity.User, error) {
	const timeFormat = "2006-01-02"
	dateOfBirth, err := time.Parse(timeFormat, r.DateOfBirth)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		Name:        r.Name,
		Address:     r.Address,
		DateOfBirth: dateOfBirth,
	}, nil
}

type UserResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
}

func NewResponseFromUser(u *entity.User) *UserResponse {
	const timeFormat = "2006-01-02"
	return &UserResponse{
		Id:          u.Id,
		Name:        u.Name,
		Address:     u.Address,
		DateOfBirth: u.DateOfBirth.Format(timeFormat),
	}
}

func NewResponsesFromUsers(us []*entity.User) []*UserResponse {
	return mapEntitiesToResponses(us, NewResponseFromUser)
}
