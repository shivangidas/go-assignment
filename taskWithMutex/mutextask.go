package taskwithmutex

import (
	"sync"

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

type TaskList struct {
	lock  sync.Mutex
	tasks map[uuid.UUID]Task
}

func NewTaskList() *TaskList {
	return &TaskList{tasks: make(map[uuid.UUID]Task)}
}
func (t *TaskList) CurrentTaskLength() int {
	t.lock.Lock()
	defer t.lock.Unlock()
	return len(t.tasks)
}

func (t *TaskList) SearchTask(id uuid.UUID) (Task, error) {
	todo, ok := t.tasks[id]
	if !ok {
		return Task{}, CannotFindTask
	}
	return todo, nil
}

func (t *TaskList) AddTask(newTodo Task) (uuid.UUID, error) {
	id := uuid.New()
	err := newTodo.Status.CheckStatus()
	if err != nil {
		return uuid.Nil, err
	}
	t.tasks[id] = newTodo
	return id, nil
}

func (t *TaskList) UpdateTaskName(id uuid.UUID, name string) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	task, err := t.SearchTask(id)

	switch err {
	case CannotFindTask:
		return CannotUpdateNonExistentTask
	case nil:
		task.Name = name
		t.tasks[id] = task
	default:
		return err
	}
	return nil
}

func (t *TaskList) UpdateStatus(id uuid.UUID, status StatusEnum) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	err := status.CheckStatus()
	if err != nil {
		return err
	}

	task, err := t.SearchTask(id)

	switch err {
	case CannotFindTask:
		return CannotUpdateNonExistentTask
	case nil:
		task.Status = status
		t.tasks[id] = task
	default:
		return err
	}
	return nil
}

func (t *TaskList) DeleteTask(id uuid.UUID) {
	t.lock.Lock()
	defer t.lock.Unlock()
	delete(t.tasks, id)
}

type Tasks interface {
	CurrentTaskLength() int
	SearchTask(id uuid.UUID) (Task, error)
	AddTask(newTodo Task) (uuid.UUID, error)
	UpdateTaskName(id uuid.UUID, name string) error
	UpdateStatus(id uuid.UUID, status StatusEnum) error
	DeleteTask(id uuid.UUID)
}

// func main() {
// 	// take input from user through scanf
// 	inMemoryTasks := TaskList{}
// 	sampleTask := Task{Name: "Hack the patriarchy", Status: StatusEnum(3)}
// 	id, _ := inMemoryTasks.AddTask(sampleTask)
// 	fmt.Println(inMemoryTasks.SearchTask(id))
// }
