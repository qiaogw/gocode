package gormx

import (
	"fmt"
	"github.com/qiaogw/gocode/gen"
	"github.com/qiaogw/gocode/global"
	"gorm.io/gorm"
	"reflect"
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
		sql = fmt.Sprintf(`CAST("%s" AS text) ~ '%s'`, table, key)
	}

	return sql
}

func MakeCondition(q interface{}, driver string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		condition := &GormCondition{
			GormPublic: GormPublic{},
			Join:       make([]*GormJoin, 0),
		}
		ResolveSearchQuery(driver, q, condition)
		for _, join := range condition.Join {
			if join == nil {
				continue
			}
			db = db.Joins(join.JoinOn)
			for k, v := range join.Where {
				db = db.Where(k, v...)
			}
			for k, v := range join.Or {
				db = db.Or(k, v...)
			}
			for _, o := range join.Order {
				db = db.Order(o)
			}
		}
		for k, v := range condition.Where {
			db = db.Where(k, v...)
		}
		for k, v := range condition.Or {
			db = db.Or(k, v...)
		}
		for _, o := range condition.Order {
			db = db.Order(o)
		}
		return db
	}
}

func Paginate(pageSize, pageIndex int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (pageIndex - 1) * pageSize
		if offset < 0 {
			offset = 0
		}
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}

func SortBy(sortBy string, descending bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var orderBy string
		if descending {
			orderBy = sortBy + " DESC"
		} else {
			orderBy = sortBy
		}
		return db.Order(orderBy)
	}
}

// GetSortBy 格式化sortBy字段
func GetSortBy(sort interface{}, fieldName string) string {
	tagJson := "json"
	rt := ""
	var t *resolveSearchTag
	// 获取结构体的类型
	qType := reflect.TypeOf(sort)
	// 遍历结构体的字段
	for i := 0; i < qType.NumField(); i++ {
		tag, ok := qType.Field(i).Tag.Lookup(FromQueryTag)
		if !ok {
			continue
		}
		switch tag {
		case "-":
			continue
		}
		t = makeTag(tag)
		if qType.Field(i).Tag.Get(tagJson) == fieldName {
			rt = t.Column
			break
		}
	}
	return rt
}
