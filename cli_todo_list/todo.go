package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (t *Todos) add(title string) {
	todo := Todo{title, false, time.Now(), nil}

	*t = append(*t, todo)
}

func (t *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*t) {
		err := errors.New("Invalid index")
		fmt.Println(err)

		return err
	}

	return nil
}

func (t *Todos) edit(index int, title string) error {
	todo := *t

	err := todo.validateIndex(index)
	if err != nil {
		return err
	}

	todo[index].Title = title

	return nil
}

func (t *Todos) delete(index int) error {
	todo := *t

	err := todo.validateIndex(index)
	if err != nil {
		return err
	}

	*t = append(todo[:index], todo[index+1:]...)

	return nil
}

func (t *Todos) toggle(index int) error {
	todo := *t

	err := t.validateIndex(index)
	if err != nil {
		return err
	}

	isCompleted := todo[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		todo[index].CompletedAt = &completionTime
	} else {
		todo[index].CompletedAt = nil
	}

	todo[index].Completed = !isCompleted

	return nil
}

func (t *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "CreatedAt", "CompletedAt")

	for i, todo := range *t {
		completed := "❌"
		completedAt := ""

		if todo.Completed {
			completed = "✅"
			completedAt = todo.CompletedAt.Format(time.RFC1123)
		}

		table.AddRow(strconv.Itoa(i), todo.Title, completed,
			todo.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}
