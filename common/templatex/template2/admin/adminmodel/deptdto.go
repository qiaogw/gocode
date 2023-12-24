package adminmodel

import (
	"github.com/qiaogw/gocode/common/modelx"
	"time"
)

type ListDeptReq struct {
	modelx.Pagination `search:"-"`
	Id                string    `json:"id" form:"id" search:"type:exact;column:id;table:admin_dept"`
	ParentId          string    `json:"parentId" form:"parentId" search:"type:exact;column:parent_id;table:admin_dept"`
	Name              string    `json:"name" form:"name" search:"type:exact;column:name;table:admin_dept"`
	Sort              int64     `json:"sort" form:"sort" search:"type:exact;column:sort;table:admin_dept"`
	Leader            string    `json:"leader" form:"leader" search:"type:exact;column:leader;table:admin_dept"`
	Phone             string    `json:"phone" form:"phone" search:"type:exact;column:phone;table:admin_dept"`
	Email             string    `json:"email" form:"email" search:"type:exact;column:email;table:admin_dept"`
	Enabled           bool      `json:"enabled" form:"enabled" search:"type:exact;column:enabled;table:admin_dept"`
	CreateBy          string    `json:"createBy" form:"createBy" search:"type:exact;column:create_by;table:admin_dept"`
	UpdateBy          string    `json:"updateBy" form:"updateBy" search:"type:exact;column:update_by;table:admin_dept"`
	CreatedAt         time.Time `json:"createdAt" form:"createdAt" search:"type:exact;column:created_at;table:admin_dept"`
	UpdatedAt         time.Time `json:"updatedAt" form:"updatedAt" search:"type:exact;column:updated_at;table:admin_dept"`
	DeletedAt         time.Time `json:"deletedAt" form:"deletedAt" search:"type:exact;column:deleted_at;table:admin_dept"`
	DeptOrder
}

type DeptOrder struct {
	IdOrder        string    `json:"id" form:"id" search:"type:order;column:id;table:admin_dept"`
	ParentIdOrder  string    `json:"parentId" form:"parentId" search:"type:order;column:parent_id;table:admin_dept"`
	NameOrder      string    `json:"name" form:"name" search:"type:order;column:name;table:admin_dept"`
	SortOrder      int64     `json:"sort" form:"sort" search:"type:order;column:sort;table:admin_dept"`
	LeaderOrder    string    `json:"leader" form:"leader" search:"type:order;column:leader;table:admin_dept"`
	PhoneOrder     string    `json:"phone" form:"phone" search:"type:order;column:phone;table:admin_dept"`
	EmailOrder     string    `json:"email" form:"email" search:"type:order;column:email;table:admin_dept"`
	EnabledOrder   bool      `json:"enabled" form:"enabled" search:"type:order;column:enabled;table:admin_dept"`
	CreateByOrder  string    `json:"createBy" form:"createBy" search:"type:order;column:create_by;table:admin_dept"`
	UpdateByOrder  string    `json:"updateBy" form:"updateBy" search:"type:order;column:update_by;table:admin_dept"`
	CreatedAtOrder time.Time `json:"createdAt" form:"createdAt" search:"type:order;column:created_at;table:admin_dept"`
	UpdatedAtOrder time.Time `json:"updatedAt" form:"updatedAt" search:"type:order;column:updated_at;table:admin_dept"`
	DeletedAtOrder time.Time `json:"deletedAt" form:"deletedAt" search:"type:order;column:deleted_at;table:admin_dept"`
}

func (m *ListDeptReq) GetNeedSearch() interface{} {
	return *m
}
