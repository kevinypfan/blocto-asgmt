package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
