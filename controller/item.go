package controller

import (
	"github.com/gin-gonic/gin"
	"project/domain"
	lojwt "project/pkg/jwt"
	"project/response"
)

type Item interface {
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
		return
	}

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
		return
	}

	if err := s.ser.CreateWareHouse(warehouse); err != nil {
		response.FailMsg(c, fc)
		return
	}
}

func (s *controller) CreateInboundRecord(c *gin.Context) {
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

	var in domain.InboundRecord
	if err := c.ShouldBindJSON(&in); err != nil {
		response.FailMsg(c, fa)
		return
	}
}

func (s *controller) DeleteItem(c *gin.Context) {}
