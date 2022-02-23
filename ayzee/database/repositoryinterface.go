package database

import (
	task "task-list-service/ayzee/database/api"
	"task-list-service/ayzee/database/repository"
)

type Repository interface {
	GetTask(string) task.Task
	GetTasksByUser(string) []task.Task
	CreateTask(task.Task) repository.PersistTaskResponse
	UpdateTask(task.Task) (int64, error)
}
