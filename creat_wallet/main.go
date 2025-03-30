package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

func main() {
	// 生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	// 私钥转16进制
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])
	// 生成公钥
	publicKey := privateKey.Public()
	// 转为ecdsa.PublicKey
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	// 公钥转16进制
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// 打印
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])
	// 生成公共地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)
	// 计算地址
	hash := sha3.NewLegacyKeccak256()
	// 删除前4个字节
	hash.Write(publicKeyBytes[1:])
	// 取后20个字节
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))
}
