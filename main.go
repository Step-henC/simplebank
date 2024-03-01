package main

import (
	"database/sql"
	"log"

	"github.com/Step-henC/simplebank/api"
	db "github.com/Step-henC/simplebank/db/sqlc"
	_ "github.com/lib/pq" //driver to talk to db
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
