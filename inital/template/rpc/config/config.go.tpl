package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	zrpc.RpcServerConf

    Database struct {
        DriverName string
        DataSource string
    }

    CacheRedis cache.ClusterConf
}
