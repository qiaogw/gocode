package gencode

import (
	"github.com/qiaogw/gocode/common/modelx"
)

type ListGenColumnReq struct {
	modelx.Pagination `search:"-"`
	Id                int64  `json:"id" form:"id" search:"type:exact;column:id;table:gen_column"`
	Name              string `json:"name" form:"name" search:"type:exact;column:name;table:gen_column"`
	GormName          string `json:"gormName" form:"gormName" search:"type:exact;column:gorm_name;table:gen_column"`
	DbType            string `json:"dbType" form:"dbType" search:"type:exact;column:db_type;table:gen_column"`
	DataType          string `json:"dataType" form:"dataType" search:"type:exact;column:data_type;table:gen_column"`
	DataTypeLong      string `json:"dataTypeLong" form:"dataTypeLong" search:"type:exact;column:data_type_long;table:gen_column"`
	Extra             string `json:"extra" form:"extra" search:"type:exact;column:extra;table:gen_column"`
	ColumnComment     string `json:"columnComment" form:"columnComment" search:"type:exact;column:column_comment;table:gen_column"`
	ColumnDefault     string `json:"columnDefault" form:"columnDefault" search:"type:exact;column:column_default;table:gen_column"`
	SiNullable        string `json:"siNullable" form:"siNullable" search:"type:exact;column:si_nullable;table:gen_column"`
	OrdinalPosition   int64  `json:"ordinalPosition" form:"ordinalPosition" search:"type:exact;column:ordinal_position;table:gen_column"`
	FieldJson         string `json:"fieldJson" form:"fieldJson" search:"type:exact;column:field_json;table:gen_column"`
	Tablename         string `json:"tablename" form:"tablename" search:"type:exact;column:tablename;table:gen_column"`
	FieldName         string `json:"fieldName" form:"fieldName" search:"type:exact;column:field_name;table:gen_column"`
	DictType          string `json:"dictType" form:"dictType" search:"type:exact;column:dict_type;table:gen_column"`
	ErrorText         string `json:"errorText" form:"errorText" search:"type:exact;column:error_text;table:gen_column"`
	IsPage            bool   `json:"isPage" form:"isPage" search:"type:exact;column:is_page;table:gen_column"`
	IsNull            bool   `json:"isNull" form:"isNull" search:"type:exact;column:is_null;table:gen_column"`
	Clearable         bool   `json:"clearable" form:"clearable" search:"type:exact;column:clearable;table:gen_column"`
	Require           bool   `json:"require" form:"require" search:"type:exact;column:require;table:gen_column"`
	IsPk              bool   `json:"isPk" form:"isPk" search:"type:exact;column:is_pk;table:gen_column"`
	IsList            bool   `json:"isList" form:"isList" search:"type:exact;column:is_list;table:gen_column"`
	Increment         bool   `json:"increment" form:"increment" search:"type:exact;column:increment;table:gen_column"`
	IsEdit            bool   `json:"isEdit" form:"isEdit" search:"type:exact;column:is_edit;table:gen_column"`
	IsSort            bool   `json:"isSort" form:"isSort" search:"type:exact;column:is_sort;table:gen_column"`
	HtmlType          string `json:"htmlType" form:"htmlType" search:"type:exact;column:html_type;table:gen_column"`
	FkTable           string `json:"fkTable" form:"fkTable" search:"type:exact;column:fk_table;table:gen_column"`
	FkTableClass      string `json:"fkTableClass" form:"fkTableClass" search:"type:exact;column:fk_table_class;table:gen_column"`
	FkTablePackage    string `json:"fkTablePackage" form:"fkTablePackage" search:"type:exact;column:fk_table_package;table:gen_column"`
	FkLabelId         string `json:"fkLabelId" form:"fkLabelId" search:"type:exact;column:fk_label_id;table:gen_column"`
	FkLabelName       string `json:"fkLabelName" form:"fkLabelName" search:"type:exact;column:fk_label_name;table:gen_column"`
	Remark            string `json:"remark" form:"remark" search:"type:exact;column:remark;table:gen_column"`
	GenColumnOrder
}

