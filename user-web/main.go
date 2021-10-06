package main

import (
	"fmt"
	"go.uber.org/zap"
	"mxshop_api/user-web/global"
	"mxshop_api/user-web/initialize"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	err := initialize.InitTrans("zh")
	if err != nil {
		panic(err)
	}

	r := initialize.Routers()
	err = r.Run(fmt.Sprintf(":%d", global.ServerConfig.Port))
	zap.S().Infof("启动服务器 端口:%d", global.ServerConfig.Port)
	if err != nil {
		zap.S().Panic("启动失败:", err.Error())
	}
}
