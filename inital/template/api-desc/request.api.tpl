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
        List []Get{{.Table}}Response `json:"list"`
    }

    //Create{{.Table}}Request 创建-{{.TableComment}}
    Create{{.Table}}Request {
        {{- range  .Columns }}
            {{- if .IsPk }}
            {{else if .IsPage}}
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

    //Delete{{.Table}}Response 删除-{{.TableComment}}
    Delete{{.Table}}Response {}
)