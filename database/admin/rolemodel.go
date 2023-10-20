// package model 自动生成模板Role(角色)
package admin

import (
	"context"
	"errors"
	"fmt"
	"github.com/qiaogw/gocode/common/casbinx"
	"github.com/qiaogw/gocode/common/gormx"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

var _ RoleModel = (*customRoleModel)(nil)

type (
	// RoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleModel.
	// 更多的自定义方法在这里添加，通过接口方法
	RoleModel interface {
		roleModel
		FindAll(ctx context.Context, query *ListRoleReq) ([]*Role, int64, error)
		SetRoleMenu(ctx context.Context, in *Role) error
		SetRoleApi(ctx context.Context, in *Role) error
	}

	customRoleModel struct {
		*defaultRoleModel
	}

	SearchRoleModel struct {
		Role
		modelx.Pagination
	}
)

// NewRoleModel returns a model for the database table.
func NewRoleModel(conn sqlx.SqlConn, c cache.CacheConf, gormX *gorm.DB) RoleModel {
	return &customRoleModel{
		defaultRoleModel: newRoleModel(conn, c, gormX),
	}
}

// FindAll 条件查询列表
func (m *customRoleModel) FindAll(ctx context.Context, query *ListRoleReq) ([]*Role, int64, error) {
	var resp []*Role
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

// SetRoleMenu 角色 授权 菜单
func (m *customRoleModel) SetRoleMenu(ctx context.Context, in *Role) error {
	err := m.gormDB.Model(in).Association("Menus").Replace(&in.Menus)
	return err
}

// SetRoleApi 角色 授权 api
func (m *customRoleModel) SetRoleApi(ctx context.Context, in *Role) error {
	var err error
	tx := m.gormDB.Begin()
	defer func() {
		if err != nil {
			logx.Error(err)
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.Model(in).Association("Apis").Replace(&in.Apis)
	if err != nil {
		return err
	}
	err = tx.Preload("Apis").First(in).Error
	if err != nil {
		return err
	}
	roleId := fmt.Sprint(in.Id)
	casbinSvc := casbinx.CasbinServiceApp(m.gormDB)

	cb := casbinSvc.Casbin()
	_, err = cb.RemoveFilteredPolicy(0, roleId)
	if err != nil {
		return err
	}
	if len(in.Apis) < 1 {
		return nil
	}
	var rules [][]string
	for _, v := range in.Apis {
		rules = append(rules, []string{roleId, v.Path, v.Method})
	}

	ok, err := cb.AddPolicies(rules)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return err
}
