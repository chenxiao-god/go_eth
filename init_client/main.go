package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 连接
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	// 账户
	account := common.HexToAddress("0xfioiuiouyogfjhjjh")
	log.Println("账户", account)
	blockNumber := big.NewInt(0)
	// 余额
	balance, _ := client.BalanceAt(context.Background(), account, blockNumber)
	log.Println("余额", balance)
	// 余额转换
	fbalance := new(big.Float)
	log.Println("fbalance Float", fbalance)
	// 单位转换
	fbalance.SetString(balance.String())
	log.Println("fbalance String", fbalance)
	// 单位转换
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	log.Println("ethValue", ethValue)
	// 最新区块高度
	nonce, err := client.NonceAt(context.Background(), account, blockNumber)
	log.Println("nonce", nonce)
	// 链ID
	chainID, err := client.ChainID(context.Background())

	log.Println("chainID:", chainID)
	pendingNonceAt, err := client.PendingNonceAt(context.Background(), account)
	log.Println("pendingNonceAt", pendingNonceAt)
}
