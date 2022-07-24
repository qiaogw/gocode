package global

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id int64 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
}

type ControlBy struct {
	CreateBy int64  `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy int64  `json:"updateBy" gorm:"index;comment:更新者"`
	Unit     string `json:"unit" gorm:"-"`
}

type ModelTime struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:最后更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"` // 软删除
}

// SetCreateBy 设置创建人id
func (e *ControlBy) SetCreateBy(createBy int64) {
	e.CreateBy = createBy
}

// SetUpdateBy 设置修改人id
func (e *ControlBy) SetUpdateBy(updateBy int64) {
	e.UpdateBy = updateBy
}

type Pagination struct {
	PageIndex int `form:"pageIndex" json:"pageIndex"`
	PageSize  int `form:"pageSize" json:"pageSize"`
}

func (m *Pagination) GetPageIndex() int {
	if m.PageIndex <= 0 {
		m.PageIndex = 1
	}
	return m.PageIndex
}

func (m *Pagination) GetPageSize() int {
	if m.PageSize <= 0 {
		m.PageSize = 10
	}
	return m.PageSize
}
