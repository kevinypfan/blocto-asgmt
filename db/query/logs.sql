-- name: CreateLog :one
INSERT INTO logs (
    "address",
    "topics",
    "data",
    "block_number",
    "transaction_hash",
    "transaction_index",
    "block_hash",
    "log_index",
    "removed"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: ListLogsByTransactionHash :many
SELECT * FROM logs
WHERE transaction_hash = $1;

-- name: ListLogsByTransactionIndex :many
SELECT * FROM logs
WHERE transaction_index = $1;