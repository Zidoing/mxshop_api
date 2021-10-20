package global

import (
	ut "github.com/go-playground/universal-translator"
	"mxshop_api/goods-web/config"
	"mxshop_api/goods-web/proto"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	NacosConfig *config.NacosConfig = &config.NacosConfig{}

	GoodsSrvClient proto.GoodsClient
)
