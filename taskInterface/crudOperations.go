package taskInterface

import (
	"github.com/google/uuid"
)

const (
	CannotAddTask               = TaskErr("Cannot add task")
	CannotFindTask              = TaskErr("Task not found")
	CannotUpdateNonExistentTask = TaskErr("Cannot update non-existent task")
	NotAValidStatus             = TaskErr("Not a valid status")
)

type TaskErr string

func (e TaskErr) Error() string {
	return string(e)
}

type StatusEnum int

const (
	Start = iota
	Ongoing
	Completed
	Ignored
)

var stateName = map[StatusEnum]string{
	Start:     "To start",
	Ongoing:   "Ongoing",
	Completed: "Completed",
	Ignored:   "Ignored",
}

func (ss StatusEnum) String() string {
	return stateName[ss]
}
func (ss StatusEnum) CheckStatus() error {
	_, ok := stateName[ss]
	if !ok {
		return NotAValidStatus
	}
	return nil
}

type Task struct {
	Name   string
	Status StatusEnum
}

type TaskList map[uuid.UUID]Task

func (tasks TaskList) CurrentTaskLength() int {
	return len(tasks)
}

func (tasks TaskList) SearchTask(id uuid.UUID) (Task, error) {
	todo, ok := tasks[id]
	if !ok {
		return Task{}, CannotFindTask
	}
	return todo, nil
}

func (tasks TaskList) AddTask(newTodo Task) (uuid.UUID, error) {
	id := uuid.New()
	err := newTodo.Status.CheckStatus()
	if err != nil {
		return uuid.Nil, err
	}
	tasks[id] = newTodo
	return id, nil
}

func (tasks TaskList) UpdateTaskName(id uuid.UUID, name string) error {
	task, err := tasks.SearchTask(id)

	switch err {
	case CannotFindTask:
		return CannotUpdateNonExistentTask
	case nil:
		task.Name = name
		tasks[id] = task
	default:
		return err
	}
	return nil
}

func (tasks TaskList) UpdateStatus(id uuid.UUID, status StatusEnum) error {
	err := status.CheckStatus()
	if err != nil {
		return err
	}

	task, err := tasks.SearchTask(id)

	switch err {
	case CannotFindTask:
		return CannotUpdateNonExistentTask
	case nil:
		task.Status = status
		tasks[id] = task
	default:
		return err
	}
	return nil
}

func (tasks TaskList) DeleteTask(id uuid.UUID) {
	delete(tasks, id)
}
