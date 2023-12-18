package svc

import (
	"{{.ParentPkg}}/api/internal/config"
	"{{.ParentPkg}}/model"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/qiaogw/gocode/common/gormx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	Cache    cache.Cache
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
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		//自动更新表，不使用外键
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	//开发模式，打印sql
	if c.Mode == service.DevMode {
		err := db.Use(&gormx.ZeroGorm{})
		if err != nil {
			logx.Errorf("配置数据库打印日志错误: %v", err)
		}
	}
	dbCache:=cache.New(c.CacheRedis, syncx.NewSingleFlight(), cache.NewStat("dc"), nil),
	return &ServiceContext{
		Config:   c,
		Cache:    dbCache,
	{{- range .Tables }}
		{{.Table}}Model: model.New{{.Table}}Model(dbCache, db),
	{{- end }}
	}
}