package userService

// func LoginService(user *model.User) (bool, error) {
// 	sault, err := userdao.GetSaultByName(user.UserName)
// 	if strings.TrimSpace(sault) == "" {
// 		return false, errors.New("查询sault失败！")
// 	}
// 	if err != nil {
// 		log.Println("查询sault失败!,详细信息：", err.Error())
// 		return false, err
// 	}
// 	user.PassWord = util.EncodingPwd(sault, user.PassWord)
// 	return userdao.UserLogin(user)
// }
