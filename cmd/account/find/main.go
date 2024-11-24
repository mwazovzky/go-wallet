package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"go-wallet/services/database"
	"go-wallet/services/repository"
	"go-wallet/services/wallet"
	"log"
	"os"
)

const databasePath = "./database/ethereum.db"
const keystorePath = "./keystore"

// go run main.go address
func main() {
	address, err := parseAddress()
	if err != nil {
		log.Fatal("missing required arguments")
	}

	db, err := database.InitConnection(databasePath)
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	defer db.Close()

	ar := repository.NewAccountRepository(db)

	model, err := ar.Find(address)
	if err != nil {
		log.Fatal("failed to query data", err)
	}

	w := wallet.NewWallet(keystorePath)
	pk, err := w.FindAccount(model.Address, model.Password)
	if err != nil {
		log.Fatal("failed to find account", err)
	}

	pkString := privateKeyToHex(pk)

	log.Println(pkString)
}

func parseAddress() (string, error) {
	args := os.Args
	if len(args) < 2 {
		return "", fmt.Errorf("missing required arguments")
	}

	return os.Args[1], nil
}

func privateKeyToHex(pk *ecdsa.PrivateKey) string {
	pkHex := make([]byte, 32)
	pkBytes := pk.D.Bytes()
	copy(pkHex[32-len(pkBytes):], pkBytes)

	return hex.EncodeToString(pkHex)
}
