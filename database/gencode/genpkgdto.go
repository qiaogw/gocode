package gencode

import (
	"time"
)

type GenPkgOrder struct {
	IdOrder        int64     `json:"id" form:"id" search:"type:order;column:id;table:gen_pkg"`
	NameOrder      string    `json:"name" form:"name" search:"type:order;column:name;table:gen_pkg"`
	LabelOrder     string    `json:"label" form:"label" search:"type:order;column:label;table:gen_pkg"`
	SortOrder      int64     `json:"sort" form:"sort" search:"type:order;column:sort;table:gen_pkg"`
	RemarkOrder    string    `json:"remark" form:"remark" search:"type:order;column:remark;table:gen_pkg"`
	CreateByOrder  int64     `json:"createBy" form:"createBy" search:"type:order;column:create_by;table:gen_pkg"`
	UpdateByOrder  int64     `json:"updateBy" form:"updateBy" search:"type:order;column:update_by;table:gen_pkg"`
	CreatedAtOrder time.Time `json:"createdAt" form:"createdAt" search:"type:order;column:created_at;table:gen_pkg"`
	UpdatedAtOrder time.Time `json:"updatedAt" form:"updatedAt" search:"type:order;column:updated_at;table:gen_pkg"`
	DeletedAtOrder time.Time `json:"deletedAt" form:"deletedAt" search:"type:order;column:deleted_at;table:gen_pkg"`
	BeginTimeOrder time.Time `json:"beginTime" form:"beginTime" search:"type:order;column:begin_time;table:gen_pkg"`
	PriceOrder     float64   `json:"price" form:"price" search:"type:order;column:price;table:gen_pkg"`
}

//
//func (m *ListGenPkgReq) GetNeedSearch() interface{} {
//	return *m
//}
