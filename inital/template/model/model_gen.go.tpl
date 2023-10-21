// Code generated by gocode. DO NOT EDIT!
package model
{{$table:=.Table}}
import (
"github.com/qiaogw/gocode/common/modelx"
	"context"

    "github.com/zeromicro/go-zero/core/stores/cache"
    "github.com/zeromicro/go-zero/core/stores/sqlc"
    "github.com/zeromicro/go-zero/core/stores/sqlx"


    "gorm.io/gorm"
	{{ if .HasTimer }}"time"{{ end }}
	{{ if .PostgreSql }}_ "github.com/lib/pq"{{ end}}
)



type (
	{{.PackageName}}Model interface {
		Insert(ctx context.Context, data *{{.Table}}) (*{{.Table}}, error)
		FindOne(ctx context.Context, id interface{}) (*{{.Table}}, error)
		Update(ctx context.Context, newData *{{.Table}}) (*{{.Table}}, error)
		Delete(ctx context.Context, id interface{}) error
	{{- range  .CacheKeys}}
		FindOneBy{{.Field}}(ctx context.Context, {{.FieldJson}} string) (*{{$table}}, error)
	{{- end }}
	}

	default{{.Table}}Model struct {
		sqlc.CachedConn
		table  string
		gormDB *gorm.DB
	}

	{{.Table}} struct {
		modelx.BaseModel
        {{- range .Columns }}
			{{- if .IsPk -}}
			{{- else if .IsModelTime }}
			{{- else if .IsControl }}
			{{- else}}
				{{- if .IsPage}}
				{{- else}}
				{{.FieldName}}  {{.DataType}} `json:"{{.FieldJson}}" comment:"{{.ColumnComment}}" {{- if ne .GormName "-" }} gorm:"column:{{.GormName}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}comment:{{.ColumnComment}};"{{- end -}}`
				{{- end -}}
			{{- end -}}
        {{- end }}
		modelx.ControlBy
        modelx.ModelTime
	}
)

{{ if .Table }}
// TableName {{.Table}} 表名
func ({{.Table}}) TableName() string {
  return "{{.Name}}"
}
{{ end }}

func new{{.Table}}Model(conn sqlx.SqlConn, c cache.CacheConf, gormx *gorm.DB) *default{{.Table}}Model {
	gormx.AutoMigrate(&{{.Table}}{})
	return &default{{.Table}}Model{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "{{.Name}}",
		gormDB:     gormx,
	}
}

func (m *default{{.Table}}Model) FindOne(ctx context.Context, id interface{}) (*{{.Table}}, error) {
	var resp {{.Table}}
	err := m.gormDB.First(&resp, id).Error
	switch err {
	case nil:
	return &resp, nil
	default:
		if  err.Error()==modelx.ErrNotFound.Error(){
			return nil, modelx.ErrNotFound
		}
		return nil, err
	}
}

{{- range  .CacheKeys}}
func (m *default{{$table}}Model) FindOneBy{{.Field}}(ctx context.Context, {{.FieldJson}} string) (*{{$table}}, error) {
	var resp {{$table}}
	err := m.gormDB.Where("{{.FieldJson}} = ?",{{.FieldJson}}).First(&resp).Error
	switch err {
	case nil:
	return &resp, nil
	default:
		if  err.Error()==modelx.ErrNotFound.Error(){
			return nil, modelx.ErrNotFound
		}
		return nil, err
	}
}
{{- end }}


func (m *default{{.Table}}Model) Insert(ctx context.Context, data *{{.Table}}) (*{{.Table}}, error) {
	err := m.gormDB.Create(data).Error
	//var re sql.Result

	return data, err
}

func (m *default{{.Table}}Model) Update(ctx context.Context, newData *{{.Table}}) (*{{.Table}}, error) {
	_, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return nil,err
	}
	err = m.gormDB.Save(newData).Error
	return newData,err
}

func (m *default{{.Table}}Model) Delete(ctx context.Context, id interface{}) error {
	_, err := m.FindOne(ctx, id)
	if err != nil {
	return err
	}

	err = m.gormDB.Delete(&{{.Table}}{}, id).Error
	return err
}


func (m *default{{.Table}}Model) tableName() string {
	return m.table
}


