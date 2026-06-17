package postgres

import (
	"database/sql"
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/omeid/conex"
)

var (
	// Image to use for the box.
	Image = "postgres:alpine"
	// Port used for connect to postgres server.
	Port = "5432"

	// PostgresUpWaitTime dictates how long we should wait for Postgresql to accept connections on {{Port}}.
	PostgresUpWaitTime = 10 * time.Second
)

func init() {
	conex.Require(func() string { return Image })
}

// Config used to connect to the database.
type Config struct {
	User     string
	Password string
	Database string // defaults to `postgres` as service db.

	host string
	port string
}

func (c *Config) url() string {

	if c.Database == "" {
		c.Database = "postgres"
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.User, c.Password, c.host, c.port, c.Database,
	)
}

// Box returns an sql.DB connection and the container running the Postgresql
// instance. It will call t.Fatal on errors.
func Box(t testing.TB, config *Config) (*sql.DB, conex.Container) {
	if !slices.Contains(sql.Drivers(), "postgres") {
		t.Fatal("No SQL driver registered for postgres. Did you forget to import one? (e.g., _ \"github.com/lib/pq\")")
	}

	if config == nil {
		config = &Config{
			Database: "postgres",
			User:     "postgres",
		}
	}

	c := conex.Box(t, &conex.Config{
		Image:  Image,
		Expose: []string{Port},
		Env:    []string{"POSTGRES_HOST_AUTH_METHOD=trust"},
	})

	config.host = c.Address()
	config.port = Port

	conex.Logf(t, "postgres", "Waiting for Postgresql to accept connections")

	err := c.Wait(Port, PostgresUpWaitTime)

	if err != nil {
		c.Drop()
		t.Fatal("Postgres failed to start:", err)
	}

	conex.Logf(t, "postgres", "Postgresql is now accepting connections")
	db, err := sql.Open("postgres", config.url())

	if err != nil {
		c.Drop()
		t.Fatal(err)
	}

	return db, c
}
