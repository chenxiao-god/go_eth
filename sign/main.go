package main

import (
	"bytes"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func main() {
	// 数据加签
	privateKey, err := crypto.HexToECDSA("")
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}
	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	sig, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatalf("Failed to sign hash: %v", err)
	}
	log.Println("Signature:", sig)
	// 验证签名 1. 使用ecrecover 从签名中 获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Failed to get public key: %v", err)
	}
	// 转为字节数组
	fromECDSAPub := crypto.FromECDSAPub(publicKeyECDSA)
	ecrecover, err := crypto.Ecrecover(hash.Bytes(), sig)
	if err != nil {
		log.Fatalf("Failed to get public key: %v", err)
	}
	if bytes.Equal(fromECDSAPub, ecrecover) {
		log.Println("Signature is valid")
	} else {
		log.Println("Signature is invalid")
	}
	// 2 使用SigToPub 获取公钥
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		log.Fatal(err)
	}
	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches := bytes.Equal(sigPublicKeyBytes, fromECDSAPub)
	log.Println(matches)
	// 3.使用VerifySignature 验证签名
	id := sig[:len(sig)-1]
	verifySignature := crypto.VerifySignature(fromECDSAPub, hash.Bytes(), id)
	log.Println("verifySignature:", verifySignature)
}
