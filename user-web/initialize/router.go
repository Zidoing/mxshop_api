package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/user-web/router"
)

func Routers() *gin.Engine {
	r := gin.Default()

	ApiGroup := r.Group("/u/v1")
	router.InitUserRouter(ApiGroup)

	return r
}
