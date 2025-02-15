package gen

import (
	"github.com/google/uuid"
	"github.com/qiaogw/gocode/common/gormx/modelx"
)

type Column struct {
	IsPk   bool
	Indexs int
	*DbColumn
	Index *DbIndex
}
type DbColumn struct {
	Id              uuid.UUID   `json:"id" comment:"主键" gorm:"primaryKey;column:id;size:256;comment:主键;"`
	TableId         string      `json:"tableId" comment:"表模块_主键" gorm:"column:table_id;size:256;comment:表模块_主键;"`
	Name            string      `json:"name" form:"name" db:"name" gorm:"column:name;size:256;comment:名称;"`
	GormName        string      `json:"gormName" form:"gormName" db:"gorm_name" gorm:"column:gorm_name;size:255;comment:gorm字段名;"`
	DbType          string      `json:"dbType" form:"dbType" db:"db_type" gorm:"column:db_type;size:256;comment:数据库字段类型;"`
	DataType        string      `json:"dataType" form:"dataType" db:"data_type" gorm:"column:data_type;size:256;comment:go数据类型;"`
	DataTypeLong    string      `json:"dataTypeLong" form:"dataTypeLong" db:"data_type_long" gorm:"column:data_type_long;size:256;comment:字段长度;"`
	DataTypeApi     string      `json:"dataTypeApi" gorm:"column:data_type_api;size:256;comment:api数据类型;"`
	DataTypeProto   string      `json:"dataTypeProto" gorm:"column:data_type_proto;size:256;comment:Proto数据类型;"`
	Extra           string      `json:"extra" form:"extra" db:"extra" gorm:"column:extra;size:255;comment:扩展;"`
	ColumnComment   string      `json:"columnComment" form:"columnComment" db:"column_comment" gorm:"column:column_comment;size:255;comment:中文名;"`
	ColumnDefault   interface{} `json:"columnDefault" form:"columnDefault" db:"column_default" gorm:"column:column_default;size:255;comment:默认数据;"`
	IsNullAble      string      `json:"isNullable" form:"isNullable" db:"is_nullable" gorm:"column:is_nullable;size:256;comment:是否为空BD;"`
	OrdinalPosition int         `json:"ordinalPosition" form:"ordinalPosition" db:"ordinal_position" gorm:"column:ordinal_position;size:4;comment:OrdinalPosition;"`
	FieldJson       string      `json:"fieldJson" form:"fieldJson" db:"field_json" gorm:"column:field_json;size:255;comment:json字段名;"`
	Tablename       string      `json:"tablename" form:"tablename" db:"tablename" gorm:"column:tablename;size:256;comment:表名;"`
	FieldName       string      `json:"fieldName" form:"fieldName" db:"field_name" gorm:"column:field_name;size:256;comment:go字段名;"`
	DictType        string      `json:"dictType" form:"dictType" db:"dict_type" gorm:"column:dict_type;size:256;comment:关联字典;"`
	ErrorText       string      `json:"errorText" form:"errorText" db:"error_text" gorm:"column:error_text;size:256;comment:验证错误;"`
	IsPage          bool        `json:"isPage" form:"isPage" db:"is_page" gorm:"column:is_page;size:1;comment:是否页码字段;"`
	IsControl       bool        `json:"isControl" gorm:"column:is_control;size:1;comment:是否控制字段;"`
	IsModelTime     bool        `json:"isModelTime" gorm:"column:is_model_time;size:1;comment:是否控制时间字段;"`
	IsNull          bool        `json:"isNull" form:"isNull" db:"is_null" gorm:"column:is_null;size:1;comment:是否为空;"`
	Clearable       bool        `json:"clearable" form:"clearable" db:"clearable" gorm:"column:clearable;size:1;comment:是否可清空;"`
	Require         bool        `json:"require" form:"require" db:"require" gorm:"column:require;size:1;comment:是否必填;"`
	IsPk            bool        `json:"isPk" form:"isPk" db:"is_pk" gorm:"column:is_pk;size:1;comment:是否主键;"`
	IsList          bool        `json:"isList" form:"isList" db:"is_list" gorm:"column:is_list;size:1;comment:是否显示;"`
	Increment       bool        `json:"increment" form:"increment" db:"increment" gorm:"column:increment;size:1;comment:是否自增;"`
	IsEdit          bool        `json:"isEdit" form:"isEdit" db:"is_edit" gorm:"column:is_edit;size:1;comment:是否编辑;"`
	IsSort          bool        `json:"isSort" form:"isSort" db:"is_sort" gorm:"column:is_sort;size:1;comment:是否排序;"`
	HtmlType        string      `json:"htmlType" form:"htmlType" db:"html_type" gorm:"column:html_type;size:256;comment:html类型;"`
	FkTableId       string      `json:"fkTableId" form:"fkTableClass" db:"fk_table_id" gorm:"column:fk_table_id;size:256;comment:关联类id;"`
	FkTable         string      `json:"fkTable" form:"fkTable" db:"fk_table" gorm:"column:fk_table;size:256;comment:关联表;"`
	FkTableClass    string      `json:"fkTableClass" form:"fkTableClass" db:"fk_table_class" gorm:"column:fk_table_class;size:256;comment:关联类;"`
	FkTablePackage  string      `json:"fkTablePackage" form:"fkTablePackage" db:"fk_table_package" gorm:"column:fk_table_package;size:256;comment:关联包;"`
	FkLabelId       string      `json:"fkLabelId" form:"fkLabelId" db:"fk_label_id" gorm:"column:fk_label_id;size:256;comment:关联id;"`
	FkLabelName     string      `json:"fkLabelName" form:"fkLabelName" db:"fk_label_name" gorm:"column:fk_label_name;size:256;comment:关联名;"`
	Remark          string      `json:"remark" form:"remark" db:"remark" gorm:"column:remark;size:256;comment:备注;"`
	Sort            int64       `json:"sort" form:"sort" db:"sort" gorm:"column:sort;comment:排序;"`
	Table           *Table      `json:"table"`
	CreateBy        string      `json:"createBy" comment:"创建者" gorm:"column:create_by;size:256;comment:创建者;"`
	UpdateBy        string      `json:"updateBy" comment:"更新者" gorm:"column:update_by;size:256;comment:更新者;"`
	modelx.ModelTime

	FormType        string `json:"formType" form:"formType" gorm:"column:form_type;size:256;comment:表单类型;"`
	FormOptions     string `json:"formOptions" form:"formOptions" gorm:"column:form_options;type:text;comment:表单选项列表;"`
	FormOptionLabel string `json:"formOptionLabel" form:"formOptionLabel" gorm:"column:form_option_label;size:256;comment:表单选项标签;"`
	FormOptionValue string `json:"formOptionValue" form:"formOptionValue" gorm:"column:form_option_value;type:text;comment:表单选项数据;"`
	FormMultiple    bool   `json:"formMultiple" form:"formMultiple" gorm:"column:form_multiple;comment:是否多选;"`
	FormDisable     bool   `json:"formDisable" form:"formDisable" gorm:"column:form_disable;comment:是否禁用;"`
	FormReadonly    bool   `json:"formReadonly" form:"formReadonly" gorm:"column:form_readonly;comment:是否只读;"`
	FormClearable   bool   `json:"formClearable" form:"formClearable" gorm:"column:form_clearable;comment:是否清空;"`
	FormRange       bool   `json:"formRange" form:"formRange" gorm:"column:form_range;comment:是否循环;"`
	FormSize        int64  `json:"formSize" form:"formSize" gorm:"column:form_size;comment:尺寸;"`
	FormColor       string `json:"formColor" form:"formColor" gorm:"column:form_color;size:256;comment:颜色;"`
	FormTextColor   string `json:"formTextColor" form:"formTextColor" gorm:"column:form_text_color;size:256;comment:字体颜色;"`
	FormSrc         string `json:"formSrc" form:"formSrc" gorm:"column:form_src;size:256;comment:文件源;"`
	FormMin         int64  `json:"formMin" form:"formMin" gorm:"column:form_min;comment:最小值;"`
	FormMax         int64  `json:"formMax" form:"formMax" gorm:"column:form_max;comment:最大值;"`
	FormClass       string `json:"formClass" form:"formClass" gorm:"column:form_class;size:2560;comment:样式类型;"`
}
type (
	DbIndex struct {
		IndexName  string `json:"indexName" gorm:"column:index_name"`
		NonUnique  int    `json:"nonUnique" gorm:"column:non_unique"`
		SeqInIndex int    `json:"seqInIndex" gorm:"column:seq_in_index"`
	}
	ColumnData struct {
		Db      string
		Table   string
		Columns []*Column
	}
	IndexType string

	// Index describes a column index
	Index struct {
		IndexType IndexType
		Columns   []*Column
	}
)
