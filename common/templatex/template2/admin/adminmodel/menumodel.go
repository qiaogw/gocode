// Package model package schema 自动生成模板Menu(菜单)
package adminmodel

import (
	"context"
	"fmt"
	"github.com/qiaogw/gocode/common/gormx"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ MenuModel = (*customMenuModel)(nil)
var cachePublicAdminMenuRoleIdPrefix = "cache:public:adminMenu:roleId:"

type (
	// MenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMenuModel.
	// 更多的自定义方法在这里添加，通过接口方法
	MenuModel interface {
		menuModel
		FindAll(ctx context.Context, query *ListMenuReq) ([]*Menu, int64, error)
		FindByRole(ctx context.Context, roleId interface{}) ([]*Menu, error)
	}

	customMenuModel struct {
		*defaultMenuModel
	}

	SearchMenuModel struct {
		Menu
		modelx.Pagination
	}
)

// NewMenuModel returns a model for the database table.
func NewMenuModel(c cache.Cache, gormX *gorm.DB) MenuModel {
	return &customMenuModel{
		defaultMenuModel: newMenuModel(c, gormX),
	}
}

// FindAll 条件查询列表
func (m *customMenuModel) FindAll(ctx context.Context, query *ListMenuReq) ([]*Menu, int64, error) {
	var resp []*Menu
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

func (m *defaultMenuModel) FindByRole(ctx context.Context, roleId interface{}) ([]*Menu, error) {

	publicAdminMenuRoleIdKey := fmt.Sprintf("%s%v", cachePublicAdminMenuRoleIdPrefix, roleId)
	var role Role
	var list []*Menu
	err := m.GetCtx(ctx, publicAdminMenuRoleIdKey, &list)
	if err != nil {
		l := make([]*Menu, 0)
		err = m.gormDB.Debug().Model(&Role{}).
			Preload("Menus", func(db *gorm.DB) *gorm.DB {
				return db.Order("sort")
			}).Where("id = ? ", roleId).First(&role).Error
		if role.Menus != nil {
			l = role.Menus
		}
		for _, v := range l {
			list = append(list, v)
		}
		_ = m.SetCtx(ctx, publicAdminMenuRoleIdKey, list)
	}

	return list, err
}
