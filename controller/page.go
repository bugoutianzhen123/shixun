package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Page interface {
	LoginH(c *gin.Context)
	AdminPanel(c *gin.Context)
}

type page struct{}

func (p *page) LoginH(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (p *page) AdminPanel(c *gin.Context) {
	c.HTML(http.StatusOK, "adminPanel.html", gin.H{})
}

func NewPage() Page {
	return &page{}
}
