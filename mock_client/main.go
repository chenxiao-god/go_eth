package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
)

func main() {
	// 生成一个新的私钥，用于签名交易
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	// 使用私钥创建一个交易发送者
	transactor := bind.NewKeyedTransactor(privateKey)

	// 初始化一个余额为100000000 wei的账户
	balance := new(big.Int)
	balance.SetString("1000000000000", 10)
	address := transactor.From

	// 创建创世区块的账户分配，指定账户的初始余额
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: balance,
		},
	}

	// 设置区块的Gas限制
	blockGasLimit := uint64(4712388)

	// 创建一个模拟的以太坊后端，用于测试和开发
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)

	// 获取发送者地址的当前nonce值，用于防止重放攻击
	fromAddress := transactor.From
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 设置交易的价值为1 ETH（以wei为单位）
	value := big.NewInt(1000000000000000000)

	// 设置交易的Gas限制为21000单位
	gasLimit := uint64(21000)

	// 获取当前建议的Gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 设置交易的接收者地址
	toAddress := common.HexToAddress("67098765799776578")

	// 创建一个新的交易对象
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// 获取当前链的ID，用于签名交易
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 使用私钥对交易进行签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发送签名后的交易到模拟的以太坊网络
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	// 打印交易的哈希值
	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())

	// 提交交易到区块链
	client.Commit()

	// 获取交易的收据，确认交易是否成功
	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt == nil {
		log.Fatal("receipt is nil. Forgot to commit?")
	}

	// 打印交易的状态，1表示成功，0表示失败
	fmt.Printf("status: %v\n", receipt.Status)
}
