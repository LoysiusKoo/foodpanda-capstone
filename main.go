package main

import (
	"database/sql"
	"log"

	"github.com/jichong-tay/foodpanda-playlist-api/api"
	db "github.com/loysiuskoo/foodpanda-capstone/database/sqlc"

	_ "github.com/lib/pq"
	"github.com/loysiuskoo/foodpanda-capstone/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
