package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"todolist/src/todo"
)

type Commands struct {
	Add string
	Del int
	Edit string
	Check int
	List bool
}

func NewCmdFlags() *Commands {
	cf := Commands{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo by title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index and specify a new title")
	flag.IntVar(&cf.Del, "delete", -1, "Delete a todo by index")
	flag.IntVar(&cf.Check, "check", -1, "mark a todo on the list as done by index")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return &cf
}

func (cf *Commands) Execute(todolist *todo.TodoList) {
	switch {
	case cf.List:
		todolist.Print()
	
	case cf.Add != "":
		todolist.Add(cf.Add)
	
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit a todo")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error, Invalid index for edit")
			os.Exit(1)
		}

		todolist.Edit(index, parts[1])
	
	case cf.Check != -1:
		todolist.Check(cf.Check)
	
	case cf.Del != -1:
		todolist.Delete(cf.Del)
	
	default:
		fmt.Println("Invalid command")
	}
}