package admin

import (
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"hertz/app/models"
	"time"
)

const AUTH_TOKEN_KEY string = "admin_auth_token"
const AUTH_TOKEN_MAX_AGE int = 3600 * 24 * 7
const LOGIN_PATH = "/admin/login"
const HOME_PATH = "/admin"

// 验证信息结构体
type Auth struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
}

func GenerateToken(id int, nickname string, phone string) (token string, err error) {
	jwt := NewJWT()
	claims := &CustomClaims{
		id,
		nickname,
		phone,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1),                     // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + TokenRefreshEffective), // 过期时间 一小时
			Issuer:    "hihozhou",                                       //签名的发行者
		},
	}
	return jwt.GenerateToken(claims)
}

//生成token
func GenerateAdminToken(admin *models.Admin) (string, error) {
	return GenerateToken(admin.ID, admin.Nickname, admin.Phone)
}

// 登录操作
func Login(c *gin.Context, admin *models.Admin) (data gin.H, err error) {
	token, err := GenerateAdminToken(admin)
	if err != nil {
		return data, err;
	}
	//保存到cookie中
	LoginSign(c, token)
	data = gin.H{"url": HOME_PATH}
	return data, nil
	//return gin.H{
	//	"admin": gin.H{
	//		"id":       admin.ID,
	//		"phone":    admin.Phone,
	//		"nickname": admin.Nickname,
	//	},
	//	"token": token,
	//}
}

func LoginSign(c *gin.Context, token string) {
	fmt.Println("AUTH_TOKEN_MAX_AGE")
	fmt.Println(AUTH_TOKEN_MAX_AGE)
	c.SetCookie(AUTH_TOKEN_KEY, token, AUTH_TOKEN_MAX_AGE, "/", c.Request.Host, false, true)
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

func Logout(c *gin.Context) {
	c.SetCookie(AUTH_TOKEN_KEY, "", -1, "/", c.Request.Host, false, true)
}
