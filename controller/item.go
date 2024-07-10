package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"project/domain"
	lojwt "project/pkg/jwt"
	"project/response"
)

type Item interface {
	CreateItem(c *gin.Context)
	CreateWarehouser(c *gin.Context)
	CreateInboundRecord(c *gin.Context)
	CreateOutboundRecord(c *gin.Context)
	DeleteItem(c *gin.Context)
	DeleteWarehouse(c *gin.Context)
	FindWarehouse(c *gin.Context)
	FindItem(c *gin.Context)
	FindInventory(c *gin.Context)
	FindInboundRecord(c *gin.Context)
	FindOutboundRecord(c *gin.Context)
}

type itemmsg struct {
	ItemId      uint `json:"itemid"`
	WarehouseId uint `json:"warehouseid"`
}

func (s *controller) CreateItem(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	msg, err := lojwt.ParseToken(tokenStr)
	if err != nil {
		response.FailMsg(c, "token解析失败")
		return
	}

	if msg.Permission != 0 {
		response.FailMsg(c, "权限不足")
		return
	}

	var item domain.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		response.FailMsg(c, "解析失败")
		fmt.Println("JSON解析失败:", err)
		return
	}

	fmt.Printf("接收到的物品数据: %+v\n", item)

	if err := s.ser.CreateItem(item); err != nil {
		response.FailMsg(c, "创建失败")
		return
	}

	response.Ok(c)
	return
}

func (s *controller) CreateWarehouser(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	msg, err := lojwt.ParseToken(tokenStr)
	if err != nil {
		response.FailMsg(c, fa)
		return
	}

	if msg.Permission != 0 {
		response.FailMsg(c, nop)
		return
	}

	var warehouse domain.Warehouse
	if err := c.ShouldBindJSON(&warehouse); err != nil {
		response.FailMsg(c, fa)
		fmt.Println("JSON解析失败:", err)
		return
	}

	fmt.Printf("接收到的仓库数据: %+v\n", warehouse)

	if err := s.ser.CreateWareHouse(warehouse); err != nil {
		response.FailMsg(c, fc)
		return
	}

	// 返回创建的仓库数据
	response.OkWithData(c, gin.H{
		"warehouse": warehouse,
	})
}

func (s *controller) CreateInboundRecord(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	msg, err := lojwt.ParseToken(tokenStr)
	if err != nil {
		response.FailMsg(c, "token解析失败")
		return
	}

	if msg.Permission != 0 {
		response.FailMsg(c, "权限不足")
		return
	}

	var in domain.InboundRecord
	if err := c.ShouldBindJSON(&in); err != nil {
		fmt.Println("JSON解析失败:", err)
		response.FailMsg(c, "解析失败")
		return
	}

	if in.WarehouseId == 0 || in.ItemId == 0 {
		response.FailMsg(c, "信息错误")
		return
	}

	// 打印解析后的数据
	fmt.Printf("接收到的入库记录数据: %+v\n", in)

	if err := s.ser.CreateInboundRecord(in); err != nil {
		fmt.Println("创建入库记录失败:", err)
		response.FailMsg(c, "创建失败")
		return
	}

	response.Ok(c)
}

func (s *controller) CreateOutboundRecord(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	msg, err := lojwt.ParseToken(tokenStr)
	if err != nil {
		response.FailMsg(c, "token解析失败")
		fmt.Printf("token解析失败: %v\n", err)
		return
	}

	if msg.Permission != 0 {
		response.FailMsg(c, "权限不足")
		fmt.Println("权限不足")
		return
	}

	var out domain.OutboundRecord
	if err := c.ShouldBindJSON(&out); err != nil {
		response.FailMsg(c, "解析失败")
		fmt.Printf("解析出库记录失败: %v\n", err)
		return
	}

	if out.WarehouseId == 0 || out.ItemId == 0 {
		response.FailMsg(c, "信息错误")
		return
	}

	fmt.Printf("接收到的出库记录数据: %+v\n", out)

	if err := s.ser.CreateOutboundRecord(out); err != nil {
		response.FailMsg(c, "创建失败")
		fmt.Printf("创建出库记录失败: %v\n", err)
		return
	}

	response.Ok(c)
	fmt.Println("创建出库记录成功")
}

