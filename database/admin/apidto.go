package admin

import "github.com/qiaogw/gocode/common/modelx"

type ListApiReq struct {
	modelx.Pagination `search:"-"`
	Id                int64  `json:"id" form:"id" search:"type:exact;column:id;table:admin_api"`
	Title             string `json:"title" form:"title" search:"type:exact;column:title;table:admin_api"`
	Path              string `json:"path" form:"path" search:"type:exact;column:path;table:admin_api"`
	Method            string `json:"method" form:"method" search:"type:exact;column:method;table:admin_api"`
	Module            string `json:"module" form:"module" search:"type:exact;column:module;table:admin_api"`
	Remark            string `json:"remark" form:"remark" search:"type:exact;column:remark;table:admin_api"`
	ApiOrder
}

type ApiOrder struct {
	IdOrder     int64  `json:"id" form:"id" search:"type:order;column:id;table:admin_api"`
	TitleOrder  string `json:"title" form:"title" search:"type:order;column:title;table:admin_api"`
	PathOrder   string `json:"path" form:"path" search:"type:order;column:path;table:admin_api"`
	MethodOrder string `json:"method" form:"method" search:"type:order;column:method;table:admin_api"`
	ModuleOrder string `json:"module" form:"module" search:"type:order;column:module;table:admin_api"`
	RemarkOrder string `json:"remark" form:"remark" search:"type:order;column:remark;table:admin_api"`
}

func (m *ListApiReq) GetNeedSearch() interface{} {
	return *m
}
