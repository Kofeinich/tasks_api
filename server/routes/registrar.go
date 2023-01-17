package routes

import "github.com/labstack/echo/v4"

type Registrar interface {
	Add(method, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route
}

type RouteHandler interface {
	Register(registrar Registrar)
}
