package dbCmd

import (
  "fmt"

  commands "totara_dev_home/cmd/db/cmd"

  "totara_dev_home/cmd"
  "github.com/spf13/cobra"
)

func init() {
  cmd.RootCmd.AddCommand(Cmd);

  Cmd.AddCommand(commands.CreateCmd)
  Cmd.AddCommand(commands.ListCmd)
}

var Cmd = &cobra.Command{
  Use:   "db",
  Short: "Handles database operations in Totara",
  Long:  `Handles database operations in Totara`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Totara Database operations")
  },
}
