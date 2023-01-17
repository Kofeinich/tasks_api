package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"simple-api/taskRepository"
)

type getAll struct {
	store taskRepository.Storage
}

func NewGetAll(store taskRepository.Storage) *getAll {
	return &getAll{store: store}
}
func (r getAll) handler(c echo.Context) error {
	all, err := r.store.GetAll(5, 0)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, all)
}
func (r getAll) Register(registrar Registrar) {
	registrar.Add(http.MethodGet, "/tasks", r.handler)
}
