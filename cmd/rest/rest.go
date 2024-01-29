package main

import (
	"log"

	"github.com/alvinfebriando/costumer-test/handler"
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

	userRepository := repository.NewUserRepository(db)
	customerUsecase := usecase.NewCustomerUsecase(userRepository)
	customerHandler := handler.NewCustomerHandler(customerUsecase)

	handlers := router.Handlers{
		Customer: customerHandler,
	}

	r := router.New(handlers)
	s := server.New(r)
	server.StartWithGracefulShutdown(s)
}
