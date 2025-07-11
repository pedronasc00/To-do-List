package todo

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title string
	Done bool
	CreatAt time.Time
	CompletedAt *time.Time
}

type TodoList []Todo

func (todolist *TodoList) Add(title string) {
	td := Todo {
		Title: title,
		Done: false,
		CreatAt: time.Now(),
		CompletedAt: nil,
	}

	*todolist = append(*todolist, td)
}

func (todolist *TodoList) ValidateIndex(index int) error {
	if index < 0 || index >= len(*todolist) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todolist *TodoList) Delete(index int) error {
	td := *todolist

	if err := td.ValidateIndex(index); err != nil {
		return err
	}
	
	*todolist = append(td[:index], td[index+1:]...)

	return nil
}

func (todolist *TodoList) Check(index int) error {
	td := *todolist
	
	if err := todolist.ValidateIndex(index); err != nil {
		return  err
	}
	
	isCompleted := td[index].Done

	if !isCompleted {
		TimeTodo := time.Now()
		td[index].CompletedAt = &TimeTodo
	}
	
	td[index].Done = !isCompleted

	return nil
}

func (todolist *TodoList) Edit(index int, title string) error {
	td := *todolist
	
	if err := todolist.ValidateIndex(index); err != nil {
		return err
	}
	
	td[index].Title = title

	return nil
}

func (todolist *TodoList) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("ID", "Tarefa", "Feita", "Criada Em", "Comcluida Em")

	for i, t := range *todolist {
		done := "❌"
		completedAt := ""

		if t.Done {
			done = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(i), t.Title, done, t.CreatAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}