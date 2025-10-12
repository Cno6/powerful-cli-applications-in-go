package main

import (
	"flag"
	"fmt"
	"go-cli/todo"
	"os"
)

const TODO_FILE_NAME = ".todo.json"

func main() {
	list := flag.Bool("list", false, "List all to-do items")
	task := flag.String("task", "", "Add new to-do item")
	complete := flag.Int("complete", 0, "Mark the to-do item as completed")

	flag.Parse()

	l := &todo.List{}

	if err := l.Get(TODO_FILE_NAME); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}
	case *complete != 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := l.Save(TODO_FILE_NAME); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		l.Add(*task)

		if err := l.Save(TODO_FILE_NAME); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

}
