// errCode generated by gocode. DO NOT EDIT!
package admin

import (
	"context"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type (
	apiModel interface {
		Insert(ctx context.Context, data *Api) (*Api, error)
		FindOne(ctx context.Context, id interface{}) (*Api, error)
		Update(ctx context.Context, newData *Api) error
		Delete(ctx context.Context, id interface{}) error
		FindOneByPath(ctx context.Context, path string) (*Api, error)
	}

	defaultApiModel struct {
		sqlc.CachedConn
		table  string
		gormDB *gorm.DB
	}

	Api struct {
		modelx.BaseModel
		Title  string  `json:"title" form:"title" db:"title" gorm:"column:title;comment:标题;"`
		Path   string  `json:"path" form:"path" db:"path" gorm:"column:path;comment:地址;"`
		Method string  `json:"method" form:"method" db:"method" gorm:"column:method;comment:请求类型;"`
		Module string  `json:"module" form:"module" db:"module" gorm:"column:module;comment:api组;"`
		Remark string  `json:"remark" form:"remark" db:"remark" gorm:"column:remark;comment:说明;"`
		Roles  []*Role `json:"roles" db:"-" gorm:"many2many:admin_role_api;"`
		modelx.ControlBy
		modelx.ModelTime
	}
)

// TableName Api 表名
func (Api) TableName() string {
	return "admin_api"
}

func newApiModel(conn sqlx.SqlConn, c cache.CacheConf, gormx *gorm.DB) *defaultApiModel {
	return &defaultApiModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "admin_api",
		gormDB:     gormx,
	}
}

func (m *defaultApiModel) FindOne(ctx context.Context, id interface{}) (*Api, error) {
	var resp Api
	err := m.gormDB.First(&resp, id).Error
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
func (m *defaultApiModel) FindOneByPath(ctx context.Context, path string) (*Api, error) {
	var resp Api
	err := m.gormDB.Where("path = ?", path).First(&resp).Error
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

func (m *defaultApiModel) Insert(ctx context.Context, data *Api) (*Api, error) {
	err := m.gormDB.Create(data).Error
	//var re sql.Result

	return data, err
}

func (m *defaultApiModel) Update(ctx context.Context, newData *Api) error {
	_, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}
	err = m.gormDB.Save(newData).Error
	return err
}

func (m *defaultApiModel) Delete(ctx context.Context, id interface{}) error {
	_, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	err = m.gormDB.Delete(&Api{}, id).Error
	return err
}

func (m *defaultApiModel) tableName() string {
	return m.table
}
