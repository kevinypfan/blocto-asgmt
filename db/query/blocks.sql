-- name: CreateBlock :one
INSERT INTO blocks (
  "block_id",
  "block_number",
  "block_hash",
  "difficulty",
  "extra_data",
  "gas_limit",
  "gas_used",
  "logs_bloom",
  "miner",
  "mix_hash",
  "nonce",
  "parent_hash",
  "receipts_root",
  "sha3_uncles",
  "size",
  "state_root",
  "timestamp",
  "total_difficulty",
  "transactions_root"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19
) RETURNING *;

-- name: GetBlockByNumber :one
SELECT * FROM blocks
WHERE block_number = $1 LIMIT 1;

-- name: GetBlockByHash :one
SELECT * FROM blocks
WHERE block_hash = $1 LIMIT 1;

-- name: GetBlockById :one
SELECT * FROM blocks
WHERE block_id = $1 LIMIT 1;

-- name: ListBlocks :many
SELECT * FROM blocks
ORDER BY block_id desc
LIMIT $1
OFFSET $2;