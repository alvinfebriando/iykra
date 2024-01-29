package main

import (
	"github.com/alvinfebriando/costumer-test/router"
	"github.com/alvinfebriando/costumer-test/server"
)

func main() {
	handlers := router.Handlers{}

	r := router.New(handlers)
	s := server.New(r)
	server.StartWithGracefulShutdown(s)
}
