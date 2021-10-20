package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"mxshop_api/goods-web/global"
	"mxshop_api/goods-web/proto"
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

	global.GoodsSrvClient = proto.NewGoodsClient(userConn)

	return userConn
}
