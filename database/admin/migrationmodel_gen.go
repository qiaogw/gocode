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
	migrationModel interface {
		Insert(ctx context.Context, data *Migration) (*Migration, error)
		FindOne(ctx context.Context, id interface{}) (*Migration, error)
		Update(ctx context.Context, newData *Migration) error
		Delete(ctx context.Context, id interface{}) error
	}

	defaultMigrationModel struct {
		sqlc.CachedConn
		table  string
		gormDB *gorm.DB
	}

	Migration struct {
		modelx.BaseModel
		Version string `json:"version" form:"version" db:"version" gorm:"column:version;comment:版本号;"`
		modelx.ControlBy
		modelx.ModelTime
	}
)

// TableName Migration 表名
func (Migration) TableName() string {
	return "admin_migration"
}

func newMigrationModel(conn sqlx.SqlConn, c cache.CacheConf, gormx *gorm.DB) *defaultMigrationModel {
	return &defaultMigrationModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "admin_migration",
		gormDB:     gormx,
	}
}

func (m *defaultMigrationModel) FindOne(ctx context.Context, id interface{}) (*Migration, error) {
	var resp Migration
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

func (m *defaultMigrationModel) Insert(ctx context.Context, data *Migration) (*Migration, error) {
	err := m.gormDB.Create(data).Error
	//var re sql.Result

	return data, err
}

func (m *defaultMigrationModel) Update(ctx context.Context, newData *Migration) error {
	_, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}
	err = m.gormDB.Save(newData).Error
	return err
}

func (m *defaultMigrationModel) Delete(ctx context.Context, id interface{}) error {
	_, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	err = m.gormDB.Delete(&Migration{}, id).Error
	return err
}

func (m *defaultMigrationModel) tableName() string {
	return m.table
}
