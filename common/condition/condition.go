package condition

import (
	"reflect"
	"strconv"
)

type Condition struct {
	Field    string      // 字段名
	Operator string      // 操作符
	Value    interface{} // 值
	Logic    string      // "AND" 或 "OR"
}

type Req struct {
	Conditions []Condition // 查询条件

}

func EvaluateCondition(link interface{}, conditions []Condition) bool {
	val := reflect.ValueOf(link).Elem()
	result := true // 如果逻辑是OR，初始化result为false，否则为true
	for _, cond := range conditions {
		fieldVal := val.FieldByName(cond.Field)
		if !fieldVal.IsValid() {
			continue // 字段不存在
		}
		match := false
		switch fieldVal.Kind() {
		case reflect.String:
			match = compareString(fieldVal.String(), cond.Operator, cond.Value.(string))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			condVal, err := strconv.ParseInt(cond.Value.(string), 10, 64)
			if err != nil {
				continue // 无效的条件值
			}
			match = compareInt(fieldVal.Int(), cond.Operator, condVal)
		}

		if cond.Logic == "AND" {
			result = result && match
			if !result {
				break // 如果逻辑是AND，任何一个条件不匹配就可以直接返回
			}
		} else if cond.Logic == "OR" {
			result = result || match
			if result {
				break // 如果逻辑是OR，任何一个条件匹配就可以直接返回
			}
		} else {
			result = match
		}
	}

	return result
}

// 比较字符串类型的值
func compareString(a, operator, b string) bool {
	switch operator {
	case "=":
		return a == b
	case "!=":
		return a != b
	// 更多运算符...
	default:
		return false
	}
}

// 比较整数类型的值
func compareInt(a int64, operator string, b int64) bool {
	switch operator {
	case "=":
		return a == b
	case "!=":
		return a != b
	case "<":
		return a < b
	case "<=":
		return a <= b
	case ">":
		return a > b
	case ">=":
		return a >= b
	// 更多运算符...
	default:
		return false
	}
}
