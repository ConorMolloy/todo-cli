package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Completes a task",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			intArg, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("%s must be an integer", arg)
			}
			ids = append(ids, intArg)
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}

		for _, id := range ids {
			if id <= 0 || id > len(ids) {
				fmt.Println("Invalid task number")
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Println("Failed to delete a task")
			} else {
				fmt.Println("Task completed")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
