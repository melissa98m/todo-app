package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add new todo")
	flag.Parse()
	todos := &todo.Todos{}
	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)

	}
	switch {
	case *add:
		todos.Add("Sample todo")
	default:
		fmt.Fprintf(os.Stdout, "invalid command")
		os.Exit(1)

	}

}
