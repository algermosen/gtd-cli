package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "gtd",
	Short: "Get some shit done",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println("Here you will process your shit")
}

func Execute() {
	err := root.Execute()
	if err != nil {
		fmt.Printf("Error running application: %v", err)
		os.Exit(1)
	}
}
