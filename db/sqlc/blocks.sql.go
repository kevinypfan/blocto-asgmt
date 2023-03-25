// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: blocks.sql

package db

import (
	"context"
)

const createBlock = `-- name: CreateBlock :one
INSERT INTO blocks (
  "block_num",
  "block_hash",
  "block_time",
  "parent_hash"
) VALUES (
  $1, $2, $3, $4
) RETURNING block_num, block_hash, block_time, parent_hash
`

type CreateBlockParams struct {
	BlockNum   int64  `json:"block_num"`
	BlockHash  string `json:"block_hash"`
	BlockTime  int64  `json:"block_time"`
	ParentHash string `json:"parent_hash"`
}

func (q *Queries) CreateBlock(ctx context.Context, arg CreateBlockParams) (Block, error) {
	row := q.db.QueryRowContext(ctx, createBlock,
		arg.BlockNum,
		arg.BlockHash,
		arg.BlockTime,
		arg.ParentHash,
	)
	var i Block
	err := row.Scan(
		&i.BlockNum,
		&i.BlockHash,
		&i.BlockTime,
		&i.ParentHash,
	)
	return i, err
}

const getBlockByHash = `-- name: GetBlockByHash :one
SELECT block_num, block_hash, block_time, parent_hash FROM blocks
WHERE block_hash = $1 LIMIT 1
`

func (q *Queries) GetBlockByHash(ctx context.Context, blockHash string) (Block, error) {
	row := q.db.QueryRowContext(ctx, getBlockByHash, blockHash)
	var i Block
	err := row.Scan(
		&i.BlockNum,
		&i.BlockHash,
		&i.BlockTime,
		&i.ParentHash,
	)
	return i, err
}

const getBlockByNumber = `-- name: GetBlockByNumber :one
SELECT block_num, block_hash, block_time, parent_hash FROM blocks
WHERE block_num = $1 LIMIT 1
`

func (q *Queries) GetBlockByNumber(ctx context.Context, blockNum int64) (Block, error) {
	row := q.db.QueryRowContext(ctx, getBlockByNumber, blockNum)
	var i Block
	err := row.Scan(
		&i.BlockNum,
		&i.BlockHash,
		&i.BlockTime,
		&i.ParentHash,
	)
	return i, err
}

const listBlocks = `-- name: ListBlocks :many
SELECT block_num, block_hash, block_time, parent_hash FROM blocks
ORDER BY block_id desc
LIMIT $1
OFFSET $2
`

type ListBlocksParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListBlocks(ctx context.Context, arg ListBlocksParams) ([]Block, error) {
	rows, err := q.db.QueryContext(ctx, listBlocks, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Block{}
	for rows.Next() {
		var i Block
		if err := rows.Scan(
			&i.BlockNum,
			&i.BlockHash,
			&i.BlockTime,
			&i.ParentHash,
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
