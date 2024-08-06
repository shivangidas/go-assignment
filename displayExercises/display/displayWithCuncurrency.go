package display

import (
	"fmt"
	"time"
)

func DisplayConcurrent() {
	tasks := []string{"Task 1", "Task 2", "Task 3", "Task 4", "Task 5", "Task 6", "Task 7", "Task 8", "Task 9", "Task 10", "Task 11", "Task 12"}
	statuses := []string{"Pending", "In Progress", "Completed", "In Progress", "In Progress", "Completed", "In Progress", "Completed", "In Progress", "Completed", "Pending", "Pending"}

	taskChan := make(chan string)
	statusChan := make(chan string)

	go func() {
		for _, task := range tasks {
			taskChan <- task
			time.Sleep(500 * time.Millisecond)
		}
		close(taskChan)
	}()

	go func() {
		for _, status := range statuses {
			statusChan <- status
			time.Sleep(500 * time.Millisecond)
		}
		close(statusChan)
	}()

	for task := range taskChan {
		fmt.Println("To Do Item:", task)
		if status, ok := <-statusChan; ok {
			fmt.Println("Status:", status)
		}
	}
}
