// errCode generated by gocode. DO NOT EDIT!
package admin

import (
	"context"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type (
	menuModel interface {
		FindOne(ctx context.Context, id interface{}) (*Menu, error)
		Update(ctx context.Context, newData *Menu) error
		Delete(ctx context.Context, id interface{}) error
		FindOneByPath(ctx context.Context, path string) (*Menu, error)
		//Upset(ctx context.Context, data *Menu) error
	}

	defaultMenuModel struct {
		sqlc.CachedConn
		table  string
		gormDB *gorm.DB
	}

	Menu struct {
		modelx.BaseModel
		Name        string  `json:"name" form:"name" db:"name" gorm:"column:name;comment:菜单名称;"`
		Title       string  `json:"title" form:"title" db:"title" gorm:"column:title;comment:菜单显示;"`
		Icon        string  `json:"icon" form:"icon" db:"icon" gorm:"column:icon;comment:图标;"`
		Path        string  `json:"path" form:"path" db:"path" gorm:"column:path;comment:路径;"`
		Type        string  `json:"type" form:"type" db:"type" gorm:"column:type;comment:菜单类型;"`
		Component   string  `json:"component" form:"component" db:"component" gorm:"column:component;comment:组件;"`
		ParentId    int64   `json:"parentId" form:"parentId" db:"parent_id" gorm:"column:parent_id;comment:父菜单;"`
		Sort        int64   `json:"sort" form:"sort" db:"sort" gorm:"column:sort;comment:排序;"`
		KeepAlive   bool    `json:"keepAlive" form:"keepAlive" db:"keep_alive" gorm:"column:keep_alive;comment:是否缓存;"`
		Hidden      bool    `json:"hidden" form:"hidden" db:"hidden" gorm:"column:hidden;comment:是否隐藏;"`
		IsFrame     bool    `json:"isFrame" form:"isFrame" db:"is_frame" gorm:"column:is_frame;comment:是否frame;"`
		Remark      string  `json:"remark" form:"remark" db:"remark" gorm:"column:remark;comment:备注;"`
		CloseTab    bool    `json:"close_tab" db:"close_tab" gorm:"column:close_tab;comment:自动关闭tab;"`
		DefaultMenu bool    `json:"default_menu" db:"default_menu" gorm:"column:default_menu;comment:是否是基础路由;"`
		Roles       []*Role `json:"roles" gorm:"many2many:admin_role_menu;" db:"-"`
		Children    []*Menu `json:"children" gorm:"-" db:"-"`
		Button      []*Menu `json:"button" gorm:"-" db:"-"`
		modelx.ControlBy
		modelx.ModelTime
		//Meta `json:"meta" gorm:"embedded;comment:附加属性"` // 附加属性
	}
	Meta struct {
		KeepAlive   bool   `json:"keepAlive" form:"keepAlive" db:"keep_alive" gorm:"column:keep_alive;comment:是否缓存;"`
		Hidden      bool   `json:"hidden" form:"hidden" db:"hidden" gorm:"column:hidden;comment:是否隐藏;"`
		DefaultMenu bool   `json:"defaultMenu" gorm:"comment:是否是基础路由（开发中）"` // 是否是基础路由（开发中）
		Title       string `json:"title" form:"title" db:"title" gorm:"column:title;comment:菜单显示;"`
		Icon        string `json:"icon" form:"icon" db:"icon" gorm:"column:icon;comment:图标;"`
		CloseTab    bool   `json:"closeTab" gorm:"comment:自动关闭tab"` // 自动关闭tab
	}
)

// TableName Menu 表名
func (Menu) TableName() string {
	return "admin_menu"
}

func newMenuModel(conn sqlx.SqlConn, c cache.CacheConf, gormx *gorm.DB) *defaultMenuModel {
	return &defaultMenuModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "admin_menu",
		gormDB:     gormx,
	}
}

func (m *defaultMenuModel) FindOne(ctx context.Context, id interface{}) (*Menu, error) {
	var resp Menu
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
func (m *defaultMenuModel) FindOneByPath(ctx context.Context, path string) (*Menu, error) {
	var resp Menu
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

func (m *defaultMenuModel) Update(ctx context.Context, newData *Menu) error {
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

	_, err = m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	err = tx.Save(newData).Error
	if err != nil {
		return err
	}

	//err = tx.Unscoped().Where("parent_id =?", newData.Id).
	//	Where("type =?", globalx.Button).Delete(&Menu{}).Error
	//if err != nil {
	//	return err
	//}
	if len(newData.Button) > 0 {
		err = tx.Create(newData.Button).Error
		if err != nil {
			return err
		}
	}

	return err
}

func (m *defaultMenuModel) Delete(ctx context.Context, id interface{}) error {
	_, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	err = m.gormDB.Delete(&Menu{}, id).Error
	return err
}

func (m *defaultMenuModel) tableName() string {
	return m.table
}
