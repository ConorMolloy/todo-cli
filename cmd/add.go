package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "adds a task to your task manager",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Added: \"%s\"\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCommand)
}