package model

import "github.com/jinzhu/gorm"

// 用户结构体
type User struct {
	gorm.Model
	UserName string `gorm:"column:uname;type:varchar(20);not null;uniqueIndex"`
	PassWord string `gorm:"column:passwd;type:varchar(35);not null"`
	Gender   string `gorm:"column:gender;type:varchar(10);default:'未知'"`
	Desc     string `gorm:"column:desc;type:varchar(100);default:'这个人很懒，什么都没有写'"`
	Phone    string `gorm:"column:phone;type:varchar(15)"`
	Mail     string `gorm:"column:mail;type:varchar(20)"`
	Sault    string `gorm:"column:sault;type:varchar(10)"`
	Level    int    `gorm:"column:level"`
	Coin     int    `gorm:"column:coin"`
}

func (u *User) Tablename() string {
	return "user_t"
}
