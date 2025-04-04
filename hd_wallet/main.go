package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"log"
	"math/big"
)

func main() {
	///***
	//  使用助记词 生成钱包 钱包中包含多个地址
	//*/
	//
	//// 生成助记词
	//mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	//// 生成钱包
	//wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// 生成地址
	//path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	//// 生成地址
	//account, err := wallet.Derive(path, false)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println(account.Address.Hex()) // 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947
	//// 生成地址
	//path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	//account, err = wallet.Derive(path, false)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(account.Address.Hex()) // 0x8230645aC28A4EdD1b0B53E7Cd8019744E9dD559

	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, true)
	if err != nil {
		log.Fatal(err)
	}

	nonce := uint64(0)
	value := big.NewInt(1000000000000000000)
	toAddress := common.HexToAddress("0x0")
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(21000000000)
	var data []byte

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedTx, err := wallet.SignTx(account, tx, nil)
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(signedTx)
}
