package route

import "github.com/gin-gonic/gin"

func tokenCheck(ctx *gin.Context) {

	ctx.Next()
}
