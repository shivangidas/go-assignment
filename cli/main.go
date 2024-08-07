package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/shivangidas/go-to-do-app/taskInterface"
)

var inMemoryTasks = taskInterface.TaskList{}

func InjectData() {
	sampleTask := []taskInterface.Task{{Name: "Buy shares", Status: taskInterface.StatusEnum(2)},
		{Name: "Check news", Status: taskInterface.StatusEnum(0)},
		{Name: "Complete assignment", Status: taskInterface.StatusEnum(1)},
		{Name: "Send email", Status: taskInterface.StatusEnum(2)},
		{Name: "Call access to work", Status: taskInterface.StatusEnum(0)}}
	for _, task := range sampleTask {
		inMemoryTasks.AddTask(task)
	}
}
func readInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
func main() {
	InjectData()
	input := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a command with inputs or press 1 for help, 2 to exit")

Loop:
	for {
		userInput := readInput(input)
		command := strings.Split(userInput, " ")[0]

		switch command {
		case "2":
			break Loop
		case "1":
			displayHelp()
		case "list":
			displayAll()
		case "add":
			addTask(userInput)
		case "show":
			showTask(userInput)
		default:
			fmt.Println("Enter a command or press 1 for help, 2 to exit")
		}

	}
}

func displayHelp() {
	fmt.Println("")
	fmt.Println("Valid commands:")
	fmt.Println("list					-Show all tasks")
	fmt.Println("add [task]				-Create a new task")
	fmt.Println("show [id]				-Show a single task")
	fmt.Println("delete [id]			-Delete task with a given ID")
	fmt.Println("update [id] [name]		-Update task of a given ID with a new name")
	fmt.Println("status [id] [status]	-Update status")
	fmt.Println("1						-Show commands")
	fmt.Println("2						-Exit")
	fmt.Println("")
}

func displayAll() {
	for id, item := range inMemoryTasks {
		fmt.Println(id.String() + ": " + item.Name + " (" + item.Status.String() + ")")
	}
}
func addTask(input string) {
	err := checkInputs(input)
	if err != nil {
		fmt.Println(err)
	} else {
		taskName := input[4:]
		task := taskInterface.Task{Name: taskName, Status: taskInterface.StatusEnum(0)}
		id, err := inMemoryTasks.AddTask(task)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(id)
		}
	}
}
func checkInputs(input string) error {
	vals := strings.Split(input, " ")
	if len(vals) == 1 || vals[1] == "" {
		return errors.New("add a value with that command")
	}
	return nil
}
func showTask(input string) {
	err := checkInputs(input)
	if err != nil {
		fmt.Println(err)
	} else {
		idString := input[5:]
		id, err := uuid.Parse(idString)
		if err != nil {
			fmt.Println(err)
		} else {
			task, err := inMemoryTasks.SearchTask(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(idString + ": " + task.Name + " (" + task.Status.String() + ")")
			}
		}
	}
}
