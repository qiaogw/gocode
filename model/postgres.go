package model

import (
	"database/sql"
	"github.com/qiaogw/gocode/global"
	"gorm.io/gorm"
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

// PostgreColumn describes a column in table
type PostgreColumn struct {
	Num               sql.NullInt32  `db:"num"`
	Field             sql.NullString `db:"field"`
	Type              sql.NullString `db:"type"`
	DataTypeLong      sql.NullString `db:"dataTypeLong"`
	NotNull           sql.NullBool   `db:"not_null"`
	Comment           sql.NullString `db:"comment"`
	ColumnDefault     sql.NullString `db:"column_default"`
	IdentityIncrement sql.NullInt32  `db:"identity_increment"`
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
// Author [qiaogw](https://github.com/qiaogw)
func (m *ModelPostgres) GetColumn(db, table string) (*ColumnData, error) {
	querySql := `
	select 
		t.num,
		t.field,
		t.type,
		t.dataTypeLong,
		t.not_null,
		t.comment, 
		c.column_default, 
		identity_increment
	from (
         SELECT a.attnum AS num,
                c.relname,
                a.attname     AS field,
                t.typname     AS type,
                a.atttypmod   AS lengthvar,
                a.attnotnull  AS not_null,
				(case
					when a.attlen > 0 then a.attlen
					when t.typname='bit' then a.atttypmod
					else a.atttypmod - 4 end) AS dataTypeLong,
                b.description AS comment
         FROM pg_class c,
              pg_attribute a
                  LEFT OUTER JOIN pg_description b ON a.attrelid = b.objoid 
					AND a.attnum = b.objsubid,
              pg_type t
         WHERE c.relname = $1
           and a.attnum > 0
           and a.attrelid = c.oid
           and a.atttypid = t.oid
 		 GROUP BY 
			a.attnum, 
			a.attlen,
			c.relname, 
			a.attname, 
			t.typname, 
			a.atttypmod, 
			a.attnotnull, 
			b.description
         ORDER BY a.attnum
	) AS t
         left join information_schema.columns AS c on t.relname = c.table_name 
		and t.field = c.column_name and c.table_schema = 'public'`

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
	//log.Printf("tableis %s,is len is %+v\n", table, len(in))
	var list []*Column
	for _, e := range in {
		//log.Printf("table is %s,FieldName is %s\n", table, e.Field)
		//log.Printf("each.name is %s,is pk is %+v\n", e.Field, e.Comment)
		var dft interface{}
		if len(e.ColumnDefault.String) > 0 {
			dft = e.ColumnDefault
		}

		isNullAble := "YES"
		if e.NotNull.Bool {
			isNullAble = "NO"
		}

		var extra string
		// when identity is true, the column is auto increment
		if e.IdentityIncrement.Int32 == 1 {
			extra = "auto_increment"
		}

		// when type is serial, it'm auto_increment. and the default value is tablename_columnname_seq
		if strings.Contains(e.ColumnDefault.String, table+"_"+e.Field.String+"_seq") {
			extra = "auto_increment"
		}

		if len(index[e.Field.String]) > 0 {
			for _, i := range index[e.Field.String] {
				list = append(list, &Column{
					DbColumn: &DbColumn{
						Name:            e.Field.String,
						DataType:        m.convertPostgreSqlTypeIntoMysqlType(e.Type.String),
						Extra:           extra,
						Comment:         e.Comment.String,
						ColumnDefault:   dft,
						IsNullAble:      isNullAble,
						DataTypeLong:    e.DataTypeLong.String,
						OrdinalPosition: int(e.Num.Int32),
					},
					Index: i,
				})
			}
		} else {
			list = append(list, &Column{
				DbColumn: &DbColumn{
					Name:            e.Field.String,
					DataType:        m.convertPostgreSqlTypeIntoMysqlType(e.Type.String),
					Extra:           extra,
					Comment:         e.Comment.String,
					ColumnDefault:   dft,
					IsNullAble:      isNullAble,
					OrdinalPosition: int(e.Num.Int32),
				},
			})
		}
	}
	//log.Println("list len is ", len(list))
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
