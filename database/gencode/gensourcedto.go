package gencode

import (
	"github.com/qiaogw/gocode/common/modelx"
)

type ListGenSourceReq struct {
	modelx.Pagination `search:"-"`
	Id                int64  `json:"id" form:"id" search:"type:exact;column:id;table:gen_source"`
	Sort              int64  `json:"sort" form:"sort" search:"type:exact;column:sort;table:gen_source"`
	Name              string `json:"name" form:"name" search:"type:exact;column:name;table:gen_source"`
	Driver            string `json:"driver" form:"driver" search:"type:exact;column:driver;table:gen_source"`
	Host              string `json:"host" form:"host" search:"type:exact;column:host;table:gen_source"`
	Port              string `json:"port" form:"port" search:"type:exact;column:port;table:gen_source"`
	User              string `json:"user" form:"user" search:"type:exact;column:user;table:gen_source"`
	Password          string `json:"password" form:"password" search:"type:exact;column:password;table:gen_source"`
	Dbname            string `json:"dbname" form:"dbname" search:"type:exact;column:dbname;table:gen_source"`
	Config            string `json:"config" form:"config" search:"type:exact;column:config;table:gen_source"`
	TablePrefix       string `json:"tablePrefix" form:"tablePrefix" search:"type:exact;column:table_prefix;table:gen_source"`
	Remark            string `json:"remark" form:"remark" search:"type:exact;column:remark;table:gen_source"`
	GenSourceOrder
}

type GenSourceOrder struct {
	IdOrder          int64  `json:"id" form:"id" search:"type:order;column:id;table:gen_source"`
	SortOrder        int64  `json:"sort" form:"sort" search:"type:order;column:sort;table:gen_source"`
	NameOrder        string `json:"name" form:"name" search:"type:order;column:name;table:gen_source"`
	DriverOrder      string `json:"driver" form:"driver" search:"type:order;column:driver;table:gen_source"`
	HostOrder        string `json:"host" form:"host" search:"type:order;column:host;table:gen_source"`
	PortOrder        string `json:"port" form:"port" search:"type:order;column:port;table:gen_source"`
	UserOrder        string `json:"user" form:"user" search:"type:order;column:user;table:gen_source"`
	PasswordOrder    string `json:"password" form:"password" search:"type:order;column:password;table:gen_source"`
	DbnameOrder      string `json:"dbname" form:"dbname" search:"type:order;column:dbname;table:gen_source"`
	ConfigOrder      string `json:"config" form:"config" search:"type:order;column:config;table:gen_source"`
	TablePrefixOrder string `json:"tablePrefix" form:"tablePrefix" search:"type:order;column:table_prefix;table:gen_source"`
	RemarkOrder      string `json:"remark" form:"remark" search:"type:order;column:remark;table:gen_source"`
}

func (m *ListGenSourceReq) GetNeedSearch() interface{} {
	return *m
}
