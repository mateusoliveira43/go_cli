package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/mateusoliveira43/go_cli/todo"
	"github.com/spf13/cobra"
)

var (
	doneOpt bool
	allOpt  bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List To Do items",
	Long:    "List To Do items of current database. By default, only list To Do items not 'Don'. Filters can be applied with flags.",
	Run:     listRun,
	Args:    cobra.NoArgs,
}

func listRun(cmd *cobra.Command, args []string) {
	fmt.Println("list called")
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(items)
	sort.Sort(todo.ByPriority(items))
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range items {
		if allOpt || i.Done == doneOpt {
			fmt.Fprintln(w, i.Label()+"\t"+i.PrettyDone()+"\t"+i.PrettyP()+"\t"+i.Description+"\t")
		}
	}
	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' Todos")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show all Todos")
}
