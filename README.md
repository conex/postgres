# Postgres [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/conex/postgres)  [![Go Report Card](https://goreportcard.com/badge/github.com/conex/postgres)](https://goreportcard.com/report/github.com/conex/postgres)

This package provides a Postgres Box using Conex.

## Usage

```go
package example_test

import (
  "testing"

  "github.com/omeid/conex"
  "github.com/conex/postgres"
  _ "github.com/lib/pq" // Bring your own driver!
)

func TestMain(m *testing.M) {
  conex.Main(m)
}

func TestPostgres(t *testing.T) {
  db, container := postgres.Box(t, nil)
  defer container.Drop()

  // use db to interact with the database
}
```

*Note: You must import your preferred sql driver (e.g., `_ "github.com/lib/pq"`) in your tests, as the plugin does not register one automatically.*
