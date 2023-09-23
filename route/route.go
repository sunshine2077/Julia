package route

import (
	"web/api"

	"github.com/gin-gonic/gin"
)

func LoadRoute() *gin.Engine {
	router := gin.New()
	rg := router.Group("/resource")
	{
		rg.POST("/upload", api.UploadSingle)
	}
	return router
}
