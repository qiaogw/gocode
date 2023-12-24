package adminmodel

import (
	"github.com/qiaogw/gocode/common/modelx"
	"time"
)

type ListRoleReq struct {
	modelx.Pagination `search:"-"`
	Id                string    `json:"id" form:"id" search:"type:exact;column:id;table:admin_role"`
	Name              string    `json:"name" form:"name" search:"type:exact;column:name;table:admin_role"`
	Code              string    `json:"code" form:"code" search:"type:exact;column:code;table:admin_role"`
	Sort              int64     `json:"sort" form:"sort" search:"type:exact;column:sort;table:admin_role"`
	IsAdmin           bool      `json:"isAdmin" form:"isAdmin" search:"type:exact;column:is_admin;table:admin_role"`
	DataScope         string    `json:"dataScope" form:"dataScope" search:"type:exact;column:data_scope;table:admin_role"`
	DataFilter        string    `json:"dataFilter" form:"dataFilter" search:"type:exact;column:data_filter;table:admin_role"`
	DefaultRoute      string    `json:"defaultRoute" form:"defaultRoute" search:"type:exact;column:default_route;table:admin_role"`
	Remark            string    `json:"remark" form:"remark" search:"type:exact;column:remark;table:admin_role"`
	CreateBy          string    `json:"createBy" form:"createBy" search:"type:exact;column:create_by;table:admin_role"`
	UpdateBy          string    `json:"updateBy" form:"updateBy" search:"type:exact;column:update_by;table:admin_role"`
	CreatedAt         time.Time `json:"createdAt" form:"createdAt" search:"type:exact;column:created_at;table:admin_role"`
	UpdatedAt         time.Time `json:"updatedAt" form:"updatedAt" search:"type:exact;column:updated_at;table:admin_role"`
	DeletedAt         time.Time `json:"deletedAt" form:"deletedAt" search:"type:exact;column:deleted_at;table:admin_role"`
	RoleOrder
}

type RoleOrder struct {
	IdOrder           string    `json:"id" form:"id" search:"type:order;column:id;table:admin_role"`
	NameOrder         string    `json:"name" form:"name" search:"type:order;column:name;table:admin_role"`
	CodeOrder         string    `json:"code" form:"code" search:"type:order;column:code;table:admin_role"`
	SortOrder         int64     `json:"sort" form:"sort" search:"type:order;column:sort;table:admin_role"`
	IsAdminOrder      bool      `json:"isAdmin" form:"isAdmin" search:"type:order;column:is_admin;table:admin_role"`
	DataScopeOrder    string    `json:"dataScope" form:"dataScope" search:"type:order;column:data_scope;table:admin_role"`
	DataFilterOrder   string    `json:"dataFilter" form:"dataFilter" search:"type:order;column:data_filter;table:admin_role"`
	DefaultRouteOrder string    `json:"defaultRoute" form:"defaultRoute" search:"type:order;column:default_route;table:admin_role"`
	RemarkOrder       string    `json:"remark" form:"remark" search:"type:order;column:remark;table:admin_role"`
	CreateByOrder     string    `json:"createBy" form:"createBy" search:"type:order;column:create_by;table:admin_role"`
	UpdateByOrder     string    `json:"updateBy" form:"updateBy" search:"type:order;column:update_by;table:admin_role"`
	CreatedAtOrder    time.Time `json:"createdAt" form:"createdAt" search:"type:order;column:created_at;table:admin_role"`
	UpdatedAtOrder    time.Time `json:"updatedAt" form:"updatedAt" search:"type:order;column:updated_at;table:admin_role"`
	DeletedAtOrder    time.Time `json:"deletedAt" form:"deletedAt" search:"type:order;column:deleted_at;table:admin_role"`
}

func (m *ListRoleReq) GetNeedSearch() interface{} {
	return *m
}
