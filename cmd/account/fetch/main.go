package main

import (
	"log"
	"wallet/services/database"
	"wallet/services/repository"
)

func main() {
	db, err := database.InitConnection()
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
