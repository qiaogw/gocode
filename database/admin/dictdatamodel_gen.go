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
	dictDataModel interface {
		Insert(ctx context.Context, data *DictData) (*DictData, error)
		FindOne(ctx context.Context, id interface{}) (*DictData, error)
		Update(ctx context.Context, newData *DictData) error
		Delete(ctx context.Context, id interface{}) error
	}

	defaultDictDataModel struct {
		sqlc.CachedConn
		table  string
		gormDB *gorm.DB
	}

	DictData struct {
		modelx.BaseModel
		DictTypeId int64     `json:"dictTypeId" form:"dictTypeId" db:"dict_type_id" gorm:"column:dict_type_id;comment:字典_主键;"`
		Sort       int64     `json:"sort" form:"sort" db:"sort" gorm:"column:sort;comment:排序;"`
		Label      string    `json:"label" form:"label" db:"label" gorm:"column:label;comment:字典标签;"`
		Value      string    `json:"value" form:"value" db:"value" gorm:"column:value;comment:字典值;"`
		IsDefault  bool      `json:"isDefault" form:"isDefault" db:"is_default" gorm:"column:is_default;comment:是否默认;"`
		Enabled    bool      `json:"enabled" form:"enabled" db:"enabled" gorm:"column:enabled;comment:可用;"`
		Remark     string    `json:"remark" form:"remark" db:"remark" gorm:"column:remark;comment:备注;"`
		DictType   *DictType `json:"dictType" db:"-" `
		modelx.ControlBy
		modelx.ModelTime
	}
)

// TableName DictData 表名
func (DictData) TableName() string {
	return "admin_dict_data"
}

func newDictDataModel(conn sqlx.SqlConn, c cache.CacheConf, gormx *gorm.DB) *defaultDictDataModel {
	return &defaultDictDataModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "admin_dict_data",
		gormDB:     gormx,
	}
}

func (m *defaultDictDataModel) FindOne(ctx context.Context, id interface{}) (*DictData, error) {
	var resp DictData
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

func (m *defaultDictDataModel) Insert(ctx context.Context, data *DictData) (*DictData, error) {
	err := m.gormDB.Create(data).Error
	//var re sql.Result

	return data, err
}

func (m *defaultDictDataModel) Update(ctx context.Context, newData *DictData) error {
	_, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}
	err = m.gormDB.Save(newData).Error
	return err
}

func (m *defaultDictDataModel) Delete(ctx context.Context, id interface{}) error {
	_, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	err = m.gormDB.Delete(&DictData{}, id).Error
	return err
}

func (m *defaultDictDataModel) tableName() string {
	return m.table
}
