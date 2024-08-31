package database

import (
  "errors"
  "totara_dev_home/internal/database/types"
  "totara_dev_home/internal/database/config"
)

type Database interface {
  Connect(config config.DatabaseConfig) error
  ListDatabases() ([]string, error)
  CreateDatabase(name string) error
}

func GetDatabase(dbtype string) (Database, error) {
  db, err := getDatabase(dbtype)

  if (err != nil) {
    return nil, err
  }

  config, err := config.GetConfig(dbtype);

  if (err != nil) {
    return nil, err
  }

  err = db.Connect(config)

  if (err != nil) {
    return nil, err
  }
  
  return db, nil
}

func getDatabase(dbtype string) (Database, error) {
  switch dbtype {
    case "pgsql15":
      return &types.Pgsql15{}, nil
    
  }

  return nil, errors.New("unknown database type")
}
