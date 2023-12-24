// package schema 自动生成模板Post(职务)
package adminmodel

import (
	"context"
	"github.com/qiaogw/gocode/common/gormx"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ PostModel = (*customPostModel)(nil)

type (
	// PostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostModel.
	// 更多的自定义方法在这里添加，通过接口方法
	PostModel interface {
		postModel
		FindAll(ctx context.Context, query *ListPostReq) ([]*Post, int64, error)
	}

	customPostModel struct {
		*defaultPostModel
	}

	SearchPostModel struct {
		Post
		modelx.Pagination
	}
)

// NewPostModel returns a model for the database table.
func NewPostModel(c cache.Cache, gormX *gorm.DB) PostModel {
	return &customPostModel{
		defaultPostModel: newPostModel(c, gormX),
	}
}

// FindAll 条件查询列表
func (m *customPostModel) FindAll(ctx context.Context, query *ListPostReq) ([]*Post, int64, error) {
	var resp []*Post
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
