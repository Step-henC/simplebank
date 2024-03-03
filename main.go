package main

import (
	"database/sql"
	"log"

	"github.com/Step-henC/simplebank/api"
	db "github.com/Step-henC/simplebank/db/sqlc"
	"github.com/Step-henC/simplebank/db/util"
	_ "github.com/lib/pq" //driver to talk to db
)

func main() {

	config, err := util.LoadConfig(".") //path in curr dir hence the dot

	if err != nil {
		log.Fatal("Cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
