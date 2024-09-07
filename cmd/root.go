package cmd

import (
	"fmt"
	"gtd/data"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
)

var store data.Store

var root = &cobra.Command{
	Use:   "gtd",
	Short: "Get some shit done",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := store.Inbox()
		size := len(tasks)
		var opt string

		for i := 0; i < size; {
			task := tasks[i]
			fmt.Println("[a - Actionable | s - Someday | t - Tickler]")
			fmt.Printf("[ ] %s: ", task.Name)
			_, err := fmt.Scan(&opt)

			if err != nil {
				fmt.Printf("Error on rune - %v", err)
				os.Exit(1)
			}

			switch opt {
			case "a", "A":
				fmt.Println("aA")
				err = store.Move(task.Id, data.ActionTask)
			case "s", "S":
				fmt.Println("sS")
				err = store.Move(task.Id, data.SomedayTask)
			case "t", "T":
				fmt.Println("tT")
				err = store.Move(task.Id, data.TicklerTask)
			default:
				fmt.Println("Try again")
				continue
			}

			if err != nil {
				fmt.Printf("Error moving task %d - %v", task.Id, err)
				os.Exit(1)
			}

			i++
		}
	},
}

func Execute() {
	dbUrl := os.Getenv("DB_URL")
	store = data.NewSqliteStore(dbUrl)
	err := root.Execute()
	if err != nil {
		fmt.Printf("Error running application: %v", err)
		os.Exit(1)
	}
}
