package config

import (
    "github.com/zeromicro/go-zero/zrpc"
    "github.com/zeromicro/go-zero/core/stores/cache"
    "github.com/zeromicro/go-zero/rest"
)

type Config struct {
    rest.RestConf
    Auth struct {
        AccessSecret string
        AccessExpire int64
    }
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

    {{.Service}}Rpc zrpc.RpcClientConf
}
