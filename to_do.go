package main

import (
	"fmt"
	"os"
)

func main() {
	showMenu()
}

func printSlice(s []string) {
	for i, todo := range s {
		fmt.Println(i, todo)
	}
}

func selectOption(todo_list []string) []string {
	var option int

	fmt.Scan(&option)
	fmt.Println("You have selected", option)
	switch option {
	case 1:
		fmt.Printf("Please, write your new ToDo\n")
		fmt.Printf("Enter a new todo, please:\n")
		var todo string
		_, err := fmt.Scan(&todo)
		if err != nil {
			fmt.Println(err)
		}

		todo_list = addTodo(todo_list, todo)

	case 2:
		showTodoList(todo_list)

	case 3:
		var todo_number int

		if len(todo_list) <= 0 {
			fmt.Printf("You have no ToDos left in your list.\n")
			return todo_list
		}

		fmt.Printf("Please enter the ToDo's number you want to delete, or write anything else to cancel:\n")
		_, err := fmt.Scan(&todo_number)

		if err != nil {
			fmt.Println(err)
		}

		todo_list = deleteTodo(todo_list, todo_number)

	case 4:
		fmt.Printf("See ya' space cowboy.")
		os.Exit(1)

	default:
		fmt.Printf("Please, enter a valid option (A number from 1 to 4). \n")
	}

	return todo_list
}

func addTodo(todo_list []string, todo string) []string {
	todo_list = append(todo_list, todo)
	fmt.Println("You've added a new ToDo", todo)
	return todo_list
}

func showTodoList(todo_list []string) {
	fmt.Printf("This is your ToDo list: \n")
	printSlice(todo_list)
	fmt.Println()
}

func deleteTodo(todo_list []string, index int) []string {
	return append(todo_list[:index], todo_list[index+1:]...)
}

func showMenu() {
	options := [4]string{
		"Add a new ToDo",
		"Show me my ToDo list",
		"Delete a ToDo",
		"Exit program",
	}

	var todo_list []string

	for {

		for index, option := range options {
			fmt.Println(index+1, option)
		}

		todo_list = selectOption(todo_list)
	}
}
