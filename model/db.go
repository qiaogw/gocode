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
		HasTimer   bool
		GitUser    string
		GitEmail   string
		Option     *config.APP
		DriverName string
		ParentPkg  string //项目路径
	}
	CacheKey struct {
		Key       string
		Value     string
		Field     string
		FieldJson string
	}
	Table struct {
		Db          string //小写服务名称
		Table       string `json:"table" gorm:"column:table_name"` //表首字母大写驼峰
		Name        string
		PackageName string //表首字母小写驼峰
		TableUrl    string //url 表全小写驼峰
		Columns     []*Column
		// Primary key not included
		UniqueIndex  map[string][]*Column
		PrimaryKey   *Column
		CacheKeys    []*CacheKey
		NormalIndex  map[string][]*Column
		HasTimer     bool
		HasCacheKey  bool //存在非主键的唯一键
		NeedValid    bool
		PostgreSql   bool
		TableComment string `json:"table_comment" gorm:"column:table_comment"`
		GitUser      string
		GitEmail     string
		ParentPkg    string //项目路径
		Service      string //模块首字母大写驼峰
	}
	Column struct {
		*DbColumn
		Index  *DbIndex
		IsPk   bool
		Indexs int
	}
	DbColumn struct {
		Name            string      `json:"name" gorm:"column:COLUMN_NAME"`
		GormName        string      `json:"gormName" gorm:"-"`
		DataType        string      `json:"dataType" gorm:"column:DATA_TYPE"`
		DataTypeLong    string      `json:"data_type_long" gorm:"column:data_type_long"`
		DataTypeProto   string      `json:"dataTypeProto" gorm:"-"`
		IsPage          bool        `json:"isPage" gorm:"-"`
		Extra           string      `json:"extra" gorm:"column:EXTRA"`
		Comment         string      `json:"comment" gorm:"column:COLUMN_COMMENT"`
		ColumnDefault   interface{} `json:"columnDefault" gorm:"column:COLUMN_DEFAULT"`
		IsNullAble      string      `json:"isNullAble" gorm:"column:IS_NULLABLE"`
		IsNull          bool        `json:"isNull" gorm:"-"`
		OrdinalPosition int         `json:"ordinalPosition" gorm:"column:ORDINAL_POSITION"`
		FieldJson       string      `json:"fieldJson"`
		FieldName       string      `json:"fieldName"`
		Clearable       bool        `json:"clearable"` // 是否可清空
		DictType        string      `json:"dictType"`  // 字典
		Require         bool        `json:"require"`   // 是否必填
		ErrorText       string      `json:"errorText"` // 校验失败文字
		TableName       string      `json:"tableName"`
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
	//log.Printf("table.Table is %s,table.PackageName is %s\n", table.Table, table.PackageName)
	//log.Printf("table.Name is %s,table.prefix is %s,table.table is %s\n", table.Name,
	//	strings.TrimPrefix(c.Table, global.GenConfig.DB.TablePrefix), table.Table)
	table.Db = strings.ToLower(util.CamelString(global.GenConfig.System.Name))
	table.Service = util.LeftUpper(table.Db)
	table.TableComment = tableComment
	//table.Columns = c.Columns
	table.UniqueIndex = map[string][]*Column{}
	table.NormalIndex = map[string][]*Column{}

	if global.GenDB.Name() == "postgres" {
		table.PostgreSql = true
	}
	table.GitUser = getGitName()
	table.GitEmail = getGitEmail()
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
			continue
		}
		if each.Name == "updated_at" {
			continue
		}
		if each.Name == "deleted_at" {
			continue
		}
		if each.Name == "create_by" {
			continue
		}
		if each.Name == "update_by" {
			continue
		}
		var isDefaultNull bool
		//isDefaultNull = each.ColumnDefault == nil && each.IsNullAble == "YES"

		dt, err := converter.ConvertStringDataType(each.DataType, isDefaultNull)
		if err != nil {
			return nil, fmt.Errorf("表： %s, 字段： %s 错误： %v", c.Table, each.Name, err)
		}
		each.DataType = dt
		each.DataTypeProto = dt
		each.IsNull = each.IsNullAble == "YES"
		if dt == "time.Time" {
			each.DataTypeProto = "string"
			table.HasTimer = true
		}
		if !each.IsPk && !each.IsNull && each.ColumnDefault != nil {
			each.ColumnDefault = converter.ConvertDefault(each.ColumnDefault)
		}
		if each.GormName == "" {
			each.GormName = each.Name
		}
		each.FieldName = util.LeftUpper(util.CamelString(each.Name))
		each.FieldJson = util.LeftLower(util.CamelString(each.Name))
		each.Comment = util.TrimNewLine(each.Comment)
		each.TableName = c.Table
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
	//log.Printf("primaryColumns:%+v\n", m)
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
	//var baseFileds = []string{"CreatedAt", "UpdatedAt", "DeletedAt", "CreateBy", "UpdateBy"}

	return &table, nil
}
