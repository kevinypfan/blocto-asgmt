package main

import (
	"database/sql"
	"log"

	"github.com/kevinypfan/blocto-asgmt/api"
	db "github.com/kevinypfan/blocto-asgmt/db/sqlc"
	"github.com/kevinypfan/blocto-asgmt/web3"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/blocto-asgmt?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

// type RPCResponse struct {
// 	Name   string `json:"name"`
// 	Values []struct {
// 		Value    int `json:"value,omitempty"`
// 		Comments int `json:"comments,omitempty"`
// 		Likes    int `json:"likes,omitempty"`
// 		Shares   int `json:"shares,omitempty"`
// 	} `json:"values"`
// }

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	go web3.RunCrawl(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
