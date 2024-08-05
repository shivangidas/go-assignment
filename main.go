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
	display.DispayConcurrent()

}
