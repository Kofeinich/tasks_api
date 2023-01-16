package repository

import (
	"errors"
	"simple-api/entity/task"
)

type Storage interface {
	Create(t task.Task) error
	Get(id int) (task.Task, error)
	Delete(id int) error
	GetAll(limit uint, offset uint) ([]task.Task, error)
}

type store struct {
	Tasks []task.Task
}

func (s *store) Create(t task.Task) error {
	s.Tasks = append(s.Tasks, t)
	return nil
}

func (s *store) Get(id int) (task.Task, error) {
	for _, t := range s.Tasks {
		if t.Id == id {
			return t, nil
		}
	}
	return task.Task{}, errors.New("cannot found task")
}

func (s *store) Delete(id int) error {
	for i, t := range s.Tasks {
		if t.Id == id {
			s.Tasks = append(s.Tasks[:i], s.Tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("cannot delete: task not found")
}

func (s *store) GetAll(limit uint, offset uint) ([]task.Task, error) {

	if limit > 100 {
		return nil, errors.New("limit bigger then 100")
	}

	if offset > 1000 {
		return nil, errors.New("offset bigger then 1000")
	}

	if len(s.Tasks) < int(limit+offset) {
		return s.Tasks[offset:], nil
	}

	return s.Tasks[offset : limit+offset], nil
}

func NewStore() Storage {
	return &store{Tasks: make([]task.Task, 0, 0)}
}
