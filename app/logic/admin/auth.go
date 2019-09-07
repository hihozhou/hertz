package admin

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"hertz/app/models"
	"time"
)

// 验证信息结构体
type Auth struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
}

func GenerateToken(admin *models.Admin) (string, error) {
	jwt := NewJWT()
	claims := CustomClaims{
		admin.ID,
		admin.Nickname,
		admin.Phone,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "hihozhou",                      //签名的发行者
		},
	}
	return jwt.GenerateToken(claims)
}

// 登录操作
func Login(admin *models.Admin) gin.H {
	token, err := GenerateToken(admin)
	if err != nil {
		panic("创建token失败：" + err.Error())
	}
	return gin.H{
		"user": gin.H{
			"id":       admin.ID,
			"phone":    admin.Phone,
			"nickname": admin.Nickname,
		},
		"token": token,
	}
}

// 获取登录的用户信息
func User(c *gin.Context) (result *CustomClaims) {
	claims, exists := c.Get("claims")
	if !exists {
		panic("找不到登录信息")
	}

	result = claims.(*CustomClaims)
	return result
}
