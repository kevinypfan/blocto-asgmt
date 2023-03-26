package web3

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	db "github.com/kevinypfan/blocto-asgmt/db/sqlc"
)

func RunCrawl(store *db.SQLStore) {
	ctx := context.Background()

	// var wg sync.WaitGroup

	blockNumChain := make(chan *big.Int)
	txChain := make(chan *types.Transaction)
	logChain := make(chan *types.Log)

	block := getLatestBlock()
	fmt.Println(block.Number().Int64())

	latestBlock, err := store.GetLatestBlock(ctx)
	fmt.Println(latestBlock.BlockNum)

	if err != nil {
		log.Println(err, "GetLatestBlock")
	}

	fmt.Println("if before")

	go func(latestBlock int64, dbBlock int64) {
		if latestBlock > dbBlock {
			for i := dbBlock + 1; i <= latestBlock; i++ {
				blockNumChain <- big.NewInt(i)
			}
		}
	}(block.Number().Int64(), latestBlock.BlockNum)

	for {
		select {
		case blockNum := <-blockNumChain:
			fmt.Printf("Save Block = %v\n", blockNum)
			block := getBlockByNumber(blockNum)
			arg := db.CreateBlockParams{
				BlockHash:  block.Hash().Hex(),
				BlockTime:  block.ReceivedAt.Unix(),
				ParentHash: block.ParentHash().Hex(),
			}

			if block.Number() != nil {
				arg.BlockNum = block.Number().Int64()
			}

			_, err := store.CreateBlock(ctx, arg)

			if err != nil {
				log.Println(err, "CreateBlock")
			}

			transactions := block.Body().Transactions

			go func(transactions []*types.Transaction) {
				for _, tx := range transactions {
					fmt.Printf("txChain Tx = %v\n", tx.Hash().Hex())
					txChain <- tx
				}
			}(transactions)

		case tx := <-txChain:
			fmt.Printf("Save Tx = %v\n", tx.Hash().Hex())
			from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
			if err != nil {
				log.Println(err, "Sender")
			}

			// fmt.Println(tx.Hash().Hex())

			receipt := getTransactionReceipt(tx.Hash())

			arg := db.CreateTransactionParams{
				TxHash:    receipt.TxHash.Hex(),
				BlockHash: receipt.BlockHash.Hex(),
				BlockNum:  receipt.BlockNumber.Int64(),
				From:      from.Hex(),
				Nonce:     int64(tx.Nonce()),
			}

			if tx.To() != nil {
				arg.To = tx.To().Hex()
			}

			if tx.Value() != nil {
				arg.Value = tx.Value().Int64()
			}

			_, err = store.CreateTransaction(ctx, arg)
			if err != nil {
				log.Println(err, "CreateTransaction")
			}

			logs := receipt.Logs

			go func(logs []*types.Log) {
				for _, lg := range logs {
					logChain <- lg
				}
			}(logs)

		// default:
		// 	fmt.Printf("Nothine input\n")

		case rlog := <-logChain:
			fmt.Printf("Save Log = %v\n", rlog.Address.Hash())

			topics := []string{}

			for _, topic := range rlog.Topics {
				topics = append(topics, topic.Hex())
			}

			arg := db.CreateLogParams{
				Address:   rlog.Address.Hex(),
				Topics:    topics,
				BlockNum:  int64(rlog.BlockNumber),
				TxHash:    rlog.TxHash.Hex(),
				BlockHash: rlog.BlockHash.Hex(),
				Removed:   rlog.Removed,
			}

			_, err = store.CreateLog(ctx, arg)
			if err != nil {
				log.Println(err, "CreateLog")
			}
		}

	}

}
