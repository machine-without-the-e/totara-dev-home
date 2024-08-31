package types

import (
  "context"
  "totara_dev_home/internal/database/config"
  "github.com/jackc/pgx/v5"
  "fmt"
)

type Pgsql15 struct {
  connection pgx.Conn
}

func (db *Pgsql15) Connect(config config.DatabaseConfig) error {
  conn, err := pgx.Connect(
    context.Background(), 
    fmt.Sprintf("postgres://%s:%s@%s/%s",
      config.User,
      config.Pass,
      config.Host,
      config.Database,
    ),
  )

  if (err != nil) {
    return err
  } 

  db.connection = *conn   
  return nil
}

func (db *Pgsql15) ListDatabases() ([]string, error) {
  rows, err := db.connection.Query(context.Background(), "SELECT datname FROM pg_database WHERE datistemplate = false;")
  
  if (err != nil) {
    return nil, err
  }

  databases, err := pgx.CollectRows(rows, pgx.RowTo[string])

  if (err != nil) {
    return nil, err
  }

  return databases, nil
}

func (db *Pgsql15) CreateDatabase(name string) error {
  _, err := db.connection.Exec(context.Background(), "CREATE DATABASE IF NOT EXISTS $1;", name)
  
  if (err != nil) {
    return err
  }

  return nil
} 
