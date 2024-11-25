package explorer

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Explorer struct {
	client *ethclient.Client
}

func NewExplorer(nodeUrl string) (*Explorer, error) {
	client, err := ethclient.Dial(nodeUrl)

	if err != nil {
		return nil, err
	}

	return &Explorer{client}, nil
}

func (e *Explorer) GetBalance(addressHex string) *big.Float {
	address := common.HexToAddress(addressHex)
	balanceWei, err := e.client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	balanceEth, ok := new(big.Float).SetString(balanceWei.String())
	if !ok {
		log.Fatal("failed to convert balance to eth")
	}

	return new(big.Float).Quo(balanceEth, big.NewFloat(1e+18))
}
