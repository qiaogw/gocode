package condition

import (
	"reflect"
	"strconv"
)

// Condition 表示一个查询条件，包括字段名、操作符、条件值和逻辑运算符（"AND" 或 "OR"）
type Condition struct {
	Field    string      // 字段名
	Operator string      // 操作符
	Value    interface{} // 条件值
	Logic    string      // 逻辑运算符："AND" 或 "OR"
}

// Req 表示一个查询请求，其中包含多个查询条件
type Req struct {
	Conditions []Condition // 查询条件列表
}

// EvaluateCondition 根据给定的条件列表对目标对象进行条件判断
// 参数 link 必须为指向结构体的指针，conditions 为条件列表
// 该函数会利用反射获取目标对象中指定字段的值，并根据条件中的操作符进行比较，
// 同时根据每个条件的逻辑运算符（AND/OR）决定整体判断结果，返回 true 表示满足条件，false 表示不满足。
func EvaluateCondition(link interface{}, conditions []Condition) bool {
	// 通过反射获取目标对象的值（需要传入指针）
	val := reflect.ValueOf(link).Elem()
	// 初始化结果为 true（默认所有条件均满足）
	result := true
	// 遍历所有条件
	for _, cond := range conditions {
		// 根据条件中的字段名获取对应字段的值
		fieldVal := val.FieldByName(cond.Field)
		if !fieldVal.IsValid() {
			// 如果字段不存在，则跳过该条件
			continue
		}
		// match 表示当前条件是否匹配
		match := false
		// 根据字段的类型调用对应的比较函数
		switch fieldVal.Kind() {
		case reflect.String:
			// 如果字段类型为字符串，直接比较字符串值
			match = compareString(fieldVal.String(), cond.Operator, cond.Value.(string))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			// 如果字段类型为整数，则将条件值转换为 int64 后比较
			condVal, err := strconv.ParseInt(cond.Value.(string), 10, 64)
			if err != nil {
				// 如果转换失败，跳过该条件
				continue
			}
			match = compareInt(fieldVal.Int(), cond.Operator, condVal)
		}

		// 根据条件的逻辑运算符更新整体结果
		if cond.Logic == "AND" {
			// 对于 AND 逻辑，所有条件必须匹配
			result = result && match
			if !result {
				// 如果有任一 AND 条件不匹配，则直接返回 false
				break
			}
		} else if cond.Logic == "OR" {
			// 对于 OR 逻辑，只要有任一条件匹配即可
			result = result || match
			if result {
				// 如果已有条件匹配，则直接返回 true
				break
			}
		} else {
			// 如果未明确指定逻辑运算符，则采用当前条件的匹配结果
			result = match
		}
	}

	return result
}

// compareString 用于比较两个字符串 a 和 b，根据指定的操作符 operator 判断是否满足条件
// 当前支持 "="（等于）和 "!="（不等于）操作符，可根据需要扩展更多运算符
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

// compareInt 用于比较两个整数 a 和 b，根据指定的操作符 operator 判断是否满足条件
// 当前支持 "="（等于）、"!="（不等于）、"<"（小于）、"<="（小于等于）、">"（大于）、">="（大于等于）操作符，可根据需要扩展更多运算符
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
