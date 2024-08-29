package dbCmd

import (
  "fmt"
  "totara_dev_home/internal/database"
  "github.com/spf13/cobra"
)


var ListCmd = &cobra.Command{
  Use:   "list",
  Short: "Lists databases",
  Long:  `Lists databases`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
    config := database.DatabaseConfig{
      Host: "localhost",
      User: "postgres",
      Pass: "",
      Database: "abc",
    }


    database.GetDatabase("pgsql", config)
  },
}
