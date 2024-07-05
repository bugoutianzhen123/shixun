package dao

import (
	"gorm.io/gorm"
	"project/domain"
)

type DAO interface {
	UserDao
}

type GORMDAO struct {
	db *gorm.DB
}

func InitTables(db *gorm.DB) error {
	db.AutoMigrate(domain.User{})
	db.AutoMigrate(domain.Warehouse{})
	db.AutoMigrate(domain.Item{})
	db.AutoMigrate(domain.Inventory{})
	db.AutoMigrate(domain.InboundRecord{})
	db.AutoMigrate(domain.OutboundRecord{})

	return db.Error
}

func NewGORMDAO(db *gorm.DB) DAO {
	return &GORMDAO{db: db}
}
