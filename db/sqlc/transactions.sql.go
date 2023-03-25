// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: transactions.sql

package db

import (
	"context"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (
  "transaction_hash",
  "block_hash",
  "block_number",
  "from",
  "gas",
  "gas_price",
  "input",
  "nonce",
  "to",
  "transaction_index",
  "value",
  "type",
  "v",
  "r",
  "s"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
) RETURNING transaction_hash, block_hash, block_number, "from", gas, gas_price, input, nonce, "to", transaction_index, value, type, v, r, s
`

type CreateTransactionParams struct {
	TransactionHash  string `json:"transaction_hash"`
	BlockHash        string `json:"block_hash"`
	BlockNumber      string `json:"block_number"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gas_price"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transaction_index"`
	Value            string `json:"value"`
	Type             string `json:"type"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createTransaction,
		arg.TransactionHash,
		arg.BlockHash,
		arg.BlockNumber,
		arg.From,
		arg.Gas,
		arg.GasPrice,
		arg.Input,
		arg.Nonce,
		arg.To,
		arg.TransactionIndex,
		arg.Value,
		arg.Type,
		arg.V,
		arg.R,
		arg.S,
	)
	var i Transaction
	err := row.Scan(
		&i.TransactionHash,
		&i.BlockHash,
		&i.BlockNumber,
		&i.From,
		&i.Gas,
		&i.GasPrice,
		&i.Input,
		&i.Nonce,
		&i.To,
		&i.TransactionIndex,
		&i.Value,
		&i.Type,
		&i.V,
		&i.R,
		&i.S,
	)
	return i, err
}

const getTransactionByHash = `-- name: GetTransactionByHash :one
SELECT transaction_hash, block_hash, block_number, "from", gas, gas_price, input, nonce, "to", transaction_index, value, type, v, r, s FROM transactions
WHERE transaction_hash = $1 LIMIT 1
`

func (q *Queries) GetTransactionByHash(ctx context.Context, transactionHash string) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, getTransactionByHash, transactionHash)
	var i Transaction
	err := row.Scan(
		&i.TransactionHash,
		&i.BlockHash,
		&i.BlockNumber,
		&i.From,
		&i.Gas,
		&i.GasPrice,
		&i.Input,
		&i.Nonce,
		&i.To,
		&i.TransactionIndex,
		&i.Value,
		&i.Type,
		&i.V,
		&i.R,
		&i.S,
	)
	return i, err
}

const listTransactionsByBlockHash = `-- name: ListTransactionsByBlockHash :many
SELECT transaction_hash, block_hash, block_number, "from", gas, gas_price, input, nonce, "to", transaction_index, value, type, v, r, s FROM transactions
WHERE block_hash = $1
`

func (q *Queries) ListTransactionsByBlockHash(ctx context.Context, blockHash string) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, listTransactionsByBlockHash, blockHash)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.TransactionHash,
			&i.BlockHash,
			&i.BlockNumber,
			&i.From,
			&i.Gas,
			&i.GasPrice,
			&i.Input,
			&i.Nonce,
			&i.To,
			&i.TransactionIndex,
			&i.Value,
			&i.Type,
			&i.V,
			&i.R,
			&i.S,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTransactionsByBlockNumber = `-- name: ListTransactionsByBlockNumber :many
SELECT transaction_hash, block_hash, block_number, "from", gas, gas_price, input, nonce, "to", transaction_index, value, type, v, r, s FROM transactions
WHERE block_number = $1
`

func (q *Queries) ListTransactionsByBlockNumber(ctx context.Context, blockNumber string) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, listTransactionsByBlockNumber, blockNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.TransactionHash,
			&i.BlockHash,
			&i.BlockNumber,
			&i.From,
			&i.Gas,
			&i.GasPrice,
			&i.Input,
			&i.Nonce,
			&i.To,
			&i.TransactionIndex,
			&i.Value,
			&i.Type,
			&i.V,
			&i.R,
			&i.S,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
