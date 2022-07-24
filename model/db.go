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
	Table struct {
		Db          string
		Table       string `json:"table" gorm:"column:table_name"`
		PackageName string //表首字母小写驼峰
		TableUrl    string //url 表全小写驼峰
		Columns     []*Column
		// Primary key not included
		UniqueIndex  map[string][]*Column
		PrimaryKey   *Column
		NormalIndex  map[string][]*Column
		HasTimer     bool
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
		DataType        string      `json:"dataType" gorm:"column:DATA_TYPE"`
		DataTypeProto   string      `json:"dataTypeProto" gorm:"-"`
		DataTypeLong    string      `json:"dataTypeLong" gorm:"-"`
		Extra           string      `json:"extra" gorm:"column:EXTRA"`
		Comment         string      `json:"comment" gorm:"column:COLUMN_COMMENT"`
		ColumnDefault   interface{} `json:"columnDefault" gorm:"column:COLUMN_DEFAULT"`
		IsNullAble      string      `json:"isNullAble" gorm:"column:IS_NULLABLE"`
		OrdinalPosition int         `json:"ordinalPosition" gorm:"column:ORDINAL_POSITION"`
		FieldJson       string      `json:"fieldJson"`
		FieldName       string      `json:"fieldName"`
		Clearable       bool        `json:"clearable"` // 是否可清空
		DictType        string      `json:"dictType"`  // 字典
		Require         bool        `json:"require"`   // 是否必填
		ErrorText       string      `json:"errorText"` // 校验失败文字
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
	table.Table = util.CamelString(strings.TrimLeft(c.Table, global.GenConfig.DB.TablePrefix))
	table.PackageName = util.LeftLower(table.Table)
	table.TableUrl = strings.ToLower(table.Table)
	//log.Printf("table.Table is %s,table.PackageName is %s\n", table.Table, table.PackageName)
	table.Db = global.GenConfig.System.Name
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

	for i, each := range c.Columns {
		//log.Printf("each.DataType is %s\n", each.DataType)
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
		isDefaultNull := each.ColumnDefault == nil && each.IsNullAble == "YES"

		dt, err := converter.ConvertStringDataType(each.DataType, isDefaultNull)
		if err != nil {
			return nil, err
		}
		each.DataType = dt
		each.DataTypeProto = dt
		if dt == "time.Time" {
			each.DataTypeProto = "string"
		}
		if each.DataType == "time.Time" {
			table.HasTimer = true
		}
		each.FieldName = util.LeftUpper(util.CamelString(each.Name))
		each.FieldJson = util.LeftLower(util.CamelString(each.Name))
		each.Comment = util.TrimNewLine(each.Comment)

		if each.Index != nil {
			m[each.Index.IndexName] = append(m[each.Index.IndexName], each)
		}
		each.Indexs = i + 1
		table.Columns = append(table.Columns, each)
	}

	primaryColumns := m[indexPri]
	//log.Printf("primaryColumns:%+v\n", m)
	if len(primaryColumns) == 0 {
		return nil, fmt.Errorf("db:%s, table:%s, missing primary key", c.Db, c.Table)
	}

	if len(primaryColumns) > 1 {
		return nil, fmt.Errorf("db:%s, table:%s, joint primary key is not supported", c.Db, c.Table)
	}

	table.PrimaryKey = primaryColumns[0]
	for indexName, columns := range m {
		//log.Printf("columns is %+v\n", len(columns))
		if indexName == indexPri {
			continue
		}

		for _, one := range columns {
			//log.Printf("each is %+v\n", columns)
			if one.Index != nil {
				if one.Index.NonUnique == 0 {
					table.UniqueIndex[indexName] = columns
				} else {
					table.NormalIndex[indexName] = columns
				}
			}
		}
	}
	//var baseFileds = []string{"CreatedAt", "UpdatedAt", "DeletedAt", "CreateBy", "UpdateBy"}

	return &table, nil
}
