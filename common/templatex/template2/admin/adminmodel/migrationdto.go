package adminmodel

import (
	"github.com/qiaogw/gocode/common/modelx"
	"time"
)

type ListMigrationReq struct {
	modelx.Pagination `search:"-"`
	Id                string    `json:"id" form:"id" search:"type:exact;column:id;table:admin_migration"`
	Version           string    `json:"version" form:"version" search:"type:exact;column:version;table:admin_migration"`
	CreateBy          string    `json:"createBy" form:"createBy" search:"type:exact;column:create_by;table:admin_migration"`
	UpdateBy          string    `json:"updateBy" form:"updateBy" search:"type:exact;column:update_by;table:admin_migration"`
	CreatedAt         time.Time `json:"createdAt" form:"createdAt" search:"type:exact;column:created_at;table:admin_migration"`
	UpdatedAt         time.Time `json:"updatedAt" form:"updatedAt" search:"type:exact;column:updated_at;table:admin_migration"`
	DeletedAt         time.Time `json:"deletedAt" form:"deletedAt" search:"type:exact;column:deleted_at;table:admin_migration"`
	MigrationOrder
}

type MigrationOrder struct {
	IdOrder        string    `json:"id" form:"id" search:"type:order;column:id;table:admin_migration"`
	VersionOrder   string    `json:"version" form:"version" search:"type:order;column:version;table:admin_migration"`
	CreateByOrder  string    `json:"createBy" form:"createBy" search:"type:order;column:create_by;table:admin_migration"`
	UpdateByOrder  string    `json:"updateBy" form:"updateBy" search:"type:order;column:update_by;table:admin_migration"`
	CreatedAtOrder time.Time `json:"createdAt" form:"createdAt" search:"type:order;column:created_at;table:admin_migration"`
	UpdatedAtOrder time.Time `json:"updatedAt" form:"updatedAt" search:"type:order;column:updated_at;table:admin_migration"`
	DeletedAtOrder time.Time `json:"deletedAt" form:"deletedAt" search:"type:order;column:deleted_at;table:admin_migration"`
}

func (m *ListMigrationReq) GetNeedSearch() interface{} {
	return *m
}
