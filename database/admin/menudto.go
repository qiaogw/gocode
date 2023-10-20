package admin

import (
	"github.com/qiaogw/gocode/common/modelx"
)

type ListMenuReq struct {
	modelx.Pagination `search:"-"`
	Id                int64  `json:"id" form:"id" search:"type:exact;column:id;table:admin_menu"`
	Name              string `json:"name" form:"name" search:"type:exact;column:name;table:admin_menu"`
	Title             string `json:"title" form:"title" search:"type:exact;column:title;table:admin_menu"`
	Icon              string `json:"icon" form:"icon" search:"type:exact;column:icon;table:admin_menu"`
	Path              string `json:"path" form:"path" search:"type:exact;column:path;table:admin_menu"`
	Type              string `json:"type" form:"type" search:"type:exact;column:type;table:admin_menu"`
	Component         string `json:"component" form:"component" search:"type:exact;column:component;table:admin_menu"`
	ParentId          int64  `json:"parentId" form:"parentId" search:"type:exact;column:parent_id;table:admin_menu"`
	Sort              int64  `json:"sort" form:"sort" search:"type:exact;column:sort;table:admin_menu"`
	KeepAlive         bool   `json:"keepAlive" form:"keepAlive" search:"type:exact;column:keep_alive;table:admin_menu"`
	Hidden            bool   `json:"hidden" form:"hidden" search:"type:exact;column:hidden;table:admin_menu"`
	IsFrame           bool   `json:"isFrame" form:"isFrame" search:"type:exact;column:is_frame;table:admin_menu"`
	Remark            string `json:"remark" form:"remark" search:"type:exact;column:remark;table:admin_menu"`
	MenuOrder
}

type MenuOrder struct {
	IdOrder        int64  `json:"id" form:"id" search:"type:order;column:id;table:admin_menu"`
	NameOrder      string `json:"name" form:"name" search:"type:order;column:name;table:admin_menu"`
	TitleOrder     string `json:"title" form:"title" search:"type:order;column:title;table:admin_menu"`
	IconOrder      string `json:"icon" form:"icon" search:"type:order;column:icon;table:admin_menu"`
	PathOrder      string `json:"path" form:"path" search:"type:order;column:path;table:admin_menu"`
	TypeOrder      string `json:"type" form:"type" search:"type:order;column:type;table:admin_menu"`
	ComponentOrder string `json:"component" form:"component" search:"type:order;column:component;table:admin_menu"`
	ParentIdOrder  int64  `json:"parentId" form:"parentId" search:"type:order;column:parent_id;table:admin_menu"`
	SortOrder      int64  `json:"sort" form:"sort" search:"type:order;column:sort;table:admin_menu"`
	KeepAliveOrder bool   `json:"keepAlive" form:"keepAlive" search:"type:order;column:keep_alive;table:admin_menu"`
	VisibleOrder   bool   `json:"visible" form:"visible" search:"type:order;column:visible;table:admin_menu"`
	IsFrameOrder   bool   `json:"isFrame" form:"isFrame" search:"type:order;column:is_frame;table:admin_menu"`
	RemarkOrder    string `json:"remark" form:"remark" search:"type:order;column:remark;table:admin_menu"`
}

func (m *ListMenuReq) GetNeedSearch() interface{} {
	return *m
}
