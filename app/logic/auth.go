package logic

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"hertz/app/models"
	"time"
)

// 验证信息结构体
type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
}

func GenerateToken(user *models.User) (string, error) {
	jwt := &JWT{
		[]byte("&*(^^fewi23&^"),
	}
	claims := CustomClaims{
		user.ID,
		user.Nickname,
		user.Phone,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "hihozhou",                      //签名的发行者
		},
	}
	return jwt.GenerateToken(claims)
}

// 登录操作
func Login(user *models.User) gin.H {
	token, err := GenerateToken(user)
	if err != nil {
		panic("创建token失败：" + err.Error())
	}
	return gin.H{
		"user":  user,
		"token": token,
	}
}
