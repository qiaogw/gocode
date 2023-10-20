package admin

import (
	"github.com/qiaogw/gocode/common/modelx"
)

type ListRoleReq struct {
	modelx.Pagination `search:"-"`
	Id                int64  `json:"id" form:"id" search:"type:exact;column:id;table:admin_role"`
	Name              string `json:"name" form:"name" search:"type:exact;column:name;table:admin_role"`
	Code              string `json:"code" form:"code" search:"type:exact;column:code;table:admin_role"`
	Sort              int64  `json:"sort" form:"sort" search:"type:exact;column:sort;table:admin_role"`
	IsAdmin           bool   `json:"isAdmin" form:"isAdmin" search:"type:exact;column:is_admin;table:admin_role"`
	DataScope         string `json:"dataScope" form:"dataScope" search:"type:exact;column:data_scope;table:admin_role"`
	DefaultRoute      string `json:"defaultRoute" form:"defaultRoute" search:"type:exact;column:default_route;table:admin_role"`
	Remark            string `json:"remark" form:"remark" search:"type:exact;column:remark;table:admin_role"`
	RoleOrder
}

type RoleOrder struct {
	IdOrder           int64  `json:"id" form:"id" search:"type:order;column:id;table:admin_role"`
	NameOrder         string `json:"name" form:"name" search:"type:order;column:name;table:admin_role"`
	CodeOrder         string `json:"code" form:"code" search:"type:order;column:code;table:admin_role"`
	SortOrder         int64  `json:"sort" form:"sort" search:"type:order;column:sort;table:admin_role"`
	IsAdminOrder      bool   `json:"isAdmin" form:"isAdmin" search:"type:order;column:is_admin;table:admin_role"`
	DataScopeOrder    string `json:"dataScope" form:"dataScope" search:"type:order;column:data_scope;table:admin_role"`
	DefaultRouteOrder string `json:"defaultRoute" form:"defaultRoute" search:"type:order;column:default_route;table:admin_role"`
	RemarkOrder       string `json:"remark" form:"remark" search:"type:order;column:remark;table:admin_role"`
}

func (m *ListRoleReq) GetNeedSearch() interface{} {
	return *m
}
