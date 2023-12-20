package toolx

import (
	"errors"
	"reflect"
)

type TagBody struct {
	Header []string `json:"header"`
	Keys   []string `json:"keys"`
	Field  []string `json:"field"`
}

// GetTag orm模型中获取tag
func GetTag(v interface{}) (ta TagBody) {
	t := reflect.TypeOf(v).Elem()
	for i := 0; i < t.NumField(); i++ {
		tg := t.Field(i).Tag.Get("json") //将tag输出出来
		if tg != "" {
			ta.Field = append(ta.Field, t.Field(i).Name)
			ta.Keys = append(ta.Keys, tg)
			key := t.Field(i).Tag.Get("comment")
			if key != "" {
				ta.Header = append(ta.Header, key)
			} else {
				ta.Header = append(ta.Header, tg)
			}
		}
	}
	return ta
}

// GetFieldByTag GetFieldByTag
func (t *TagBody) GetFieldByTag(tag string) (string, error) {
	for i, v := range t.Keys {
		if v == tag {
			return t.Field[i], nil
		}
	}
	return "", errors.New("不存在")
}
