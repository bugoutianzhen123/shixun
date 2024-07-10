package repository

import (
	"project/domain"
	"project/pkg/logger"
	"project/repository/cache"
)

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
	u, err := repo.cache.GetUserById(userid)
	if err == nil {
		return u, err
	}
	if err != cache.ErrKeyNotExists {
		repo.l.Error("访问Redis失败，查询用户缓存", logger.Error(err), logger.Uint("userid", userid))
	}
	user, err := repo.dao.GetUserById(userid)
	if err != nil {
		return domain.User{}, err
	}
	//异步回写
	go func() {
		err := repo.cache.SetUserById(user)
		if err != nil {
			repo.l.Error("User回写失败", logger.Error(err), logger.Uint("userid", userid))
		}
	}()
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
