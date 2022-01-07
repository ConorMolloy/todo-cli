package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "adds a task to your task manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	RootCmd.AddCommand(addCommand)
}