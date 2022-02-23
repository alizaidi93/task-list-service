package service

import (
	"task-list-service/ayzee/database"
	task "task-list-service/ayzee/database/api"

	"github.com/sirupsen/logrus"
)

type Service struct {
	Repository database.Repository
	log        *logrus.Logger
}

func ProvideService(repository database.Repository, log *logrus.Logger) *Service {
	return &Service{Repository: repository, log: log}
}

func (s *Service) GetTask(uuid string) task.Task {
	task := s.Repository.GetTask(uuid)
	return task
}

func (s *Service) GetTasksByUser(user string) []task.Task {
	tasks := s.Repository.GetTasksByUser(user)
	return tasks
}

func (s *Service) CreateTask(task task.Task) string {
	taskPersistResponse := s.Repository.CreateTask(task)

	if taskPersistResponse.RowsAffected != 1 {
		s.log.Fatal("Task unable to be created.")
	}
	return taskPersistResponse.Uuid
}

func (s *Service) UpdateTask(task task.Task) {
	recordsAffected, _ := s.Repository.UpdateTask(task)
	if recordsAffected != 1 {
		s.log.Fatal("Task unable to be updated.")
	}
}
