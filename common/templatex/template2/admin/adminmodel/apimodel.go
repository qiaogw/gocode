// Package model package schema 自动生成模板Api(api)
package adminmodel

import (
	"context"
	"github.com/google/uuid"
	"github.com/qiaogw/gocode/common/gormx"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ ApiModel = (*customApiModel)(nil)

type (
	// ApiModel is an interface to be customized, add more methods here,
	// and implement the added methods in customApiModel.
	// 更多的自定义方法在这里添加，通过接口方法
	ApiModel interface {
		apiModel
		FindAll(ctx context.Context, query *ListApiReq) ([]*Api, int64, error)
		FindByRole(ctx context.Context, roleId interface{}) ([]*Api, error)
		UpdateAll(ctx context.Context, data *Api) error
	}

	customApiModel struct {
		*defaultApiModel
	}

	SearchApiModel struct {
		Api
		modelx.Pagination
	}
)

// NewApiModel returns a model for the database table.
func NewApiModel(c cache.Cache, gormX *gorm.DB) ApiModel {
	return &customApiModel{
		defaultApiModel: newApiModel(c, gormX),
	}
}

// FindAll 条件查询列表
func (m *customApiModel) FindAll(ctx context.Context, query *ListApiReq) ([]*Api, int64, error) {
	var resp []*Api
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

func (m *customApiModel) FindByRole(ctx context.Context, roleId interface{}) ([]*Api, error) {
	var role Role
	var list []*Api
	l := make([]*Api, 0)
	err := m.gormDB.Model(&Role{}).Where("id = ? ", roleId).
		Preload("Apis").Find(&role).Error
	if role.Apis != nil {
		l = role.Apis
	}
	for _, v := range l {
		list = append(list, v)
	}
	return list, err
}

func (m *customApiModel) UpdateAll(ctx context.Context, v *Api) error {
	// 生成一个新的 UUID
	newUUID := uuid.New()
	//err := m.gormDB.Model(v).
	//	Where(Api{Path: v.Path, Method: v.Method}).
	//	Attrs(Api{Module: v.Module, Title: v.Title}).
	//	FirstOrCreate(&Api{}).
	//	Updates(Api{Module: v.Module, Title: v.Title, Method: v.Method}).
	//	Error
	// 创建或更新对象
	//logx.Infof("【api】%+v", v)
	err := m.gormDB.Model(v).
		Where(Api{Path: v.Path, Method: v.Method}).
		Attrs(Api{Id: newUUID, Module: v.Module, Title: v.Title}).
		FirstOrCreate(&Api{}).
		Error
	return err
}
