package controller

import (
	"project/pkg/logger"
	"project/server"
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
