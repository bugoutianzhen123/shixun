package cache

import (
	"github.com/go-redis/redis"
)

var ErrKeyNotExists = redis.Nil

type CacheRe interface {
	UserCache
	ItemCache
}

type RedisCache struct {
	cmd redis.Cmdable
}

func NewRedisCache(cmd redis.Cmdable) CacheRe {
	return &RedisCache{cmd: cmd}
}
