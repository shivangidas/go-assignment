package display

import (
	"encoding/json"
	"io"

	"github.com/shivangidas/go-to-do-app/displayExercises/display/model"
)

func PrintListJSON(out io.Writer, items ...model.Todo) { //same as the write file
	jsonList, err := json.Marshal(items)
	if err != nil {
		panic("we failed to make that into a json")
	}
	Print(out, string(jsonList))
}
