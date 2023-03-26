package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	db "github.com/kevinypfan/blocto-asgmt/db/sqlc"
	"github.com/kevinypfan/blocto-asgmt/util"
	web3 "github.com/kevinypfan/blocto-asgmt/web3"
	"github.com/segmentio/kafka-go"
)

func newKafkaReader(kafkaURL, topic string, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		GroupID:  groupID,
	})
}

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	fmt.Println(config.KafkaBrokers)

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	ctx := context.Background()

	store := db.NewStore(conn)

	reader := newKafkaReader(config.KafkaBrokers, config.KafkaTxTopic, config.KafkaGroup)

	defer reader.Close()

	fmt.Println("start consuming ... !!")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("message at topic: %v partition: %v offset: %v value: %s\n", m.Topic, m.Partition, m.Offset, string(m.Value))

		txHash := common.HexToHash(string(m.Value))

		fmt.Printf("Save Tx = %v\n", txHash.Hex())

		tx := web3.GetTransactionByHash(txHash)
		receipt := web3.GetTransactionReceipt(txHash)

		from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
		if err != nil {
			log.Println(err, "Sender")
		}

		// fmt.Println(tx.Hash().Hex())

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

		_, err = store.CreateTransaction(ctx, arg)
		if err != nil {
			log.Println(err, "CreateTransaction")
		}

		logs := receipt.Logs

		for _, rlog := range logs {
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

			_, err = store.CreateLog(ctx, arg)
			if err != nil {
				log.Println(err, "CreateLog")
			}
		}
	}
}
