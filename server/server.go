package server

import "project/repository"

type Server interface {
	UserServer
	ItemServer
}

type server struct {
	rep repository.Repository
}

func NewServer(rep repository.Repository) Server {
	return &server{rep}
}
