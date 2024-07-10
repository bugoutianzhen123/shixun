package domain

import (
	"gorm.io/gorm"
)

// 仓库
type Warehouse struct {
	gorm.Model
	Name        string `json:"warehousename"`
	Location    string `json:"warehouselocation"`
	Description string `json:"warehousedescription"`
}

// 物品
type Item struct {
	gorm.Model
	Name        string `json:"itemname"`
	TotalNumber int
	Description string `json:"itemdescription"`
}

// 库存
type Inventory struct {
	gorm.Model
	Number      int
	WarehouseId uint      `json:"warehouseid"`
	ItemId      uint      `json:"itemid"`
	Item        Item      `gorm:"foreignKey:ItemId;references:ID"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseId;references:ID"`
}

// 入库记录
type InboundRecord struct {
	gorm.Model
	WarehouseId uint      `json:"warehouseid"`
	ItemId      uint      `json:"itemid"`
	Number      int       `json:"innumber"`
	Description string    `json:"indescription"`
	Item        Item      `gorm:"foreignKey:ItemId;references:ID"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseId;references:ID"`
}

// 出库记录
type OutboundRecord struct {
	gorm.Model
	WarehouseId uint      `json:"warehouseid"`
	ItemId      uint      `json:"itemid"`
	Number      int       `json:"outnumber"`
	Description string    `json:"outdescription"`
	Item        Item      `gorm:"foreignKey:ItemId;references:ID"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseId;references:ID"`
}
