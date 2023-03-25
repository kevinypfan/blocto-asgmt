-- name: CreateTransaction :one
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
) RETURNING *;

-- name: GetTransactionByHash :one
SELECT * FROM transactions
WHERE transaction_hash = $1 LIMIT 1;

-- name: ListTransactionsByBlockNumber :many
SELECT * FROM transactions
WHERE block_number = $1;

-- name: ListTransactionsByBlockHash :many
SELECT * FROM transactions
WHERE block_hash = $1;