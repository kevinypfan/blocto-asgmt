package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionByHashRequest struct {
	TxHash string `uri:"txHash" binding:"required"`
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

	ctx.JSON(http.StatusOK, transaction)
}
