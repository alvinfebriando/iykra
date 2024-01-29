package usecase

import (
	"context"
	"errors"

	"github.com/alvinfebriando/costumer-test/apperror"
	"github.com/alvinfebriando/costumer-test/appjwt"
	"github.com/alvinfebriando/costumer-test/entity"
	hasher "github.com/alvinfebriando/costumer-test/hash"
	"github.com/alvinfebriando/costumer-test/repository"
)

type AuthUsecase interface {
	Register(context.Context, *entity.User) (*entity.User, error)
	Login(context.Context, *entity.User) (string, error)
}

type authUsecase struct {
	userRepository repository.UserRepository
	hash           hasher.Hasher
	jwt            appjwt.Jwt
}

func NewAuthUsecase(userRepository repository.UserRepository, hash hasher.Hasher, jwt appjwt.Jwt) AuthUsecase {
	return &authUsecase{userRepository: userRepository, hash: hash, jwt: jwt}
}

func (u *authUsecase) Register(ctx context.Context, user *entity.User) (*entity.User, error) {
	fetchedUser, err := u.userRepository.FindByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	if fetchedUser != nil {
		return nil, apperror.NewClientError(errors.New("user already exists"))
	}

	hashedPassword, err := u.hash.Hash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	createdUser, err := u.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
func (u *authUsecase) Login(ctx context.Context, user *entity.User) (string, error) {
	fetchedUser, err := u.userRepository.FindByEmail(ctx, user.Email)
	if err != nil {
		return "", err
	}
	if fetchedUser == nil {
		return "", apperror.NewResourceNotFoundError("user")
	}

	if !(u.hash.Compare(fetchedUser.Password, user.Password)) {
		return "", apperror.NewInvalidCredentialsError()
	}
	token, err := u.jwt.GenerateToken(fetchedUser)
	if err != nil {
		return "", err
	}

	return token, err

}
