package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	store "go_eth/deploy_contract"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io")
	if err != nil {
		log.Println("Failed to connect to Ethereum client:", err)
	}
	privateKey, err := crypto.HexToECDSA("99d1d5bce8ba237975387d009183ad02b4b5cc868dd57c7540e6eb2a4f5e89be")
	if err != nil {
		log.Println("err", err)
	}
	// 3.获取公钥
	publicKey := privateKey.Public()
	// 4.将公钥转换为*ecdsa.PublicKey
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("err", err)
	}
	// 5.生成公共地址-账户地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonceAt, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println("err", err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println("err", err)
	}
	networkID, err := client.NetworkID(context.Background())
	// 5.交易额度
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, networkID)
	if err != nil {
		log.Println("err", err)
	}
	opts.Nonce = big.NewInt(int64(nonceAt))
	opts.Value = big.NewInt(0)
	opts.GasLimit = uint64(10000000)
	opts.GasPrice = gasPrice
	// 6.部署合约
	//input := "1.0"
	//address, transaction, newStore, err := store.DeployStore(opts, client, input)
	//if err != nil {
	//	log.Println("err", err)
	//}
	//log.Println("address", address.Hex())
	//log.Println("transaction", transaction.Hash().Hex())
	//log.Println("newStore1", newStore)
	// 7.调用合约
	//toAddress := common.HexToAddress("0x755e689807D3b20a52613F29fD0Eb7FcE3ec2ec1")
	//newStore1, err := store.NewStore(toAddress, client)
	//if err != nil {
	//	log.Println("Failed to create new Store newStore1:", err)
	//}
	//log.Println("Store value:", newStore1.)
	//version, err := newStore1.Version(nil)
	//if err != nil {
	//	log.Println("err", err)
	//}
	//log.Println("version", version)
	// 7.1 合约交互
	toAddress := common.HexToAddress("0x6eE4eb7A07666d0408e469088B9535390Ce8821a")
	newStore1, err := store.NewStore(toAddress, client)
	if err != nil {
		log.Println("Failed to create new Store newStore1:", err)
	}
	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))
	//tx, err := newStore1.SetItem(opts, key, value)
	//if err != nil {
	//	log.Println("err", err)
	//}
	//log.Println("tx", tx.Hash().Hex())
	items, err := newStore1.Items(nil, key)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("items", string(bytes.TrimRight(items[:], "\x00")))
	codeAt, err := client.CodeAt(context.Background(), toAddress, nil)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("codeAt", hex.EncodeToString(codeAt))
}
