package gormx

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	// FromQueryTag tag标记
	FromQueryTag = "search"
	// Mysql 数据库标识
	Mysql = "mysql"
	// Postgres 数据库标识
	Postgres = "postgres"
)

// ResolveSearchQuery 解析
/**
 * 	exact / iexact 等于
 * 	contains / icontains 包含
 *	gt / gte 大于 / 大于等于
 *	lt / lte 小于 / 小于等于
 *	startswith / istartswith 以…起始
 *	endswith / iendswith 以…结束
 *	in
 *	isnull
 *  order 排序		e.g. order[key]=desc     order[key]=asc
 */
func ResolveSearchQuery(driver string, q interface{}, condition Condition) {
	qType := reflect.TypeOf(q)
	qValue := reflect.ValueOf(q)
	var tag string
	var ok bool
	var t *resolveSearchTag
	for i := 0; i < qType.NumField(); i++ {
		tag, ok = "", false
		tag, ok = qType.Field(i).Tag.Lookup(FromQueryTag)
		if !ok {
			//递归调用
			ResolveSearchQuery(driver, qValue.Field(i).Interface(), condition)
			continue
		}
		switch tag {
		case "-":
			continue
		}
		t = makeTag(tag)
		if qValue.Field(i).IsZero() {
			continue
		}

		//解析
		switch t.Type {
		//fix Postgres 双引号
		case "left":
			//左关联
			if driver == Postgres {
				join := condition.SetJoinOn(t.Type, fmt.Sprintf(
					`left join "%s" on "%s"."%s" = "%s"."%s"`,
					t.Join,
					t.Join,
					t.On[0],
					t.Table,
					t.On[1],
				))
				ResolveSearchQuery(driver, qValue.Field(i).Interface(), join)
			} else {
				join := condition.SetJoinOn(t.Type, fmt.Sprintf(
					"left join `%s` on `%s`.`%s` = `%s`.`%s`",
					t.Join,
					t.Join,
					t.On[0],
					t.Table,
					t.On[1],
				))
				ResolveSearchQuery(driver, qValue.Field(i).Interface(), join)
			}
		case "exact", "iexact":
			if driver == Postgres {
				condition.SetWhere(fmt.Sprintf(` "%s"."%s" = ?`, t.Table, t.Column), []interface{}{qValue.Field(i).Interface()})
			} else {
				condition.SetWhere(fmt.Sprintf("`%s`.`%s` = ?", t.Table, t.Column), []interface{}{qValue.Field(i).Interface()})
			}
		case "contains", "icontains":
			if driver == Postgres {
				condition.SetWhere(fmt.Sprintf(` "%s"."%s" ilike ?`, t.Table, t.Column), []interface{}{"%" + qValue.Field(i).String() + "%"})
			} else {
				condition.SetWhere(fmt.Sprintf("`%s`.`%s` like ?", t.Table, t.Column), []interface{}{"%" + qValue.Field(i).String() + "%"})
			}
		case "gt":
			if driver == Postgres {
				condition.SetWhere(fmt.Sprintf(` "%s"."%s" > ?`, t.Table, t.Column), []interface{}{qValue.Field(i).Interface()})
			} else {
				condition.SetWhere(fmt.Sprintf("`%s`.`%s` > ?", t.Table, t.Column), []interface{}{qValue.Field(i).Interface()})
			}
		case "gte":
			if driver == Postgres {
				condition.SetWhere(fmt.Sprintf(` "%s"."%s" >= ?`, t.Table, t.Column), []interface{}{qValue.Field(i).Interface()})
			} else {
				condition.SetWhere(fmt.Sprintf("`%s`.`%s` >= ?", t.Table, t.Column), []interface{}{qValue.Field(i).Interface()})
			}
		case "lt":
			if driver == Postgres {
				condition.SetWhere(fmt.Sprintf(` "%s"."%s"  < ?`, t.Table, t.Column), []interface{}{qValue.Field(i).Interface()})
			} else {
				condition.SetWhere(fmt.Sprintf("`%s`.`%s` < ?", t.Table, t.Column), []interface{}{qValue.Field(i).Interface()})
			}
		case "lte":
			if driver == Postgres {
				condition.SetWhere(fmt.Sprintf(` "%s"."%s"  <=  ?`, t.Table, t.Column), []interface{}{qValue.Field(i).Interface()})
			} else {
				condition.SetWhere(fmt.Sprintf("`%s`.`%s` <= ?", t.Table, t.Column), []interface{}{qValue.Field(i).Interface()})
			}
		case "startswith", "istartswith":
			if driver == Postgres {
				condition.SetWhere(fmt.Sprintf(`"%s"."%s" ilike ?`, t.Table, t.Column), []interface{}{qValue.Field(i).String() + "%"})
			} else {
				condition.SetWhere(fmt.Sprintf("`%s`.`%s` like ?", t.Table, t.Column), []interface{}{qValue.Field(i).String() + "%"})
			}
		case "endswith", "iendswith":
			if driver == Postgres {
				condition.SetWhere(fmt.Sprintf(`"%s"."%s" ilike ?`, t.Table, t.Column), []interface{}{"%" + qValue.Field(i).String()})
			} else {
				condition.SetWhere(fmt.Sprintf("`%s`.`%s` like ?", t.Table, t.Column), []interface{}{"%" + qValue.Field(i).String()})
			}
		case "in":
			if driver == Postgres {
				condition.SetWhere(fmt.Sprintf(`"%s"."%s" in (?)`, t.Table, t.Column), []interface{}{qValue.Field(i).Interface()})
			} else {
				condition.SetWhere(fmt.Sprintf("`%s`.`%s` in (?)", t.Table, t.Column), []interface{}{qValue.Field(i).Interface()})
			}
		case "isnull":
			if !(qValue.Field(i).IsZero() && qValue.Field(i).IsNil()) {
				if driver == Postgres {
					condition.SetWhere(fmt.Sprintf(` "%s"."%s" isnull`, t.Table, t.Column), make([]interface{}, 0))
				} else {
					condition.SetWhere(fmt.Sprintf("`%s`.`%s` isnull", t.Table, t.Column), make([]interface{}, 0))
				}
			}
		case "order":
			switch strings.ToLower(qValue.Field(i).String()) {
			case "desc", "asc":
				if driver == Postgres {
					condition.SetOrder(fmt.Sprintf(`"%s"."%s" %s`, t.Table, t.Column, qValue.Field(i).String()))
				} else {
					condition.SetOrder(fmt.Sprintf("`%s`.`%s` %s", t.Table, t.Column, qValue.Field(i).String()))
				}
			}
		}
	}
}
