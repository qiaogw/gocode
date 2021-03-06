package model

import (
    "{{.ParentPkg}}/common/global"
{{ if .HasTimer }}"time"{{ end }}
)

type List{{.Table}}Req struct {
    global.Pagination `search:"-"`
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
