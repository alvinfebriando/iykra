package usecase

import (
	"context"

	"github.com/alvinfebriando/costumer-test/entity"
	"github.com/alvinfebriando/costumer-test/repository"
	"github.com/alvinfebriando/costumer-test/valueobject"
)

type CustomerUsecase interface {
	GetAllCustomers(ctx context.Context, query *valueobject.Query) ([]*entity.User, error)
}

type customerUsecase struct {
	userRepo repository.UserRepository
}

func NewCustomerUsecase(userRepo repository.UserRepository) CustomerUsecase {
	return &customerUsecase{userRepo: userRepo}
}

func (u *customerUsecase) GetAllCustomers(ctx context.Context, query *valueobject.Query) ([]*entity.User, error) {
	return u.userRepo.Find(ctx, query)
}
