package main

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {

	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		panic(err)
	}
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Panicln(err)
	}
	transactions := block.Transactions()
	for _, transaction := range transactions {
		log.Println(transaction.Hash().Hex())
		log.Println(transaction.Value().String())
		log.Println(transaction.Gas())
		log.Println(transaction.GasPrice())
		log.Println(transaction.Nonce())
		log.Println(transaction.To().Hex())
		networkID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Println(err)
		}
		log.Println(networkID)
		sender, err := types.Sender(types.NewEIP155Signer(networkID), transaction)
		if err != nil {
			log.Println(err)
		}
		log.Println(sender.Hex())
	}

}
