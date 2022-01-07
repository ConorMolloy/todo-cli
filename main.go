package main

import (
	"fmt"
	"os"
	"path/filepath"
	"task/cmd"
	"task/db"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Printf("Failed to find the home directary")
		os.Exit(1)
	}
	dbHome := filepath.Join(home, "tasks.db")
	err = db.Init(dbHome)
	if err != nil {
		fmt.Printf("Filed to initialise the db")
		os.Exit(1)
	}
	cmd.RootCmd.Execute()
}
