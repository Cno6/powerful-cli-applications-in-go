package main

import (
	"bufio"
	"flag"
	"fmt"
	"go-cli/todo"
	"io"
	"os"
	"strings"
)

const (
	DefaultToDoFileName = ".todo.json"
	FileNameEnv         = "TODO_FILENAME"
)

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("task cannot be blank")
	}

	return s.Text(), nil
}

func main() {
	list := flag.Bool("list", false, "List all to-do items")
	add := flag.Bool("add", false, "Add new to-do item")
	complete := flag.Int("complete", 0, "Mark the to-do item as completed")

	flag.Parse()

	todoFileName := DefaultToDoFileName

	if os.Getenv(FileNameEnv) != "" {
		todoFileName = os.Getenv(FileNameEnv)
	}

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		fmt.Print(l)
	case *complete != 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

}
