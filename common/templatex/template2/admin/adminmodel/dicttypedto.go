package adminmodel

import (
	"github.com/qiaogw/gocode/common/modelx"
	"time"
)

type ListDictTypeReq struct {
	modelx.Pagination `search:"-"`
	Id                string    `json:"id" form:"id" search:"type:exact;column:id;table:admin_dict_type"`
	Name              string    `json:"name" form:"name" search:"type:exact;column:name;table:admin_dict_type"`
	Type              string    `json:"type" form:"type" search:"type:exact;column:type;table:admin_dict_type"`
	Enabled           bool      `json:"enabled" form:"enabled" search:"type:exact;column:enabled;table:admin_dict_type"`
	Remark            string    `json:"remark" form:"remark" search:"type:exact;column:remark;table:admin_dict_type"`
	CreateBy          string    `json:"createBy" form:"createBy" search:"type:exact;column:create_by;table:admin_dict_type"`
	UpdateBy          string    `json:"updateBy" form:"updateBy" search:"type:exact;column:update_by;table:admin_dict_type"`
	CreatedAt         time.Time `json:"createdAt" form:"createdAt" search:"type:exact;column:created_at;table:admin_dict_type"`
	UpdatedAt         time.Time `json:"updatedAt" form:"updatedAt" search:"type:exact;column:updated_at;table:admin_dict_type"`
	DeletedAt         time.Time `json:"deletedAt" form:"deletedAt" search:"type:exact;column:deleted_at;table:admin_dict_type"`
	DictTypeOrder
}

type DictTypeOrder struct {
	IdOrder        string    `json:"id" form:"id" search:"type:order;column:id;table:admin_dict_type"`
	NameOrder      string    `json:"name" form:"name" search:"type:order;column:name;table:admin_dict_type"`
	TypeOrder      string    `json:"type" form:"type" search:"type:order;column:type;table:admin_dict_type"`
	EnabledOrder   bool      `json:"enabled" form:"enabled" search:"type:order;column:enabled;table:admin_dict_type"`
	RemarkOrder    string    `json:"remark" form:"remark" search:"type:order;column:remark;table:admin_dict_type"`
	CreateByOrder  string    `json:"createBy" form:"createBy" search:"type:order;column:create_by;table:admin_dict_type"`
	UpdateByOrder  string    `json:"updateBy" form:"updateBy" search:"type:order;column:update_by;table:admin_dict_type"`
	CreatedAtOrder time.Time `json:"createdAt" form:"createdAt" search:"type:order;column:created_at;table:admin_dict_type"`
	UpdatedAtOrder time.Time `json:"updatedAt" form:"updatedAt" search:"type:order;column:updated_at;table:admin_dict_type"`
	DeletedAtOrder time.Time `json:"deletedAt" form:"deletedAt" search:"type:order;column:deleted_at;table:admin_dict_type"`
}

func (m *ListDictTypeReq) GetNeedSearch() interface{} {
	return *m
}
