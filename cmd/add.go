package cmd

import (
	"fmt"
	"log"

	"github.com/mateusoliveira43/go_cli/todo"
	"github.com/spf13/cobra"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new To Do",
	Long:  "Add will create a new To Do item to the list",
	Run:   addRun,
	// TODO does not appear in help
	// https://stackoverflow.com/questions/69167939/named-positional-arguments-in-cobra
	// https://cobra.dev/#defining-your-own-usage
	Args: cobra.MinimumNArgs(1),
}

// Add each argument to the list
//
// # Parameters
//
// ## args : `[]string`
//
//	Items to be added to the list
func addRun(cmd *cobra.Command, args []string) {
	fmt.Println("add called")
	// var items = []todo.Item{}
	items, err := todo.ReadItems(dataFile)
	// items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		// log.Printf("%v", err)
		log.Println("Creating Data Base")
	}
	currentId := len(items) + 1
	for index, x := range args {
		fmt.Println(x)
		item := todo.ToDoItem{Index: currentId + index, Description: x}
		item.SetPriority(priority)
		items = append(items, item)
	}
	fmt.Println(items)
	fmt.Printf("%#v\n", items)
	err = todo.SaveItems(dataFile, items)
	// err = todo.SaveItems(viper.GetString("datafile"), items)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "To Do priority, one of: 1 (High), 2 (Normal), 3 (Low)")
}
