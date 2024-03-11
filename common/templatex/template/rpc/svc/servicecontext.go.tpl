package svc

import (
	"github.com/qiaogw/gocode/common/gormx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/syncx"

	"gorm.io/driver/{{.DriverName}}"
	"gorm.io/gorm"

	"{{.ParentPkg}}/model"
	"{{.ParentPkg}}/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
  Cache     cache.Cache
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

	return &ServiceContext{
		Config:         c,
		Cache:          cache.New(c.CacheRedis, syncx.NewSingleFlight(), cache.NewStat("sub-{{.Package}}-rpc"), nil),
{{- range .Tables }}
			{{.Table}}Model: model.New{{.Table}}Model(conn, c.CacheRedis, db),
		{{- end }}
	}
}
