package sqlc

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	// Initialize the test database connection
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// Ensure the connection is available
	if err := testDB.Ping(); err != nil {
		log.Fatal("cannot ping db:", err)
	}

	// Initialize the test queries using the connected database
	testQueries = New(testDB)

	// Run all the tests
	code := m.Run()

	// Clean up and close the database connection
	if err := testDB.Close(); err != nil {
		log.Fatal("failed to close the database:", err)
	}

	// Exit with the code returned by m.Run()
	os.Exit(code)
}
