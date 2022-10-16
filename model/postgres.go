package model

import (
	"database/sql"
	"github.com/qiaogw/gocode/global"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

var ModelPostgresApp = new(ModelPostgres)

type ModelPostgres struct {
	DB *gorm.DB
}

func (m *ModelPostgres) Init() {
	m.DB = global.GenDB
}

var p2m = map[string]string{
	"int8":        "bigint",
	"numeric":     "bigint",
	"float8":      "double",
	"float4":      "float",
	"int2":        "smallint",
	"int4":        "integer",
	"timestamptz": "timestamp",
}

type PostgreColumn struct {
	Num               int32  `json:"num" gorm:"column:num"`
	Field             string `json:"field" gorm:"column:field" `
	Type              string `json:"type" gorm:"column:type" `
	DataTypeLong      int    `json:"data_type_long" gorm:"data_type_long"`
	NotNull           bool   `json:"not_null" gorm:"not_null"`
	Comment           string `json:"comment" gorm:"comment"`
	ColumnDefault     string `json:"column_default" gorm:"column_default"`
	IdentityIncrement int32  `json:"identity_increment" gorm:"identity_increment"`
	IsPk              bool   `json:"is_pk" gorm:"is_pk"`
	Extra             string `json:"extra" gorm:"extra"`
}

// PostgreIndex describes an index for a column
type PostgreIndex struct {
	IndexName  sql.NullString `db:"index_name"`
	IndexId    sql.NullInt32  `db:"index_id"`
	IsUnique   sql.NullBool   `db:"is_unique"`
	IsPrimary  sql.NullBool   `db:"is_primary"`
	ColumnName sql.NullString `db:"column_name"`
	IndexSort  sql.NullInt32  `db:"index_sort"`
}

// GetDB 获取数据库的所有数据库名
func (m *ModelPostgres) GetDB() (data []Db, err error) {
	var entities []Db
	sql := `SELECT datname as database FROM pg_database WHERE datistemplate = false`
	err = global.GenDB.Raw(sql).Scan(&entities).Error
	return entities, err
}

type tb struct {
	TableName    string `json:"table_name" gorm:"table_name"`
	TableComment string `json:"table_comment" gorm:"table_comment"`
}

// GetTables 获取数据库的所有表名
func (m *ModelPostgres) GetTables(db string) ([]Table, error) {
	var entities []Table
	var tables []tb
	sql := `
		SELECT
			relname AS table_name,
			cast( obj_description ( relfilenode, 'pg_class' ) AS VARCHAR ) AS table_comment 
		FROM
			pg_class c 
		WHERE
			relkind = 'r' 
			AND relname NOT LIKE 'pg_%' 
			AND relname NOT LIKE 'sql_%' 
		ORDER BY
			relname`
	err := global.GenDB.Raw(sql).Scan(&tables).Error
	for _, v := range tables {
		entities = append(entities, Table{
			TableComment: v.TableComment,
			Table:        v.TableName,
		})
	}

	return entities, err
}

// GetColumn 获取指定数据库和指定数据表的所有字段名,类型值等
// Author [qiaogw](https://github.com/qiaogw)
func (m *ModelPostgres) GetColumn(db, table string) (*ColumnData, error) {
	querySql := `
	select
	c.relname as "table_name",
	a.attname as field,
	(case
		when a.attnotnull = true then true
		else false end) as not_null,
	(case
		when (
		select
			count(pg_constraint.*)
		from
			pg_constraint
		inner join pg_class on
			pg_constraint.conrelid = pg_class.oid
		inner join pg_attribute on
			pg_attribute.attrelid = pg_class.oid
			and pg_attribute.attnum = any(pg_constraint.conkey)
		inner join pg_type on
			pg_type.oid = pg_attribute.atttypid
		where
			pg_class.relname = c.relname
			and pg_constraint.contype = 'p'
			and pg_attribute.attname = a.attname) > 0 then true
		else false end) as is_pk,
	concat_ws('', t.typname) as type,
	(case
		when a.attlen > 0 then a.attlen
		when t.typname='bit' then a.atttypmod
		else a.atttypmod - 4 end) as data_type_long,
	 col.is_identity	as extra,
	 col.column_default	as column_default,
	(select description from pg_description where objoid = a.attrelid
	and objsubid = a.attnum) as comment
from
	pg_class c,
	pg_attribute a ,
	pg_type t,
	information_schema.columns as col
where
	c.relname = $1
	and a.attnum>0
	and a.attrelid = c.oid
	and a.atttypid = t.oid
	and col.table_name=c.relname and col.column_name=a.attname
order by
	c.relname desc,
	a.attnum asc
`

	var reply []*PostgreColumn
	err := m.DB.Raw(querySql, table).Scan(&reply).Error
	if err != nil {
		return nil, err
	}
	schame := "public"

	list, err := m.getColumns(schame, table, reply)
	if err != nil {
		return nil, err
	}

	var columnData ColumnData
	columnData.Db = db
	columnData.Table = table
	columnData.Columns = list
	return &columnData, nil
}

func (m *ModelPostgres) getColumns(schema, table string, in []*PostgreColumn) ([]*Column, error) {
	index, err := m.getIndex(schema, table)
	if err != nil {
		return nil, err
	}
	var list []*Column
	for _, e := range in {
		//log.Printf("column is %+v\n", e)
		var dft interface{}
		if len(e.ColumnDefault) > 0 {
			dft = e.ColumnDefault
		}

		isNullAble := "YES"
		if e.NotNull {
			isNullAble = "NO"
		}
		var extra string
		// when identity is true, the column is auto increment
		if e.IdentityIncrement == 1 {
			extra = "auto_increment"
		}
		// when type is serial, it'm auto_increment. and the default value is tablename_columnname_seq
		if strings.Contains(e.ColumnDefault, table+"_"+e.Field+"_seq") {
			extra = "auto_increment"
		}

		if len(index[e.Field]) > 0 {
			for _, i := range index[e.Field] {
				list = append(list, &Column{
					DbColumn: &DbColumn{
						Name:            e.Field,
						DataType:        m.convertPostgreSqlTypeIntoMysqlType(e.Type),
						Extra:           extra,
						Comment:         e.Comment,
						ColumnDefault:   dft,
						IsNullAble:      isNullAble,
						DataTypeLong:    strconv.Itoa(e.DataTypeLong),
						OrdinalPosition: int(e.Num),
					},
					Index: i,
					IsPk:  e.IsPk,
				})
			}
		} else {
			list = append(list, &Column{
				DbColumn: &DbColumn{
					Name:            e.Field,
					DataType:        m.convertPostgreSqlTypeIntoMysqlType(e.Type),
					Extra:           extra,
					Comment:         e.Comment,
					ColumnDefault:   dft,
					IsNullAble:      isNullAble,
					OrdinalPosition: int(e.Num),
					DataTypeLong:    strconv.Itoa(e.DataTypeLong),
				},
				IsPk: e.IsPk,
			})
		}
	}
	//for _, o := range list {
	//	log.Printf("cl is %+v\n", o.DbColumn)
	//}
	return list, nil
}

func (m *ModelPostgres) convertPostgreSqlTypeIntoMysqlType(in string) string {
	r, ok := p2m[strings.ToLower(in)]
	if ok {
		return r
	}

	return in
}

func (m *ModelPostgres) getIndex(schema, table string) (map[string][]*DbIndex, error) {
	indexes, err := m.FindIndex(schema, table)
	if err != nil {
		return nil, err
	}

	index := make(map[string][]*DbIndex)
	var pkName string
	for _, e := range indexes {
		if e.IsPrimary.Bool {
			pkName = e.ColumnName.String
			index[e.ColumnName.String] = append(index[e.ColumnName.String], &DbIndex{
				IndexName:  indexPri,
				SeqInIndex: int(e.IndexSort.Int32),
			})
			continue
		}
		if e.ColumnName.String == pkName {
			continue
		}
		nonUnique := 0
		if !e.IsUnique.Bool {
			nonUnique = 1
		}

		index[e.ColumnName.String] = append(index[e.ColumnName.String], &DbIndex{
			IndexName:  e.IndexName.String,
			NonUnique:  nonUnique,
			SeqInIndex: int(e.IndexSort.Int32),
		})
	}

	return index, nil
}

// FindIndex finds index with given schema, table and column.
func (m *ModelPostgres) FindIndex(schema, table string) ([]*PostgreIndex, error) {
	querySql := `
	select A.INDEXNAME AS index_name,
		   C.INDEXRELID AS index_id,
		   C.INDISUNIQUE AS is_unique,
		   C.INDISPRIMARY AS is_primary,
		   G.ATTNAME AS column_name,
		   G.attnum AS index_sort
	from PG_AM B
			 left join PG_CLASS F on
		B.OID = F.RELAM
			 left join PG_STAT_ALL_INDEXES E on
		F.OID = E.INDEXRELID
			 left join PG_INDEX C on
		E.INDEXRELID = C.INDEXRELID
			 left outer join PG_DESCRIPTION D on
		C.INDEXRELID = D.OBJOID,
		 PG_INDEXES A,
		 pg_attribute G
	where A.SCHEMANAME = E.SCHEMANAME
	  and A.TABLENAME = E.RELNAME
	  and A.INDEXNAME = E.INDEXRELNAME
	  and F.oid = G.attrelid
	  and E.SCHEMANAME = $1
	  and E.RELNAME = $2
    order by C.INDEXRELID,G.attnum`

	var reply []*PostgreIndex
	err := m.DB.Raw(querySql, schema, table).Scan(&reply).Error
	if err != nil {
		return nil, err
	}

	return reply, nil
}
