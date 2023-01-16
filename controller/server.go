package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"simple-api/repository"
)

type Server struct {
	store repository.Storage
}

func NewServer(store repository.Storage) *Server {
	return &Server{store: store}
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, From GO!")
}

func (s *Server) Start() error {
	e := echo.New()
	e.GET("/", hello)
	e.GET("/tasks", func(c echo.Context) error {
		all, err := s.store.GetAll(5, 0)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, all)
	})

	e.Start(":8080")
	return nil
}
