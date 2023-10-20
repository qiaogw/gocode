package admin

import (
	"github.com/qiaogw/gocode/common/modelx"
)

type ListPostReq struct {
	modelx.Pagination `search:"-"`
	Id                int64  `json:"id" form:"id" search:"type:exact;column:id;table:admin_post"`
	Name              string `json:"name" form:"name" search:"type:exact;column:name;table:admin_post"`
	Code              string `json:"code" form:"code" search:"type:exact;column:code;table:admin_post"`
	Sort              int64  `json:"sort" form:"sort" search:"type:exact;column:sort;table:admin_post"`
	Remark            string `json:"remark" form:"remark" search:"type:exact;column:remark;table:admin_post"`
	PostOrder
}

type PostOrder struct {
	IdOrder     int64  `json:"id" form:"id" search:"type:order;column:id;table:admin_post"`
	NameOrder   string `json:"name" form:"name" search:"type:order;column:name;table:admin_post"`
	CodeOrder   string `json:"code" form:"code" search:"type:order;column:code;table:admin_post"`
	SortOrder   int64  `json:"sort" form:"sort" search:"type:order;column:sort;table:admin_post"`
	RemarkOrder string `json:"remark" form:"remark" search:"type:order;column:remark;table:admin_post"`
}

func (m *ListPostReq) GetNeedSearch() interface{} {
	return *m
}
