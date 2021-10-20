package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/user-web/middlewares"
	"mxshop_api/user-web/router"
	"net/http"
)

func Routers() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	r.Use(middlewares.Cors())

	ApiGroup := r.Group("/u/v1")

	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)

	return r
}
