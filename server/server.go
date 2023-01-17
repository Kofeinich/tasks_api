package server

import (
	"github.com/labstack/echo/v4"
	"simple-api/server/routes"
)

type server struct {
	e *echo.Echo
}

func (s *server) E() *echo.Echo {
	return s.e
}

func NewServer() *server {
	e := echo.New()
	return &server{e: e}
}
func (s server) RegisterHandlers(routeHandlers []routes.RouteHandler) {
	for _, handler := range routeHandlers {
		handler.Register(s.E())
	}
}
func (s *server) Start() error {
	s.e.Start(":8080")
	return nil
}
