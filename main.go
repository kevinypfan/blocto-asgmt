package main

import (
	"database/sql"
	"log"

	"github.com/kevinypfan/blocto-asgmt/api"
	db "github.com/kevinypfan/blocto-asgmt/db/sqlc"
	"github.com/kevinypfan/blocto-asgmt/util"
	"github.com/kevinypfan/blocto-asgmt/web3"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(config, store)

	go web3.RunCrawl(config, store)

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
