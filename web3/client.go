package web3

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ctx         = context.Background()
	url         = "https://data-seed-prebsc-2-s3.binance.org:8545"
	client, err = ethclient.DialContext(ctx, url)
)

func GetBlockByNumber(number *big.Int) *types.Block {
	block, err := client.BlockByNumber(ctx, number)
	if err != nil {
		log.Println(err)
	}

	return block
}

func GetLatestBlock() *types.Block {
	block, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		log.Println(err)
	}

	return block
}

func GetHeaderByNumber(number *big.Int) *types.Header {
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Println(err)
	}

	return header
}

func GetTransactionByHash(hash common.Hash) *types.Transaction {
	tx, isPending, err := client.TransactionByHash(ctx, hash)
	if err != nil {
		log.Println(err)
	}

	if isPending {
		fmt.Println("pending")
	}

	return tx
}

func GetTransactionReceipt(txHash common.Hash) *types.Receipt {
	receipt, err := client.TransactionReceipt(ctx, txHash)
	if err != nil {
		log.Println(err)
	}

	return receipt
}
