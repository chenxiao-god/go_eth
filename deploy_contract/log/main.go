package main

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws")
	if err != nil {
		log.Println("Failed to connect to Ethereum client:", err)
	}
	address := common.HexToAddress("0x6eE4eb7A07666d0408e469088B9535390Ce8821a")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}
	logs := make(chan types.Log)
	filterLogs, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Println("Failed to subscribe to logs:", err)
	}
	for {
		select {
		case err := <-filterLogs.Err():
			log.Println("Subscription error:", err)
		case vlog := <-logs:
			log.Println("New log event:", vlog)
		}
	}
}
