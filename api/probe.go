package api

import (
	"net/http"
	"web/api/common"

	"github.com/gin-gonic/gin"
)

// Liveness Probe
func Liveness(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, &common.CommonResponse{
		ErrCode: common.CODE_SUCCESS,
		Msg:     "ok",
	})
	return
}
