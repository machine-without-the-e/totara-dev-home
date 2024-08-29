package database

import "errors"


type Connection struct {
  connection any
}

type Database interface {
  getConfig() DatabaseConfig
  getConnection() Connection 
  listDatabases() []string
  createDatabase(name string)
}

func GetDatabase(dbtype string, config DatabaseConfig) (Database, error) {
  switch dbtype {
    case "pgsql":
      return &Pgsql{
        config: config,
      }, nil
    
  }

  return nil, errors.New("unknown database type")
}

