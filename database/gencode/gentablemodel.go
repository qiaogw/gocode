// Package model 自动生成模板GenTable(表模块)
package gencode

import (
	"context"
	"github.com/qiaogw/gocode/common/gormx"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ GenTableModel = (*customGenTableModel)(nil)

type (
	// GenTableModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGenTableModel.
	// 更多的自定义方法在这里添加，通过接口方法
	GenTableModel interface {
		genTableModel
		FindAll(ctx context.Context, query *ListGenTableReq) ([]*GenTable, int64, error)
		UpSetAll(ctx context.Context, newData *GenTable) error
	}

	customGenTableModel struct {
		*defaultGenTableModel
	}

	SearchGenTableModel struct {
		GenTable
		modelx.Pagination
	}
)

// NewGenTableModel returns a model for the database table.
func NewGenTableModel(conn sqlx.SqlConn, c cache.CacheConf, gormX *gorm.DB) GenTableModel {
	return &customGenTableModel{
		defaultGenTableModel: newGenTableModel(conn, c, gormX),
	}
}

// FindAll 条件查询列表
func (m *customGenTableModel) FindAll(ctx context.Context, query *ListGenTableReq) ([]*GenTable, int64, error) {
	var resp []*GenTable
	var count int64
	findKeySql := gormx.SearchKey(m.gormDB, m.tableName(), query.SearchKey)
	res := m.gormDB.Scopes(
		gormx.MakeCondition(query.GetNeedSearch(), m.gormDB.Name()),
		gormx.Paginate(query.GetPageSize(), query.GetPageIndex()),
		gormx.SortBy(query.SortBY, query.Descending),
	).Preload("Source").Where(findKeySql).Find(&resp).Limit(-1).Offset(-1)
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

func (m *customGenTableModel) UpSetAll(ctx context.Context, v *GenTable) error {
	var err error
	tx := m.gormDB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.Debug().Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "db"},
			{Name: "table"},
		},
		DoUpdates: clause.AssignmentColumns([]string{"db", "table", "name",
			"package_name", "table_url", "service", "is_auth", "is_import"}),
	}).Create(v).Error
	if err != nil {
		return err
	}

	err = tx.Debug().Model(v).Association("Columns").Replace(v.Columns)
	if err != nil {
		return err
	}
	return err
}
