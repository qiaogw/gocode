// Package model 自动生成模板{{.Table}}({{.TableComment}})
package model

import (
	"context"
"io"
"{{.PKG}}/common/modelx"
"{{.PKG}}/common/toolx"
	"{{.PKG}}/common/gormx"
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
//var data []{{.Table}}
//err = ex.ReadExcelIo(tx, reader)
//if err != nil {
//	return err
//}
//err = json.Unmarshal(ex.Content, &data)
//if err != nil {
//	return err
//}
//tx=tx.Clauses(clause.OnConflict{
//	Columns: []clause.Column{
//		{Name: "driver"},
//		{Name: "host"},
//		{Name: "dbname"},
//	},
//	UpdateAll: true,
//})
//for i := 0; i < len(data); i += 1000 {
//	end := i + 1000
//	if end > len(data) {
//		end = len(data)
//	}
//	err = tx.CreateInBatches(data[i:end], len(data[i:end])).Error
//	if err != nil {
//		return err
//	}
//}
return nil
}
