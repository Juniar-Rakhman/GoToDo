package handlers

import (
	"fmt"
	"github.com/jrakhman/todo"
)

var currentId int

var todoList todo.Todos

// Give us some seed data
func init() {
	RepoCreateTodo(todo.Todo{Name: "Write presentation"})
	RepoCreateTodo(todo.Todo{Name: "Host meetup"})

}

func RepoFindTodo(id int) todo.Todo {
	for _, t := range todoList {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return todo.Todo{}
}

func RepoCreateTodo(t todo.Todo) todo.Todo {
	currentId += 1
	t.Id = currentId
	todoList = append(todoList, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range todoList {
		if t.Id == id {
			todoList = append(todoList[:i], todoList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
