package admin

import (
	"github.com/qiaogw/gocode/common/modelx"
)

type ListMigrationReq struct {
	modelx.Pagination `search:"-"`
	Id                int64  `json:"id" form:"id" search:"type:exact;column:id;table:admin_migration"`
	Version           string `json:"version" form:"version" search:"type:exact;column:version;table:admin_migration"`
	MigrationOrder
}

type MigrationOrder struct {
	IdOrder      int64  `json:"id" form:"id" search:"type:order;column:id;table:admin_migration"`
	VersionOrder string `json:"version" form:"version" search:"type:order;column:version;table:admin_migration"`
}

func (m *ListMigrationReq) GetNeedSearch() interface{} {
	return *m
}
