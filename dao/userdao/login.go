package userdao

// import (
// 	"log"
// 	"web/model"
// )

// // Get Sault of User
// func GetSaultByName(name string) (string, error) {
// 	var user model.User
// 	err := config.MysqlClient.Debug().Select("sault").First(&user, "uname=?", name).Error
// 	return user.Sault, err
// }

// // Login
// func UserLogin(user *model.User) (bool, error) {
// 	var count int
// 	err := config.MysqlClient.Debug().Model(&model.User{}).Where("uname=? and passwd=?", user.UserName, user.PassWord).Count(&count).Error
// 	log.Println("$$", count)
// 	return count == 1, err
// }
