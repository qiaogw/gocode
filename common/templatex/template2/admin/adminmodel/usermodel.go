// Package model package schema 自动生成模板User(用户)
package adminmodel

import (
	"context"
	"github.com/qiaogw/gocode/common/gormx"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	// 更多的自定义方法在这里添加，通过接口方法
	UserModel interface {
		userModel
		FindAll(ctx context.Context, query *ListUserReq) ([]*User, int64, error)
		FindOneByMobile(ctx context.Context, mobile string) (*User, error)
		FindOneByName(ctx context.Context, name string) (*User, error)
		FindOneByEmail(ctx context.Context, email string) (*User, error)
		ResetPassword(ctx context.Context, ids []int64, pwd string) error
		ListNoUser(ctx context.Context, query *ListUserReq) ([]*User, int64, error)
		TreeList(ctx context.Context, query *ListUserReq) ([]*User, int64, error)
	}

	customUserModel struct {
		*defaultUserModel
	}

	SearchUserModel struct {
		User
		modelx.Pagination
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(c cache.Cache, gormX *gorm.DB) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(c, gormX),
	}
}

// FindAll 条件查询列表
func (m *customUserModel) FindAll(ctx context.Context, query *ListUserReq) ([]*User, int64, error) {
	var resp []*User
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

func (m *defaultUserModel) FindOneByMobile(ctx context.Context, mobile string) (*User, error) {
	var resp User
	err := m.gormDB.Preload("Roles").Preload("Post").
		Preload("Dept").Where("mobile = ?", mobile).First(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	default:
		if err.Error() == modelx.ErrNotFound.Error() {
			return nil, modelx.ErrNotFound
		}
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByName(ctx context.Context, name string) (*User, error) {
	var resp User
	err := m.gormDB.Preload("Roles").Preload("Post").
		Preload("Dept").Where("username = ?", name).First(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	default:
		if err.Error() == modelx.ErrNotFound.Error() {
			return nil, modelx.ErrNotFound
		}
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByEmail(ctx context.Context, email string) (*User, error) {
	var resp User
	err := m.gormDB.Preload("Roles").Preload("Post").
		Preload("Dept").Where("email = ?", email).First(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	default:
		if err.Error() == modelx.ErrNotFound.Error() {
			return nil, modelx.ErrNotFound
		}
		return nil, err
	}
}
func (m *defaultUserModel) ResetPassword(ctx context.Context, ids []int64, pwd string) error {
	err := m.gormDB.Model(User{}).
		Where("id in ?", ids).Updates(User{Password: pwd}).Error
	return err
}

// ListNoUser 角色 无权用户
func (m *defaultUserModel) ListNoUser(ctx context.Context, query *ListUserReq) ([]*User,
	int64, error) {
	var resp []*User
	var count int64
	sql := gormx.SearchKey(m.gormDB, m.tableName(), query.SearchKey)
	res := m.gormDB.Model(User{}).Scopes(
		gormx.Paginate(query.GetPageSize(), query.GetPageIndex()),
		gormx.SortBy(query.SortBY, query.Descending),
	).
		Joins(`left join ( SELECT * FROM admin_user_role 
WHERE admin_user_role.role_id = ? ) AS b
 	ON admin_user."id" = b.user_id `, query.Id).
		Where("b.user_id is null").
		Where(sql).
		Find(&resp)
	res.Count(&count)
	err := res.Error
	return resp, count, err
}

// TreeList 树列表
func (m *customUserModel) TreeList(ctx context.Context, query *ListUserReq) ([]*User, int64, error) {
	var resp []*User
	var count int64
	res := m.gormDB.Scopes(
		gormx.MakeCondition(query.GetNeedSearch(), m.gormDB.Name()),
		gormx.Paginate(query.GetPageSize(), query.GetPageIndex()),
		gormx.SortBy(query.SortBY, query.Descending),
	).Preload("Dept").Find(&resp).Limit(-1).Offset(-1)
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
