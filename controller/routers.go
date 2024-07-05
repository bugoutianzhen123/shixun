package controller

import "github.com/gin-gonic/gin"

type ServiceStart interface {
	InitServer() error
}

type service struct {
	con User
}

func (cont *service) InitServer() error {

	r := gin.Default()

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

	err := r.Run(":8088")
	return err
}

func NewService(con Controller) ServiceStart {
	return &service{con: con}
}
