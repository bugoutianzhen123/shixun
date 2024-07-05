package ioc

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func InitRedis() redis.Cmdable {
	type Config struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
	}
	var cfg Config
	err := viper.UnmarshalKey("redis", &cfg)
	if err != nil {
		panic(err)
	}
	return redis.NewClient(&redis.Options{Addr: cfg.Addr, Password: cfg.Password})
}
