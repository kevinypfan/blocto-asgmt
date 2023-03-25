package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/kevinypfan/blocto-asgmt/db/sqlc"
)

type Server struct {
	router *gin.Engine
	store  *db.SQLStore
}

func NewServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/ping", server.ping)
	router.GET("/block/:id", server.getBlockById)
	router.GET("/blocks", server.getListBlocks)
	router.GET("/transaction/:txHash", server.getTransactionByHash)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
