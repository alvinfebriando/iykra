package repository

import (
	"context"
	"errors"

	"github.com/alvinfebriando/costumer-test/entity"
	"github.com/alvinfebriando/costumer-test/valueobject"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	Find(ctx context.Context, query *valueobject.Query) ([]*entity.User, error)
	FindById(ctx context.Context, userId uint) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) (*entity.User, error)
	Delete(ctx context.Context, user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Find(ctx context.Context, query *valueobject.Query) ([]*entity.User, error) {
	var users []*entity.User
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
func (r *userRepository) FindById(ctx context.Context, userId uint) (*entity.User, error) {
	var user *entity.User
	err := r.db.
		WithContext(ctx).
		Where("id", userId).
		First(&user).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user *entity.User
	err := r.db.
		WithContext(ctx).
		Where("email = ?", email).
		First(&user).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (r *userRepository) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := r.db.
		WithContext(ctx).
		Model(user).
		Clauses(clause.Returning{}).
		Select("*").
		Updates(user).
		Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (r *userRepository) Delete(ctx context.Context, user *entity.User) error {
	err := r.db.WithContext(ctx).Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}
