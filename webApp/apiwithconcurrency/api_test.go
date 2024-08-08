package apiwithconcurrency

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"

	tasks "github.com/shivangidas/go-to-do-app/taskWithMutex"
)

func BenchmarkCreateHandler(b *testing.B) {
	taskList := tasks.NewTaskList()
	taskServer := NewTaskServer(taskList)

	form := url.Values{}
	form.Add("todo", "Test task")
	form.Add("status", "0")

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodPost, "/task", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		response := httptest.NewRecorder()

		taskServer.CreateHandler(response, req)

		if response.Result().StatusCode != http.StatusFound {
			b.Fatalf("Expected status 302, got %d", response.Result().StatusCode)
		}
	}
}

// Does not work
func BenchmarkCreateHandlerConcurrent(b *testing.B) {
	taskList := tasks.NewTaskList()
	taskServer := NewTaskServer(taskList)

	form := url.Values{}
	form.Add("todo", "Test task")
	form.Add("status", "0")

	var wg sync.WaitGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			req := httptest.NewRequest(http.MethodPost, "/task", strings.NewReader(form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			response := httptest.NewRecorder()

			taskServer.CreateHandler(response, req)

		}()
	}

	wg.Wait()
}
