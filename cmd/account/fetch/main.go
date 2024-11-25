package main

import (
	"go-wallet/services/config"
	"go-wallet/services/database"
	"go-wallet/services/repository"
	"log"
)

func main() {
	cfg := config.Load()

	db, err := database.InitConnection(cfg.DatabasePath)
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	defer db.Close()

	ar := repository.NewAccountRepository(db)

	data, err := ar.Fetch()
	if err != nil {
		log.Fatal("failed to query data", err)
	}

	log.Printf("%+v", data)
}
