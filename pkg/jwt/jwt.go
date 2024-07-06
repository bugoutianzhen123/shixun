package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

var stSignKey = []byte(viper.GetString("jwt.signkey"))

type JwtCustomClaims struct {
	ID         uint
	Permission uint
	jwt.RegisteredClaims
}

// GenerateToken 生成Token
func GenerateToken(id uint, permission uint) (string, error) {
	// 初始化
	iJwtCustomClaims := JwtCustomClaims{
		ID:         id,
		Permission: permission,
		RegisteredClaims: jwt.RegisteredClaims{
			// 设置过期时间 在当前基础上 添加一个小时后 过期
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.TokenExpire") * time.Millisecond)),
			// 颁发时间 也就是生成时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			//主题
			Subject: "Token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustomClaims)
	return token.SignedString(stSignKey)
}

// ParseToken 获取并解析token
func ParseToken(tokenStr string) (JwtCustomClaims, error) {
	// 声明一个空的数据声明
	iJwtCustomClaims := JwtCustomClaims{}
	//ParseWithClaims是NewParser().ParseWithClaims()的快捷方式
	//第一个值是token ，
	//第二个值是我们之后需要把解析的数据放入的地方，
	//第三个值是Keyfunc将被Parse方法用作回调函数，以提供用于验证的键。函数接收已解析但未验证的令牌。
	token, err := jwt.ParseWithClaims(tokenStr, &iJwtCustomClaims, func(token *jwt.Token) (interface{}, error) {
		return stSignKey, nil
	})

	// 判断 是否为空 或者是否无效只要两边有一处是错误 就返回无效token
	if err != nil && !token.Valid {
		err = errors.New("invalid Token")
	}
	return iJwtCustomClaims, err
}
