-- name: CreateBlock :one
INSERT INTO blocks (
  "block_num",
  "block_hash",
  "block_time",
  "parent_hash"
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetBlockByNumber :one
SELECT * FROM blocks
WHERE block_num = $1 LIMIT 1;

-- name: GetBlockByHash :one
SELECT * FROM blocks
WHERE block_hash = $1 LIMIT 1;

-- name: GetLatestBlock :one
SELECT * FROM blocks
ORDER BY block_num desc LIMIT 1;

-- name: ListBlocks :many
SELECT * FROM blocks
ORDER BY block_num desc
LIMIT $1
OFFSET $2;