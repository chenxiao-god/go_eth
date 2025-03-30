package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

func main() {

	hexToAddress := common.HexToAddress("**************")
	log.Println("hexToAddress", hexToAddress)
	hex := hexToAddress.Hex()
	log.Println("hex", hex)
	log.Println("hexToAddress", hexToAddress.String())
	fmt.Println(common.BytesToHash(hexToAddress.Bytes()).Hex())
	log.Println("hexToAddress", hexToAddress.Bytes())

}
