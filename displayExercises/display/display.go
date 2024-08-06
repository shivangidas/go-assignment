package display

import (
	"fmt"
	"io"

	"github.com/shivangidas/go-to-do-app/displayExercises/display/model"
)

func Print(out io.Writer, item string) {
	fmt.Fprintln(out, item)
}
func PrintList(out io.Writer, listOfWork ...model.Todo) {
	for _, item := range listOfWork {
		Print(out, item.Name+" "+item.Status.String())
	}
}
