package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop_api/user-web/global"
	"mxshop_api/user-web/proto"
)

func InitSrvConn() {

	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d",
		global.ServerConfig.ConsulInfo.Host,
		global.ServerConfig.ConsulInfo.Port,
	)

	userSrvHost := ""
	userSrvPort := 0
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().ServicesWithFilter(
		fmt.Sprintf("Service == \"%s\"", global.ServerConfig.UserSrvInfo.Name),
	)

	if err != nil {
		panic(err)
	}

	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}

	if userSrvHost == "" {
		zap.S().Fatal("[InitSrvConn fail]")
		return
	}

	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d",
		userSrvHost,
		userSrvPort,
	), grpc.WithInsecure())

	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 [用户服务失败]", "msg", err.Error())
	}

	global.UserSrvClient = proto.NewUserClient(userConn)
}
