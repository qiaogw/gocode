package adminmodel

import (
	"github.com/qiaogw/gocode/common/modelx"
	"time"
)

type ListApiReq struct {
	modelx.Pagination `search:"-"`
	Id                string    `json:"id" form:"id" search:"type:exact;column:id;table:admin_api"`
	Title             string    `json:"title" form:"title" search:"type:exact;column:title;table:admin_api"`
	Path              string    `json:"path" form:"path" search:"type:exact;column:path;table:admin_api"`
	Method            string    `json:"method" form:"method" search:"type:exact;column:method;table:admin_api"`
	Module            string    `json:"module" form:"module" search:"type:exact;column:module;table:admin_api"`
	Remark            string    `json:"remark" form:"remark" search:"type:exact;column:remark;table:admin_api"`
	CreatedAt         time.Time `json:"createdAt" form:"createdAt" search:"type:exact;column:created_at;table:admin_api"`
	UpdatedAt         time.Time `json:"updatedAt" form:"updatedAt" search:"type:exact;column:updated_at;table:admin_api"`
	DeletedAt         time.Time `json:"deletedAt" form:"deletedAt" search:"type:exact;column:deleted_at;table:admin_api"`
	CreateBy          string    `json:"createBy" form:"createBy" search:"type:exact;column:create_by;table:admin_api"`
	UpdateBy          string    `json:"updateBy" form:"updateBy" search:"type:exact;column:update_by;table:admin_api"`
	ApiOrder
}

type ApiOrder struct {
	IdOrder        string    `json:"id" form:"id" search:"type:order;column:id;table:admin_api"`
	TitleOrder     string    `json:"title" form:"title" search:"type:order;column:title;table:admin_api"`
	PathOrder      string    `json:"path" form:"path" search:"type:order;column:path;table:admin_api"`
	MethodOrder    string    `json:"method" form:"method" search:"type:order;column:method;table:admin_api"`
	ModuleOrder    string    `json:"module" form:"module" search:"type:order;column:module;table:admin_api"`
	RemarkOrder    string    `json:"remark" form:"remark" search:"type:order;column:remark;table:admin_api"`
	CreatedAtOrder time.Time `json:"createdAt" form:"createdAt" search:"type:order;column:created_at;table:admin_api"`
	UpdatedAtOrder time.Time `json:"updatedAt" form:"updatedAt" search:"type:order;column:updated_at;table:admin_api"`
	DeletedAtOrder time.Time `json:"deletedAt" form:"deletedAt" search:"type:order;column:deleted_at;table:admin_api"`
	CreateByOrder  string    `json:"createBy" form:"createBy" search:"type:order;column:create_by;table:admin_api"`
	UpdateByOrder  string    `json:"updateBy" form:"updateBy" search:"type:order;column:update_by;table:admin_api"`
}

func (m *ListApiReq) GetNeedSearch() interface{} {
	return *m
}
