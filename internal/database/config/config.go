package config

import(
  "github.com/goccy/go-yaml"
  "os"
  "fmt"
)

type DatabaseConfig struct {
  Host string `yaml:"host"`
  User string `yaml:"user"`
  Pass string `yaml:"pass"`
  Database string `yaml:"database"`
}

func GetConfig(dbtype string) (DatabaseConfig, error) {
  file, err := os.ReadFile(fmt.Sprintf("config/database/types/%s.yaml", dbtype))
  
  if (err != nil) {
    return DatabaseConfig{}, err
  }

  config := DatabaseConfig{}

  err = yaml.Unmarshal(file, &config)

  if (err != nil) {
    return DatabaseConfig{}, err
  }

  return config, nil
}
