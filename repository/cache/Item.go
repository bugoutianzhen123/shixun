package cache

import (
	"encoding/json"
	"project/domain"
	"time"
)

type ItemCache interface {
	GetItem() ([]domain.Item, error)
	SetItem(item []domain.Item) error
	GetWarehouse() ([]domain.Warehouse, error)
	SetWarehouse(ware []domain.Warehouse) error
}

func (redis *RedisCache) GetItem() ([]domain.Item, error) {
	key := "totalitem"
	i, err := redis.cmd.Get(key).Bytes()
	if err != nil {
		return []domain.Item{}, err
	}
	var item []domain.Item
	err = json.Unmarshal(i, &item)
	if err != nil {
		return []domain.Item{}, err
	}
	return item, err

}

func (redis *RedisCache) SetItem(item []domain.Item) error {
	key := "totalitem"
	val, err := json.Marshal(&item)
	if err != nil {
		return err
	}
	err = redis.cmd.Set(key, val, 30*time.Minute).Err()
	return err
}

func (redis *RedisCache) GetWarehouse() ([]domain.Warehouse, error) {
	key := "totalwarehouse"
	w, err := redis.cmd.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	var warehouse []domain.Warehouse
	err = json.Unmarshal(w, &warehouse)
	if err != nil {
		return nil, err
	}
	return warehouse, err
}

func (redis *RedisCache) SetWarehouse(ware []domain.Warehouse) error {
	key := "totalwarehouse"
	val, err := json.Marshal(&ware)
	if err != nil {
		return err
	}
	err = redis.cmd.Set(key, val, 30*time.Minute).Err()
	return err
}
