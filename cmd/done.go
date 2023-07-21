package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mateusoliveira43/go_cli/todo"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "Mark To Do Item as Done",
	Long:    "Mark a To Do Item as Done using its index",
	Run:     doneRun,
	// TODO does not appear in help
	Args: cobra.ExactArgs(1),
}

func doneRun(cmd *cobra.Command, args []string) {
	fmt.Println("done called")
	items, err := todo.ReadItems(dataFile)
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid To Do item Index.")
	}
	if i > 0 && i < len(items) {
		items[i-1].Done = true
		fmt.Printf("%q %v\n", items[i-1].Description, "marked done")
		todo.SaveItems(dataFile, items)
	} else {
		log.Fatalln("Index", i, "does not match any To Do item.")
	}
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
