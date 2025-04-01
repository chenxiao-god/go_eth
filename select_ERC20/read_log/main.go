package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go_eth/select_ERC20/token"

	"log"
	"math/big"
	"strings"
)

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func main() {
	client, err := ethclient.Dial("wss://sepolia.infura.io/")
	if err != nil {
		log.Println("Failed to connect to Ethereum client:", err)
	}
	address := common.HexToAddress("45789907654534567890765")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(7962782),
		ToBlock:   big.NewInt(8011954),
		Addresses: []common.Address{
			address,
		},
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(token.ERC20ABI)))
	if err != nil {
		log.Println("Failed to parse contract ABI:", err)
	}
	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Println("Failed to filter logs:", err)
	}
	for _, vLog := range logs {
		log.Println("New log event:", vLog)
		log.Println("Block Number:", vLog.BlockNumber)
		log.Println("Block Hash:", vLog.BlockHash.Hex())
		log.Println("Tx Hash:", vLog.TxHash.Hex())
		log.Println("Tx Index:", vLog.TxIndex)
		log.Println("Topics:", vLog.Topics)

		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")
			var transferEvent LogTransfer
			_, err := contractAbi.Unpack("Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())

		case logApprovalSigHash.Hex():
			fmt.Printf("Log Name: Approval\n")

			var approvalEvent LogApproval

			_, err := contractAbi.Unpack("Approval", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
			fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
			fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())
		}

	}
}
