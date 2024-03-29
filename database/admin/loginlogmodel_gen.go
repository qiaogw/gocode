// errCode generated by gocode. DO NOT EDIT!
package admin

import (
	"context"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type (
	loginLogModel interface {
		Insert(ctx context.Context, data *LoginLog) (*LoginLog, error)
		FindOne(ctx context.Context, id interface{}) (*LoginLog, error)
		Update(ctx context.Context, newData *LoginLog) error
		Delete(ctx context.Context, id interface{}) error
	}

	defaultLoginLogModel struct {
		sqlc.CachedConn
		table  string
		gormDB *gorm.DB
	}

	LoginLog struct {
		modelx.BaseModel
		Username      string    `json:"username" form:"username" db:"username" gorm:"column:username;comment:用户名;"`
		Status        string    `json:"status" form:"status" db:"status" gorm:"column:status;comment:状态;"`
		Ipaddr        string    `json:"ipaddr" form:"ipaddr" db:"ipaddr" gorm:"column:ipaddr;comment:ip地址;"`
		LoginLocation string    `json:"loginLocation" form:"loginLocation" db:"login_location" gorm:"column:login_location;comment:归属地;"`
		Browser       string    `json:"browser" form:"browser" db:"browser" gorm:"column:browser;comment:浏览器;"`
		Os            string    `json:"os" form:"os" db:"os" gorm:"column:os;comment:系统;"`
		Platform      string    `json:"platform" form:"platform" db:"platform" gorm:"column:platform;comment:固件;"`
		LoginTime     time.Time `json:"loginTime" form:"loginTime" db:"login_time" gorm:"column:login_time;comment:登录时间;"`
		Remark        string    `json:"remark" form:"remark" db:"remark" gorm:"column:remark;comment:备注;"`
		Msg           string    `json:"msg" form:"msg" db:"msg" gorm:"column:msg;comment:信息;"`
		modelx.ControlBy
		modelx.ModelTime
	}
)

// TableName LoginLog 表名
func (LoginLog) TableName() string {
	return "admin_login_log"
}

func newLoginLogModel(conn sqlx.SqlConn, c cache.CacheConf, gormx *gorm.DB) *defaultLoginLogModel {
	return &defaultLoginLogModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "admin_login_log",
		gormDB:     gormx,
	}
}

func (m *defaultLoginLogModel) FindOne(ctx context.Context, id interface{}) (*LoginLog, error) {
	var resp LoginLog
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

func (m *defaultLoginLogModel) Insert(ctx context.Context, data *LoginLog) (*LoginLog, error) {
	err := m.gormDB.Create(data).Error
	//var re sql.Result

	return data, err
}

func (m *defaultLoginLogModel) Update(ctx context.Context, newData *LoginLog) error {
	_, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}
	err = m.gormDB.Save(newData).Error
	return err
}

func (m *defaultLoginLogModel) Delete(ctx context.Context, id interface{}) error {
	_, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	err = m.gormDB.Delete(&LoginLog{}, id).Error
	return err
}

func (m *defaultLoginLogModel) tableName() string {
	return m.table
}
