package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop_api/user-web/global"
	"mxshop_api/user-web/initialize"
	"mxshop_api/user-web/utils"
	mx_validator "mxshop_api/user-web/validator"
)

func main() {
	// 初始化日志
	initialize.InitLogger()
	// 初始化配置
	initialize.InitConfig()
	// 初始化服务
	userConn := initialize.InitSrvConn()

	defer func(userConn *grpc.ClientConn) {
		err := userConn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(userConn)

	debug := initialize.GetEnvInfo("MXSHOP_DEBUG")
	if !debug {
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}

	// 初始化翻译
	err := initialize.InitTrans("zh")
	if err != nil {
		panic(err)
	}
	// 注册验证器
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		_ = v.RegisterValidation("mobile", mx_validator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	// 初始化路由
	r := initialize.Routers()
	err = r.Run(fmt.Sprintf(":%d", global.ServerConfig.Port))
	zap.S().Infof("启动服务器 端口:%d", global.ServerConfig.Port)
	if err != nil {
		zap.S().Panic("启动失败:", err.Error())
	}
}
