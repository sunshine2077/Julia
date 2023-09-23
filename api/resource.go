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
	name := ctx.PostForm("name")
	file, err := ctx.FormFile(name)
	if err != nil {
		log.Printf("upload file error, msg=%s\n", err.Error())
		ctx.JSON(http.StatusInternalServerError, &common.CommonResponse{
			ErrCode: common.UPLOAD_FAILED,
			Msg:     "err",
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
		})
		return
	}
	// Step3. translate to the embedding:

	// Step4. Save to the qdrant
	fmt.Println(content)
	ctx.JSON(http.StatusOK, &common.CommonResponse{
		ErrCode: common.CODE_SUCCESS,
		Msg:     "ok",
	})
}
