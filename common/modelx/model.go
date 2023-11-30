package modelx

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var ErrNotFound = gorm.ErrRecordNotFound

type BaseModel struct {
	Id uuid.UUID `json:"id" db:"id" gorm:"column:id;primaryKey;comment:主键编码"`
}

type ControlBy struct {
	CreateBy string `json:"createBy" db:"create_by" gorm:"column:create_by;size:255;index;comment:创建者"`
	UpdateBy string `json:"updateBy" db:"update_by" gorm:"column:update_by;size:255;index;comment:更新者"`
}

type ModelTime struct {
	CreatedAt time.Time      `json:"createdAt"  db:"created_at" gorm:"column:created_at;comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt"  db:"updated_at" gorm:"column:updated_at;comment:最后更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" db:"deleted_at" gorm:"index;comment:删除时间"` // 软删除
}

// SetCreateBy 设置创建人id
func (e *ControlBy) SetCreateBy(createBy string) {
	e.CreateBy = createBy
}

// SetUpdateBy 设置修改人id
func (e *ControlBy) SetUpdateBy(updateBy string) {
	e.UpdateBy = updateBy
}

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

type ActiveRecord interface {
	schema.Tabler
	SetCreateBy(createBy int)
	SetUpdateBy(updateBy int)
	Generate() ActiveRecord
	GetId() interface{}
}
