package model

import (
    "github.com/qiaogw/gocode/common/modelx"
{{ if .HasTimer }}"time"{{ end }}
)

type List{{.Table}}Req struct {
    modelx.Pagination `search:"-"`
{{- range .Columns }}
    {{- if .IsPage}}
    {{- else}}
        {{.FieldName}}  {{.DataType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" search:"type:exact;column:{{.GormName}};table:{{.TableName}}"`
    {{- end -}}
{{- end }}
    {{.Table}}Order
}


type {{.Table}}Order struct {
{{- range .Columns }}
    {{- if .IsPage}}
    {{- else}}
        {{.FieldName}}Order  {{.DataType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" search:"type:order;column:{{.GormName}};table:{{.TableName}}"`
    {{- end -}}
{{- end }}
}


func (m *List{{.Table}}Req) GetNeedSearch() interface{} {
    return *m
}
