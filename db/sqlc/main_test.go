package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	DB_DRIVER = "postgres"
	DB_SOURCE = "postgresql://bank_owner:npg_ngiKW0RH9XYw@ep-lucky-truth-a8wth6db-pooler.eastus2.azure.neon.tech/bank?sslmode=require"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(DB_DRIVER, DB_SOURCE)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
