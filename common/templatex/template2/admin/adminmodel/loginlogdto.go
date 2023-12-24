package adminmodel

import (
	"github.com/qiaogw/gocode/common/modelx"
	"time"
)

type ListLoginLogReq struct {
	modelx.Pagination `search:"-"`
	Id                string    `json:"id" form:"id" search:"type:exact;column:id;table:admin_login_log"`
	Username          string    `json:"username" form:"username" search:"type:exact;column:username;table:admin_login_log"`
	Status            string    `json:"status" form:"status" search:"type:exact;column:status;table:admin_login_log"`
	Ipaddr            string    `json:"ipaddr" form:"ipaddr" search:"type:exact;column:ipaddr;table:admin_login_log"`
	LoginLocation     string    `json:"loginLocation" form:"loginLocation" search:"type:exact;column:login_location;table:admin_login_log"`
	Browser           string    `json:"browser" form:"browser" search:"type:exact;column:browser;table:admin_login_log"`
	Os                string    `json:"os" form:"os" search:"type:exact;column:os;table:admin_login_log"`
	Platform          string    `json:"platform" form:"platform" search:"type:exact;column:platform;table:admin_login_log"`
	LoginTime         time.Time `json:"loginTime" form:"loginTime" search:"type:exact;column:login_time;table:admin_login_log"`
	Remark            string    `json:"remark" form:"remark" search:"type:exact;column:remark;table:admin_login_log"`
	Msg               string    `json:"msg" form:"msg" search:"type:exact;column:msg;table:admin_login_log"`
	CreatedAt         time.Time `json:"createdAt" form:"createdAt" search:"type:exact;column:created_at;table:admin_login_log"`
	UpdatedAt         time.Time `json:"updatedAt" form:"updatedAt" search:"type:exact;column:updated_at;table:admin_login_log"`
	CreateBy          string    `json:"createBy" form:"createBy" search:"type:exact;column:create_by;table:admin_login_log"`
	UpdateBy          string    `json:"updateBy" form:"updateBy" search:"type:exact;column:update_by;table:admin_login_log"`
	LoginLogOrder
}

type LoginLogOrder struct {
	IdOrder            string    `json:"id" form:"id" search:"type:order;column:id;table:admin_login_log"`
	UsernameOrder      string    `json:"username" form:"username" search:"type:order;column:username;table:admin_login_log"`
	StatusOrder        string    `json:"status" form:"status" search:"type:order;column:status;table:admin_login_log"`
	IpaddrOrder        string    `json:"ipaddr" form:"ipaddr" search:"type:order;column:ipaddr;table:admin_login_log"`
	LoginLocationOrder string    `json:"loginLocation" form:"loginLocation" search:"type:order;column:login_location;table:admin_login_log"`
	BrowserOrder       string    `json:"browser" form:"browser" search:"type:order;column:browser;table:admin_login_log"`
	OsOrder            string    `json:"os" form:"os" search:"type:order;column:os;table:admin_login_log"`
	PlatformOrder      string    `json:"platform" form:"platform" search:"type:order;column:platform;table:admin_login_log"`
	LoginTimeOrder     time.Time `json:"loginTime" form:"loginTime" search:"type:order;column:login_time;table:admin_login_log"`
	RemarkOrder        string    `json:"remark" form:"remark" search:"type:order;column:remark;table:admin_login_log"`
	MsgOrder           string    `json:"msg" form:"msg" search:"type:order;column:msg;table:admin_login_log"`
	CreatedAtOrder     time.Time `json:"createdAt" form:"createdAt" search:"type:order;column:created_at;table:admin_login_log"`
	UpdatedAtOrder     time.Time `json:"updatedAt" form:"updatedAt" search:"type:order;column:updated_at;table:admin_login_log"`
	CreateByOrder      string    `json:"createBy" form:"createBy" search:"type:order;column:create_by;table:admin_login_log"`
	UpdateByOrder      string    `json:"updateBy" form:"updateBy" search:"type:order;column:update_by;table:admin_login_log"`
}

func (m *ListLoginLogReq) GetNeedSearch() interface{} {
	return *m
}
