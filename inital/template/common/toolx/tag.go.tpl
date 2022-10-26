package toolx

import "reflect"

type TagBody struct {
Header []string `json:"header"`
Keys   []string `json:"keys"`
Field  []string `json:"field"`
}

// GetTag orm模型中获取tag
func GetTag(v interface{}) (ta TagBody) {
t := reflect.TypeOf(v).Elem()
//beego.Debug(t)
// var tb TagBody
for i := 0; i < t.NumField(); i++ {
tg := t.Field(i).Tag.Get("json") //将tag输出出来
if tg != "" {
ta.Field = append(ta.Field, t.Field(i).Name)
ta.Keys = append(ta.Keys, tg)
key := t.Field(i).Tag.Get("comment")
if key != "" {
//arr := strings.Split(key, ";")
//for _, v := range arr {
//	c := strings.Split(v, ":")
//	if c[0] == "comment" {
//		ta.Header = append(ta.Header, c[1])
//	}
//}
ta.Header = append(ta.Header, key)
} else {
ta.Header = append(ta.Header, tg)
}
}
}
return ta
}

// GetTagSelf kv获取tag
func GetTagSelf(data []map[string]interface{}) (ta TagBody) {
var tb TagBody
for k, _ := range data[0] {
tb.Header = append(tb.Header, k)
tb.Keys = append(tb.Keys, k)
}
return ta
}
