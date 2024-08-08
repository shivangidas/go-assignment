package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/google/uuid"
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
func assertError(t testing.TB, err, want error) {
	t.Helper()
	if err != want {
		t.Errorf("got %s, want %s", err, want)
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
		request, _ := http.NewRequest(http.MethodGet, "/delete/1232", nil)
		response := httptest.NewRecorder()

		deleteHandler(response, request)

		got := response.Body.String()
		want := "Invalid task ID\n"
		assertStrings(t, got, want)
	})
	t.Run("Test delete function with good id", func(t *testing.T) {

		id, _ := inMemoryTasks.AddTask(taskInterface.Task{Name: "Pass this test", Status: taskInterface.StatusEnum(0)})

		url := "/delete?id=" + id.String()

		request, _ := http.NewRequest(http.MethodGet, url, nil)
		response := httptest.NewRecorder()

		deleteHandler(response, request)

		got := response.Code
		want := 302
		if got != want {
			t.Errorf(" got %q, want %q", got, want)
		}
	})

}

func TestUpdateHandler(t *testing.T) {
	task := taskInterface.Task{Name: "Old Task", Status: taskInterface.Start}
	id, _ := inMemoryTasks.AddTask(task)

	// Try a table format
	tests := []struct {
		name           string
		id             string
		todo           string
		status         int64
		expectedCode   int
		expectedName   string
		expectedStatus taskInterface.StatusEnum
	}{
		{
			name:           "valid update",
			id:             id.String(),
			todo:           "Updated Task",
			status:         int64(taskInterface.Ongoing),
			expectedCode:   http.StatusFound,
			expectedName:   "Updated Task",
			expectedStatus: taskInterface.Ongoing,
		},
		{
			name:           "invalid UUID",
			id:             "invalid-uuid",
			todo:           "Updated Task",
			status:         int64(taskInterface.Ongoing),
			expectedCode:   http.StatusBadRequest,
			expectedName:   "Old Task",
			expectedStatus: taskInterface.Start,
		},
		{
			name:           "non-existent task",
			id:             uuid.New().String(),
			todo:           "Updated Task",
			status:         int64(taskInterface.Ongoing),
			expectedCode:   http.StatusNotFound,
			expectedName:   "Old Task",
			expectedStatus: taskInterface.Start,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("id", tt.id)
			form.Add("todo", tt.todo)
			form.Add("status", strconv.FormatInt(tt.status, 10))
			request, _ := http.NewRequest(http.MethodPost, "/edit", strings.NewReader(form.Encode()))
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			response := httptest.NewRecorder()

			updateTaskHandler(response, request)

			assertCode(t, tt.expectedCode, response.Code)

			if tt.expectedCode == http.StatusFound {
				updatedTask, err := inMemoryTasks.SearchTask(id)
				assertError(t, err, nil)
				assertStrings(t, tt.expectedName, updatedTask.Name)
				assertCode(t, int(tt.expectedStatus), int(updatedTask.Status))
			}
		})
	}
}
