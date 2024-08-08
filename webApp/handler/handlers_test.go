package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/shivangidas/go-to-do-app/taskInterface"
)

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func assertCode(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("HTTP code got %d, want %d", got, want)
	}
}

// TODO Use interface to mock the crud operations when using DB
func TestCreateHandler(t *testing.T) {
	t.Run("Test create function success", func(t *testing.T) {
		form := url.Values{}
		form.Add("todo", "Test Task")
		form.Add("status", strconv.Itoa(int(taskInterface.Completed)))
		request, _ := http.NewRequest(http.MethodPost, "/task", strings.NewReader(form.Encode()))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		response := httptest.NewRecorder()

		addTaskHandler(response, request)

		got := response.Code
		if got != http.StatusFound {
			t.Errorf("got %q want %q", got, http.StatusFound)
		}

	})

	t.Run("Test create function", func(t *testing.T) {
		form := url.Values{}
		form.Add("todo", "Test Task")
		form.Add("status", "5")
		request, _ := http.NewRequest(http.MethodPost, "/task", strings.NewReader(form.Encode()))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		response := httptest.NewRecorder()

		addTaskHandler(response, request)
		got := response.Body.String()
		want := "Did not add task\n"
		assertStrings(t, got, want)
		assertCode(t, response.Code, http.StatusBadRequest)
	})
}

func TestDeleteHandler(t *testing.T) {
	t.Run("Test delete function with bad id", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/delete/1232", nil)
		response := httptest.NewRecorder()

		deleteHandler(response, request)

		got := response.Body.String()
		want := "Invalid task ID\n"
		assertStrings(t, got, want)
	})
	t.Run("Test delete function with good id", func(t *testing.T) {
		var mockTask = taskInterface.TaskList{}
		id, _ := mockTask.AddTask(taskInterface.Task{Name: "Pass this test", Status: taskInterface.StatusEnum(0)})

		url := "/delete?id=" + id.String()

		request, _ := http.NewRequest(http.MethodPost, url, nil)
		response := httptest.NewRecorder()

		deleteHandler(response, request)

		got := response.Code
		want := 302
		if got != want {
			t.Errorf(" got %q, want %q", got, want)
		}
	})

}
