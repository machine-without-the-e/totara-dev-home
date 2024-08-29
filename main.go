package main

import (
  "totara_dev_home/cmd"
  _ "totara_dev_home/cmd/db"
  _ "totara_dev_home/internal/database"
)

func main() {
  cmd.Execute()
}
