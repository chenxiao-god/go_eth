package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	dial, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		panic(err)
	}
	block, err := dial.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	log.Println("block", block.Number.String())
	newInt := big.NewInt(block.Number.Int64())
	blockByNumber, err := dial.BlockByNumber(context.Background(), newInt)
	if err != nil {
		panic(err)
	}
	log.Println("blockByNumber", blockByNumber.Number().String())
	log.Println("blockByNumber", blockByNumber.Time())
	log.Println("blockByNumber", blockByNumber.Number().Int64())
	log.Println("blockByNumber", blockByNumber.Hash().Hex())
	log.Println("blockByNumber", blockByNumber.Number().Uint64())
	log.Println("blockByNumber", len(blockByNumber.Transactions()))
	log.Println("blockByNumber", blockByNumber.Number().Cmp(big.NewInt(2)))
	count, err := dial.TransactionCount(context.Background(), blockByNumber.Hash())
	if err != nil {
		panic(err)
	}
	log.Println("count", count)
	transactions := blockByNumber.Transactions()
	for i, transaction := range transactions {
		log.Println("transaction 第", i)
		hex := transaction.Hash().Hex()
		log.Println("hex", hex)
		log.Println("transaction", transaction.Gas())
		log.Println("transaction", transaction.GasPrice())
		log.Println("transaction", transaction.Nonce())
		log.Println("transaction", transaction.Data())
		log.Println("transaction", transaction.To().Hex())
		log.Println("transaction", transaction.Value())
	}
}
