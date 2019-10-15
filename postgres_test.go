package postgres_test

import (
	"os"
	"testing"
	"time"

	"github.com/conex/postgres"
	"github.com/omeid/conex"
)

func TestMain(m *testing.M) {
	os.Exit(conex.Run(m))
}

func init() {
	postgres.PostgresUpWaitTime = 20 * time.Second
}

func TestPostgres(t *testing.T) {

	sql, con := postgres.Box(t, &postgres.Config{
		Database: "postgres",
		User:     "postgres",
	})
	defer con.Drop()

	var resp int
	err := sql.QueryRow("SELECT 1").Scan(&resp)

	if err != nil {
		t.Fatal(err)
	}

	if resp != 1 {
		t.Fatalf("Unexpected response: %v", resp)
	}

}

func TestPostgresNilConf(t *testing.T) {

	sql, con := postgres.Box(t, nil)
	defer con.Drop()

	var resp int
	err := sql.QueryRow("SELECT 1").Scan(&resp)

	if err != nil {
		t.Fatal(err)
	}

	if resp != 1 {
		t.Fatalf("Unexpected response: %v", resp)
	}

}
