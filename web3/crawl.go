package web3

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	db "github.com/kevinypfan/blocto-asgmt/db/sqlc"
)

const poolSize = 10
const chunk = 30
const startNum = 28377631

// Consumer struct
type Consumer struct {
	blockNumChain chan *big.Int
	txChain       chan *types.Transaction
	logChain      chan *types.Log
	jobsChan      chan int
	store         *db.SQLStore
}

func (c *Consumer) startTraceBlocks(num int64) {
	block := getLatestBlock()
	fmt.Println(block.Number().Int64())

	latestBlock, err := c.store.GetLatestBlock(ctx)
	fmt.Println(latestBlock.BlockNum)

	if err != nil {
		log.Println(err, "GetLatestBlock")
	}

	startBlockNum := num - 1
	if latestBlock.BlockNum > startBlockNum {
		startBlockNum = latestBlock.BlockNum
	}

	go func(latestBlock int64, dbBlock int64) {
		if latestBlock > startBlockNum {
			for i := startBlockNum + 1; i <= latestBlock; i++ {
				c.blockNumChain <- big.NewInt(i)
			}
		}

		firstBlock, err := c.store.GetFirstBlock(ctx)
		fmt.Println(firstBlock.BlockNum)

		if err != nil {
			log.Println(err, "GetFirstBlock")
		}

		var minNum int64

		if firstBlock.BlockNum-chunk > 0 {
			minNum = firstBlock.BlockNum - chunk
		} else {
			minNum = 0
		}

		if firstBlock.BlockNum == 0 {
			return
		}

		for i := firstBlock.BlockNum - 1; i >= minNum; i-- {
			c.blockNumChain <- big.NewInt(i)
		}
	}(block.Number().Int64(), latestBlock.BlockNum)
}

func (c *Consumer) saveBlockProcess(blockNum *big.Int) {
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

	_, err := c.store.CreateBlock(ctx, arg)

	if err != nil {
		log.Println(err, "CreateBlock")
	}

	transactions := block.Body().Transactions

	go func(transactions []*types.Transaction) {
		for _, tx := range transactions {
			fmt.Printf("txChain Tx = %v\n", tx.Hash().Hex())
			c.txChain <- tx
		}
	}(transactions)
}

func (c *Consumer) saveTxProcess(tx *types.Transaction) {
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
		Gas:       int64(receipt.GasUsed),
		TxIndex:   receipt.TxHash.Big().Int64(),
	}

	if tx.To() != nil {
		arg.To = tx.To().Hex()
	}

	if tx.Value() != nil {
		arg.Value = tx.Value().Int64()
	}

	_, err = c.store.CreateTransaction(ctx, arg)
	if err != nil {
		log.Println(err, "CreateTransaction")
	}

	logs := receipt.Logs

	go func(logs []*types.Log) {
		for _, lg := range logs {
			c.logChain <- lg
		}
	}(logs)
}

func (c *Consumer) saveLogProcess(rlog *types.Log) {
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
		LogIndex:  int64(rlog.Index),
	}

	_, err = c.store.CreateLog(ctx, arg)
	if err != nil {
		log.Println(err, "CreateLog")
	}
}

func (c *Consumer) worker(ctx context.Context, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("start the worker", num)
	for {
		select {
		case blockNum := <-c.blockNumChain:
			if ctx.Err() != nil {
				log.Println("get next job blockNum", blockNum, "and close the worker", num)
				return
			}

			c.saveBlockProcess(blockNum)

			fmt.Printf("worker %v: ", num)

		case tx := <-c.txChain:
			if ctx.Err() != nil {
				log.Println("get next job tx", tx.Hash(), "and close the worker", num)
				return
			}
			fmt.Printf("worker %v: ", num)
			c.saveTxProcess(tx)
		case rlog := <-c.logChain:
			if ctx.Err() != nil {
				log.Println("get next job rlog", rlog.Address.Hash(), "and close the worker", num)
				return
			}
			fmt.Printf("worker %v: ", num)
			c.saveLogProcess(rlog)

		case <-ctx.Done():
			log.Println("close the worker", num)
			return
		default:
			if num == 0 {
				log.Println("defautl", num)
				c.startTraceBlocks(startNum)
			}
		}
	}
}

func RunCrawl(store *db.SQLStore) {

	finished := make(chan bool)
	wg := &sync.WaitGroup{}
	wg.Add(poolSize)

	consumer := Consumer{
		jobsChan:      make(chan int, poolSize),
		blockNumChain: make(chan *big.Int),
		txChain:       make(chan *types.Transaction),
		logChain:      make(chan *types.Log),
		store:         store,
	}

	ctx := context.Background()

	for i := 0; i < poolSize; i++ {
		go consumer.worker(ctx, i, wg)
	}

	// consumer.startTraceBlocks()

	<-finished
	log.Println("Done")
}
