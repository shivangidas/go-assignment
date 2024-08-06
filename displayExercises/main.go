package main

import (
	"github.com/shivangidas/go-to-do-app/displayExercises/display"
	"github.com/shivangidas/go-to-do-app/displayExercises/display/model"
)

func main() {
	var items = make([]model.Todo, 10)
	for i := 0; i < 10; i++ {
		items[i] = model.Todo{Name: "Test", Status: 0}
	}
	//display.PrintList(os.Stdout, items...)
	//PrintListJSON(os.Stdout, items...)
	display.CreateFileAndWrite(items...)
	// display.DisplayConcurrent()

}
