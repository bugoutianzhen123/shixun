package dao

import (
	"fmt"
	"gorm.io/gorm"
	"project/domain"
)

type ItemDao interface {
	CreateWarehouse(ware domain.Warehouse) error
	DeleteWarehouse(ware domain.Warehouse) error
	CreateItem(item domain.Item) error
	DeleteItem(item domain.Item) error
	CreateInventory(Inv domain.Inventory) error
	CreateInboundRecord(In domain.InboundRecord) error
	CreateOutboundRecord(Out domain.OutboundRecord) error
	GetWarehouse() ([]domain.Warehouse, error)
	GetWarehouseById(wareid uint) (domain.Warehouse, error)
	GetItem() ([]domain.Item, error)
	GetItemById(itemId uint) (domain.Item, error)
	GetInventory() ([]domain.Inventory, error)
	GetInboundRecord() ([]domain.InboundRecord, error)
	GetOutboundRecord() ([]domain.OutboundRecord, error)
	GetInventoryOfWarehouseId(warehouseid uint) ([]domain.Inventory, error)
	GetInventoryOfItemId(itemId uint) ([]domain.Inventory, error)
	GetInboundRecordOfWarehouseId(warehouseid uint) ([]domain.InboundRecord, error)
	GetOutboundRecordOfWarehouseId(warehouseid uint) ([]domain.OutboundRecord, error)
	GetInventoryOfWarehouseIdAndItemId(warehouseid uint, itemid uint) (domain.Inventory, error)
	GetInboundRecordOfWarehouseIdAndItemId(warehouseid uint, itemid uint) ([]domain.InboundRecord, error)
	GetOutboundRecordOfWarehouseIdAndItemId(warehouseid uint, itemid uint) ([]domain.OutboundRecord, error)
}

func (dao *GORMDAO) CreateWarehouse(ware domain.Warehouse) error {
	err := dao.db.Create(&ware).Error
	return err
}

func (dao *GORMDAO) DeleteWarehouse(ware domain.Warehouse) error {
	err := dao.db.Delete(&ware).Error
	return err
}

func (dao *GORMDAO) CreateItem(item domain.Item) error {
	err := dao.db.Create(&item).Error
	return err
}

func (dao *GORMDAO) DeleteItem(item domain.Item) error {
	err := dao.db.Delete(&item).Error
	return err
}

func (dao *GORMDAO) CreateInboundRecord(In domain.InboundRecord) error {
	err := dao.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&In).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Model(&domain.Inventory{}).Where("Warehouse_ID = ?", In.WarehouseId).Update("Number", gorm.Expr("number + ?", In.Number)).Error; err != nil {
			return err
		}

		if err := tx.Model(&domain.Item{}).Where("ID = ?", In.ItemId).Update("Total_Number", gorm.Expr("total_number + ?", In.Number)).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	return err
}

func (dao *GORMDAO) CreateOutboundRecord(Out domain.OutboundRecord) error {
	err := dao.db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&Out).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Model(&domain.Inventory{}).Where("Warehouse_ID = ?", Out.WarehouseId).Update("Number", gorm.Expr("number - ?", Out.Number)).Error; err != nil {
			return err
		}

		if err := tx.Model(&domain.Item{}).Where("ID = ?", Out.ItemId).Update("Total_Number", gorm.Expr("total_number - ?", Out.Number)).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	return err
}

func (dao *GORMDAO) CreateInventory(Inv domain.Inventory) error {
	err := dao.db.Create(&Inv).Error
	return err
}

func (dao *GORMDAO) GetWarehouse() ([]domain.Warehouse, error) {
	var warehouse []domain.Warehouse
	err := dao.db.Find(&warehouse).Error
	return warehouse, err
}

func (dao *GORMDAO) GetWarehouseById(wareid uint) (domain.Warehouse, error) {
	var warehouse domain.Warehouse
	err := dao.db.First(&warehouse, "Id = ?", wareid).Error
	return warehouse, err
}

func (dao *GORMDAO) GetItem() ([]domain.Item, error) {
	var item []domain.Item
	err := dao.db.Find(&item).Error
	return item, err
}

func (dao *GORMDAO) GetItemById(itemId uint) (domain.Item, error) {
	var item domain.Item
	err := dao.db.First(&item, "Id = ?", itemId).Error
	return item, err
}

func (dao *GORMDAO) GetInventory() ([]domain.Inventory, error) {
	var inventory []domain.Inventory
	err := dao.db.Find(&inventory).Error
	return inventory, err
}

func (dao *GORMDAO) GetInboundRecord() ([]domain.InboundRecord, error) {
	var Inbound []domain.InboundRecord
	err := dao.db.Find(&Inbound).Error
	return Inbound, err
}

func (dao *GORMDAO) GetOutboundRecord() ([]domain.OutboundRecord, error) {
	var Outbound []domain.OutboundRecord
	err := dao.db.Find(&Outbound).Error
	return Outbound, err
}

func (dao *GORMDAO) GetInventoryOfWarehouseId(warehouseid uint) ([]domain.Inventory, error) {
	var inventories []domain.Inventory
	err := dao.db.Preload("Item").Where("Warehouse_id = ?", warehouseid).Find(&inventories).Error
	return inventories, err
}

func (dao *GORMDAO) GetInventoryOfItemId(itemId uint) ([]domain.Inventory, error) {
	var inventories []domain.Inventory
	err := dao.db.Preload("Item").Where("Item_id = ?", itemId).Find(&inventories).Error
	return inventories, err
}

func (dao *GORMDAO) GetInventoryOfWarehouseIdAndItemId(warehouseid uint, itemid uint) (domain.Inventory, error) {
	var inventories domain.Inventory
	err := dao.db.Preload("Item").Preload("Warehouse").Where("warehouse_id = ? and item_id = ?", warehouseid, itemid).First(&inventories).Error
	fmt.Println("find", inventories, err)
	return inventories, err
}

func (dao *GORMDAO) GetInboundRecordOfWarehouseId(warehouseid uint) ([]domain.InboundRecord, error) {
	var InboundRecord []domain.InboundRecord
	err := dao.db.Preload("Item").Where("Warehouse_id = ?", warehouseid).Find(&InboundRecord).Error
	return InboundRecord, err
}

func (dao *GORMDAO) GetInboundRecordOfWarehouseIdAndItemId(warehouseid uint, itemid uint) ([]domain.InboundRecord, error) {
	var InboundRecord []domain.InboundRecord
	err := dao.db.Preload("Item").Where("Warehouse_id = ? and Item_Id = ?", warehouseid, itemid).Find(&InboundRecord).Error
	return InboundRecord, err
}

func (dao *GORMDAO) GetOutboundRecordOfWarehouseId(warehouseid uint) ([]domain.OutboundRecord, error) {
	var OutboundRecord []domain.OutboundRecord
	err := dao.db.Preload("Item").Where("Warehouse_id = ?", warehouseid).Find(&OutboundRecord).Error
	return OutboundRecord, err
}

func (dao *GORMDAO) GetOutboundRecordOfWarehouseIdAndItemId(warehouseid uint, itemid uint) ([]domain.OutboundRecord, error) {
	var OutboundRecord []domain.OutboundRecord
	err := dao.db.Preload("Item").Where("Warehouse_id = ? and Item_Id = ?", warehouseid, itemid).Find(&OutboundRecord).Error
	return OutboundRecord, err
}
