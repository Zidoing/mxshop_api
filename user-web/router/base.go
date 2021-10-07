package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mxshop_api/user-web/api"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base")
	zap.S().Info("配置用户相关的router")
	{
		BaseRouter.GET("captcha", api.GetCaptCha)
		BaseRouter.POST("send_sms", api.SendSms)
	}
}
