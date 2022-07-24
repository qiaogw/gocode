package model

import (
	"fmt"
	"github.com/qiaogw/gocode/global"
	"gorm.io/gorm"
	"log"
	"sort"
)

var ModelMysqlApp = new(ModelMysql)

type ModelMysql struct {
	DB *gorm.DB
}

func (m *ModelMysql) Init() {
	m.DB = global.GenDB
}

// GetDB 获取数据库的所有数据库名
func (m *ModelMysql) GetDB() (data []Db, err error) {
	var entities []Db
	sql := "SELECT SCHEMA_NAME AS `database` FROM INFORMATION_SCHEMA.SCHEMATA;"
	err = global.GenDB.Raw(sql).Scan(&entities).Error
	return entities, err
}

// GetTables 获取数据库的所有表名
func (m *ModelMysql) GetTables(db string) ([]Table, error) {
	var entities []Table
	sql := `
		select table_name as table_name ,
		table_comment
		from information_schema.tables 
		where table_schema = ?`
	err := global.GenDB.Raw(sql, db).Scan(&entities).Error
	return entities, err
}

// GetColumn 获取指定数据库和指定数据表的所有字段名,类型值等
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (m *ModelMysql) GetColumn(db, table string) (*ColumnData, error) {
	var reply []*DbColumn
	sql := `
	SELECT c.COLUMN_NAME,
		c.DATA_TYPE,
		c.EXTRA,
		c.COLUMN_COMMENT,
		c.COLUMN_DEFAULT,
		c.IS_NULLABLE,
		c.ORDINAL_POSITION ,
       CASE DATA_TYPE
           WHEN 'longtext' THEN c.CHARACTER_MAXIMUM_LENGTH
           WHEN 'varchar' THEN c.CHARACTER_MAXIMUM_LENGTH
           WHEN 'double' THEN CONCAT_WS(',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE)
           WHEN 'decimal' THEN CONCAT_WS(',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE)
           WHEN 'int' THEN c.NUMERIC_PRECISION
           WHEN 'bigint' THEN c.NUMERIC_PRECISION
           ELSE '' END AS data_type_long
		from INFORMATION_SCHEMA.COLUMNS c 
		WHERE c.TABLE_SCHEMA = ? 
		and c.TABLE_NAME = ? `
	err := m.DB.Raw(sql, db, table).Scan(&reply).Error
	if err != nil {
		log.Printf("getclumn err is %v\n", err)
		return nil, err
	}
	var list []*Column
	for _, item := range reply {
		index, err := m.FindIndex(db, table, item.Name)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return nil, err
			}
			continue
		}

		if len(index) > 0 {
			for _, i := range index {
				list = append(list, &Column{
					DbColumn: item,
					Index:    i,
				})
			}
		} else {
			list = append(list, &Column{
				DbColumn: item,
			})
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].OrdinalPosition < list[j].OrdinalPosition
	})

	var columnData ColumnData
	columnData.Db = db
	columnData.Table = table
	columnData.Columns = list
	return &columnData, nil
}

// FindIndex 获取索引
func (m *ModelMysql) FindIndex(db, table, column string) ([]*DbIndex, error) {
	querySql := `SELECT 
		m.INDEX_NAME as index_name,
		m.NON_UNIQUE as non_unique,
		m.SEQ_IN_INDEX  as seq_in_index
		from  INFORMATION_SCHEMA.STATISTICS m  
		WHERE  m.TABLE_SCHEMA = ? 
		and m.TABLE_NAME = ? 
		and m.COLUMN_NAME = ?`
	var reply []*DbIndex
	err := m.DB.Raw(querySql, db, table, column).Scan(&reply).Error
	if err != nil {
		fmt.Printf("getindex err is %v\n", err)
		return nil, err
	}

	return reply, nil
}
