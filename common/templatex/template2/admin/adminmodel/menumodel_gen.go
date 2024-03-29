// Package model  generated by gocode. DO NOT EDIT!
package adminmodel

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/qiaogw/gocode/common/globalx"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var cachePublicAdminMenuIdPrefix = "cache:public:adminMenu:id:"

type (
	menuModel interface {
		Insert(ctx context.Context, data *Menu) (*Menu, error)
		FindOne(ctx context.Context, id interface{}) (*Menu, error)
		Update(ctx context.Context, newData *Menu) error
		Delete(ctx context.Context, id interface{}) error
		FindOneByPath(ctx context.Context, path string) (*Menu, error)
		Upset(ctx context.Context, data *Menu) error
	}

	defaultMenuModel struct {
		cache.Cache
		table  string
		gormDB *gorm.DB
	}

	Menu struct {
		modelx.ModelWithCommon
		Name        string  `json:"name" comment:"名称" gorm:"column:name;size:256;comment:名称;"`
		Title       string  `json:"title" comment:"标题" gorm:"column:title;size:256;comment:标题;"`
		Icon        string  `json:"icon" comment:"图标" gorm:"column:icon;size:256;comment:图标;"`
		Path        string  `json:"path" comment:"路径" gorm:"column:path;size:256;comment:路径;"`
		Type        string  `json:"type" comment:"菜单类型" gorm:"column:type;size:128;comment:菜单类型;"`
		Component   string  `json:"component" comment:"组件" gorm:"column:component;size:256;comment:组件;"`
		ParentId    string  `json:"parentId" comment:"父菜单" gorm:"column:parent_id;size:256;comment:父菜单;"`
		Sort        int64   `json:"sort" comment:"排序" gorm:"column:sort;size:2;comment:排序;"`
		DefaultMenu bool    `json:"defaultMenu" comment:"是否基础路由" gorm:"column:default_menu;size:1;comment:是否基础路由;"`
		CloseTab    bool    `json:"closeTab" comment:"自动关闭tab" gorm:"column:close_tab;size:1;comment:自动关闭tab;"`
		KeepAlive   bool    `json:"keepAlive" comment:"是否缓存" gorm:"column:keep_alive;size:1;comment:是否缓存;"`
		Hidden      bool    `json:"hidden" comment:"是否隐藏" gorm:"column:hidden;size:1;comment:是否隐藏;"`
		IsFrame     bool    `json:"isFrame" comment:"是否frame" gorm:"column:is_frame;size:1;comment:是否frame;"`
		Roles       []*Role `json:"roles" gorm:"many2many:admin_role_menu;" db:"-"`
		Children    []*Menu `json:"children" gorm:"-" db:"-"`
		Button      []*Menu `json:"button" gorm:"-" db:"-"`
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

func newMenuModel(c cache.Cache, gormx *gorm.DB) *defaultMenuModel {
	return &defaultMenuModel{
		Cache:  c,
		table:  "admin_menu",
		gormDB: gormx,
	}
}

func (m *defaultMenuModel) FindOne(ctx context.Context, id interface{}) (*Menu, error) {
	publicAdminMenuIdKey := fmt.Sprintf("%s%v", cachePublicAdminMenuIdPrefix, id)
	var resp Menu
	err := m.GetCtx(ctx, publicAdminMenuIdKey, &resp)
	if err != nil {
		err = m.gormDB.Where("id = ?", id).First(&resp).Error
		switch err {
		case nil:
			err = m.SetCtx(ctx, publicAdminMenuIdKey, resp)
		default:
			if err.Error() == modelx.ErrNotFound.Error() {
				return nil, modelx.ErrNotFound
			}
			return nil, err
		}
	}
	return &resp, nil
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

func (m *defaultMenuModel) Insert(ctx context.Context, data *Menu) (*Menu, error) {
	newUUID := uuid.New()
	data.Id = newUUID
	var err error
	publicAdminMenuIdKey := fmt.Sprintf("%s%v", cachePublicAdminMenuIdPrefix, data.Id)
	tx := m.gormDB.Begin()
	defer func() {
		if err != nil {
			logx.Error(err)
			tx.Rollback()
		} else {
			_ = m.DelCtx(ctx, publicAdminMenuIdKey)
			tx.Commit()
		}
	}()
	var menu Menu
	_ = copier.Copy(&menu, data)

	err = tx.Create(&menu).Error
	if err != nil {
		return nil, err
	}

	for i, _ := range data.Button {
		data.Button[i].ParentId = menu.Id.String()
		data.Button[i].Id = uuid.New()
		data.Button[i].UpdateBy = menu.UpdateBy
		data.Button[i].Type = globalx.Button
	}
	err = tx.Unscoped().Where("parent_id =?", menu.Id).
		Where("type =?", globalx.Button).Delete(&Menu{}).Error
	if err != nil {
		return nil, err
	}
	if len(data.Button) > 0 {
		err = tx.Create(data.Button).Error
		if err != nil {
			return nil, err
		}
	}

	return data, err
}

func (m *defaultMenuModel) Update(ctx context.Context, newData *Menu) error {
	var err error
	publicAdminMenuIdKey := fmt.Sprintf("%s%v", cachePublicAdminMenuIdPrefix, newData.Id)

	tx := m.gormDB.Begin()
	defer func() {
		if err != nil {
			logx.Error(err)
			tx.Rollback()
		} else {
			_ = m.DelCtx(ctx, publicAdminMenuIdKey)
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

	err = tx.Unscoped().Where("parent_id =?", newData.Id).
		Where("type =?", globalx.Button).Delete(&Menu{}).Error
	if err != nil {
		return err
	}
	if len(newData.Button) > 0 {
		err = tx.Create(newData.Button).Error
		if err != nil {
			return err
		}
	}
	return err
}

func (m *defaultMenuModel) Delete(ctx context.Context, id interface{}) error {
	publicAdminMenuIdKey := fmt.Sprintf("%s%v", cachePublicAdminMenuIdPrefix, id)

	_, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	err = m.gormDB.Delete(&Menu{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	_ = m.DelCtx(ctx, publicAdminMenuIdKey)
	return nil
}

func (m *defaultMenuModel) tableName() string {
	return m.table
}

func (m *defaultMenuModel) Upset(ctx context.Context, data *Menu) error {
	var err error
	publicAdminMenuIdKey := fmt.Sprintf("%s%v", cachePublicAdminMenuIdPrefix, data.Id)

	tx := m.gormDB.Begin()
	defer func() {
		if err != nil {
			logx.Error(err)
			tx.Rollback()
		} else {
			_ = m.DelCtx(ctx, publicAdminMenuIdKey)
			tx.Commit()
		}
	}()

	var menu Menu
	data.Id = uuid.New()
	_ = copier.Copy(&menu, data)

	err = tx.Model(Menu{}).Where(Menu{Path: menu.Path}).
		Assign(menu).
		FirstOrCreate(&menu).
		Error
	if err != nil {
		return err
	}

	for i, _ := range data.Button {
		data.Button[i].ParentId = menu.Id.String()
		data.Button[i].Id = uuid.New()
		data.Button[i].UpdateBy = menu.UpdateBy
		data.Button[i].Type = globalx.Button
	}
	err = tx.Unscoped().Where("parent_id =?", menu.Id).
		Where("type =?", globalx.Button).Delete(&Menu{}).Error
	if err != nil {
		return err
	}
	if len(data.Button) > 0 {
		err = tx.Create(data.Button).Error
		if err != nil {
			return err
		}
	}

	return err
}