func (s *controller) DeleteItem(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	msg, err := lojwt.ParseToken(tokenStr)
	if err != nil {
		response.FailMsg(c, fa)
		return
	}

	if msg.Permission != 0 {
		response.FailMsg(c, nop)
		return
	}

	var item domain.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		response.FailMsg(c, fa)
		return
	}

	if err := s.ser.DeleteItem(item); err != nil {
		response.FailMsg(c, fd)
		return
	}

	response.Ok(c)
	return
}

func (s *controller) DeleteWarehouse(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	msg, err := lojwt.ParseToken(tokenStr)
	if err != nil {
		response.FailMsg(c, fa)
		return
	}

	if msg.Permission != 0 {
		response.FailMsg(c, nop)
		return
	}

	var warehouse domain.Warehouse
	if err := c.ShouldBindJSON(&warehouse); err != nil {
		response.FailMsg(c, fa)
		return
	}

	if err := s.ser.DeleteWarehouse(warehouse); err != nil {
		response.FailMsg(c, fd)
		return
	}

	response.Ok(c)
	return
}

func (s *controller) FindWarehouse(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	_, err := lojwt.ParseToken(tokenStr)
	if err != nil {
		response.FailMsg(c, "认证失败")
		return
	}

	var waremsg itemmsg
	if err := c.ShouldBindJSON(&waremsg); err != nil {
		response.FailMsg(c, "解析失败")
		return
	}

	warehouses, err := s.ser.FindWarehouse(waremsg.WarehouseId)
	if err != nil {
		response.FailMsg(c, "查找仓库失败")
		return
	}

	response.OkData(c, warehouses)
}

func (s *controller) FindItem(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	_, err := lojwt.ParseToken(tokenStr)
	if err != nil {
		response.FailMsg(c, fa)
		return
	}

	msg := itemmsg{0, 0}
	if err := c.ShouldBindJSON(&msg); err != nil {
		response.FailMsg(c, fa)
		return
	}

	items, err := s.ser.FindItem(msg.ItemId)
	if err != nil {
		response.FailMsg(c, fs)
		return
	}

	response.OkData(c, items)
	return
}

func (s *controller) FindInventory(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	_, err := lojwt.ParseToken(tokenStr)
	if err != nil {
		response.FailMsg(c, fa)
		return
	}

	msg := itemmsg{0, 0}
	if err := c.ShouldBindJSON(&msg); err != nil {
		response.FailMsg(c, fa)
		return
	}

	inventory, err := s.ser.FindInventory(msg.WarehouseId, msg.ItemId)
	if err != nil {
		response.FailMsg(c, fs)
		return
	}

	response.OkData(c, inventory)
	return
}

func (s *controller) FindInboundRecord(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	per, err := lojwt.ParseToken(tokenStr)
	if err != nil {
		response.FailMsg(c, fa)
		return
	}

	if per.Permission != 0 {
		response.FailMsg(c, nop)
		return
	}

	msg := itemmsg{0, 0}
	if err := c.ShouldBindJSON(&msg); err != nil {
		response.FailMsg(c, fa)
		return
	}

	InboundRecords, err := s.ser.FindInboundRecord(msg.WarehouseId, msg.ItemId)
	if err != nil {
		response.FailMsg(c, fs)
		return
	}

	response.OkData(c, InboundRecords)
	return
}

func (s *controller) FindOutboundRecord(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	per, err := lojwt.ParseToken(tokenStr)
	if err != nil {
		response.FailMsg(c, fa)
		return
	}

	if per.Permission != 0 {
		response.FailMsg(c, nop)
		return
	}

	msg := itemmsg{0, 0}
	if err := c.ShouldBindJSON(&msg); err != nil {
		response.FailMsg(c, fa)
		return
	}

	OutboundRecords, err := s.ser.FindOutboundRecord(msg.WarehouseId, msg.ItemId)
	if err != nil {
		response.FailMsg(c, fs)
		return
	}

	response.OkData(c, OutboundRecords)
	return
}
