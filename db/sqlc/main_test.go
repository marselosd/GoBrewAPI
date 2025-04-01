package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_"github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secretpwd@localhost:5432/go_brew?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error 
	
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("ERROR: cannot connect to database")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}