package authx

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
)

var configFile = "common/authx/auth.yaml"
var authConfig = new()

type Config struct {
	AdminRpc zrpc.RpcClientConf
}

func new() *Config {
	var c Config

	conf.MustLoad(configFile, &c)
	return &c
}
