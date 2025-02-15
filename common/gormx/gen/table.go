package gen

import (
	"github.com/google/uuid"
	"github.com/qiaogw/gocode/common/gormx/modelx"
	"gorm.io/datatypes"
)

type Table struct {
	Id           uuid.UUID `json:"id" comment:"主键" gorm:"primaryKey;column:id;size:256;comment:主键;"`
	DatabaseId   string    `json:"databaseId" comment:"数据库" gorm:"column:database_id;size:256;comment:数据库;"`
	Name         string    `json:"name" form:"name" db:"name" gorm:"column:name;size:256;comment:名称;"`
	TableComment string    `json:"tableComment" form:"tableComment" db:"table_comment" gorm:"column:table_comment;size:255;comment:中文名;"`
	CreateBy     string    `json:"createBy" comment:"创建者" gorm:"column:create_by;size:256;comment:创建者;"`
	UpdateBy     string    `json:"updateBy" comment:"更新者" gorm:"column:update_by;size:256;comment:更新者;"`
	Remark       string    `json:"remark" gorm:"column:remark;size:256;comment:备注;"`
	modelx.ModelTime

	AppName       string `json:"appName" comment:"项目名称" gorm:"column:app_name;size:256;comment:项目名称;"`
	Db            string `json:"db" gorm:"column:db;size:256;comment:数据库名称全小写;index:idx_table_unique,unique"`
	Table         string `json:"table" gorm:"column:table;size:256;comment:表名首字母大写驼峰;index:idx_table_unique,unique"`
	TableUrl      string `json:"tableUrl" gorm:"column:table_url;size:256;comment:表全小写;"`
	Package       string `json:"package"  gorm:"column:package;size:256;comment:本包名小写驼峰;"`
	ParentPackage string `json:"parentPackage"  gorm:"column:parent_package;size:256;comment:父项目名全小写;"`
	Service       string `json:"service" gorm:"column:service;size:256;comment:服务名首字母大写驼峰;"`
	Dir           string `json:"dir" gorm:"column:dir;size:256;comment:业务目录;"`
	Author        string `json:"author" gorm:"column:author;size:255;comment:作者;"`
	Email         string `json:"email" gorm:"column:email;size:255;comment:作者邮箱;"`
	HasTimer      bool   `json:"hasTimer" gorm:"column:has_timer;type:bool;comment:存在时间;"`
	HasCacheKey   bool   `json:"hasCacheKey" gorm:"column:has_cache_key;type:bool;comment:存在非主键的唯一键;"`
	NeedValid     bool   `json:"needValid" gorm:"column:need_valid;type:bool;comment:需要验证;"`
	Postgres      bool   `json:"postgres" gorm:"column:postgres;type:bool;comment:postgres;"`
	IsCurd        bool   `json:"isCurd" gorm:"column:is_curd;size:1;comment:开启crud;"`
	IsDataScope   bool   `json:"isDataScope" gorm:"column:is_data_scope;size:1;comment:开启数据权限;"`
	IsAuth        bool   `json:"isAuth" gorm:"column:is_auth;size:1;comment:开启用户认证;"`
	IsImport      bool   `json:"isImport"  gorm:"column:is_import;size:1;comment:是否开启导入;"`
	IsFlow        bool   `json:"isFlow"  gorm:"column:is_flow;size:1;comment:是否流程审批;"`
	PkIsChar      bool   `json:"pkIsChar" form:"postgres" db:"postgres" gorm:"column:pk_is_char;type:bool;comment:主键为uuid;"`

	Columns  []*Column `json:"columns" gorm:"foreignKey:TableId"`
	Database *Database `json:"database"`

	CacheIndexKeys datatypes.JSONSlice[CacheKey] `json:"cacheIndexKeys" gorm:"column:cache_keys;comment:唯一键;"`
	CacheKeys      []CacheKey                    `json:"cacheKeys" gorm:"-"`
}
