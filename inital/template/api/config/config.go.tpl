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
    CacheRedis cache.CacheConf

    {{.Service}}Rpc zrpc.RpcClientConf
    Captcha struct {
        KeyLong   int // 验证码长度
        ImgWidth  int // 验证码宽度
        ImgHeight int // 验证码高度
    }
}
