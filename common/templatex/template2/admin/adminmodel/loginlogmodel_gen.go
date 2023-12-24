// Package model errCode generated by gocode. DO NOT EDIT!
package adminmodel

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
	"time"
)

var cacheAdminLoginLogIdPrefix = "cache:admin:loginLog:id:"

type (
	loginLogModel interface {
		Insert(ctx context.Context, data *LoginLog) (*LoginLog, error)
		FindOne(ctx context.Context, id interface{}) (*LoginLog, error)
		Update(ctx context.Context, newData *LoginLog) error
		Delete(ctx context.Context, id interface{}) error
	}

	defaultLoginLogModel struct {
		cache.Cache
		table  string
		gormDB *gorm.DB
	}

	LoginLog struct {
		modelx.ModelWithCommon
		Username      string    `json:"username" comment:"用户名" gorm:"column:username;size:128;comment:用户名;"`
		Status        string    `json:"status" comment:"状态" gorm:"column:status;size:4;comment:状态;"`
		Ipaddr        string    `json:"ipaddr" comment:"ip地址" gorm:"column:ipaddr;size:255;comment:ip地址;"`
		LoginLocation string    `json:"loginLocation" comment:"归属地" gorm:"column:login_location;size:255;comment:归属地;"`
		Browser       string    `json:"browser" comment:"浏览器" gorm:"column:browser;size:255;comment:浏览器;"`
		Os            string    `json:"os" comment:"系统" gorm:"column:os;size:255;comment:系统;"`
		Platform      string    `json:"platform" comment:"固件" gorm:"column:platform;size:255;comment:固件;"`
		LoginTime     time.Time `json:"loginTime" comment:"登录时间" gorm:"column:login_time;size:8;comment:登录时间;"`
		Msg           string    `json:"msg" comment:"信息" gorm:"column:msg;size:255;comment:信息;"`
	}
)

// TableName LoginLog 表名
func (LoginLog) TableName() string {
	return "admin_login_log"
}

func newLoginLogModel(c cache.Cache, gormx *gorm.DB) *defaultLoginLogModel {
	return &defaultLoginLogModel{
		Cache:  c,
		table:  "admin_login_log",
		gormDB: gormx,
	}
}

func (m *defaultLoginLogModel) FindOne(ctx context.Context, id interface{}) (*LoginLog, error) {
	publicAdminLoginLogIdKey := fmt.Sprintf("%s%v", cacheAdminLoginLogIdPrefix, id)
	var resp LoginLog

	err := m.GetCtx(ctx, publicAdminLoginLogIdKey, &resp)
	if err != nil {
		err := m.gormDB.First(&resp, "id = ?", id).Error
		switch err {
		case nil:
			err = m.SetCtx(ctx, publicAdminLoginLogIdKey, resp)
		default:
			if err.Error() == modelx.ErrNotFound.Error() {
				return nil, modelx.ErrNotFound
			}
			return nil, err
		}
	}
	return &resp, nil
}

func (m *defaultLoginLogModel) Insert(ctx context.Context, data *LoginLog) (*LoginLog, error) {
	newUUID := uuid.New()
	data.Id = newUUID
	publicAdminLoginLogIdKey := fmt.Sprintf("%s%v", cacheAdminLoginLogIdPrefix, data.Id)
	err := m.gormDB.Create(data).Error
	if err != nil {
		return nil, err
	}
	_ = m.DelCtx(ctx, publicAdminLoginLogIdKey)
	return data, nil
}

func (m *defaultLoginLogModel) Update(ctx context.Context, newData *LoginLog) error {
	err := m.gormDB.Save(newData).Error
	if err != nil {
		return err
	}
	publicAdminLoginLogIdKey := fmt.Sprintf("%s%v", cacheAdminLoginLogIdPrefix, newData.Id)
	_ = m.DelCtx(ctx, publicAdminLoginLogIdKey)
	return nil
}

func (m *defaultLoginLogModel) Delete(ctx context.Context, id interface{}) error {
	publicAdminLoginLogIdKey := fmt.Sprintf("%s%v", cacheAdminLoginLogIdPrefix, id)
	_, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	err = m.gormDB.Delete(&LoginLog{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	_ = m.DelCtx(ctx, publicAdminLoginLogIdKey)
	return nil
}

func (m *defaultLoginLogModel) tableName() string {
	return m.table
}
