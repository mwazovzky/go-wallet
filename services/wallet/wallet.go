package wallet

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	keystorePath string
}

func NewWallet(keystorePath string) *Wallet {
	return &Wallet{keystorePath}
}

func (w *Wallet) CreateAccount(password string) (*accounts.Account, error) {
	ks := keystore.NewKeyStore(w.keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := ks.NewAccount(password)

	return &account, err
}

func (w *Wallet) FindAccount(address string, password string) (*ecdsa.PrivateKey, error) {
	file, err := findKeystore(w.keystorePath, address)
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

func (w *Wallet) ImportAccount(privateKeyHex string, password string) (*accounts.Account, error) {
	pkBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return nil, err
	}

	pk, err := crypto.ToECDSA(pkBytes)
	if err != nil {
		return nil, err
	}

	ks := keystore.NewKeyStore(w.keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := ks.ImportECDSA(pk, password)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func findKeystore(keystorePath string, address string) (string, error) {
	files, err := os.ReadDir(keystorePath)
	if err != nil {
		log.Println("failed to read keystore directory", err)
		return "", err
	}

	for _, file := range files {
		suffix := strings.ToLower(strings.TrimPrefix(address, "0x"))
		if strings.HasSuffix(file.Name(), suffix) {
			return filepath.Join(keystorePath, file.Name()), nil
		}
	}

	return "", fmt.Errorf("keystore file not found")
}
