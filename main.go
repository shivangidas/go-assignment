package main

import (
	"github.com/shivangidas/go-to-do-app/display"
	"github.com/shivangidas/go-to-do-app/model"
)

func main() {
	var items = make([]model.Todo, 10)
	for i := 0; i < 10; i++ {
		items[i] = model.Todo{Name: "Test", Status: 0}
	}
	//display.PrintList(os.Stdout, items...)
	//PrintListJSON(os.Stdout, items...)
	//save.CreateFileAndWrite(items...)
	display.DisplayConcurrent()
	// cmdline tasks
	// inMemoryTasks := cmdLineApp.TaskList{}
	// sampleTask := cmdLineApp.Task{Name: "Hack the patriarchy", Status: cmdLineApp.StatusEnum(3)}
	// id, _ := inMemoryTasks.AddTask(sampleTask)
	// fmt.Println(inMemoryTasks.SearchTask(id))

}
