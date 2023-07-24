package cmd

import (
	"fmt"
	"os"

	"github.com/mateusoliveira43/go_cli/pkg/database"
	"github.com/mateusoliveira43/go_cli/pkg/util"
	"github.com/spf13/cobra"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add TODO [...TODO]",
	Short:   "Add new To Do items",
	Long:    "Add new To Do items to the database",
	Example: "  go_cli add test",
	Args:    cobra.MinimumNArgs(1),
	PreRun:  DebugFlagsValues,
	Run:     addRun,
}

// Add each argument to the list
//
// # Parameters
//
// ## args : `[]string`
//
//	Items to be added to the list
func addRun(cmd *cobra.Command, args []string) {
	if debug {
		util.Debug("Add command called")
	}
	var items = []database.ToDoItem{}
	if _, err := os.ReadFile(dataFile); err != nil {
		util.Warn(fmt.Sprintf("No database found for %v, creating it", dataFile))
	} else {
		items = database.LoadItems(dataFile, debug)
		// items, err := todo.ReadItems(viper.GetString("datafile"))
	}
	currentIndex := len(items) + 1
	for index, arg := range args {
		if debug {
			util.Debug(fmt.Sprintf("Adding %q to To Do list", arg))
		}
		item := database.ToDoItem{Index: currentIndex + index, Description: arg}
		item.SetPriority(priority)
		items = append(items, item)
	}
	database.SaveItems(dataFile, items, debug)
	// err = todo.SaveItems(viper.GetString("datafile"), items)
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "To Do priority, one of: 1 (High), 2 (Normal), 3 (Low)")
}
