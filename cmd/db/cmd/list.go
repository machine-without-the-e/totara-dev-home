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
    numberTheList, err := cmd.Flags().GetBool("numbered-list")

    if (err != nil) {
      fmt.Printf("Issue with cobra. Error: \n %v", err)
    }

    fmt.Printf("args: %v", args)

    db, err := database.GetDatabase("pgsql15")

    if (err != nil) {
      fmt.Printf("Issue initalising database. Error: \n %v", err)
      return
    }

    databases, err := db.ListDatabases()

    for number, database := range databases {
      if (numberTheList) {
        print(fmt.Sprintf("%d - %v\n", number, database))
        continue
      }

      print(fmt.Sprintf("%v\n", database))
    }
  },
}

func init() {
  ListCmd.Flags().BoolP(
    "numbered-list",
    "n",
    false,
    "Formats the list of databases as a numbered list.",
  )
}
