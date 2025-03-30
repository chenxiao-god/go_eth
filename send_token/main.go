package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
)

func main() {
	// 初始化客户端
	client, err := ethclient.Dial("https://sepolia.infura.io")
	if err != nil {
		log.Fatal(err)
	}
	// 获取ECDSA对象
	privateKey, err := crypto.HexToECDSA("00000000000000000000000")
	if err != nil {
		log.Fatal(err)
	}
	// 获取公钥
	publicKey := privateKey.Public()
	// 将公钥转换为*ecdsa.PublicKey
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	// 生成公共地址 只有公共地址可以发起交易
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// 获取nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	// 设置交易金额
	value := big.NewInt(0)
	// 获取gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 设置接收地址
	toAddress := common.HexToAddress("00000000000000000000000")
	// 获取token合约地址
	tokenAddress := common.HexToAddress("00000000000000000000000")
	// 获取发送token的函数选择器
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	log.Println(hexutil.Encode(methodID)) // 0xa9059cbb
	// 获取接收地址 填充地址
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	log.Println(hexutil.Encode(paddedAddress))
	// 获取发送token的数量 填充数量
	amount := new(big.Int)
	amount.SetString("1", 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	log.Println(hexutil.Encode(paddedAmount))
	// 拼接数据
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	// 获取gas
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	gasLimit = 50000
	log.Println(gasLimit)
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &tokenAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})
	// 获取chainID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("tx sent: %s", signedTx.Hash().Hex())
}
