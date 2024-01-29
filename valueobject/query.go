package valueobject

import "fmt"

type Query struct {
	page   int
	limit  *int
	sortBy *string
	order  Order
}

func NewQuery() *Query {
	return &Query{
		page:   1,
		limit:  nil,
		sortBy: nil,
		order:  OrderAsc,
	}
}

func (q *Query) WithPage(page int) *Query {
	q.page = page
	return q
}

func (q *Query) GetPage() int {
	return q.page
}

func (q *Query) WithLimit(limit int) *Query {
	q.limit = &limit
	return q
}

func (q *Query) GetLimit() *int {
	return q.limit
}

func (q *Query) WithSortBy(field string) *Query {
	q.sortBy = &field
	return q
}

func (q *Query) WithOrder(order Order) *Query {
	q.order = order
	return q
}

func (q *Query) GetOrder() string {
	if q.sortBy == nil {
		return ""
	}
	return fmt.Sprintf("%s %s", *q.sortBy, q.order)
}
