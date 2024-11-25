package main

import (
	"fmt"
	"go-wallet/services/config"
	"go-wallet/services/explorer"
	"log"
	"os"
)

// go run main.go address
func main() {
	address, err := parseAddress()
	if err != nil {
		log.Fatal("missing required arguments")
	}

	cfg := config.Load()

	ex, err := explorer.NewExplorer(cfg.NodeUrl)
	if err != nil {
		log.Fatal(err)
	}

	value := ex.GetBalance(address)

	log.Printf("Account balance: %f ETH", value)
}

func parseAddress() (string, error) {
	args := os.Args
	if len(args) < 2 {
		return "", fmt.Errorf("missing required arguments")
	}

	return os.Args[1], nil
}
