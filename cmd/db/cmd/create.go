package dbCmd

import (
 // "fmt"

  //"totara_dev_home/internal/database"
  "github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
  Use:   "create",
  Short: "Creates a database",
  Long:  `Creates databases within the currently used database`,
  Run: func(cmd *cobra.Command, args []string) {

  },
}
