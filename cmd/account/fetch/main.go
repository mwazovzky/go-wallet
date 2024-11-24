package main

import (
	"go-wallet/services/database"
	"go-wallet/services/repository"
	"log"
)

const databasePath = "./database/ethereum.db"

func main() {
	db, err := database.InitConnection(databasePath)
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
