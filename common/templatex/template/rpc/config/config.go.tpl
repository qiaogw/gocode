package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

    DbConf struct {
		Driver   string
		Host     string
		Port     int64
		User     string
		Password string
		Db       string
		Schema   string `json:"schema,optional"`
		Config   string `json:"config,optional"`
	}

    CacheRedis cache.ClusterConf

}
