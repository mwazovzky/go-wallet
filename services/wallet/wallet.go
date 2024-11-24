package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type Wallet struct {
	path string
}

func NewWallet(path string) *Wallet {
	return &Wallet{path}
}

func (w *Wallet) CreateAccount(password string) (accounts.Account, error) {
	ks := keystore.NewKeyStore(w.path, keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := ks.NewAccount(password)
	if err != nil {
		log.Println("failed to create account", err)
	}

	log.Println("account created, address", account.Address.Hex())

	return account, err
}

func (w *Wallet) FindAccount(address string, password string) (*ecdsa.PrivateKey, error) {
	file, err := findKeystore(w.path, address)
	if err != nil {
		log.Println("failed to find keystore", err)
		return nil, err
	}

	data, err := os.ReadFile(file)
	if err != nil {
		log.Println("failed to read keystore", err)
		return nil, err
	}

	key, err := keystore.DecryptKey(data, password)
	if err != nil {
		log.Println("failed to decrypt keystore", err)
		return nil, err
	}

	return key.PrivateKey, nil
}

func findKeystore(path string, address string) (string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Println("failed to read keystore directory", err)
		return "", err
	}

	for _, file := range files {
		suffix := strings.ToLower(strings.TrimPrefix(address, "0x"))
		if strings.HasSuffix(file.Name(), suffix) {
			return filepath.Join(path, file.Name()), nil
		}
	}

	return "", fmt.Errorf("keystore file not found")
}
