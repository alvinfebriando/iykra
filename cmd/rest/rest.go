package main

import (
	"log"

	"github.com/alvinfebriando/costumer-test/appjwt"
	"github.com/alvinfebriando/costumer-test/handler"
	hasher "github.com/alvinfebriando/costumer-test/hash"
	"github.com/alvinfebriando/costumer-test/repository"
	"github.com/alvinfebriando/costumer-test/router"
	"github.com/alvinfebriando/costumer-test/server"
	"github.com/alvinfebriando/costumer-test/usecase"
)

func main() {
	db, err := repository.GetConnection()
	if err != nil {
		log.Printf("failed to get db connection: %s\n", err)
	}

	hash := hasher.NewHasher()
	jwt := appjwt.NewJwt()

	userRepository := repository.NewUserRepository(db)
	customerUsecase := usecase.NewCustomerUsecase(userRepository)
	customerHandler := handler.NewCustomerHandler(customerUsecase)

	authUsecase := usecase.NewAuthUsecase(userRepository, hash, jwt)
	authHandler := handler.NewAuthHandler(authUsecase)

	handlers := router.Handlers{
		Auth:     authHandler,
		Customer: customerHandler,
	}

	r := router.New(handlers)
	s := server.New(r)
	server.StartWithGracefulShutdown(s)
}
