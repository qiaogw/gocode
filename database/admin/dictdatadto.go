package admin

import (
	"github.com/qiaogw/gocode/common/modelx"
)

type ListDictDataReq struct {
	modelx.Pagination `search:"-"`
	Id                int64  `json:"id" form:"id" search:"type:exact;column:id;table:admin_dict_data"`
	DictTypeId        int64  `json:"dictTypeId" form:"dictTypeId" search:"type:exact;column:dict_type_id;table:admin_dict_data"`
	Sort              int64  `json:"sort" form:"sort" search:"type:exact;column:sort;table:admin_dict_data"`
	Label             string `json:"label" form:"label" search:"type:exact;column:label;table:admin_dict_data"`
	Value             string `json:"value" form:"value" search:"type:exact;column:value;table:admin_dict_data"`
	IsDefault         bool   `json:"isDefault" form:"isDefault" search:"type:exact;column:is_default;table:admin_dict_data"`
	Enabled           bool   `json:"enabled" form:"enabled" search:"type:exact;column:enabled;table:admin_dict_data"`
	Remark            string `json:"remark" form:"remark" search:"type:exact;column:remark;table:admin_dict_data"`
	DictDataOrder
}

type DictDataOrder struct {
	IdOrder         int64  `json:"id" form:"id" search:"type:order;column:id;table:admin_dict_data"`
	DictTypeIdOrder int64  `json:"dictTypeId" form:"dictTypeId" search:"type:order;column:dict_type_id;table:admin_dict_data"`
	SortOrder       int64  `json:"sort" form:"sort" search:"type:order;column:sort;table:admin_dict_data"`
	LabelOrder      string `json:"label" form:"label" search:"type:order;column:label;table:admin_dict_data"`
	ValueOrder      string `json:"value" form:"value" search:"type:order;column:value;table:admin_dict_data"`
	IsDefaultOrder  bool   `json:"isDefault" form:"isDefault" search:"type:order;column:is_default;table:admin_dict_data"`
	EnabledOrder    bool   `json:"enabled" form:"enabled" search:"type:order;column:enabled;table:admin_dict_data"`
	RemarkOrder     string `json:"remark" form:"remark" search:"type:order;column:remark;table:admin_dict_data"`
}

func (m *ListDictDataReq) GetNeedSearch() interface{} {
	return *m
}
