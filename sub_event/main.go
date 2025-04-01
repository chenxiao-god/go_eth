package main

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {

	client, err := ethclient.Dial("wss://mainnet.infura.io/ws")
	if err != nil {
		log.Println("Failed to connect to Ethereum client:", err)
	}
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Println("Failed to subscribe to new head events:", err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Println("Subscription error:", err)
		case header := <-headers:
			log.Println("New block header:", header)
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Println("Failed to fetch block by hash:", err)
			}
			log.Println("Block number:", block.Number().Uint64())
			log.Println("Block hash:", block.Hash().Hex())
			log.Println("Block timestamp:", block.Time())
			log.Println("Block transactions:", len(block.Transactions()))
		}
	}
}
