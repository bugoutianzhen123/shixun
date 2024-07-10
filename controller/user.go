package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"project/domain"
	lojwt "project/pkg/jwt"
	"project/response"
	"time"
)

type User interface {
	CreateUser(c *gin.Context)
	GetUserById(c *gin.Context)
	ChangeUserName(c *gin.Context)
	ChangeUserPassword(c *gin.Context)
	ChangeUserPermission(c *gin.Context)
	LoginUser(c *gin.Context)
	RefreshHandler(c *gin.Context)
}

type usermsg struct {
	Name        string `json:"username"`
	Password    string `json:"password"`
	Prepassword string `json:"prepassword"`
	Newpassword string `json:"newPassword"`
}

func (s *controller) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailMsg(c, "获取用户信息失败/用户信息出错")
		return
	}

	if user.Name == "" || user.Password == "" {
		response.FailMsg(c, "用户名/密码不能为空")
		return
	}

	if _, err := s.ser.GetUserByName(user.Name); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有找到匹配的记录
			//进行注册操作
			// 加密密码
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				response.FailMsg(c, "加密失败")
				return
			}
			user.Password = string(hashedPassword)

			if err := s.ser.CreateUser(user); err != nil {
				response.FailMsg(c, "创建用户失败")
				return
			}
			response.Ok(c)
			return
		} else {
			// 其他查询错误
			fmt.Printf("查询错误: %s\n", err.Error())
			response.FailMsg(c, "未知错误")
			return
		}
	} else {
		//找到匹配用户
		response.FailMsg(c, "该用户已存在")
		return
	}
}

func (s *controller) GetUserById(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	msg, err := lojwt.ParseToken(tokenStr)
	if err != nil {
		response.FailMsg(c, "token解析失败")
		return
	}

	user, err := s.ser.GetUserById(msg.ID)
	if err != nil {
		response.FailMsg(c, "查询信息失败")
		return
	}

	response.OkData(c, user)
	return
}

func (s *controller) ChangeUserName(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	msg, err := lojwt.ParseToken(tokenStr)
	if err != nil {
		response.FailMsg(c, fa)
		return
	}

	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailMsg(c, "获取用户信息失败")
		return
	}

	user.ID = msg.ID
	if err := s.ser.ChangeUserName(user); err != nil {
		response.FailMsg(c, "更改失败")
		return
	}

	response.Ok(c)
	return
}

func (s *controller) ChangeUserPassword(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	msg, err := lojwt.ParseToken(tokenStr)
	if err != nil {
		response.FailMsg(c, "token解析失败")
		return
	}

	var usermsg usermsg
	if err := c.ShouldBindJSON(&usermsg); err != nil {
		response.FailMsg(c, "获取用户信息失败")
		return
	}

	user, err := s.ser.GetUserById(msg.ID)
	if err != nil {
		response.FailMsg(c, "查询用户信息失败")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(usermsg.Prepassword))
	if err != nil {
		response.FailMsg(c, "原密码错误")
		return
	}

	newpassword, err := bcrypt.GenerateFromPassword([]byte(usermsg.Newpassword), bcrypt.DefaultCost)
	if err != nil {
		response.FailMsg(c, "加密失败")
		return
	}

	user.Password = string(newpassword)
	if err := s.ser.ChangeUserPassword(user); err != nil {
		response.FailMsg(c, "更改失败")
		return
	}

	response.Ok(c)
	return
}

func (s *controller) ChangeUserPermission(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	msg, err := lojwt.ParseToken(tokenStr)
	if err != nil {
		response.FailMsg(c, "token解析失败")
		return
	}

	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailMsg(c, "获取用户信息失败")
		return
	}

	user.ID = msg.ID
	if err := s.ser.ChangeUserPermission(user); err != nil {
		response.FailMsg(c, "更改失败")
		return
	}

	response.Ok(c)
	return
}

func (s *controller) LoginUser(c *gin.Context) {
	var user usermsg
	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailMsg(c, "获取用户信息失败")
		return
	}
	fmt.Println(user)
	u, err := s.ser.GetUserByName(user.Name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有找到匹配的记录
			response.FailMsg(c, "该用户不存在")
			return
		} else {
			// 其他查询错误
			response.FailMsg(c, "other")
			return
		}
	}
	fmt.Println(u.Password)
	fmt.Println(user.Password)

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		response.FailMsg(c, "密码错误")
		return
	}

	if token, err := lojwt.GenerateToken(u.ID, u.Permission); err != nil {
		response.FailMsg(c, "token生成失败")
		return
	} else {
		c.Header("Authorization", token)

		response.Ok(c)
		return
	}
}

func (s *controller) RefreshHandler(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	claims, err := lojwt.ParseToken(tokenStr)
	if err != nil {
		response.FailMsg(c, "解析失败")
		return
	}

	// 检查是否快要过期，如果是，则生成一个新的token
	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		response.FailMsg(c, "未临近过期")
		return
	}

	newToken, err := lojwt.GenerateToken(claims.ID, claims.Permission)
	if err != nil {
		response.FailMsg(c, "无法重新生成roken")
		return
	}
	c.Header("Authorization", newToken)

	response.Ok(c)
	return
}
