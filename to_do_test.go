package main

import "testing"

func TestAddTodo(t *testing.T) {
	var todo_list []string
	addTodo(todo_list, "New Todo")
	if len(todo_list) > 0 {
		t.Error("No ToDos were added, length expected to be 1, got ", len(todo_list))
	}
}
