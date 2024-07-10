package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Page interface {
	Login(c *gin.Context)
}

type page struct{}

func (p *page) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func NewPage() Page {
	return &page{}
}
