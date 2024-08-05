package display

import (
	"bytes"
	"testing"

	"github.com/shivangidas/go-to-do-app/model"
)

func TestPrintListJSON(t *testing.T) {
	t.Run("convert to json format", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		item := model.Todo{Name: "Order food", Status: 2}

		PrintListJSON(buffer, item)

		got := buffer.String()
		want := "[{\"Name\":\"Order food\",\"Status\":2}]\n"
		assertString(t, got, want)
	})
}
