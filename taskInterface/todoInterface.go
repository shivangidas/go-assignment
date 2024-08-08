package taskInterface

import "github.com/google/uuid"

type TaskService interface {
	CurrentTaskLength() int
	SearchTask(id uuid.UUID) (Task, error)
	AddTask(newTodo Task) (uuid.UUID, error)
	UpdateTaskName(id uuid.UUID, name string) error
	UpdateStatus(id uuid.UUID, status StatusEnum) error
	DeleteTask(id uuid.UUID)
}

type TaskServer struct {
	service TaskService
}

func NewTaskServer(service TaskService) *TaskServer {
	return &TaskServer{service: service}
}
