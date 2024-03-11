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
    CacheRedis cache.ClusterConf

    {{.Service}}Rpc zrpc.RpcClientConf
    AdminRpc zrpc.RpcClientConf
    FsmRpc   zrpc.RpcClientConf

}
