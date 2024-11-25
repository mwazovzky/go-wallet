package main

import (
	"go-wallet/services/config"
	"go-wallet/services/database"
	"go-wallet/services/repository"
	"go-wallet/services/wallet"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		log.Fatal("missing required arguments")
	}

	cfg := config.Load()

	pkHex := os.Args[1]
	password := os.Args[2]

	db, err := database.InitConnection(cfg.DatabasePath)
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	defer db.Close()

	database.SetupDatabase(db)
	ar := repository.NewAccountRepository(db)

	w := wallet.NewWallet(cfg.KeystorePath)
	account, err := w.ImportAccount(pkHex, password)
	if err != nil {
		log.Fatal("failed to create keystore: ", err)
	}

	model := repository.Account{
		Address:  account.Address.Hex(),
		Password: password,
	}

	err = ar.Create(model)
	if err != nil {
		log.Fatal("failed to create account record", err)
	}
}
