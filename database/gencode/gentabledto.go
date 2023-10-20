package gencode

import (
	"github.com/qiaogw/gocode/common/modelx"
)

type ListGenTableReq struct {
	modelx.Pagination `search:"-"`
	Id                int64  `json:"id" form:"id" search:"type:exact;column:id;table:gen_table"`
	Db                string `json:"db" form:"db" search:"type:exact;column:db;table:gen_table"`
	Table             string `json:"table" form:"table" search:"type:exact;column:table;table:gen_table"`
	Name              string `json:"name" form:"name" search:"type:exact;column:name;table:gen_table"`
	PackageName       string `json:"packageName" form:"packageName" search:"type:exact;column:package_name;table:gen_table"`
	TableUrl          string `json:"tableUrl" form:"tableUrl" search:"type:exact;column:table_url;table:gen_table"`
	HasTimer          bool   `json:"hasTimer" form:"hasTimer" search:"type:exact;column:has_timer;table:gen_table"`
	HasCacheKey       bool   `json:"hasCacheKey" form:"hasCacheKey" search:"type:exact;column:has_cache_key;table:gen_table"`
	NeedValid         bool   `json:"needValid" form:"needValid" search:"type:exact;column:need_valid;table:gen_table"`
	PostgreSql        bool   `json:"postgreSql" form:"postgreSql" search:"type:exact;column:postgre_sql;table:gen_table"`
	TableComment      string `json:"tableComment" form:"tableComment" search:"type:exact;column:table_comment;table:gen_table"`
	Author            string `json:"author" form:"author" search:"type:exact;column:author;table:gen_table"`
	Email             string `json:"email" form:"email" search:"type:exact;column:email;table:gen_table"`
	Pkg               string `json:"pkg" form:"pkg" search:"type:exact;column:pkg;table:gen_table"`
	Service           string `json:"service" form:"service" search:"type:exact;column:service;table:gen_table"`
	ParentPkg         string `json:"parentPkg" form:"parentPkg" search:"type:exact;column:parent_pkg;table:gen_table"`
	IsCurd            bool   `json:"isCurd" form:"isCurd" search:"type:exact;column:is_curd;table:gen_table"`
	IsDataScope       bool   `json:"isDataScope" form:"isDataScope" search:"type:exact;column:is_data_scope;table:gen_table"`
	IsAuth            bool   `json:"isAuth" form:"isAuth" search:"type:exact;column:is_auth;table:gen_table"`
	Remark            string `json:"remark" form:"remark" search:"type:exact;column:remark;table:gen_table"`
	GenTableOrder
}

type GenTableOrder struct {
	IdOrder           int64  `json:"id" form:"id" search:"type:order;column:id;table:gen_table"`
	DbOrder           string `json:"db" form:"db" search:"type:order;column:db;table:gen_table"`
	TableOrder        string `json:"table" form:"table" search:"type:order;column:table;table:gen_table"`
	NameOrder         string `json:"name" form:"name" search:"type:order;column:name;table:gen_table"`
	PackageNameOrder  string `json:"packageName" form:"packageName" search:"type:order;column:package_name;table:gen_table"`
	TableUrlOrder     string `json:"tableUrl" form:"tableUrl" search:"type:order;column:table_url;table:gen_table"`
	HasTimerOrder     bool   `json:"hasTimer" form:"hasTimer" search:"type:order;column:has_timer;table:gen_table"`
	HasCacheKeyOrder  bool   `json:"hasCacheKey" form:"hasCacheKey" search:"type:order;column:has_cache_key;table:gen_table"`
	NeedValidOrder    bool   `json:"needValid" form:"needValid" search:"type:order;column:need_valid;table:gen_table"`
	PostgreSqlOrder   bool   `json:"postgreSql" form:"postgreSql" search:"type:order;column:postgre_sql;table:gen_table"`
	TableCommentOrder string `json:"tableComment" form:"tableComment" search:"type:order;column:table_comment;table:gen_table"`
	AuthorOrder       string `json:"author" form:"author" search:"type:order;column:author;table:gen_table"`
	EmailOrder        string `json:"email" form:"email" search:"type:order;column:email;table:gen_table"`
	PkgOrder          string `json:"pkg" form:"pkg" search:"type:order;column:pkg;table:gen_table"`
	ServiceOrder      string `json:"service" form:"service" search:"type:order;column:service;table:gen_table"`
	ParentPkgOrder    string `json:"parentPkg" form:"parentPkg" search:"type:order;column:parent_pkg;table:gen_table"`
	IsCurdOrder       bool   `json:"isCurd" form:"isCurd" search:"type:order;column:is_curd;table:gen_table"`
	IsDataScopeOrder  bool   `json:"isDataScope" form:"isDataScope" search:"type:order;column:is_data_scope;table:gen_table"`
	IsAuthOrder       bool   `json:"isAuth" form:"isAuth" search:"type:order;column:is_auth;table:gen_table"`
	RemarkOrder       string `json:"remark" form:"remark" search:"type:order;column:remark;table:gen_table"`
}

func (m *ListGenTableReq) GetNeedSearch() interface{} {
	return *m
}
