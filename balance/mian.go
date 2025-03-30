package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/b3cca88747014b828db633bfc327473e")
	if err != nil {
		log.Println("err", err)
	}
	address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balanceAt, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("balanceAt", balanceAt)
	newInt := big.NewInt(5532993)
	b, err := client.BalanceAt(context.Background(), address, newInt)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("b", b)
	float := new(big.Float)
	float.SetInt(balanceAt)
	quo := new(big.Float).Quo(float, big.NewFloat(math.Pow10(18)))
	log.Println("quo", quo)
	pendingBalance, err := client.PendingBalanceAt(context.Background(), address)
	if err != nil {
		log.Println("err", err)
	}

	log.Println("pendingBalance", pendingBalance)

}
