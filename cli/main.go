package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
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

func printOneTask(id string, task taskInterface.Task) {
	fmt.Println(id + ": " + task.Name + " (" + task.Status.String() + ")")
}
func displayAll() {
	for id, item := range inMemoryTasks {
		printOneTask(id.String(), item)
	}
}
func addTask(input string) {
	err := checkInputs(input, 2)
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
func checkInputs(input string, expectedFields int) error {
	vals := strings.Split(input, " ")
	if len(vals) < expectedFields || vals[expectedFields-1] == "" {
		return errors.New("add a value with that command")
	}
	return nil
}
func showTask(input string) {
	err := checkInputs(input, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	idString := input[5:]
	id, err := uuid.Parse(idString)
	if err != nil {
		fmt.Println(err)
		return
	}
	task, err := inMemoryTasks.SearchTask(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	printOneTask(idString, task)
}

func deleteTask(input string) {
	err := checkInputs(input, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	idString := input[5:]
	id, err := uuid.Parse(idString)
	if err != nil {
		fmt.Println(err)
		return
	}
	inMemoryTasks.DeleteTask(id)
}

func updateTask(input string) {
	err := checkInputs(input, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	idString := strings.Split(input, " ")[1]
	id, err := uuid.Parse(idString)
	if err != nil {
		fmt.Println(err)
		return
	}
	taskName := input[8+len(idString):]
	updateErr := inMemoryTasks.UpdateTaskName(id, taskName)
	if updateErr != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Updated task " + idString)
}

func updateStatus(input string) {
	err := checkInputs(input, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	idString := strings.Split(input, " ")[1]
	id, err := uuid.Parse(idString)
	if err != nil {
		fmt.Println(err)
		return
	}
	status := strings.Split(input, " ")[2]
	statusInt, err := strconv.ParseInt(status, 10, 64)
	if err != nil {
		fmt.Println("Wrong input for status ", err)
		return
	}
	updateErr := inMemoryTasks.UpdateStatus(id, taskInterface.StatusEnum(statusInt))
	if updateErr != nil {
		fmt.Println("Cannot update", err)
		return
	}
	fmt.Println("Updated task status " + idString)
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
		case "delete":
			deleteTask(userInput)
		case "update":
			updateTask(userInput)
		case "status":
			updateStatus(userInput)
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
	fmt.Println("status [id] [0/1/2/3] 	-Update status, 0-To start, 1-Ongoing, 2-Completed, 3-Ignored")
	fmt.Println("1						-Show commands")
	fmt.Println("2						-Exit")
	fmt.Println("")
}
