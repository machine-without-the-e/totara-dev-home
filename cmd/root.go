package cmd

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
  Use:   "td",
  Short: "Totara Dev Home is a toolkit for developing Totara",
  Long: `
  Totara Dev Home :)
  `,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}

func Execute() {
  if err := RootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
