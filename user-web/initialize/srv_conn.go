package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"mxshop_api/user-web/global"
	"mxshop_api/user-web/proto"
)

type grpcLog struct {
	*logrus.Logger
}

func (l *grpcLog) V(lvl int) bool {
	return true
}

func InitSrvConn() *grpc.ClientConn {
	logger := logrus.New()
	grpclog.SetLoggerV2(&grpcLog{logger})

	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&healthy=true",
			global.ServerConfig.ConsulInfo.Host,
			global.ServerConfig.ConsulInfo.Port,
			global.ServerConfig.UserSrvInfo.Name,
		),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)

	if err != nil {
		zap.S().Fatalf("[InitSrvConn] 连接 [用户服务失败]")
	}

	global.UserSrvClient = proto.NewUserClient(userConn)

	return userConn
}

func InitSrvConn2() {

	logger := logrus.New()
	grpclog.SetLoggerV2(&grpcLog{logger})

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
