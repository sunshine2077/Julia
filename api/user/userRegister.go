package user

type UserRegister struct {
	UserName string `form:"uname"  json:"uname"  binding:"required,min=5,max=10"`
	UserPwd  string `form:"upwd"   json:"upwd"   binding:"required,min=5,max=20"`
	Gender   string `form:"gender" json:"gender"`
	Desc     string `form:"desc"   json:"desc"`
	Phone    string `form:"phone"  json:"phone"  binding:"len=11|len=0"`
	Mail     string `form:"mail"   json:"mail"`
}
