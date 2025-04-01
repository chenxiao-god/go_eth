package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"regexp"
)

func main() {
	// 创建正则表达式
	compile := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	fmt.Printf("is valid: %v\n", compile.MatchString("333333333333333333333333333")) // is valid: true
	fmt.Printf("is valid: %v\n", compile.MatchString("333333333333333333333333333333"))

	client, err := ethclient.Dial("xxxxxxxxx")
	if err != nil {
		log.Println("err", err)
	}
	// 判断地址智能合约 还是以太坊账户
	address := common.HexToAddress("0x9472B68ba946858FF58bdcEe4992A5CB3739260F")
	bytes, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Println("err", err)
	}
	flag := len(bytes) > 0
	log.Println("flag", flag)
	// 判断地址智能合约 还是以太坊账户
	address = common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
	bytes, err = client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Println("err", err)
	}
	flag = len(bytes) > 0
	log.Println("flag", flag)
}
