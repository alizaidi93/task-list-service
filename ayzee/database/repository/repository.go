package repository

import (
	"database/sql"
	task "task-list-service/ayzee/database/api"

	"github.com/sirupsen/logrus"
	"github.com/twinj/uuid"
)

type Repository struct {
	db  *sql.DB
	log *logrus.Logger
}

type PersistTaskResponse struct {
	RowsAffected int64
	Uuid         string
}

func ProvideRepository(db *sql.DB, log *logrus.Logger) Repository {
	return Repository{db: db, log: log}
}

func (r Repository) GetTask(uuid string) task.Task {
	var task task.Task
	taskSqlStatement := "SELECT * FROM todo_list WHERE uuid=$1"
	err := r.db.QueryRow(taskSqlStatement, uuid).Scan(&task.Id, &task.Uuid, &task.Task, &task.Completed, &task.CreatedBy, &task.ModifiedBy, &task.ModifiedOn)
	if err != nil {
		r.log.Error("Failed to execute query: ", err)
	}
	return task
}

func (r Repository) GetTasksByUser(user string) []task.Task {
	var tasks = make([]task.Task, 0, 10)
	taskSqlStatement := "SELECT * FROM todo_list WHERE createdBy=$1"
	rows, err := r.db.Query(taskSqlStatement, user)
	for rows.Next() {
		task := task.Task{}
		err := rows.Scan(&task.Id, &task.Uuid, &task.Task, &task.Completed, &task.CreatedBy, &task.ModifiedBy, &task.ModifiedOn)
		if err != nil {
			r.log.Error("Failed to return tasks: ", err)
		}
		tasks = append(tasks, task)
	}
	if err != nil {
		r.log.Error("Failed to execute query: ", err)
	}
	return tasks
}

func (r Repository) CreateTask(task task.Task) PersistTaskResponse {
	task.Uuid = uuid.NewV4().String()
	task.CreatedBy = "AyZee"
	task.ModifiedBy = "AyZee"

	taskRawSql := "INSERT INTO todo_list (uuid, task, completed, createdby, modifiedby, modifiedon) VALUES($1, $2, $3, $4, $5, NOW())"
	result, err := r.db.Exec(taskRawSql, task.Uuid, task.Task, false, task.CreatedBy, task.ModifiedBy)
	if err != nil {
		r.log.Error("Failed to execute query: ", err)
	}
	rowsAffected, _ := result.RowsAffected()
	return PersistTaskResponse{rowsAffected, task.Uuid}
}

func (r Repository) UpdateTask(task task.Task) (int64, error) {
	taskRawSql := "UPDATE todo_list SET completed = $1, modifiedOn = NOW() WHERE uuid = $2"
	result, err := r.db.Exec(taskRawSql, task.Completed, task.Uuid)
	if err != nil {
		r.log.Error("Failed to upadte row: ", err)
	}

	return result.RowsAffected()
}
