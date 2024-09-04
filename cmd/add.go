package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(args)
		// err := store.Add("")
		// if err != nil {
		// 	fmt.Print("Task not created")
		// }
	},
}
