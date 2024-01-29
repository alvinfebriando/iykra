package usecase

import (
	"context"

	"github.com/alvinfebriando/costumer-test/apperror"
	"github.com/alvinfebriando/costumer-test/entity"
	"github.com/alvinfebriando/costumer-test/repository"
	"github.com/alvinfebriando/costumer-test/valueobject"
)

type CustomerUsecase interface {
	GetAllCustomers(ctx context.Context, query *valueobject.Query) ([]*entity.User, error)
	AddCustomer(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdateCustomer(ctx context.Context, user *entity.User) (*entity.User, error)
	DeleteCustomer(ctx context.Context, userId uint) error
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

func (u *customerUsecase) AddCustomer(ctx context.Context, user *entity.User) (*entity.User, error) {
	return u.userRepo.Create(ctx, user)
}
func (u *customerUsecase) UpdateCustomer(ctx context.Context, user *entity.User) (*entity.User, error) {
	fetchedCustomer, err := u.userRepo.FindById(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	if fetchedCustomer == nil {
		return nil, apperror.NewResourceNotFoundError("customer")
	}

	fetchedCustomer.Name = user.Name
	fetchedCustomer.Address = user.Address
	fetchedCustomer.DateOfBirth = user.DateOfBirth

	updatedCustomer, err := u.userRepo.Update(ctx, fetchedCustomer)
	if err != nil {
		return nil, err
	}

	return updatedCustomer, nil
}
func (u *customerUsecase) DeleteCustomer(ctx context.Context, userId uint) error {
	fetchedCustomer, err := u.userRepo.FindById(ctx, userId)
	if err != nil {
		return err
	}
	if fetchedCustomer == nil {
		return apperror.NewResourceNotFoundError("customer")
	}

	customer := &entity.User{Id: userId}
	return u.userRepo.Delete(ctx, customer)
}
