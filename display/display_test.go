package display

import (
	"bytes"
	"testing"

	"github.com/shivangidas/go-to-do-app/model"
)

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestPrint(t *testing.T) {
	t.Run("print 1 thing", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		item := model.Todo{Name: "Order food", Status: 2}

		Print(buffer, item.Name+" "+item.Status.String())

		got := buffer.String()
		want := "Order food completed\n"
		assertString(t, got, want)
	})
}
func TestPrintList(t *testing.T) {
	t.Run("print 1 todo item with PrintList function", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		item := model.Todo{Name: "Order food", Status: 2}

		PrintList(buffer, item)

		got := buffer.String()
		want := "Order food completed\n"
		assertString(t, got, want)
	})

	t.Run("print 10 todo items", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		var items = make([]model.Todo, 10)
		for i := 0; i < 10; i++ {
			items[i] = model.Todo{Name: "Test", Status: 0}
		}

		PrintList(buffer, items...)

		got := buffer.String()
		want := "Test started\nTest started\nTest started\nTest started\nTest started\nTest started\nTest started\nTest started\nTest started\nTest started\n"
		assertString(t, got, want)
	})
}
