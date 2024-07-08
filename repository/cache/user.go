package cache

import (
	"encoding/json"
	"project/domain"
	"time"
)

type UserCache interface {
	GetUserById(userid uint) (domain.User, error)
	SetUserById(u domain.User) error
}

func (redis *RedisCache) GetUserById(userid uint) (domain.User, error) {
	key := "userid" + string(userid)
	u, err := redis.cmd.Get(key).Bytes()
	if err != nil {
		return domain.User{}, err
	}
	var user domain.User
	err = json.Unmarshal(u, &user)
	if err != nil {
		return domain.User{}, err
	}
	return user, err
}

func (redis *RedisCache) SetUserById(u domain.User) error {
	key := "userid" + string(u.ID)
	val, err := json.Marshal(&u)
	if err != nil {
		return err
	}
	err = redis.cmd.Set(key, val, 30*time.Minute).Err()
	return err
}
