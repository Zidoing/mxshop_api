package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/user-web/middlewares"
	"mxshop_api/user-web/router"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())

	ApiGroup := r.Group("/u/v1")

	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)

	return r
}
