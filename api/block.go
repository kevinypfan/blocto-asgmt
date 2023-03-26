package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kevinypfan/blocto-asgmt/db/sqlc"
)

type listBlocksRequest struct {
	Limit int32 `form:"limit" binding:"required,min=1"`
}

func (server *Server) getListBlocks(ctx *gin.Context) {
	var req listBlocksRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListBlocksParams{
		Limit:  int32(req.Limit),
		Offset: 0,
	}

	accounts, err := server.store.ListBlocks(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)

}

type blockByIdRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type blockByIdResponse struct {
	BlockNum     int64            `json:"block_num"`
	BlockHash    string           `json:"block_hash"`
	BlockTime    int64            `json:"block_time"`
	ParentHash   string           `json:"parent_hash"`
	Transactions []db.Transaction `json:"transactions"`
}

func (server *Server) getBlockById(ctx *gin.Context) {
	var req blockByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	block, err := server.store.GetBlockByNumber(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	transactions, err := server.store.ListTransactionsByBlockNumber(ctx, req.ID)

	blockByIdResponse := blockByIdResponse{
		BlockNum:     block.BlockNum,
		BlockHash:    block.BlockHash,
		BlockTime:    block.BlockTime,
		ParentHash:   block.ParentHash,
		Transactions: transactions,
	}

	ctx.JSON(http.StatusOK, blockByIdResponse)
}
