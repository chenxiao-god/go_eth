package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {

	// 1.初始化客户端
	client, err := ethclient.Dial("https://sepolia.infura.io")
	if err != nil {
		log.Println("err", err)
	}
	// 2.获取ECDSA对象
	privateKey, err := crypto.HexToECDSA("0000000000000000000000000000")
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
	// 5.交易额度
	value := big.NewInt(1e18)
	// 6.获取gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println("err", err)
	}
	// 7.设定gas最大限制
	gasLimit := uint64(3000000)
	// 8.设定接收地址
	toAddress := common.HexToAddress("00000000000000000000000000000")
	networkID, _ := client.NetworkID(context.Background())
	// 9.生成交易 参数1：nonceAt 参数2：toAddress 参数3：value 参数4：gasLimit 参数5：gasPrice 参数6：data
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonceAt,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     nil,
	})
	// 10.签名交易 参数1：交易对象 参数2：签名对象 参数3：私钥·
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(networkID), privateKey)
	if err != nil {
		log.Println("err", err)
	}
	// 11发送交易
	if err := client.SendTransaction(context.Background(), signedTx); err != nil {
		log.Println("err", err)
	}
}
