package database

import (
  "context"
  "github.com/jackc/pgx/v5"
  "fmt"
  "os"
)

type Pgsql struct {
  config DatabaseConfig
  connection Connection
}

func (db *Pgsql) connect() {
  conn, err := pgx.Connect(
    context.Background(), 
    fmt.Sprintf("postgres://%s:%s@%s",
      db.config.User,
      db.config.Pass,
      db.config.Host,
    ),
  )
    
if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
defer conn.Close(context.Background())

var id int
	err = conn.QueryRow(context.Background(), "select id from public.ttr_course").Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

  fmt.Println(id)
}

func (db *Pgsql) getConnection() Connection {
  return db.connection
}

func (db *Pgsql) getConfig() DatabaseConfig {
  return db.config
}

func (db *Pgsql) listDatabases() []string {
  return []string{"a", "b"}
}

func (db *Pgsql) createDatabase(name string) {
}
