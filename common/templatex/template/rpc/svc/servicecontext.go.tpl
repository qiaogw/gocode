package svc

import (
	"github.com/go-redis/redis"
	"github.com/qiaogw/gorm-cache/cache"
	cacheConfig "github.com/qiaogw/gorm-cache/config"
	"{{.ParentPkg}}/model"
	"github.com/qiaogw/gocode/common/gormx"
	"{{.ParentPkg}}/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/{{.DriverName}}"
	redisX "github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	CacheRedis     *redisX.Redis
	{{- range .Tables }}
	{{.Table}}Model model.{{.Table}}Model
	{{- end }}
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := gormx.GetDsn(c.DbConf.Driver,
		c.DbConf.Host,
		c.DbConf.Port,
		c.DbConf.User,
		c.DbConf.Password,
		c.DbConf.Db,
		c.DbConf.Schema,
		c.DbConf.Config)

	conn := sqlx.NewSqlConn(c.DbConf.Driver, dsn)
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	redisClient := redis.NewClient(&redis.Options{
		Addr: c.Redis.Host,
		Password: c.Redis.Pass,
	})
	_ = db.Use(&gormx.ZeroGorm{})
	caches, err := cache.NewGorm2Cache(&cacheConfig.CacheConfig{
	CacheLevel:           cacheConfig.CacheLevelAll,
	CacheStorage:         cacheConfig.CacheStorageRedis,
	RedisConfig:          cache.NewRedisConfigWithClient(redisClient),
	InvalidateWhenUpdate: true,       // when you create/update/delete objects, invalidate cache
	CacheTTL:             5000 * 3600, // 5000 ms
	CacheMaxItemCnt:      5,           // if length of objects retrieved one single time
	// exceeds this number, then don't cache
	})
	if err != nil {
	logx.Errorf("setup all cache error: %v", err)
	} else {
	err = db.Use(caches) // use gorm plugin
	if err != nil {
	logx.Errorf("setup all cache error: %v", err)
	}
	}




	redisConf := redisX.RedisConf{
	Host: c.Redis.Host,
	Pass: c.Redis.Pass,
	Type: c.Redis.Type,
	}
	return &ServiceContext{
	Config:         c,
	CacheRedis:     redisX.MustNewRedis(redisConf),
		{{- range .Tables }}
			{{.Table}}Model: model.New{{.Table}}Model(conn, c.CacheRedis, db),
		{{- end }}
		}
}
