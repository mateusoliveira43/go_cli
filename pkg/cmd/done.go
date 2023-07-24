package cmd

import (
	"fmt"
	"strconv"

	"github.com/mateusoliveira43/go_cli/pkg/database"
	"github.com/mateusoliveira43/go_cli/pkg/util"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done INDEX [...INDEX]",
	Aliases: []string{"do"},
	Short:   "Mark To Do Items as Done",
	Long:    "Mark To Do Items as Done using its index",
	Example: "  go_cli done 3",
	Args:    cobra.MinimumNArgs(1),
	PreRun:  DebugFlagsValues,
	Run:     doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {
	if debug {
		util.Debug("Done command called")
	}
	var doneIndexes = []int{}
	items := database.LoadItems(dataFile, debug)
	for _, arg := range args {
		index, err := strconv.Atoi(arg)
		if err != nil {
			util.Fatal(fmt.Sprintf("%q is not a valid To Do item Index.", arg))
		}
		if index > 0 && index <= len(items) {
			// TODO ignore already done items
			items[index-1].Done = true
			doneIndexes = append(doneIndexes, index)
		} else {
			util.Fatal(fmt.Sprintf("Index %v does not match any To Do item.", index))
		}
	}
	database.SaveItems(dataFile, items, debug)
	for _, index := range doneIndexes {
		util.Info(fmt.Sprintf("To Do item %v. %q marked as done", index, items[index-1].Description))
	}
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
