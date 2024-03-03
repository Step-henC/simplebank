package db

import (
	"database/sql" //must use with add'l driver
	"log"
	"os"
	"testing"

	"github.com/Step-henC/simplebank/db/util"
	_ "github.com/lib/pq" //must keep empty identifier of driver so go keeps it on save
)

var testQueries *Queries
var testDB *sql.DB //created for store transactions class

func TestMain(m *testing.M) {
	config, prob := util.LoadConfig("../..")
	if prob != nil {
		log.Fatal("Cannot load config", prob)
	}

	var err error
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	testQueries = New(testDB)
	m.Run()

	os.Exit(m.Run())
}
