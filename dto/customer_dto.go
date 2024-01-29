package dto

import "github.com/alvinfebriando/costumer-test/valueobject"

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
