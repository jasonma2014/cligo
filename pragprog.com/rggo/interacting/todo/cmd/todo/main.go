package main

import (
	"flag"
	"fmt"
	"os"
	"pragprog.com/rggo/interacting/todo"
)

// hard coding the file name
const todoFileName = ".todo.json"

func main() {
	// Parse command line flags
	task := flag.String("task", "", "Task to be in included in the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	flag.Parse()

	// Define an item list
	l := &todo.List{}

	// Use the Get command to read to do items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// decide what to do based on the number of arguments provided
	switch {
	// Decide what to do based on the provided flags
	case *list:
		// list current todo items
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}
	case *complete > 0:
		// Complete the given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// Save the list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		// Add the task
		l.Add(*task)

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// Invalid option
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
