package adminmodel

import (
	"github.com/qiaogw/gocode/common/modelx"
	"time"
)

type ListAdminMenuReq struct {
	modelx.Pagination `search:"-"`
	Id                string    `json:"id" form:"id" search:"type:exact;column:id;table:admin_menu"`
	Name              string    `json:"name" form:"name" search:"type:exact;column:name;table:admin_menu"`
	Title             string    `json:"title" form:"title" search:"type:exact;column:title;table:admin_menu"`
	Icon              string    `json:"icon" form:"icon" search:"type:exact;column:icon;table:admin_menu"`
	Path              string    `json:"path" form:"path" search:"type:exact;column:path;table:admin_menu"`
	Type              string    `json:"type" form:"type" search:"type:exact;column:type;table:admin_menu"`
	Component         string    `json:"component" form:"component" search:"type:exact;column:component;table:admin_menu"`
	ParentId          string    `json:"parentId" form:"parentId" search:"type:exact;column:parent_id;table:admin_menu"`
	Sort              int64     `json:"sort" form:"sort" search:"type:exact;column:sort;table:admin_menu"`
	DefaultMenu       bool      `json:"defaultMenu" form:"defaultMenu" search:"type:exact;column:default_menu;table:admin_menu"`
	CloseTab          bool      `json:"closeTab" form:"closeTab" search:"type:exact;column:close_tab;table:admin_menu"`
	KeepAlive         bool      `json:"keepAlive" form:"keepAlive" search:"type:exact;column:keep_alive;table:admin_menu"`
	Hidden            bool      `json:"hidden" form:"hidden" search:"type:exact;column:hidden;table:admin_menu"`
	IsFrame           bool      `json:"isFrame" form:"isFrame" search:"type:exact;column:is_frame;table:admin_menu"`
	Remark            string    `json:"remark" form:"remark" search:"type:exact;column:remark;table:admin_menu"`
	CreateBy          string    `json:"createBy" form:"createBy" search:"type:exact;column:create_by;table:admin_menu"`
	UpdateBy          string    `json:"updateBy" form:"updateBy" search:"type:exact;column:update_by;table:admin_menu"`
	CreatedAt         time.Time `json:"createdAt" form:"createdAt" search:"type:exact;column:created_at;table:admin_menu"`
	UpdatedAt         time.Time `json:"updatedAt" form:"updatedAt" search:"type:exact;column:updated_at;table:admin_menu"`
	DeletedAt         time.Time `json:"deletedAt" form:"deletedAt" search:"type:exact;column:deleted_at;table:admin_menu"`
	MenuOrder
}

type AdminMenuOrder struct {
	IdOrder          string    `json:"id" form:"id" search:"type:order;column:id;table:admin_menu"`
	NameOrder        string    `json:"name" form:"name" search:"type:order;column:name;table:admin_menu"`
	TitleOrder       string    `json:"title" form:"title" search:"type:order;column:title;table:admin_menu"`
	IconOrder        string    `json:"icon" form:"icon" search:"type:order;column:icon;table:admin_menu"`
	PathOrder        string    `json:"path" form:"path" search:"type:order;column:path;table:admin_menu"`
	TypeOrder        string    `json:"type" form:"type" search:"type:order;column:type;table:admin_menu"`
	ComponentOrder   string    `json:"component" form:"component" search:"type:order;column:component;table:admin_menu"`
	ParentIdOrder    string    `json:"parentId" form:"parentId" search:"type:order;column:parent_id;table:admin_menu"`
	SortOrder        int64     `json:"sort" form:"sort" search:"type:order;column:sort;table:admin_menu"`
	DefaultMenuOrder bool      `json:"defaultMenu" form:"defaultMenu" search:"type:order;column:default_menu;table:admin_menu"`
	CloseTabOrder    bool      `json:"closeTab" form:"closeTab" search:"type:order;column:close_tab;table:admin_menu"`
	KeepAliveOrder   bool      `json:"keepAlive" form:"keepAlive" search:"type:order;column:keep_alive;table:admin_menu"`
	HiddenOrder      bool      `json:"hidden" form:"hidden" search:"type:order;column:hidden;table:admin_menu"`
	IsFrameOrder     bool      `json:"isFrame" form:"isFrame" search:"type:order;column:is_frame;table:admin_menu"`
	RemarkOrder      string    `json:"remark" form:"remark" search:"type:order;column:remark;table:admin_menu"`
	CreateByOrder    string    `json:"createBy" form:"createBy" search:"type:order;column:create_by;table:admin_menu"`
	UpdateByOrder    string    `json:"updateBy" form:"updateBy" search:"type:order;column:update_by;table:admin_menu"`
	CreatedAtOrder   time.Time `json:"createdAt" form:"createdAt" search:"type:order;column:created_at;table:admin_menu"`
	UpdatedAtOrder   time.Time `json:"updatedAt" form:"updatedAt" search:"type:order;column:updated_at;table:admin_menu"`
	DeletedAtOrder   time.Time `json:"deletedAt" form:"deletedAt" search:"type:order;column:deleted_at;table:admin_menu"`
}

func (m *ListMenuReq) GetNeedSearch() interface{} {
	return *m
}
