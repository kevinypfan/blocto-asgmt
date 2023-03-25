package web3

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	db "github.com/kevinypfan/blocto-asgmt/db/sqlc"
)

func RunCrawl(store *db.SQLStore) {
	ctx := context.Background()

	block := getLatestBlock()

	fmt.Println(block.Number())

	arg := db.CreateBlockParams{
		BlockNum:   block.Number().Int64(),
		BlockHash:  block.Hash().Hex(),
		BlockTime:  block.ReceivedAt.Unix(),
		ParentHash: block.ParentHash().Hex(),
	}

	_, err := store.CreateBlock(ctx, arg)

	if err != nil {
		log.Println(err)
	}

	transactions := block.Body().Transactions

	for _, tx := range transactions {
		go func(tx *types.Transaction) {
			from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
			if err != nil {
				log.Println(err)
			}

			fmt.Println(tx.Hash().Hex())
			arg := db.CreateTransactionParams{
				TxHash:    tx.Hash().Hex(),
				BlockHash: block.Hash().Hex(),
				BlockNum:  block.Number().Int64(),
				From:      from.Hex(),
				To:        tx.To().Hex(),
				Nonce:     int64(tx.Nonce()),
				Value:     tx.Value().Int64(),
			}
			_, err = store.CreateTransaction(ctx, arg)
			if err != nil {
				log.Println(err)
			}

			// receipt := getTransactionReceipt(tx.Hash())
			// logs := receipt.Logs

			// for _, rlog := range logs {

			// 	topics := []string{}

			// 	for _, topic := range rlog.Topics {
			// 		topics = append(topics, topic.Hex())
			// 	}

			// 	arg := db.CreateLogParams{
			// 		Address:   rlog.Address.Hex(),
			// 		Topics:    topics,
			// 		Data:      string(rlog.Data),
			// 		BlockNum:  block.Number().Int64(),
			// 		TxHash:    tx.Hash().Hex(),
			// 		BlockHash: block.Hash().Hex(),
			// 		Removed:   rlog.Removed,
			// 	}

			// 	_, err = store.CreateLog(ctx, arg)
			// 	if err != nil {
			// 		log.Println(err)
			// 	}

			// }
		}(tx)
	}

}
