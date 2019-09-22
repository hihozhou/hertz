package admin

import (
	"errors"
	jwtgo "github.com/dgrijalva/jwt-go"
	"hertz/pkg/setting"
	"time"
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenVerifyEffective  int64 = 3600
	TokenRefreshEffective int64 = 3600 * 24 * 7
	TokenExpired          error = errors.New("Token is expired")        //token过期
	TokenNotValidYet      error = errors.New("Token not active yet")    //token 已经失效
	TokenMalformed        error = errors.New("That's not even a token") //
	TokenInvalid          error = errors.New("Couldn't handle this token:")
)

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// 获取signKey
func GetSignKey() string {
	return setting.JwtSecret
}

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	jwtgo.StandardClaims
}

// GenerateToken 生成一个token
func (jwt *JWT) GenerateToken(claims *CustomClaims) (string, error) {
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	return token.SignedString(jwt.SigningKey)
}

// 解析token
func (jwt *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwtgo.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwtgo.Token) (interface{}, error) {
			return jwt.SigningKey, nil
		},
	)
	if err != nil {
		if ve, ok := err.(*jwtgo.ValidationError); ok {
			if ve.Errors&jwtgo.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwtgo.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwtgo.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
		return nil, err

	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

//刷新claims刷新token
func (jwt *JWT) RefreshClaimsToken(claims *CustomClaims) (string, error) {
	claims.StandardClaims.NotBefore = time.Now().Add(-1 * time.Second).Unix()
	claims.StandardClaims.ExpiresAt = time.Now().Add(time.Duration(TokenRefreshEffective)).Unix()
	return jwt.GenerateToken(claims)
}

// 更新token
func (jwt *JWT) RefreshToken(tokenString string) (string, error) {

	claims, err := jwt.ParseToken(tokenString)
	if err != nil {
		return "", err
	}
	token, err := jwt.RefreshClaimsToken(claims)
	return token, err
}
