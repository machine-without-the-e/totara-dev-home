package database

type Pgsql struct {
  config DatabaseConfig
  connection Connection
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
