package global

import (
	ut "github.com/go-playground/universal-translator"
	"mxshop_api/user-web/config"
	"mxshop_api/user-web/proto"
)

var (
	ServerConfig  = &config.ServerConfig{}
	NacosConfig   = &config.NacosConfig{}
	Trans         ut.Translator
	UserSrvClient proto.UserClient
)
