package main

import (
	"log"
	"wallet/services/database"
	"wallet/services/repository"

	_ "modernc.org/sqlite"
)

const ethereum = "ethereum"

func main() {
	db, err := database.InitConnection()
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	defer db.Close()

	database.SetupDatabase(db)

	ar := repository.NewAccountRepository(db)

	a := repository.Account{
		Address:  "hash-123",
		Password: "pwd",
		Chain:    ethereum,
	}

	err = ar.Create(a)
	if err != nil {
		log.Fatal("failed to create account record", err)
	}
}
