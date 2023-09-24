package api

import (
	"fmt"
	"log"
	"net/http"
	"web/api/common"
	"web/db"

	"github.com/gin-gonic/gin"
)

// File upload and save to qdrant:
func UploadSingle(ctx *gin.Context) {
	// Step1. load the file
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Printf("upload file error, msg=%s\n", err.Error())
		ctx.JSON(http.StatusInternalServerError, &common.CommonResponse{
			ErrCode: common.UPLOAD_FAILED,
			Msg:     "err",
			Data:    err.Error(),
		})
		return
	}
	// Step2. open the file
	reader, err := file.Open()
	defer reader.Close()
	content, err := db.TikaClient.Parse(ctx, reader)
	if err != nil {
		log.Printf("analysis file error, msg=%s\n", err.Error())
		ctx.JSON(http.StatusInternalServerError, &common.CommonResponse{
			ErrCode: common.UPLOAD_FAILED,
			Msg:     "err",
			Data:    err.Error(),
		})
		return
	}
	// Step3. translate to the embedding:
	fmt.Println(content)

	// Step4. Save to the qdrant
	ctx.JSON(http.StatusOK, &common.CommonResponse{
		ErrCode: common.CODE_SUCCESS,
		Msg:     "ok",
	})
}
