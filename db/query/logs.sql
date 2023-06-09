-- name: CreateLog :one
INSERT INTO logs (
    "address",
    "topics",
    "block_num",
    "tx_hash",
    "block_hash",
    "removed",
    "log_index"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: ListLogsByTransactionHash :many
SELECT * FROM logs
WHERE tx_hash = $1 ORDER BY log_index asc;
