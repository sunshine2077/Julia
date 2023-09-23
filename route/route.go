package route

import (
	"web/api"

	"github.com/gin-gonic/gin"
)

func LoadRoute() *gin.Engine {
	router := gin.New()
	resource := router.Group("/resource")
	{
		resource.POST("/upload", api.UploadSingle)
	}
	probe := router.Group("/probe")
	{
		probe.GET("/liveness", api.Liveness)
	}
	return router
}
