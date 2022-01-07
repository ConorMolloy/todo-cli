package cmd

import (
	"fmt"
	"os"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "adds a task to your task manager",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Could not create task:", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Added: \"%s\"\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCommand)
}