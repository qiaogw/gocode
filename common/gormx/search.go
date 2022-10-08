package gormx

import (
	"fmt"
	"github.com/qiaogw/gocode/gen"
	"github.com/qiaogw/gocode/global"
	"gorm.io/gorm"
)

func SearchKey(db *gorm.DB, table, key string) string {
	var sql string
	//
	genApp := gen.AutoCodeServiceApp
	global.GenDB = db

	genApp.Init()

	database := db.Config.NamingStrategy.SchemaName(table)
	field, err := genApp.DB.GetColumn(database, table)
	if err != nil {
		return sql
	}
	// sql
	switch db.Name() {
	case "mysql":
		sql = fmt.Sprintf("concat(%v) like '%%%s%%'", field, key)
	case "postgres":
		sql = fmt.Sprintf("record_to_text(%s) ~ '%s'", table, key)
	}

	return sql
}
