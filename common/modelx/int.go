package modelx

type BaseModelInt struct {
	Id int64 `json:"id" comment:"主键编码" gorm:"column:id;primaryKey;comment:主键编码"`
}

type ControlByInt struct {
	CreateBy int64 `json:"createBy" comment:"创建者" gorm:"column:create_by;index;comment:创建者"`
	UpdateBy int64 `json:"updateBy" comment:"更新者" gorm:"column:update_by;index;comment:更新者"`
}

// SetCreateBy 设置创建人id
func (e *ControlByInt) SetCreateBy(createBy int64) {
	e.CreateBy = createBy
}

// SetUpdateBy 设置修改人id
func (e *ControlByInt) SetUpdateBy(updateBy int64) {
	e.UpdateBy = updateBy
}
