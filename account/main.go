package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

func main() {

	hexToAddress := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	log.Println("hexToAddress", hexToAddress)
	hex := hexToAddress.Hex()
	log.Println("hex", hex)
	log.Println("hexToAddress", hexToAddress.String())
	fmt.Println(common.BytesToHash(hexToAddress.Bytes()).Hex())
	log.Println("hexToAddress", hexToAddress.Bytes())

}
