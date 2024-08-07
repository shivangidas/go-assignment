package main

import (
	"fmt"
	"os"

	"github.com/shivangidas/go-to-do-app/displayExercises/display"
	"github.com/shivangidas/go-to-do-app/displayExercises/display/model"
)

func main() {
	var items = make([]model.Todo, 10)
	for i := 0; i < 10; i++ {
		items[i] = model.Todo{Name: "Test " + fmt.Sprint(i), Status: 0}
	}
	display.PrintList(os.Stdout, items...)
	display.PrintListJSON(os.Stdout, items...)
	display.CreateFileAndWrite(items...)
	display.DisplayConcurrent()
}
