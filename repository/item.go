package repository

import (
	"errors"
	"gorm.io/gorm"
	"project/domain"
)

type Item interface {
	CreateItem(item domain.Item) error
	CreateWareHouse(warehouse domain.Warehouse) error
	CreateInboundRecord(in domain.InboundRecord) error
	CreateOutboundRecord(out domain.OutboundRecord) error
	DeleteItem(item domain.Item) error
	DeleteWarehouse(warehouse domain.Warehouse) error
	FindWarehouse() ([]domain.Warehouse, error)
	FindItem() ([]domain.Item, error)
	FindInventory() ([]domain.Inventory, error)
	FindInboundRecord() ([]domain.InboundRecord, error)
	FindOutboundRecord() ([]domain.OutboundRecord, error)
	FindWarehouseById(wareid uint) (domain.Warehouse, error)
	FindItemById(itemid uint) (domain.Item, error)
	FindInventoryByWarehouseId(wareid uint) ([]domain.Inventory, error)
	FindInventoryByItemId(itemid uint) ([]domain.Inventory, error)
	FindInventoryByWarehouseIdAndItemId(wareid, itemid uint) (domain.Inventory, error)
	FindInboundRecordByWarehouseId(wareid uint) ([]domain.InboundRecord, error)
	FindInboundRecordByWarehouseIdAndItemId(wareid, itemid uint) ([]domain.InboundRecord, error)
	FindOutboundRecordByWarehouseId(wareid uint) ([]domain.OutboundRecord, error)
	FindOutboundRecordByWarehouseIdAndItemId(wareid, itemid uint) ([]domain.OutboundRecord, error)
}

func (repo *CachedDaoRepository) CreateItem(item domain.Item) error {
	err := repo.dao.CreateItem(item)
	return err
}

func (repo *CachedDaoRepository) CreateWareHouse(warehouse domain.Warehouse) error {
	err := repo.dao.CreateWarehouse(warehouse)
	return err
}

func (repo *CachedDaoRepository) CreateInboundRecord(in domain.InboundRecord) error {
	_, err := repo.dao.GetInventoryOfWarehouseIdAndItemId(in.WarehouseId, in.ItemId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有找到匹配的记录
			inventory := domain.Inventory{
				WarehouseId: in.WarehouseId,
				ItemId:      in.ItemId,
			}
			if err := repo.dao.CreateInventory(inventory); err != nil {
				return err
			}
		} else {
			// 其他查询错误
			return err
		}
	}
	err = repo.dao.CreateInboundRecord(in)
	return err
}

func (repo *CachedDaoRepository) CreateOutboundRecord(out domain.OutboundRecord) error {
	err := repo.dao.CreateOutboundRecord(out)
	return err
}

func (repo *CachedDaoRepository) DeleteItem(item domain.Item) error {
	err := repo.dao.DeleteItem(item)
	return err
}

func (repo *CachedDaoRepository) DeleteWarehouse(warehouse domain.Warehouse) error {
	err := repo.dao.DeleteWarehouse(warehouse)
	return err
}

func (repo *CachedDaoRepository) FindWarehouse() ([]domain.Warehouse, error) {
	ware, err := repo.dao.GetWarehouse()
	return ware, err
}

func (repo *CachedDaoRepository) FindItem() ([]domain.Item, error) {
	item, err := repo.dao.GetItem()
	return item, err
}

func (repo *CachedDaoRepository) FindInventory() ([]domain.Inventory, error) {
	inventory, err := repo.dao.GetInventory()
	return inventory, err
}

func (repo *CachedDaoRepository) FindInboundRecord() ([]domain.InboundRecord, error) {
	inboundRecord, err := repo.dao.GetInboundRecord()
	return inboundRecord, err
}

func (repo *CachedDaoRepository) FindOutboundRecord() ([]domain.OutboundRecord, error) {
	outboundRecord, err := repo.dao.GetOutboundRecord()
	return outboundRecord, err
}

func (repo *CachedDaoRepository) FindWarehouseById(wareid uint) (domain.Warehouse, error) {
	ware, err := repo.dao.GetWarehouseById(wareid)
	return ware, err
}

func (repo *CachedDaoRepository) FindItemById(itemid uint) (domain.Item, error) {
	item, err := repo.dao.GetItemById(itemid)
	return item, err
}

func (repo *CachedDaoRepository) FindInventoryByWarehouseId(wareid uint) ([]domain.Inventory, error) {
	inventory, err := repo.dao.GetInventoryOfWarehouseId(wareid)
	return inventory, err
}

func (repo *CachedDaoRepository) FindInventoryByItemId(itemid uint) ([]domain.Inventory, error) {
	inventory, err := repo.dao.GetInventoryOfItemId(itemid)
	return inventory, err
}

func (repo *CachedDaoRepository) FindInventoryByWarehouseIdAndItemId(wareid, itemid uint) (domain.Inventory, error) {
	inventory, err := repo.dao.GetInventoryOfWarehouseIdAndItemId(wareid, itemid)
	return inventory, err
}

func (repo *CachedDaoRepository) FindInboundRecordByWarehouseId(wareid uint) ([]domain.InboundRecord, error) {
	InboundRecord, err := repo.dao.GetInboundRecordOfWarehouseId(wareid)
	return InboundRecord, err
}

func (repo *CachedDaoRepository) FindInboundRecordByWarehouseIdAndItemId(wareid, itemid uint) ([]domain.InboundRecord, error) {
	InboundRecord, err := repo.dao.GetInboundRecordOfWarehouseIdAndItemId(wareid, itemid)
	return InboundRecord, err
}

func (repo *CachedDaoRepository) FindOutboundRecordByWarehouseId(wareid uint) ([]domain.OutboundRecord, error) {
	OutboundRecord, err := repo.dao.GetOutboundRecordOfWarehouseId(wareid)
	return OutboundRecord, err
}

func (repo *CachedDaoRepository) FindOutboundRecordByWarehouseIdAndItemId(wareid, itemid uint) ([]domain.OutboundRecord, error) {
	OutboundRecord, err := repo.dao.GetOutboundRecordOfWarehouseIdAndItemId(wareid, itemid)
	return OutboundRecord, err
}
