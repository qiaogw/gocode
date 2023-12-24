package adminmodel

import (
	"github.com/qiaogw/gocode/common/modelx"
	"time"
)

type ListPostReq struct {
	modelx.Pagination `search:"-"`
	Id                string    `json:"id" form:"id" search:"type:exact;column:id;table:admin_post"`
	Name              string    `json:"name" form:"name" search:"type:exact;column:name;table:admin_post"`
	Code              string    `json:"code" form:"code" search:"type:exact;column:code;table:admin_post"`
	Sort              int64     `json:"sort" form:"sort" search:"type:exact;column:sort;table:admin_post"`
	Remark            string    `json:"remark" form:"remark" search:"type:exact;column:remark;table:admin_post"`
	CreateBy          string    `json:"createBy" form:"createBy" search:"type:exact;column:create_by;table:admin_post"`
	UpdateBy          string    `json:"updateBy" form:"updateBy" search:"type:exact;column:update_by;table:admin_post"`
	CreatedAt         time.Time `json:"createdAt" form:"createdAt" search:"type:exact;column:created_at;table:admin_post"`
	UpdatedAt         time.Time `json:"updatedAt" form:"updatedAt" search:"type:exact;column:updated_at;table:admin_post"`
	DeletedAt         time.Time `json:"deletedAt" form:"deletedAt" search:"type:exact;column:deleted_at;table:admin_post"`
	PostOrder
}

type PostOrder struct {
	IdOrder        string    `json:"id" form:"id" search:"type:order;column:id;table:admin_post"`
	NameOrder      string    `json:"name" form:"name" search:"type:order;column:name;table:admin_post"`
	CodeOrder      string    `json:"code" form:"code" search:"type:order;column:code;table:admin_post"`
	SortOrder      int64     `json:"sort" form:"sort" search:"type:order;column:sort;table:admin_post"`
	RemarkOrder    string    `json:"remark" form:"remark" search:"type:order;column:remark;table:admin_post"`
	CreateByOrder  string    `json:"createBy" form:"createBy" search:"type:order;column:create_by;table:admin_post"`
	UpdateByOrder  string    `json:"updateBy" form:"updateBy" search:"type:order;column:update_by;table:admin_post"`
	CreatedAtOrder time.Time `json:"createdAt" form:"createdAt" search:"type:order;column:created_at;table:admin_post"`
	UpdatedAtOrder time.Time `json:"updatedAt" form:"updatedAt" search:"type:order;column:updated_at;table:admin_post"`
	DeletedAtOrder time.Time `json:"deletedAt" form:"deletedAt" search:"type:order;column:deleted_at;table:admin_post"`
}

func (m *ListPostReq) GetNeedSearch() interface{} {
	return *m
}
