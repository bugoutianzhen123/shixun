package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"os"
)

type ServiceStart interface {
	InitServer() error
}

type service struct {
	con Controller
	p   Page
}

func (cont *service) InitServer() error {
	r := gin.Default()
	initlog(r)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:63342"}, // 允许前端地址
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.LoadHTMLGlob("frontend/*")

	HtmlGroup := r.Group("/page")
	{
		HtmlGroup.GET("/login", cont.p.Login)
	}

	UserGroup := r.Group("/user")
	{
		UserGroup.POST("/register", cont.con.CreateUser)
		UserGroup.POST("/login", cont.con.LoginUser)
		UserGroup.POST("/changepassword", cont.con.ChangeUserPassword)
		UserGroup.POST("/changeusername", cont.con.ChangeUserName)
		UserGroup.POST("/changepermission", cont.con.ChangeUserPermission)
		UserGroup.GET("/getinfo", cont.con.GetUserById)
		UserGroup.GET("/refreshHandler", cont.con.RefreshHandler) //刷新token
	}

	ItemGroup := r.Group("/item")
	{
		ItemGroup.POST("/createitem", cont.con.CreateItem)
		ItemGroup.POST("/createware", cont.con.CreateWarehouser)
		ItemGroup.POST("/createinb", cont.con.CreateInboundRecord)
		ItemGroup.POST("/createoutb", cont.con.CreateOutboundRecord)
		ItemGroup.POST("/deleteitem", cont.con.DeleteItem)
		ItemGroup.POST("/deleteware", cont.con.DeleteWarehouse)
		ItemGroup.POST("/getitem", cont.con.FindItem)
		ItemGroup.POST("/getware", cont.con.FindWarehouse)
		ItemGroup.POST("/getinb", cont.con.FindInboundRecord)
		ItemGroup.POST("/getoutb", cont.con.FindOutboundRecord)
		ItemGroup.POST("/getinven", cont.con.FindInventory)
	}

	err := r.Run(":8088")
	return err
}

func NewService(con Controller, p Page) ServiceStart {
	return &service{con: con, p: p}
}

func initlog(r *gin.Engine) {
	// 创建一个日志记录器
	logger := zap.NewExample().Sugar()
	defer logger.Sync() // 释放资源
	// 将GIN框架日志记录到文件中
	file, err := os.OpenFile("gin.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Fatalf("无法打开或创建日志文件：%s", err.Error())
	}
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout) // 将日志同时输出到文件和控制台
}
