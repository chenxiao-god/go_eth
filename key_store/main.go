package main

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"log"
	"os"
)

func createKeyStore() {
	keyStore := keystore.NewKeyStore("/Users/xiaochen/go/src/go_eth/key_store", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := keyStore.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(account.Address.Hex()) //0xc927A9ed212937598e47828335E52772364C74b2
}

func importKeyStore() {
	file := "/Users/xiaochen/go/src/go_eth/key_store/UTC--2025-03-24T13-49-41.729215000Z--9ff7386a5ac8a23f5f2e7417cd14b10cf508289e"
	ks := keystore.NewKeyStore("/Users/xiaochen/go/src/go_eth/key_store/new_key_store", keystore.StandardScryptN, keystore.StandardScryptP)
	bytes, err := ioutil.ReadFile(file)
	log.Println("bytes", string(bytes))
	if err != nil {
		log.Fatal(err)
	}
	// 导入
	account, err := ks.Import(bytes, "secret", "secret")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(account.Address.Hex()) //0x9Ff7386a5AC8a23F5F2E7417Cd14b10cF508289E
	// 删除源文件
	err = os.Remove(file)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//createKeyStore()
	importKeyStore()
}
