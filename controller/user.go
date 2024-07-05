package controller

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"project/domain"
	"project/response"
)

type User interface {
	CreateUser(c *gin.Context)
	GetUserById(c *gin.Context)
	ChangeUserName(c *gin.Context)
	ChangeUserPassword(c *gin.Context)
	ChangeUserPermission(c *gin.Context)
	LoginUser(c *gin.Context)
}

func (s *controller) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailMsg(c, "获取用户信息失败")
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	user.Password = string(hashedPassword)

	if err := s.ser.CreateUser(user); err != nil {
		response.FailMsg(c, "创建用户失败")
		return
	}

	response.Ok(c)
	return
}

func (s *controller) GetUserById(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailMsg(c, "获取用户信息失败")
		return
	}

	user, err := s.ser.GetUserById(user.ID)
	if err != nil {
		response.FailMsg(c, "查询信息失败")
		return
	}

	response.OkData(c, user)
	return
}

func (s *controller) ChangeUserName(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailMsg(c, "获取用户信息失败")
		return
	}

	if err := s.ser.ChangeUserName(user); err != nil {
		response.FailMsg(c, "更改失败")
		return
	}

	response.Ok(c)
	return
}

func (s *controller) ChangeUserPassword(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailMsg(c, "获取用户信息失败")
		return
	}

	if err := s.ser.ChangeUserPassword(user); err != nil {
		response.FailMsg(c, "更改失败")
		return
	}

	response.Ok(c)
	return
}

func (s *controller) ChangeUserPermission(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailMsg(c, "获取用户信息失败")
		return
	}

	if err := s.ser.ChangeUserPermission(user); err != nil {
		response.FailMsg(c, "更改失败")
		return
	}

	response.Ok(c)
	return
}

func (s *controller) LoginUser(c *gin.Context) {
	var user domain.User
	c.ShouldBindJSON(&user)
	u, err := s.ser.GetUserByName(user.Name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有找到匹配的记录
			response.FailMsg(c, "该用户不存在")
			return
		} else {
			// 其他查询错误
			response.FailMsg(c, "other")
			//fmt.Printf("查询错误: %s\n", err.Error())
			//fmt.Println(user)
			return
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		response.FailMsg(c, "密码错误")
		return
	} else {
		response.Ok(c)
		return
	}
}
