package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "cosmart",
	Short: "cosmart",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use serve to start a server")
		fmt.Println("Use -h to see the list of command")
	},
}

func Run() {
	rootCommand.AddCommand(serverCommand)
	if err := rootCommand.Execute(); err != nil {
		panic(err)
	}
}
