package repository

import (
	"project/pkg/logger"
	"project/repository/cache"
	"project/repository/dao"
)

type Repository interface {
	User
}

type CachedDaoRepository struct {
	cache cache.CacheRe
	dao   dao.DAO
	l     logger.Logger
}

func NewCachedUserRepository(cache cache.CacheRe, dao dao.DAO, l logger.Logger) Repository {
	return &CachedDaoRepository{cache: cache, dao: dao, l: l}
}
