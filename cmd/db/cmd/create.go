package dbCmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
  Use:   "create",
  Short: "Creates a database",
  Long:  `All software has versions. This is Hugo's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
  },
}
