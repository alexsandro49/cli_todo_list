package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CommandFlags struct {
	Add    string
	Edit   string
	Delete int
	Toggle int
	List   bool
}

func newCommandFlags() *CommandFlags {
	commandFlags := CommandFlags{}

	flag.StringVar(&commandFlags.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&commandFlags.Edit, "edit", "", "Edit a todo by index & specify a new title. id:new_title")
	flag.IntVar(&commandFlags.Delete, "delete", -1, "Specify a todo by index to delete")
	flag.IntVar(&commandFlags.Toggle, "toggle", -1, "Specify a todo by index to toggle")
	flag.BoolVar(&commandFlags.List, "list", false, "List all todos")

	flag.Parse()

	return &commandFlags
}

func (cf *CommandFlags) execute(todos *Todos) {
	switch {
	case cf.Add != "":
		todos.add(cf.Add)

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: Invalid format for edit! Please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid index for edit!")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Delete != -1:
		todos.delete(cf.Delete)

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.List:
		todos.print()

	default:
		fmt.Println("Error: Invalid command!")
	}
}
