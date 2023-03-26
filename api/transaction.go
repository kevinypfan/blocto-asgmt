package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kevinypfan/blocto-asgmt/db/sqlc"
)

type transactionByHashRequest struct {
	TxHash string `uri:"txHash" binding:"required"`
}

type transactionByHashResponse struct {
	TxHash    string   `json:"tx_hash"`
	BlockHash string   `json:"block_hash"`
	BlockNum  int64    `json:"block_num"`
	From      string   `json:"from"`
	To        string   `json:"to"`
	Nonce     int64    `json:"nonce"`
	Value     int64    `json:"value"`
	Gas       int64    `json:"gas"`
	TxIndex   int64    `json:"tx_index"`
	Logs      []db.Log `json:"logs"`
}

func (server *Server) getTransactionByHash(ctx *gin.Context) {
	var req transactionByHashRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	transaction, err := server.store.GetTransactionByHash(ctx, req.TxHash)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	logs, err := server.store.ListLogsByTransactionHash(ctx, req.TxHash)

	transactionByHashResponse := transactionByHashResponse{
		TxHash:    transaction.TxHash,
		BlockHash: transaction.BlockHash,
		BlockNum:  transaction.BlockNum,
		From:      transaction.From,
		To:        transaction.To,
		Nonce:     transaction.Nonce,
		Value:     transaction.Value,
		Gas:       transaction.Gas,
		TxIndex:   transaction.TxIndex,
		Logs:      logs,
	}

	ctx.JSON(http.StatusOK, transactionByHashResponse)
}
