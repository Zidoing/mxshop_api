package main

import (
	"fmt"
	"go.uber.org/zap"
	"mxshop_api/user-web/initialize"
)

func main() {
	port := 8021
	initialize.InitLogger()

	r := initialize.Routers()
	err := r.Run(fmt.Sprintf(":%d", port))
	zap.S().Infof("启动服务器 端口:%d", port)
	if err != nil {
		zap.S().Panic("启动失败:", err.Error())
	}
}
