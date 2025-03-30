package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go_eth/select_ERC20/token"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/XXXXXXX")
	if err != nil {
		panic(err)
	}

	contractAddress := common.HexToAddress("XXXXXXXXXX")
	fromAddress := common.HexToAddress("XXXX")
	toAddress := common.HexToAddress("XXXXXXXX")
	erc20, err := token.NewERC20(contractAddress, client)
	if err != nil {
		panic(err)
	}
	// 检查 fromAddress 的余额
	balance, err := erc20.BalanceOf(&bind.CallOpts{}, fromAddress)
	if err != nil {
		log.Println("BalanceOf failed:", err)
		panic(err)
	}
	log.Println("fromAddress balance:", balance)

	// 检查授权额度
	allowance, err := erc20.Allowance(&bind.CallOpts{}, fromAddress, toAddress)
	if err != nil {
		log.Println("Allowance failed:", err)
		panic(err)
	}
	log.Println("Allowance:", allowance)
	// 加载私钥
	privateKey, err := crypto.HexToECDSA("XXXXXXXXXXXXXXXX")
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// 创建 TransactOpts
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatalf("Failed to create transact opts: %v", err)
	}
	// 执行 TransferFrom
	from, err := erc20.TransferFrom(opts, fromAddress, toAddress, big.NewInt(1))
	if err != nil {
		log.Println("TransferFrom failed:", err)
		panic(err)
	}
	log.Println("TransferFrom success:", from)

	// 检查 toAddress 的余额
	balanceOf, err := erc20.BalanceOf(&bind.CallOpts{}, toAddress)
	if err != nil {
		log.Println("BalanceOf failed:", err)
		panic(err)
	}
	log.Println("toAddress balance:", balanceOf)

}
