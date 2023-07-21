// TODO DOC COMMENTS everything
// TODO standardize use of print and logs: use only one
package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

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

func SaveItems(filename string, items []ToDoItem) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadItems(filename string) ([]ToDoItem, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return []ToDoItem{}, err
	}
	var items []ToDoItem
	if err := json.Unmarshal(b, &items); err != nil {
		return []ToDoItem{}, err
	}
	return items, nil
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
