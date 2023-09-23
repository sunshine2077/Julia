package userService

// import (
// 	"log"
// 	"web/dao/userdao"
// 	"web/model"
// 	"web/util"
// )

// func RegisterService(user *model.User) error {
// 	user.Sault = util.GetSault(5)
// 	user.PassWord = util.EncodingPwd(user.Sault, user.PassWord)
// 	err := userdao.RegisterDao(user)
// 	if err != nil {
// 		log.Println("用户注册失败！信息:", err.Error())
// 		return err
// 	} else {
// 		return nil
// 	}
// }
