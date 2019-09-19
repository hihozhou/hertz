package admin

//登录验证器
type LoginValidator struct {
	Phone    string `form:"phone" json:"phone" xml:"phone"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}
