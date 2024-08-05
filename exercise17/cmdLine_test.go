package main

import (
	"testing"

	"github.com/google/uuid"
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
		mockTasks := TaskList{}
		newTask := Task{Name: "Pass this test", Status: StatusEnum(Pending)}
		id, _ := mockTasks.AddTask(newTask)
		assertStrings(t, mockTasks[id].Name, newTask.Name)
		assertStrings(t, mockTasks[id].Status.String(), "pending")
	})
	t.Run("Non existent status", func(t *testing.T) {
		mockTasks := TaskList{}
		newTask := Task{Name: "Pass this test", Status: 5}
		_, err := mockTasks.AddTask(newTask)
		assertError(t, err, NotAValidStatus)
	})
}
func TestSearchTask(t *testing.T) {
	ids := []uuid.UUID{uuid.New(), uuid.New()}
	mockTasks := TaskList{ids[0]: {Name: "Write the test", Status: StatusEnum(Pending)},
		ids[1]: {Name: "Pass this test", Status: StatusEnum(Pending)}}
	t.Run("Search task by id", func(t *testing.T) {

		got, err := mockTasks.SearchTask(ids[1])
		assertError(t, err, nil)
		assertStrings(t, got.Name, mockTasks[ids[1]].Name)
	})
	t.Run("Fail search task if not present", func(t *testing.T) {

		_, err := mockTasks.SearchTask(uuid.New())
		assertError(t, err, CannotFindTask)
	})
}

func TestUpdateTaskName(t *testing.T) {
	ids := []uuid.UUID{uuid.New(), uuid.New()}
	mockTasks := TaskList{ids[0]: {Name: "Write the test", Status: StatusEnum(Pending)},
		ids[1]: {Name: "Pass this test", Status: StatusEnum(Pending)}}
	t.Run("Update task name", func(t *testing.T) {
		err := mockTasks.UpdateTaskName(ids[1], "Pass all tests")
		assertError(t, err, nil)
		assertStrings(t, mockTasks[ids[1]].Name, "Pass all tests")
	})
	t.Run("Error in Update task name", func(t *testing.T) {
		err := mockTasks.UpdateTaskName(uuid.New(), "Fail this test")
		assertError(t, err, CannotUpdateNonExistentTask)
	})
}

func TestUpdateStatus(t *testing.T) {
	ids := []uuid.UUID{uuid.New(), uuid.New()}
	mockTasks := TaskList{ids[0]: {Name: "Write the test", Status: StatusEnum(Pending)},
		ids[1]: {Name: "Pass this test", Status: StatusEnum(Pending)}}
	t.Run("Update task status", func(t *testing.T) {
		err := mockTasks.UpdateStatus(ids[1], 1)
		assertError(t, err, nil)
		assertStrings(t, mockTasks[ids[1]].Status.String(), "ongoing")
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
	ids := []uuid.UUID{uuid.New(), uuid.New()}
	mockTasks := TaskList{ids[0]: {Name: "Write the test", Status: StatusEnum(Pending)},
		ids[1]: {Name: "Pass this test", Status: StatusEnum(Pending)}}
	mockTasks.DeleteTask(ids[0])
	_, err := mockTasks.SearchTask(ids[0])
	assertError(t, err, CannotFindTask)
}
