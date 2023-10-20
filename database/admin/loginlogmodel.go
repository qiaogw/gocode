// package model 自动生成模板LoginLog(登录日志)
package admin

import (
	"context"
	"github.com/qiaogw/gocode/common/gormx"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

var _ LoginLogModel = (*customLoginLogModel)(nil)

type (
	// LoginLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLoginLogModel.
	// 更多的自定义方法在这里添加，通过接口方法
	LoginLogModel interface {
		loginLogModel
		FindAll(ctx context.Context, query *ListLoginLogReq) ([]*LoginLog, int64, error)
	}

	customLoginLogModel struct {
		*defaultLoginLogModel
	}

	SearchLoginLogModel struct {
		LoginLog
		modelx.Pagination
	}
)

// NewLoginLogModel returns a model for the database table.
func NewLoginLogModel(conn sqlx.SqlConn, c cache.CacheConf, gormX *gorm.DB) LoginLogModel {
	return &customLoginLogModel{
		defaultLoginLogModel: newLoginLogModel(conn, c, gormX),
	}
}

// FindAll 条件查询列表
func (m *customLoginLogModel) FindAll(ctx context.Context, query *ListLoginLogReq) ([]*LoginLog, int64, error) {
	var resp []*LoginLog
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
