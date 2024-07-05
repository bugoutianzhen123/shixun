package cache

import "github.com/go-redis/redis"

type CacheRe interface {
	test()
}

type RedisCache struct {
	cmd redis.Cmdable
}

func NewRedisCache(cmd redis.Cmdable) CacheRe {
	return &RedisCache{cmd: cmd}
}
