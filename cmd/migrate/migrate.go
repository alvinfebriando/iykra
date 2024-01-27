package main

import (
	"log"

	"github.com/alvinfebriando/costumer-test/migration"
	"github.com/alvinfebriando/costumer-test/repository"
)

func main() {
	db, err := repository.GetConnection()
	if err != nil {
		log.Println(err)
	}

	migration.Migrate(db)
}
