package server

import "project/domain"

type UserServer interface {
	CreateUser(user domain.User) error
	GetUserById(userid uint) (domain.User, error)
	GetUserByName(username string) (domain.User, error)
	ChangeUserPassword(user domain.User) error
	ChangeUserName(user domain.User) error
	ChangeUserPermission(user domain.User) error
}

func (s *server) CreateUser(user domain.User) error {
	return s.rep.CreateUser(user)
}

func (s *server) GetUserById(userid uint) (domain.User, error) {
	return s.rep.GetUserById(userid)
}

func (s *server) GetUserByName(username string) (domain.User, error) {
	return s.rep.GetUserByName(username)
}

func (s *server) ChangeUserPassword(user domain.User) error {
	return s.rep.ChangePassword(user)
}

func (s *server) ChangeUserName(user domain.User) error {
	return s.rep.ChangePassword(user)
}

func (s *server) ChangeUserPermission(user domain.User) error {
	return s.rep.ChangePassword(user)
}
