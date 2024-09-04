package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add NAME",
	Short: "Add new tasks",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		err := store.Add(name)
		if err != nil {
			fmt.Println("Task not created")
			fmt.Println(err)
		}
	},
}
