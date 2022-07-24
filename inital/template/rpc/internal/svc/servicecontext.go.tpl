package svc

import (
	"{{.ParentPkg}}/model"
	"{{.ParentPkg}}/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/{{.DriverName}}"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	{{- range .Tables }}
	{{.Table}}Model model.{{.Table}}Model
	{{- end }}
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn(c.Database.DriverName, c.Database.DataSource) //.NewMysql(c.Mysql.DataSource)
	dsn := c.Database.DataSource
	gormX, _ := gorm.Open({{.DriverName}}.Open(dsn), &gorm.Config{})

	return &ServiceContext{
		Config:    c,
    {{- range .Tables }}
        {{.Table}}Model: model.New{{.Table}}Model(conn, c.CacheRedis, gormX),
    {{- end }}
	}
}
