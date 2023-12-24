package adminmodel

import (
	"github.com/qiaogw/gocode/common/modelx"
	"time"
)

type ListOperaLogReq struct {
	modelx.Pagination `search:"-"`
	Id                string    `json:"id" form:"id" search:"type:exact;column:id;table:admin_opera_log"`
	Title             string    `json:"title" form:"title" search:"type:exact;column:title;table:admin_opera_log"`
	BusinessType      string    `json:"businessType" form:"businessType" search:"type:exact;column:business_type;table:admin_opera_log"`
	BusinessTypes     string    `json:"businessTypes" form:"businessTypes" search:"type:exact;column:business_types;table:admin_opera_log"`
	Method            string    `json:"method" form:"method" search:"type:exact;column:method;table:admin_opera_log"`
	RequestMethod     string    `json:"requestMethod" form:"requestMethod" search:"type:exact;column:request_method;table:admin_opera_log"`
	OperatorType      string    `json:"operatorType" form:"operatorType" search:"type:exact;column:operator_type;table:admin_opera_log"`
	OperName          string    `json:"operName" form:"operName" search:"type:exact;column:oper_name;table:admin_opera_log"`
	DeptName          string    `json:"deptName" form:"deptName" search:"type:exact;column:dept_name;table:admin_opera_log"`
	OperUrl           string    `json:"operUrl" form:"operUrl" search:"type:exact;column:oper_url;table:admin_opera_log"`
	OperIp            string    `json:"operIp" form:"operIp" search:"type:exact;column:oper_ip;table:admin_opera_log"`
	OperLocation      string    `json:"operLocation" form:"operLocation" search:"type:exact;column:oper_location;table:admin_opera_log"`
	OperParam         string    `json:"operParam" form:"operParam" search:"type:exact;column:oper_param;table:admin_opera_log"`
	Status            string    `json:"status" form:"status" search:"type:exact;column:status;table:admin_opera_log"`
	OperTime          time.Time `json:"operTime" form:"operTime" search:"type:exact;column:oper_time;table:admin_opera_log"`
	JsonResult        string    `json:"jsonResult" form:"jsonResult" search:"type:exact;column:json_result;table:admin_opera_log"`
	Remark            string    `json:"remark" form:"remark" search:"type:exact;column:remark;table:admin_opera_log"`
	LatencyTime       string    `json:"latencyTime" form:"latencyTime" search:"type:exact;column:latency_time;table:admin_opera_log"`
	UserAgent         string    `json:"userAgent" form:"userAgent" search:"type:exact;column:user_agent;table:admin_opera_log"`
	CreatedAt         time.Time `json:"createdAt" form:"createdAt" search:"type:exact;column:created_at;table:admin_opera_log"`
	UpdatedAt         time.Time `json:"updatedAt" form:"updatedAt" search:"type:exact;column:updated_at;table:admin_opera_log"`
	CreateBy          string    `json:"createBy" form:"createBy" search:"type:exact;column:create_by;table:admin_opera_log"`
	UpdateBy          string    `json:"updateBy" form:"updateBy" search:"type:exact;column:update_by;table:admin_opera_log"`
	OperaLogOrder
}

type OperaLogOrder struct {
	IdOrder            string    `json:"id" form:"id" search:"type:order;column:id;table:admin_opera_log"`
	TitleOrder         string    `json:"title" form:"title" search:"type:order;column:title;table:admin_opera_log"`
	BusinessTypeOrder  string    `json:"businessType" form:"businessType" search:"type:order;column:business_type;table:admin_opera_log"`
	BusinessTypesOrder string    `json:"businessTypes" form:"businessTypes" search:"type:order;column:business_types;table:admin_opera_log"`
	MethodOrder        string    `json:"method" form:"method" search:"type:order;column:method;table:admin_opera_log"`
	RequestMethodOrder string    `json:"requestMethod" form:"requestMethod" search:"type:order;column:request_method;table:admin_opera_log"`
	OperatorTypeOrder  string    `json:"operatorType" form:"operatorType" search:"type:order;column:operator_type;table:admin_opera_log"`
	OperNameOrder      string    `json:"operName" form:"operName" search:"type:order;column:oper_name;table:admin_opera_log"`
	DeptNameOrder      string    `json:"deptName" form:"deptName" search:"type:order;column:dept_name;table:admin_opera_log"`
	OperUrlOrder       string    `json:"operUrl" form:"operUrl" search:"type:order;column:oper_url;table:admin_opera_log"`
	OperIpOrder        string    `json:"operIp" form:"operIp" search:"type:order;column:oper_ip;table:admin_opera_log"`
	OperLocationOrder  string    `json:"operLocation" form:"operLocation" search:"type:order;column:oper_location;table:admin_opera_log"`
	OperParamOrder     string    `json:"operParam" form:"operParam" search:"type:order;column:oper_param;table:admin_opera_log"`
	StatusOrder        string    `json:"status" form:"status" search:"type:order;column:status;table:admin_opera_log"`
	OperTimeOrder      time.Time `json:"operTime" form:"operTime" search:"type:order;column:oper_time;table:admin_opera_log"`
	JsonResultOrder    string    `json:"jsonResult" form:"jsonResult" search:"type:order;column:json_result;table:admin_opera_log"`
	RemarkOrder        string    `json:"remark" form:"remark" search:"type:order;column:remark;table:admin_opera_log"`
	LatencyTimeOrder   string    `json:"latencyTime" form:"latencyTime" search:"type:order;column:latency_time;table:admin_opera_log"`
	UserAgentOrder     string    `json:"userAgent" form:"userAgent" search:"type:order;column:user_agent;table:admin_opera_log"`
	CreatedAtOrder     time.Time `json:"createdAt" form:"createdAt" search:"type:order;column:created_at;table:admin_opera_log"`
	UpdatedAtOrder     time.Time `json:"updatedAt" form:"updatedAt" search:"type:order;column:updated_at;table:admin_opera_log"`
	CreateByOrder      string    `json:"createBy" form:"createBy" search:"type:order;column:create_by;table:admin_opera_log"`
	UpdateByOrder      string    `json:"updateBy" form:"updateBy" search:"type:order;column:update_by;table:admin_opera_log"`
}

func (m *ListOperaLogReq) GetNeedSearch() interface{} {
	return *m
}
