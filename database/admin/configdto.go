package admin

import (
	"github.com/qiaogw/gocode/common/modelx"
)

type ListConfigReq struct {
	modelx.Pagination `search:"-"`
	Id                int64  `json:"id" form:"id" search:"type:exact;column:id;table:admin_config"`
	ConfigName        string `json:"configName" form:"configName" search:"type:exact;column:config_name;table:admin_config"`
	ConfigKey         string `json:"configKey" form:"configKey" search:"type:exact;column:config_key;table:admin_config"`
	ConfigValue       string `json:"configValue" form:"configValue" search:"type:exact;column:config_value;table:admin_config"`
	ConfigType        string `json:"configType" form:"configType" search:"type:exact;column:config_type;table:admin_config"`
	IsFrontend        string `json:"isFrontend" form:"isFrontend" search:"type:exact;column:is_frontend;table:admin_config"`
	Remark            string `json:"remark" form:"remark" search:"type:exact;column:remark;table:admin_config"`
	ConfigOrder
}

type ConfigOrder struct {
	IdOrder          int64  `json:"id" form:"id" search:"type:order;column:id;table:admin_config"`
	ConfigNameOrder  string `json:"configName" form:"configName" search:"type:order;column:config_name;table:admin_config"`
	ConfigKeyOrder   string `json:"configKey" form:"configKey" search:"type:order;column:config_key;table:admin_config"`
	ConfigValueOrder string `json:"configValue" form:"configValue" search:"type:order;column:config_value;table:admin_config"`
	ConfigTypeOrder  string `json:"configType" form:"configType" search:"type:order;column:config_type;table:admin_config"`
	IsFrontendOrder  string `json:"isFrontend" form:"isFrontend" search:"type:order;column:is_frontend;table:admin_config"`
	RemarkOrder      string `json:"remark" form:"remark" search:"type:order;column:remark;table:admin_config"`
}

func (m *ListConfigReq) GetNeedSearch() interface{} {
	return *m
}
