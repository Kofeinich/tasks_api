package main

import (
	server "simple-api/server"
	"simple-api/server/routes"
	"simple-api/taskRepository"
)

func main() {
	store := taskRepository.NewStore()
	getAllHandler := routes.NewGetAll(store)
	s := server.NewServer()
	s.RegisterHandlers([]routes.RouteHandler{getAllHandler})
	s.Start()
}
