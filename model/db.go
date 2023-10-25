package model

import (
	"fmt"
	"github.com/qiaogw/gocode/config"
	"github.com/qiaogw/gocode/converter"
	"github.com/qiaogw/gocode/global"
	"github.com/qiaogw/gocode/util"
	"strings"
)

type (
	Db struct {
		Database   string `json:"database" gorm:"column:database"`
		Tables     []*Table
		Package    string //首字母小写驼峰
		Service    string //首字母大写驼峰
		FileName   string //全小写，web生成
		HasTimer   bool
		Author     string
		Email      string
		Option     *config.APP
		DriverName string
		ParentPkg  string //项目路径
		PKG        string //根目录
	}
	CacheKey struct {
		Key       string
		Value     string
		Field     string
		FieldJson string
	}
	Table struct {
		Db           string //小写服务名称
		Table        string `json:"table" gorm:"column:table_name"` //表首字母大写驼峰
		Name         string //表名
		PackageName  string //表首字母小写驼峰
		TableUrl     string //url 表全小写
		HasTimer     bool   //存在时间
		HasCacheKey  bool   //存在非主键的唯一键
		NeedValid    bool
		PostgreSql   bool
		TableComment string `json:"table_comment" gorm:"column:table_comment"`
		Author       string
		Email        string
		ParentPkg    string //项目路径
		PKG          string //根目录
		Service      string //服务名，全部小写
		IsCurd       bool
		IsAuth       bool
		IsImport     bool
		IsDataScope  bool
		Columns      []*Column
		// Primary key not included
		UniqueIndex map[string][]*Column
		PrimaryKey  *Column
		CacheKeys   []*CacheKey
		NormalIndex map[string][]*Column
	}
	Column struct {
		IsPk   bool
		Indexs int
		*DbColumn
		Index *DbIndex
	}
	DbColumn struct {
		Name            string      `json:"name" gorm:"column:COLUMN_NAME"`
		GormName        string      `json:"gormName" gorm:"-"`
		DbType          string      `json:"dbType" gorm:"column:db_TYPE"`
		DataType        string      `json:"dataType" gorm:"column:DATA_TYPE"`
		DataTypeLong    string      `json:"data_type_long" gorm:"column:data_type_long"`
		DataTypeProto   string      `json:"dataTypeProto" gorm:"-"`
		DataTypeApi     string      `json:"dataTypeApi" `
		Extra           string      `json:"extra" gorm:"column:EXTRA"`
		ColumnComment   string      `json:"comment" gorm:"column:COLUMN_COMMENT"`
		ColumnDefault   interface{} `json:"columnDefault" gorm:"column:COLUMN_DEFAULT"`
		IsNullAble      string      `json:"isNullAble" gorm:"column:IS_NULLABLE"`
		OrdinalPosition int         `json:"ordinalPosition" gorm:"column:ORDINAL_POSITION"`
		FieldJson       string      `json:"fieldJson"`
		FieldName       string      `json:"fieldName"`
		DictType        string      `json:"dictType"`  // 字典
		ErrorText       string      `json:"errorText"` // 校验失败文字
		Tablename       string      `json:"tablename"`
		IsPage          bool        `json:"isPage" gorm:"-"`
		IsControl       bool        `json:"isControl" gorm:"-"`
		IsModelTime     bool        `json:"isModelTime" gorm:"-"`
		IsNull          bool        `json:"isNull" gorm:"-"`
		Clearable       bool        `json:"clearable"` // 是否可清空
		Require         bool        `json:"require"`   // 是否必填
		IsPk            bool        `json:"is_pk"`
		IsList          bool        `json:"isList" form:"isList" db:"is_list" gorm:"column:is_list;size:1;comment:是否显示;"`
		Increment       bool        `json:"increment" form:"increment" db:"increment" gorm:"column:increment;size:1;comment:是否自增;"`
		IsEdit          bool        `json:"isEdit" form:"isEdit" db:"is_edit" gorm:"column:is_edit;size:1;comment:是否编辑;"`
		IsSort          bool        `json:"isSort" form:"isSort" db:"is_sort" gorm:"column:is_sort;size:1;comment:是否排序;"`
		HtmlType        string      `json:"htmlType" form:"htmlType" db:"html_type" gorm:"column:html_type;size:255;comment:html类型;"`
		FkTable         string      `json:"fkTable" form:"fkTable" db:"fk_table" gorm:"column:fk_table;size:256;comment:关联表;"`
		FkTableClass    string      `json:"fkTableClass" form:"fkTableClass" db:"fk_table_class" gorm:"column:fk_table_class;size:256;comment:关联类;"`
		FkTablePackage  string      `json:"fkTablePackage" form:"fkTablePackage" db:"fk_table_package" gorm:"column:fk_table_package;size:256;comment:关联包;"`
		FkLabelId       string      `json:"fkLabelId" form:"fkLabelId" db:"fk_label_id" gorm:"column:fk_label_id;size:256;comment:关联id;"`
		FkLabelName     string      `json:"fkLabelName" form:"fkLabelName" db:"fk_label_name" gorm:"column:fk_label_name;size:256;comment:关联名;"`
	}

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

// Convert converts column data into Table
func (c *ColumnData) Convert(tableComment string) (*Table, error) {
	var table Table
	table.Name = c.Table
	table.Table = util.CamelString(strings.TrimPrefix(c.Table, global.GenConfig.DB.TablePrefix))
	table.PackageName = util.LeftLower(table.Table)
	table.TableUrl = strings.ToLower(table.Table)

	table.Db = strings.ToLower(util.CamelString(global.GenConfig.System.Name))
	table.Service = strings.ToLower(table.Db)
	table.TableComment = tableComment
	//table.Columns = c.Columns
	table.UniqueIndex = map[string][]*Column{}
	table.NormalIndex = map[string][]*Column{}

	if global.GenDB.Name() == "postgres" {
		table.PostgreSql = true
	}
	table.Author = getGitName()
	table.Email = getGitEmail()
	m := make(map[string][]*Column)
	var pageIndex, pageSize DbColumn

	pageIndex.Name = "PageIndex"
	pageIndex.DataType = "int"
	pageIndex.GormName = "-"
	pageIndex.IsPage = true
	pageSize.Name = "PageSize"
	pageSize.DataType = "int"
	pageSize.GormName = "-"
	pageSize.IsPage = true
	c.Columns = append(c.Columns, &Column{DbColumn: &pageIndex}, &Column{DbColumn: &pageSize})
	ct := 0
	for _, each := range c.Columns {
		//log.Printf("each.name is %s,is pk is %+v\n", each.Name, each.IsPk)
		if each.Index != nil {
			if each.Index.IndexName == "PRIMARY" {
				each.IsPk = true
			}
		}
		if each.Name == "created_at" {
			each.IsModelTime = true
		}
		if each.Name == "updated_at" {
			each.IsModelTime = true
		}
		if each.Name == "deleted_at" {
			each.IsModelTime = true
		}
		if each.Name == "create_by" {
			each.IsControl = true
		}
		if each.Name == "update_by" {
			each.IsControl = true
		}
		var isDefaultNull bool
		each.DbType = each.DataType
		dt, err := converter.ConvertStringDataType(each.DataType, isDefaultNull)
		if err != nil {
			return nil, fmt.Errorf("表： %s, 字段： %s 错误： %v", c.Table, each.Name, err)
		}
		if dt == "int64" || dt == "float64" {
			each.HtmlType = "number"
		}
		each.DataType = dt
		each.DataTypeProto = dt
		each.DataTypeApi = dt
		if dt == "float64" {
			each.DataTypeProto = "double"
		}
		each.IsNull = each.IsNullAble == "YES"
		each.Require = !each.IsNull
		if dt == "time.Time" {
			each.DataTypeProto = "string"
			each.DataTypeApi = "string"
			if !each.IsModelTime {
				table.HasTimer = true
			}
		}
		if !each.IsPk && !each.IsNull && each.ColumnDefault != nil {
			each.ColumnDefault = converter.ConvertDefault(each.ColumnDefault)
		}
		if each.GormName == "" {
			each.GormName = each.Name
		}
		each.FieldName = util.LeftUpper(util.CamelString(each.Name))
		each.FieldJson = util.LeftLower(util.CamelString(each.Name))
		each.ColumnComment = util.TrimNewLine(each.ColumnComment)
		each.Tablename = c.Table
		if each.Index != nil {
			//log.Printf("each.Index is %+v\n", each.Index)
			m[each.Index.IndexName] = append(m[each.Index.IndexName], each)
		}
		ct++
		each.Indexs = ct
		//log.Printf("table is %s,FieldName is %s\n", c.Table, each.FieldName)
		table.Columns = append(table.Columns, each)
	}

	primaryColumns := m[indexPri]
	if len(primaryColumns) == 0 {
		return nil, fmt.Errorf("db:%s, table:%s, 缺失主键", c.Db, c.Table)
	}

	if len(primaryColumns) > 1 {
		return nil, fmt.Errorf("db:%s, table:%s, 程序不支持联合主键", c.Db, c.Table)
	}
	if primaryColumns[0].Name != "id" {
		return nil, fmt.Errorf("gocode 要求表主键唯一，且主键名称为\"id\",表%s 主键为%s,请更新表！", c.Table, primaryColumns[0].Name)
	}

	table.PrimaryKey = primaryColumns[0]
	for indexName, columns := range m {

		if indexName == indexPri {
			continue
		}

		for _, one := range columns {
			//log.Printf("table is %s ,columns is %+v,pk is %v\n", table.Table, one.Name, table.PrimaryKey.Name)
			if one.Index != nil {
				if one.Index.NonUnique == 0 && one.Name != table.PrimaryKey.Name {
					table.UniqueIndex[indexName] = columns
					ck := new(CacheKey)
					ck.Key = "cache" + table.Service + table.Table + one.FieldName
					ck.Value = "cache:" + table.Service + ":" + table.PackageName + ":" + one.FieldJson + ":"
					ck.Field = one.FieldName
					ck.FieldJson = one.FieldJson
					//log.Printf("%s = %s\n", k, v)
					table.CacheKeys = append(table.CacheKeys, ck)
				} else {
					table.NormalIndex[indexName] = columns
				}
			}
		}
	}
	if len(table.CacheKeys) > 0 {
		table.HasCacheKey = true
	}
	return &table, nil
}
