package main

import (
	"todolist/src/storage"
	"todolist/src/todo"
)

func main() {
	todolist := todo.TodoList{}
	storage := storage.NewStorage[todo.TodoList]("todos.json")
	storage.Load(&todolist)
	cmd := NewCmdFlags()
	cmd.Execute(&todolist)
	storage.Save(todolist)
}