package adminmodel

import (
	"github.com/qiaogw/gocode/common/modelx"
	"time"
)

type ListDictDataReq struct {
	modelx.Pagination `search:"-"`
	Id                string    `json:"id" form:"id" search:"type:exact;column:id;table:admin_dict_data"`
	DictTypeId        string    `json:"dictTypeId" form:"dictTypeId" search:"type:exact;column:dict_type_id;table:admin_dict_data"`
	Sort              int64     `json:"sort" form:"sort" search:"type:exact;column:sort;table:admin_dict_data"`
	Label             string    `json:"label" form:"label" search:"type:exact;column:label;table:admin_dict_data"`
	Value             string    `json:"value" form:"value" search:"type:exact;column:value;table:admin_dict_data"`
	IsDefault         bool      `json:"isDefault" form:"isDefault" search:"type:exact;column:is_default;table:admin_dict_data"`
	Enabled           bool      `json:"enabled" form:"enabled" search:"type:exact;column:enabled;table:admin_dict_data"`
	Remark            string    `json:"remark" form:"remark" search:"type:exact;column:remark;table:admin_dict_data"`
	CreateBy          string    `json:"createBy" form:"createBy" search:"type:exact;column:create_by;table:admin_dict_data"`
	UpdateBy          string    `json:"updateBy" form:"updateBy" search:"type:exact;column:update_by;table:admin_dict_data"`
	CreatedAt         time.Time `json:"createdAt" form:"createdAt" search:"type:exact;column:created_at;table:admin_dict_data"`
	UpdatedAt         time.Time `json:"updatedAt" form:"updatedAt" search:"type:exact;column:updated_at;table:admin_dict_data"`
	DeletedAt         time.Time `json:"deletedAt" form:"deletedAt" search:"type:exact;column:deleted_at;table:admin_dict_data"`
	DictDataOrder
}

type DictDataOrder struct {
	IdOrder         string    `json:"id" form:"id" search:"type:order;column:id;table:admin_dict_data"`
	DictTypeIdOrder string    `json:"dictTypeId" form:"dictTypeId" search:"type:order;column:dict_type_id;table:admin_dict_data"`
	SortOrder       int64     `json:"sort" form:"sort" search:"type:order;column:sort;table:admin_dict_data"`
	LabelOrder      string    `json:"label" form:"label" search:"type:order;column:label;table:admin_dict_data"`
	ValueOrder      string    `json:"value" form:"value" search:"type:order;column:value;table:admin_dict_data"`
	IsDefaultOrder  bool      `json:"isDefault" form:"isDefault" search:"type:order;column:is_default;table:admin_dict_data"`
	EnabledOrder    bool      `json:"enabled" form:"enabled" search:"type:order;column:enabled;table:admin_dict_data"`
	RemarkOrder     string    `json:"remark" form:"remark" search:"type:order;column:remark;table:admin_dict_data"`
	CreateByOrder   string    `json:"createBy" form:"createBy" search:"type:order;column:create_by;table:admin_dict_data"`
	UpdateByOrder   string    `json:"updateBy" form:"updateBy" search:"type:order;column:update_by;table:admin_dict_data"`
	CreatedAtOrder  time.Time `json:"createdAt" form:"createdAt" search:"type:order;column:created_at;table:admin_dict_data"`
	UpdatedAtOrder  time.Time `json:"updatedAt" form:"updatedAt" search:"type:order;column:updated_at;table:admin_dict_data"`
	DeletedAtOrder  time.Time `json:"deletedAt" form:"deletedAt" search:"type:order;column:deleted_at;table:admin_dict_data"`
}

func (m *ListDictDataReq) GetNeedSearch() interface{} {
	return *m
}
