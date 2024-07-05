package domain

import (
	"gorm.io/gorm"
	"time"
)

// 仓库
type Warehouse struct {
	gorm.Model
	Name string
	time.Location
}

// 物品
type Item struct {
	gorm.Model
	Name        string
	Description string
}

// 库存
type Inventory struct {
	gorm.Model
	Number      int
	WarehouseId uint
	ItemId      uint
}

// 入库记录
type InboundRecord struct {
	gorm.Model
	WarehouseId uint
	ItemId      uint
	Number      int
}

// 出库记录
type OutboundRecord struct {
	gorm.Model
	WarehouseId uint
	ItemId      uint
	Number      int
}
