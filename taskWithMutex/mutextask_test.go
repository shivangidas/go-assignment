package taskwithmutex

import (
	"fmt"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertError(t testing.TB, err, want error) {
	t.Helper()
	if err != want {
		t.Errorf("got %s, want %s", err, want)
	}
}
func TestAddTask(t *testing.T) {
	t.Run("Add new task", func(t *testing.T) {
		mockTasks := NewTaskList()
		newTask := Task{Name: "Pass this test", Status: StatusEnum(Start)}
		id, _ := mockTasks.AddTask(newTask)
		assertStrings(t, mockTasks.tasks[id].Name, newTask.Name)
		assertStrings(t, mockTasks.tasks[id].Status.String(), "To start")
	})
	t.Run("Non existent status", func(t *testing.T) {
		mockTasks := NewTaskList()
		newTask := Task{Name: "Pass this test", Status: 5}
		_, err := mockTasks.AddTask(newTask)
		assertError(t, err, NotAValidStatus)
	})
}
func TestSearchTask(t *testing.T) {
	mockTasks := NewTaskList()
	ids := []uuid.UUID{uuid.New(), uuid.New()}

	tasks := []Task{
		{Name: "Write the test", Status: StatusEnum(Start)},
		{Name: "Pass this test", Status: StatusEnum(Start)},
	}
	mockTasks.tasks[ids[0]] = tasks[0]
	mockTasks.tasks[ids[1]] = tasks[1]
	t.Run("Search task by id", func(t *testing.T) {

		got, err := mockTasks.SearchTask(ids[1])
		assertError(t, err, nil)
		assertStrings(t, got.Name, mockTasks.tasks[ids[1]].Name)
	})
	t.Run("Fail search task if not present", func(t *testing.T) {

		_, err := mockTasks.SearchTask(uuid.New())
		assertError(t, err, CannotFindTask)
	})
}

func TestUpdateTaskName(t *testing.T) {
	mockTasks := NewTaskList()
	ids := []uuid.UUID{uuid.New(), uuid.New()}

	tasks := []Task{
		{Name: "Write the test", Status: StatusEnum(Start)},
		{Name: "Pass this test", Status: StatusEnum(Start)},
	}
	mockTasks.tasks[ids[0]] = tasks[0]
	mockTasks.tasks[ids[1]] = tasks[1]
	t.Run("Update task name", func(t *testing.T) {
		err := mockTasks.UpdateTaskName(ids[1], "Pass all tests")
		assertError(t, err, nil)
		assertStrings(t, mockTasks.tasks[ids[1]].Name, "Pass all tests")
	})
	t.Run("Error in Update task name", func(t *testing.T) {
		err := mockTasks.UpdateTaskName(uuid.New(), "Fail this test")
		assertError(t, err, CannotUpdateNonExistentTask)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		newMockTasks := NewTaskList()
		newTask := Task{Name: "Pass this test", Status: StatusEnum(Start)}
		id, _ := newMockTasks.AddTask(newTask)
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(i int) {
				defer wg.Done()
				err := newMockTasks.UpdateTaskName(id, "Pass all tests "+fmt.Sprint(i))
				assertError(t, err, nil)
			}(i)
		}
		wg.Wait()
		finalName := newMockTasks.tasks[id].Name
		expectedPrefix := "Pass all tests "
		assert.Contains(t, finalName, expectedPrefix, "final task name %q does not start with expected prefix %q", finalName, expectedPrefix)
	})
}

func TestUpdateStatus(t *testing.T) {
	mockTasks := NewTaskList()
	ids := []uuid.UUID{uuid.New(), uuid.New()}

	tasks := []Task{
		{Name: "Write the test", Status: StatusEnum(Start)},
		{Name: "Pass this test", Status: StatusEnum(Start)},
	}
	mockTasks.tasks[ids[0]] = tasks[0]
	mockTasks.tasks[ids[1]] = tasks[1]
	t.Run("Update task status", func(t *testing.T) {
		err := mockTasks.UpdateStatus(ids[1], 1)
		assertError(t, err, nil)
		assertStrings(t, mockTasks.tasks[ids[1]].Status.String(), "Ongoing")
	})
	t.Run("Error in Update task status", func(t *testing.T) {
		err := mockTasks.UpdateStatus(uuid.New(), 1)
		assertError(t, err, CannotUpdateNonExistentTask)
	})
	t.Run("Non existent status", func(t *testing.T) {
		err := mockTasks.UpdateStatus(ids[0], 5)
		assertError(t, err, NotAValidStatus)
	})
}

func TestDeleteTask(t *testing.T) {
	mockTasks := NewTaskList()
	ids := []uuid.UUID{uuid.New(), uuid.New()}

	tasks := []Task{
		{Name: "Write the test", Status: StatusEnum(Start)},
		{Name: "Pass this test", Status: StatusEnum(Start)},
	}
	mockTasks.tasks[ids[0]] = tasks[0]
	mockTasks.tasks[ids[1]] = tasks[1]
	mockTasks.DeleteTask(ids[0])
	_, err := mockTasks.SearchTask(ids[0])
	assertError(t, err, CannotFindTask)
}
