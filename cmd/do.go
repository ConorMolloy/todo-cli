package cmd

import (
	"fmt"
	"strconv"

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
		fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
