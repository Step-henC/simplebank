package db

import (
	"database/sql" //must use with add'l driver
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" //must keep empty identifier of driver so go keeps it on save
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	testQueries = New(conn)
	m.Run()

	os.Exit(m.Run())
}
