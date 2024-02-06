// Package model 自动生成模板{{.Table}}({{.TableComment}})

{{$table:=.Table}}
package model

import (
	"context"
	"io"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/qiaogw/gocode/common/toolx"
	"github.com/qiaogw/gocode/common/gormx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

var _ {{.Table}}Model = (*custom{{.Table}}Model)(nil)

type (
	// {{.Table}}Model is an interface to be customized, add more methods here,
	// and implement the added methods in custom{{.Table}}Model.
	// 更多的自定义方法在这里添加，通过接口方法
	{{.Table}}Model interface {
		{{.PackageName}}Model
		FindAll(ctx context.Context, query *List{{.Table}}Req) ([]*{{.Table}},int64, error)
		Import(reader io.Reader) error
		GetName() string
		DeleteList(ctx context.Context, list []string) error
	{{- range  .CacheKeys}}
		FindOneBy{{.Field}}(ctx context.Context, {{.FieldJson}} {{.DataType}}) (*{{$table}}, error)
	{{- end }}
	}

	custom{{.Table}}Model struct {
		*default{{.Table}}Model
	}

	Search{{.Table}}Model struct {
		{{.Table}}
		modelx.Pagination
	}
)

// New{{.Table}}Model returns a model for the database table.
func New{{.Table}}Model(conn sqlx.SqlConn, c cache.CacheConf, gormX *gorm.DB) {{.Table}}Model {
	return &custom{{.Table}}Model{
		default{{.Table}}Model: new{{.Table}}Model(conn, c, gormX),
	}
}
// GetName 获取业务表名
func (m *custom{{.Table}}Model) GetName() string {
		return m.tableName()
}

// DeleteList 批量删除
func (m *custom{{.Table}}Model) DeleteList(ctx context.Context, list []string) error {
	err := m.gormDB.Debug().Where("id in ?", list).Delete(&{{.Table}}{}).Error
	if err != nil {
		return err
	}
	return nil
}

// FindAll 条件查询列表
func (m *custom{{.Table}}Model) FindAll(ctx context.Context, query *List{{.Table}}Req) ([]*{{.Table}},int64, error) {
	var resp []*{{.Table}}
	var count int64
	findKeySql := gormx.SearchKey(m.gormDB, m.tableName(), query.SearchKey)
	res:= m.gormDB.Scopes(
			gormx.MakeCondition(query.GetNeedSearch(), m.gormDB.Name()),
			gormx.Paginate(query.GetPageSize(), query.GetPageIndex()),
			gormx.SortBy(query.SortBY, query.Descending),
		).Where(findKeySql).Find(&resp).Limit(-1).Offset(-1)
	res.Count(&count)
	err := res.Error
	switch err {
	case nil:
		return resp,count, nil
	case modelx.ErrNotFound:
		return nil,0, modelx.ErrNotFound
	default:
		return nil,0, err
	}
}


// Import 导入
func (m *custom{{.Table}}Model) Import(reader io.Reader) error {
	var err error
	tx := m.gormDB
	temp := new({{.Table}})
	ex := new(toolx.ExcelStruct)
	ex.Model = temp
	err = ex.SaveDb(tx, reader)
	if err != nil {
		return err
	}
	return nil
}

{{- range  .CacheKeys}}
	func (m *default{{$table}}Model) FindOneBy{{.Field}}(ctx context.Context, {{.FieldJson}} {{.DataType}}) (*{{$table}}, error) {
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