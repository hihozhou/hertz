package logic

import jwtgo "github.com/dgrijalva/jwt-go"

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	jwtgo.StandardClaims
}

// GenerateToken 生成一个token
func (jwt *JWT) GenerateToken(claims CustomClaims) (string, error) {
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	return token.SignedString(jwt.SigningKey)
}
