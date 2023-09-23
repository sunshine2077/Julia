package user

// import (
// 	"log"
// 	"net/http"
// 	"web/model"
// 	"web/service/userService"

// 	"github.com/gin-gonic/gin"
// )

// func RegisterAPI(ctx *gin.Context) {
// 	register := UserRegister{}
// 	err := ctx.ShouldBind(&register)
// 	if err != nil {
// 		log.Println("注册失败，详细信息:", err.Error())
// 		ctx.JSON(http.StatusBadRequest, "注册失败！填写信息有误")
// 		return
// 	}
// 	err = userService.RegisterService(&model.User{
// 		UserName: register.UserName,
// 		PassWord: register.UserPwd,
// 		Gender:   register.Gender,
// 		Desc:     register.Desc,
// 		Phone:    register.Phone,
// 		Mail:     register.Mail,
// 	})
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, "注册失败")
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, "注册成功")
// }
