package modelx

import (
	"reflect"
)

type TableTag struct {
	Label string `json:"label"`
	Field string `json:"field"`
}

func GetTableTag(obj interface{}) []*TableTag {
	var repList []*TableTag
	objType := reflect.TypeOf(obj)
	if objType.Kind() == reflect.Struct {
		for i := 0; i < objType.NumField(); i++ {
			field := objType.Field(i)
			rep := TableTag{
				Label: field.Tag.Get("comment"),
				Field: field.Tag.Get("json"),
			}
			repList = append(repList, &rep)
		}
	}
	return repList
}
