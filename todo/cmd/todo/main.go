package main

import (
	"fmt"
	"go-cli/todo"
	"os"
	"strings"
)

const TODO_FILE_NAME = ".todo.json"

func main() {
	l := &todo.List{}

	if err := l.Get(TODO_FILE_NAME); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")

		l.Add(item)

		if err := l.Save(TODO_FILE_NAME); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
