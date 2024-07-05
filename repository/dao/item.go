package dao

import "project/domain"

type ItemDao interface {
}

func (dao *GORMDAO) CreateWarehouse(ware domain.Warehouse) error {
	err := dao.db.Create(&ware).Error
	return err
}

func (dao *GORMDAO) CreateItem(item domain.Item) error {
	err := dao.db.Create(&item).Error
	return err
}
