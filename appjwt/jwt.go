package appjwt

import (
	"errors"
	"strconv"
	"time"

	"github.com/alvinfebriando/costumer-test/config"
	"github.com/alvinfebriando/costumer-test/entity"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Id uint `json:"id"`
}

type Jwt interface {
	GenerateToken(user *entity.User) (string, error)
	ValidateToken(tokenString string) (*entity.User, error)
}

type jwtImpl struct {
	secretKey []byte
}

func NewJwt() Jwt {
	jwtConfig := config.NewJwtConfig()
	return &jwtImpl{
		secretKey: []byte(jwtConfig.Secret),
	}
}

func (j *jwtImpl) GenerateToken(user *entity.User) (string, error) {
	const jwtExpiryDuration = 5 * time.Hour

	userId := strconv.Itoa(int(user.Id))

	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtExpiryDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   userId,
		},
		Id: user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(j.secretKey)
	return signedString, err
}

func (j *jwtImpl) ValidateToken(tokenString string) (*entity.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok {
		user := &entity.User{
			Id: claims.Id,
		}
		return user, nil
	}
	return nil, errors.New("invalid claims type")
}
