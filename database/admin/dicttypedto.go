package admin

import (
	"github.com/qiaogw/gocode/common/modelx"
)

type ListDictTypeReq struct {
	modelx.Pagination `search:"-"`
	Id                int64  `json:"id" form:"id" search:"type:exact;column:id;table:admin_dict_type"`
	Name              string `json:"name" form:"name" search:"type:exact;column:name;table:admin_dict_type"`
	Type              string `json:"type" form:"type" search:"type:exact;column:type;table:admin_dict_type"`
	Enabled           bool   `json:"enabled" form:"enabled" search:"type:exact;column:enabled;table:admin_dict_type"`
	Remark            string `json:"remark" form:"remark" search:"type:exact;column:remark;table:admin_dict_type"`

	DictTypeOrder
}

type DictTypeOrder struct {
	IdOrder      int64  `json:"id" form:"id" search:"type:order;column:id;table:admin_dict_type"`
	NameOrder    string `json:"name" form:"name" search:"type:order;column:name;table:admin_dict_type"`
	TypeOrder    string `json:"type" form:"type" search:"type:order;column:type;table:admin_dict_type"`
	EnabledOrder bool   `json:"enabled" form:"enabled" search:"type:order;column:enabled;table:admin_dict_type"`
	RemarkOrder  string `json:"remark" form:"remark" search:"type:order;column:remark;table:admin_dict_type"`
}

func (m *ListDictTypeReq) GetNeedSearch() interface{} {
	return *m
}
