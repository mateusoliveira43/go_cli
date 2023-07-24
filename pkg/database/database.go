// TODO DOC COMMENTS everything
package database

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/mateusoliveira43/go_cli/pkg/util"
)

const ErrorMessage = "An error ocurred while trying to %v database items"

type ToDoItem struct {
	Index       int
	Description string
	Priority    int
	Done        bool
}

func (item *ToDoItem) SetPriority(priority int) {
	// TODO instead of setting right value, avoid wrong inputs
	switch priority {
	case 1:
		item.Priority = 1
	case 3:
		item.Priority = 3
	default:
		item.Priority = 2
	}
}

func (item *ToDoItem) PrettyP() string {
	if item.Priority == 1 {
		return "(High)"
	}
	if item.Priority == 3 {
		return "(Low)"
	}
	return " "
}

func (item *ToDoItem) Label() string {
	return strconv.Itoa(item.Index) + "."
}

func (item *ToDoItem) PrettyDone() string {
	if item.Done {
		return "[X]"
	}
	return "[ ]"
}

func SaveItems(filename string, items []ToDoItem, debug bool) {
	if debug {
		util.Debug(fmt.Sprintf("Saving database items: %+v", items))
	}
	bytesContent, err := json.Marshal(items)
	if err != nil {
		util.Error(fmt.Sprintf(ErrorMessage, "save"))
		util.Fatal(fmt.Sprintf("%v", err))
	}
	err = os.WriteFile(filename, bytesContent, 0644)
	if err != nil {
		util.Error(fmt.Sprintf(ErrorMessage, "save"))
		util.Fatal(fmt.Sprintf("%v", err))
	}
}

func LoadItems(filename string, debug bool) []ToDoItem {
	bytesContent, err := os.ReadFile(filename)
	if err != nil {
		util.Error(fmt.Sprintf(ErrorMessage, "load"))
		util.Fatal(fmt.Sprintf("%v", err))
	}
	var items []ToDoItem
	if err = json.Unmarshal(bytesContent, &items); err != nil {
		util.Error(fmt.Sprintf(ErrorMessage, "load"))
		util.Fatal(fmt.Sprintf("%v", err))
	}
	if debug {
		util.Debug(fmt.Sprintf("Loaded database items: %+v", items))
	}
	return items
}

// Implements sort.Interface for an array of To Do Items based mainly on the Priority field.
type ByPriority []ToDoItem

func (array ByPriority) Len() int {
	return len(array)
}

// TODO read to see if "current" and "next" are good name for the array indexes
func (array ByPriority) Swap(current, next int) {
	array[current], array[next] = array[next], array[current]
}

// TODO read about sort interface, it seems it it implemented in reverse
func (array ByPriority) Less(current, next int) bool {
	if array[current].Done != array[next].Done {
		return array[current].Done
	}
	if array[current].Priority != array[next].Priority {
		return array[current].Priority < array[next].Priority
	}
	return array[current].Index < array[next].Index
}
