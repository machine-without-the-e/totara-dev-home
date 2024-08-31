package dbCmd

import (
  "fmt"
  "totara_dev_home/internal/database/config"
  "totara_dev_home/internal/database"
  "github.com/spf13/cobra"
)


var ListCmd = &cobra.Command{
  Use:   "list",
  Short: "Lists databases",
  Long:  `Lists databases`,
  Run: func(cmd *cobra.Command, args []string) {
    config := config.DatabaseConfig{
      Host: "0.0.0.0:5444",
      User: "postgres",
      Pass: "",
      Database: "learn_evergreen-dev_gb_566e5c33b7df67f7d0552a2f0a2a426ba8fe9470",
    }

    numberTheList, err := cmd.Flags().GetBool("numbered-list")

    if (err != nil) {
      fmt.Printf("Issue with cobra. Error: \n %v", err)
    }

    fmt.Printf("args: %v", args)

    db, err := database.GetDatabase("pgsql15", config)

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
