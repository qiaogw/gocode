package converter

import (
	"reflect"
)

func ConvertDefault(c interface{}) string {
	cType := reflect.TypeOf(c)
	cValue := reflect.ValueOf(c)
	//log.Printf("mValue.FieldByName(mType.Field(i).Name) is %s\n", cValue.FieldByName(cType.Field(0).Name).String())
	if cType.Name() == "NullString" {
		return cValue.FieldByName("String").String()
	}
	return ""
}
