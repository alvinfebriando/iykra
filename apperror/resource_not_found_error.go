package apperror

import "fmt"

type ResourceNotFoundError struct {
	Name string
}

func NewResourceNotFoundError(name string) error {
	return NewClientError(&ResourceNotFoundError{Name: name}).NotFound()
}

func (e *ResourceNotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Name)
}