type GenColumnOrder struct {
	IdOrder              int64  `json:"id" form:"id" search:"type:order;column:id;table:gen_column"`
	NameOrder            string `json:"name" form:"name" search:"type:order;column:name;table:gen_column"`
	GormNameOrder        string `json:"gormName" form:"gormName" search:"type:order;column:gorm_name;table:gen_column"`
	DbTypeOrder          string `json:"dbType" form:"dbType" search:"type:order;column:db_type;table:gen_column"`
	DataTypeOrder        string `json:"dataType" form:"dataType" search:"type:order;column:data_type;table:gen_column"`
	DataTypeLongOrder    string `json:"dataTypeLong" form:"dataTypeLong" search:"type:order;column:data_type_long;table:gen_column"`
	ExtraOrder           string `json:"extra" form:"extra" search:"type:order;column:extra;table:gen_column"`
	ColumnCommentOrder   string `json:"columnComment" form:"columnComment" search:"type:order;column:column_comment;table:gen_column"`
	ColumnDefaultOrder   string `json:"columnDefault" form:"columnDefault" search:"type:order;column:column_default;table:gen_column"`
	SiNullableOrder      string `json:"siNullable" form:"siNullable" search:"type:order;column:si_nullable;table:gen_column"`
	OrdinalPositionOrder int64  `json:"ordinalPosition" form:"ordinalPosition" search:"type:order;column:ordinal_position;table:gen_column"`
	FieldJsonOrder       string `json:"fieldJson" form:"fieldJson" search:"type:order;column:field_json;table:gen_column"`
	TablenameOrder       string `json:"tablename" form:"tablename" search:"type:order;column:tablename;table:gen_column"`
	FieldNameOrder       string `json:"fieldName" form:"fieldName" search:"type:order;column:field_name;table:gen_column"`
	DictTypeOrder        string `json:"dictType" form:"dictType" search:"type:order;column:dict_type;table:gen_column"`
	ErrorTextOrder       string `json:"errorText" form:"errorText" search:"type:order;column:error_text;table:gen_column"`
	IsPageOrder          bool   `json:"isPage" form:"isPage" search:"type:order;column:is_page;table:gen_column"`
	IsNullOrder          bool   `json:"isNull" form:"isNull" search:"type:order;column:is_null;table:gen_column"`
	ClearableOrder       bool   `json:"clearable" form:"clearable" search:"type:order;column:clearable;table:gen_column"`
	RequireOrder         bool   `json:"require" form:"require" search:"type:order;column:require;table:gen_column"`
	IsPkOrder            bool   `json:"isPk" form:"isPk" search:"type:order;column:is_pk;table:gen_column"`
	IsListOrder          bool   `json:"isList" form:"isList" search:"type:order;column:is_list;table:gen_column"`
	IncrementOrder       bool   `json:"increment" form:"increment" search:"type:order;column:increment;table:gen_column"`
	IsEditOrder          bool   `json:"isEdit" form:"isEdit" search:"type:order;column:is_edit;table:gen_column"`
	IsSortOrder          bool   `json:"isSort" form:"isSort" search:"type:order;column:is_sort;table:gen_column"`
	HtmlTypeOrder        string `json:"htmlType" form:"htmlType" search:"type:order;column:html_type;table:gen_column"`
	FkTableOrder         string `json:"fkTable" form:"fkTable" search:"type:order;column:fk_table;table:gen_column"`
	FkTableClassOrder    string `json:"fkTableClass" form:"fkTableClass" search:"type:order;column:fk_table_class;table:gen_column"`
	FkTablePackageOrder  string `json:"fkTablePackage" form:"fkTablePackage" search:"type:order;column:fk_table_package;table:gen_column"`
	FkLabelIdOrder       string `json:"fkLabelId" form:"fkLabelId" search:"type:order;column:fk_label_id;table:gen_column"`
	FkLabelNameOrder     string `json:"fkLabelName" form:"fkLabelName" search:"type:order;column:fk_label_name;table:gen_column"`
	RemarkOrder          string `json:"remark" form:"remark" search:"type:order;column:remark;table:gen_column"`
}

func (m *ListGenColumnReq) GetNeedSearch() interface{} {
	return *m
}
