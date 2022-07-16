package gen

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

type FieldType struct {
	Type  string `json:"type"`
	Label string `json:"label"`
}

//go:embed type.json
var filedType string

func getField() (data []FieldType) {

	_ = json.Unmarshal([]byte(filedType), &data)
	return data
}

func getType() map[string]string {
	data := getField()
	ret := make(map[string]string)
	for _, v := range data {
		fmt.Println(v.Label, v.Type)
		ret[v.Label] = v.Type
	}
	return ret
}
