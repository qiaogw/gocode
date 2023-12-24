// package schema 自动生成模板Dept(部门)
package adminmodel

import (
	"context"

	"github.com/qiaogw/gocode/common/gormx"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ DeptModel = (*customDeptModel)(nil)

type (
	// DeptModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeptModel.
	// 更多的自定义方法在这里添加，通过接口方法
	DeptModel interface {
		deptModel
		FindAll(ctx context.Context, query *ListDeptReq) ([]*Dept, int64, error)
	}

	customDeptModel struct {
		*defaultDeptModel
	}

	SearchDeptModel struct {
		Dept
		modelx.Pagination
	}
)

// NewDeptModel returns a model for the database table.
func NewDeptModel(c cache.Cache, gormX *gorm.DB) DeptModel {
	return &customDeptModel{
		defaultDeptModel: newDeptModel(c, gormX),
	}
}

// FindAll 条件查询列表
func (m *customDeptModel) FindAll(ctx context.Context, query *ListDeptReq) ([]*Dept, int64, error) {
	var resp []*Dept
	var count int64
	sql := gormx.SearchKey(m.gormDB, m.tableName(), query.SearchKey)
	res := m.gormDB.Scopes(
		gormx.MakeCondition(query.GetNeedSearch(), m.gormDB.Name()),
		gormx.Paginate(query.GetPageSize(), query.GetPageIndex()),
		gormx.SortBy(query.SortBY, query.Descending),
	).Where(sql).Find(&resp).Limit(-1).Offset(-1)
	res.Count(&count)
	err := res.Error
	switch err {
	case nil:
		return resp, count, nil
	case modelx.ErrNotFound:
		return nil, 0, modelx.ErrNotFound
	default:
		return nil, 0, err
	}
}
