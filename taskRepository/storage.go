package taskRepository

import "simple-api/entity/db"

type Storage interface {
	Create(t db.Task) error
	Get(id int) (db.Task, error)
	Delete(id int) error
	GetAll(limit uint, offset uint) ([]db.Task, error)
}
