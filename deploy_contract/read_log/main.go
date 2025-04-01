package main

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	store "go_eth/deploy_contract"
	"log"
	"math/big"
	"strings"
)

func main() {
	// 连接节点
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/b3cca88747014b828db633bfc327473e")
	if err != nil {
		log.Println("Failed to connect to Ethereum client:", err)
	}
	address := common.HexToAddress("0x6eE4eb7A07666d0408e469088B9535390Ce8821a")
	// 过滤器查询
	filterQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(8014266),
		ToBlock:   big.NewInt(8014266),
		Addresses: []common.Address{
			address,
		},
	}
	// 获取日志
	logs, err := client.FilterLogs(context.Background(), filterQuery)
	if err != nil {
		log.Println("Failed to filter logs:", err)
	}
	// 解析合约abi
	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		log.Println("Failed to parse contract ABI:", err)
	}
	// 遍历日志 解析事件
	for _, vLog := range logs {
		log.Println("New log event:", vLog)
		log.Println("Block Number:", vLog.BlockNumber)
		log.Println("Block Hash:", vLog.BlockHash.Hex())
		log.Println("Tx Hash:", vLog.TxHash.Hex())
		log.Println("Tx Index:", vLog.TxIndex)
		log.Println("Topics:", vLog.Topics)
		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}
		//  解析事件
		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Println("Failed to unpack log data:", err)
		}
		log.Println("Key:", string(event.Key[:]))
		log.Println("Value:", string(event.Value[:]))
		// 解析topic
		var topics [4]string
		for i, topic := range vLog.Topics {
			topics[i] = topic.Hex()
		}
		log.Println("Topics:", topics[0])
	}
	// 计算event签名
	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	// 取签名
	hash := crypto.Keccak256Hash(eventSignature)
	log.Println("Hash:", hash.Hex())
}
