package modelx

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var ErrNotFound = gorm.ErrRecordNotFound

// BaseModel 主键
type BaseModel struct {
	Id uuid.UUID `json:"id" comment:"主键编码" gorm:"column:id;primaryKey;comment:主键编码"`
}

// ControlBy 控制字段
type ControlBy struct {
	CreateBy string `json:"createBy" comment:"创建者" gorm:"column:create_by;size:255;index;comment:创建者"`
	UpdateBy string `json:"updateBy" comment:"更新者" gorm:"column:update_by;size:255;index;comment:更新者"`
}

// ModelTime 时间字段
type ModelTime struct {
	CreatedAt time.Time      `json:"createdAt"  comment:"创建时间" gorm:"column:created_at;comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt"  comment:"最后更新时间" gorm:"column:updated_at;comment:最后更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" comment:"删除时间" gorm:"index;comment:删除时间"` // 软删除
}

// ModelUtils 排序、备注、可用
type ModelUtils struct {
	Sort    int64  `json:"sort" comment:"排序" gorm:"column:sort;comment:排序;"`
	Remark  string `json:"remark" comment:"备注" gorm:"column:remark;type:varchar(1024);comment:备注;"`
	Enabled bool   `json:"enabled" comment:"可用" gorm:"column:enabled;comment:可用状态;"`
}

// ModelWithCommon 通用模型
type ModelWithCommon struct {
	BaseModel
	ModelUtils
	ControlBy
	ModelTime
}

// Pagination 分页
type Pagination struct {
	PageIndex  int64  `form:"pageIndex" json:"pageIndex" gorm:"-"`
	PageSize   int64  `form:"pageSize" json:"pageSize"  gorm:"-"`
	SortBY     string `json:"sortBy,optional" gorm:"-"`
	Descending bool   `json:"descending,optional" gorm:"-"`
	SearchKey  string `json:"searchKey"  gorm:"-"`
}

func (m *Pagination) GetPageIndex() int64 {
	if m.PageIndex <= 0 {
		m.PageIndex = 1
	}
	return m.PageIndex
}

func (m *Pagination) GetPageSize() int64 {
	if m.PageSize <= 0 {
		m.PageSize = 10
	}
	return m.PageSize
}

// SetCreateBy 设置创建人id
func (e *ControlBy) SetCreateBy(createBy string) {
	e.CreateBy = createBy
}

// SetUpdateBy 设置修改人id
func (e *ControlBy) SetUpdateBy(updateBy string) {
	e.UpdateBy = updateBy
}

type ActiveRecord interface {
	schema.Tabler
	SetCreateBy(createBy int)
	SetUpdateBy(updateBy int)
	Generate() ActiveRecord
	GetId() interface{}
}
