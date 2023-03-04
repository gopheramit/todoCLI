package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/gopheramit/todoCLI"
)

const todoFileName = ".todo.json"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Devloped for The Pragmatic Bookshellf\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "copyright 2020\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage Information\n")
		flag.PrintDefaults()

	}
	task := flag.String("task", "", "task to be included in to do list")
	list := flag.Bool("list", false, "list all the tasks")
	complete := flag.Int("complete", 0, "item to be completed")
	flag.Parse()

	l := &todo.List{}
	if err := l.Get(todoFileName); err != nil {

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

	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *task != "":
		l.Add(*task)

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)

		}
	default:
		fmt.Fprintln(os.Stderr, "invalid option")
		os.Exit(1)

	}

}