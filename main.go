package main

import (
	server "simple-api/controller"
	"simple-api/repository"
)

func main() {
	store := repository.NewStore()
	s := server.NewServer(store)
	s.Start()
}
