package server

import "project/domain"

type ItemServer interface {
	CreateItem(item domain.Item) error
	CreateWareHouse(warehouse domain.Warehouse) error
	CreateInboundRecord(in domain.InboundRecord) error
	CreateOutboundRecord(out domain.OutboundRecord) error
	DeleteItem(item domain.Item) error
	DeleteWarehouse(warehouse domain.Warehouse) error
	FindWarehouse(wareid uint, warename string) ([]domain.Warehouse, error)
	FindItem(itemid uint, itemname string) ([]domain.Item, error)
	FindInventory(wareid, itemid uint) ([]domain.Inventory, error)
	FindInboundRecord(wareid, itemid uint) ([]domain.InboundRecord, error)
	FindOutboundRecord(wareid, itemid uint) ([]domain.OutboundRecord, error)
}

func (s *server) CreateItem(item domain.Item) error {
	err := s.rep.CreateItem(item)
	return err
}

func (s *server) CreateWareHouse(warehouse domain.Warehouse) error {
	err := s.rep.CreateWareHouse(warehouse)
	return err
}

func (s *server) CreateInboundRecord(in domain.InboundRecord) error {
	err := s.rep.CreateInboundRecord(in)
	return err
}

func (s *server) CreateOutboundRecord(out domain.OutboundRecord) error {
	err := s.rep.CreateOutboundRecord(out)
	return err
}

func (s *server) DeleteItem(item domain.Item) error {
	err := s.rep.DeleteItem(item)
	return err
}

func (s *server) DeleteWarehouse(warehouse domain.Warehouse) error {
	err := s.rep.DeleteWarehouse(warehouse)
	return err
}

func (s *server) FindWarehouse(wareid uint, warename string) ([]domain.Warehouse, error) {
	if warename != "" {
		return s.rep.FindWarehouseByName(warename)
	}
	if wareid != 0 {
		ware, err := s.rep.FindWarehouseById(wareid)
		warehouses := make([]domain.Warehouse, 0)
		warehouses = append(warehouses, ware)
		return warehouses, err
	} else {
		return s.rep.FindWarehouse()
	}
}

func (s *server) FindItem(itemid uint, itemname string) ([]domain.Item, error) {
	if itemname != "" {
		return s.rep.FindItemByName(itemname)
	}
	if itemid != 0 {
		item, err := s.rep.FindItemById(itemid)
		items := make([]domain.Item, 0)
		items = append(items, item)
		return items, err
	} else {
		return s.rep.FindItem()
	}
}

func (s *server) FindInventory(wareid, itemid uint) ([]domain.Inventory, error) {
	if wareid != 0 {
		if itemid != 0 {
			Inventory, err := s.rep.FindInventoryByWarehouseIdAndItemId(wareid, itemid)
			inventorys := make([]domain.Inventory, 0)
			inventorys = append(inventorys, Inventory)
			return inventorys, err
		} else {
			return s.rep.FindInventoryByWarehouseId(wareid)
		}
	} else {
		if itemid != 0 {
			return s.rep.FindInventoryByItemId(itemid)
		} else {
			return s.rep.FindInventory()
		}
	}
}

func (s *server) FindInboundRecord(wareid, itemid uint) ([]domain.InboundRecord, error) {
	if wareid != 0 {
		if itemid != 0 {
			return s.rep.FindInboundRecordByWarehouseIdAndItemId(wareid, itemid)
		} else {
			return s.rep.FindInboundRecordByWarehouseId(wareid)
		}
	} else {
		return s.rep.FindInboundRecord()
	}
}

func (s *server) FindOutboundRecord(wareid, itemid uint) ([]domain.OutboundRecord, error) {
	if wareid != 0 {
		if itemid != 0 {
			return s.rep.FindOutboundRecordByWarehouseIdAndItemId(wareid, itemid)
		} else {
			return s.rep.FindOutboundRecordByWarehouseId(wareid)
		}
	} else {
		return s.rep.FindOutboundRecord()
	}
}
