-- name: CreateTransaction :one
INSERT INTO transactions (
  "tx_hash",
  "block_hash",
  "block_num",
  "from",
  "to",
  "nonce",
  "value"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetTransactionByHash :one
SELECT * FROM transactions
WHERE tx_hash = $1 LIMIT 1;

-- name: ListTransactionsByBlockNumber :many
SELECT * FROM transactions
WHERE block_num = $1;

-- name: ListTransactionsByBlockHash :many
SELECT * FROM transactions
WHERE block_hash = $1;
