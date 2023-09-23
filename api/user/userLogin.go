package user

type UserLogin struct {
	UserName string `form:"uname"  json:"uname"  binding:"required,min=5,max=10"`
	UserPwd  string `form:"upwd"   json:"upwd"   binding:"required,min=5,max=20"`
}
