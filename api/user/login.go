package user

// func LoginAPI(ctx *gin.Context) {
// 	login := UserLogin{}
// 	log.Println("$$", ctx.PostForm("uname"))
// 	err := ctx.ShouldBind(&login)
// 	if err != nil {
// 		log.Println("登录失败，详细信息:", err.Error())
// 		ctx.JSON(http.StatusOK, common.LoginParaUnMatched(err))
// 		return
// 	}
// 	succ, err := userService.LoginService(&model.User{
// 		UserName: login.UserName,
// 		PassWord: login.UserPwd,
// 	})
// 	if err != nil {
// 		ctx.JSON(http.StatusOK, common.LoginFailed(err))
// 		return
// 	}
// 	if !succ {
// 		ctx.JSON(http.StatusOK, common.LoginFailed(errors.New("未检索到有效密码!")))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, common.LoginSuccess())
// }
