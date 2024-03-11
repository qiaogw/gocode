type (
    //Get{{.Table}}Request 提取-{{.TableComment}}
    Get{{.Table}}Request {
{{- range  .Columns}}
    {{- if .IsPk }}
        {{.FieldName}} {{.DataTypeApi}} `json:"{{.FieldJson}}"`
    {{- end }}
{{- end }}
    }

    //Get{{.Table}}Response 提取-{{.TableComment}}
    Get{{.Table}}Response {
    {{- range  .Columns }}
        {{- if .IsPage}}
        {{- else}}
        {{.FieldName}} {{.DataTypeApi}} `json:"{{.FieldJson}},omitempty"`
        {{- end }}
    {{- end }}
    {{- if .IsFlow }}
        FlowInstance GetFlowInstanceResponse `json:"flowInstance,omitempty"`
        CurrentNode GetNodeInstanceResponse `json:"currentNode,omitempty"`
        FlowInstanceHistory []*GetFlowInstanceResponse `json:"flowInstanceHistory,omitempty"`
    {{- end}}
    }

    //List{{.Table}}Request 列表-{{.TableComment}}-
    List{{.Table}}Request {
{{- range  .Columns }}
    {{- if .IsPk }}
    {{else}}
        {{.FieldName}} {{.DataTypeApi}} `json:"{{.FieldJson}},optional"`
    {{- end }}
{{- end }}
        SearchKey string  `json:"searchKey,optional"`
        SortBY string `json:"sortBy,optional"`
        Descending bool `json:"descending,optional"`
    }

    //List{{.Table}}Response 列表-{{.TableComment}}
    List{{.Table}}Response {
        Count int64 `json:"count"`
        List []*Get{{.Table}}Response `json:"list"`
    }

    //Create{{.Table}}Request 创建-{{.TableComment}}
    Create{{.Table}}Request {
{{- range  .Columns }}
    {{- if .IsPk }}
    {{- else if .IsPage}}
    {{- else if .IsModelTime}}
    {{- else if .IsControl}}
    {{- else}}
        {{.FieldName}} {{.DataTypeApi}} `json:"{{.FieldJson}}{{- if .IsNull -}},optional{{- end -}}{{- if .ColumnDefault -}},default={{.ColumnDefault}}{{- end -}}"`
    {{- end }}
{{- end }}
    }

    //Create{{.Table}}Response 创建-{{.TableComment}}
    Create{{.Table}}Response {
{{- range  .Columns }}
    {{- if .IsPk }}
        {{.FieldName}} {{.DataTypeApi}} `json:"{{.FieldJson}},omitempty"`
    {{- end }}
{{- end }}
    }

    //Update{{.Table}}Request 修改-{{.TableComment}}
    Update{{.Table}}Request {
    {{- range  .Columns }}
        {{- if .IsPage}}
        {{- else if .IsPage}}
        {{- else if .IsModelTime}}
        {{- else if .IsControl}}
        {{- else}}
       {{.FieldName}} {{.DataTypeApi}} `json:"{{.FieldJson}}{{- if .IsNull -}},optional{{- end -}}{{- if .ColumnDefault -}},default={{.ColumnDefault}}{{- end -}}"`
        {{- end }}
    {{- end }}
    }

    // Update{{.Table}}Response 修改-{{.TableComment}}
    Update{{.Table}}Response {}

    //Delete{{.Table}}Request 删除-{{.TableComment}}
    Delete{{.Table}}Request {
{{- range  .Columns }}
    {{- if .IsPk }}
        {{.FieldName}} {{.DataTypeApi}} `json:"{{.FieldJson}}"`
    {{- end }}
{{- end }}
    }

    //DeleteList{{.Table}}Request 批量删除-{{.TableComment}}
    DeleteList{{.Table}}Request {
        IdList []string `json:"idList"`
    }

    //Delete{{.Table}}Response 删除-{{.TableComment}}
    Delete{{.Table}}Response {}
)