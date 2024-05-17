package routers

import (
	"qrninja/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/qr", controllers.StoreQRCode)
	r.GET("/qr", controllers.FetchQRCodes)
	r.DELETE("/qr/:id", controllers.DeleteQRCode)

	return r
}
