package validator

//登录验证器
type Login struct {
	Phone    string `form:"phone" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}
