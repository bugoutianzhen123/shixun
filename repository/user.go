package repository

import "project/domain"

type User interface {
	CreateUser(user domain.User) error
	GetUserById(userid uint) (domain.User, error)
	ChangePassword(user domain.User) error
	ChangePermission(user domain.User) error
	ChangeName(user domain.User) error
	GetUserByName(username string) (domain.User, error)
}

func (repo *CachedDaoRepository) CreateUser(user domain.User) error { return repo.dao.CreateUser(user) }

func (repo *CachedDaoRepository) GetUserById(userid uint) (domain.User, error) {
	user, err := repo.dao.GetUserById(userid)
	return user, err
}

func (repo *CachedDaoRepository) ChangePassword(user domain.User) error {
	return repo.dao.ChangePassword(user)
}

func (repo *CachedDaoRepository) ChangePermission(user domain.User) error {
	return repo.dao.ChangePermission(user)
}

func (repo *CachedDaoRepository) ChangeName(user domain.User) error { return repo.dao.ChangeName(user) }

func (repo *CachedDaoRepository) GetUserByName(username string) (domain.User, error) {
	return repo.dao.GetUserByName(username)
}
