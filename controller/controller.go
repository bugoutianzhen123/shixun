package controller

import (
	"project/pkg/logger"
	"project/server"
)

const (
	nop = "权限不足"
	fc  = "创建失败"
	fa  = "解析失败"
)

type Controller interface {
	User
}

type controller struct {
	ser server.Server
	l   logger.Logger
}

func NewContrpller(ser server.Server, l logger.Logger) Controller {
	return &controller{ser: ser, l: l}
}
