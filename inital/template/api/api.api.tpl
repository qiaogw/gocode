syntax = "v1"

info (
	title: "{{.Service}}"//  add title
	desc: "{{.Service}}"//  add description
	author: "{{.GitUser}}"
	email: "{{.GitEmail}}"
)
type (
{{range .Tables }}
    //Get{{.Table}}Request 提取-{{.TableComment}}
    Get{{.Table}}Request {
        {{- range  .Columns}}
            {{- if .IsPk }}
                {{.FieldName}} {{.DataTypeProto}} `json:"{{.FieldJson}}"`
            {{- end }}
        {{- end }}
    }
    //Get{{.Table}}Response 提取-{{.TableComment}}
    Get{{.Table}}Response {
        {{- range  .Columns }}
            {{.FieldName}} {{.DataTypeProto}} `json:"{{.FieldJson}}"`
        {{- end }}
    }
    //List{{.Table}}Request 列表-{{.TableComment}}-
    List{{.Table}}Request {
    {{- range  .Columns }}
        {{- if .IsPk }}
        {{else}}
            {{.FieldName}} {{.DataTypeProto}} `json:"{{.FieldJson}}"`
        {{- end }}
    {{- end }}
    }
    //List{{.Table}}Response 列表-{{.TableComment}}
    List{{.Table}}Response {
        {{- range  .Columns }}
            {{.FieldName}} {{.DataTypeProto}} `json:"{{.FieldJson}}"`
        {{- end }}
    }
    //Create{{.Table}}Request 创建-{{.TableComment}}
    Create{{.Table}}Request {
        {{- range  .Columns }}
            {{- if .IsPk }}
            {{else}}
                {{.FieldName}} {{.DataTypeProto}} `json:"{{.FieldJson}}"`
            {{- end }}
        {{- end }}
    }
    //Create{{.Table}}Response 创建-{{.TableComment}}
    Create{{.Table}}Response {
        {{- range  .Columns }}
            {{- if .IsPk }}
                {{.FieldName}} {{.DataTypeProto}} `json:"{{.FieldJson}}"`
            {{- end }}
        {{- end }}
    }
    //Update{{.Table}}Request 修改-{{.TableComment}}
    Update{{.Table}}Request {
        {{- range  .Columns }}
            {{.FieldName}} {{.DataTypeProto}} `json:"{{.FieldJson}}"`
        {{- end }}
    }

    // Update{{.Table}}Response 修改-{{.TableComment}}
    Update{{.Table}}Response {}

    //Delete{{.Table}}Request 删除-{{.TableComment}}
    Delete{{.Table}}Request {
        {{- range  .Columns }}
            {{- if .IsPk }}
                {{.FieldName}} {{.DataTypeProto}} `json:"{{.FieldJson}}"`
            {{- end }}
        {{- end }}
    }
    //Delete{{.Table}}Response 删除-{{.TableComment}}
    Delete{{.Table}}Response {
    }
{{- end }}
)

service {{.Package}} {
{{- range .Tables }}
    @doc(
        summary: "提取-{{.TableComment}}"
    )
	@handler Get{{.Table}} //  set handler name and delete this comment
	post  /{{.TableUrl}}/get (Get{{.Table}}Request) returns(Get{{.Table}}Response)

    @doc(
        summary: "列表-{{.TableComment}}"
    )
	@handler List{{.Table}} //  set handler name and delete this comment
	post /{{.TableUrl}}/list (List{{.Table}}Request) returns(List{{.Table}}Response)

    @doc(
        summary: "创建-{{.TableComment}}"
    )
	@handler Create{{.Table}} //  set handler name and delete this comment
	post  /{{.TableUrl}}/create (Create{{.Table}}Request) returns(Create{{.Table}}Response)

    @doc(
        summary: "更新-{{.TableComment}}"
    )
	@handler Update{{.Table}} //  set handler name and delete this comment
	post /{{.TableUrl}}/update (Update{{.Table}}Request) returns(Update{{.Table}}Response)

    @doc(
        summary: "删除-{{.TableComment}}"
    )
	@handler Delete{{.Table}} //  set handler name and delete this comment
	post  /{{.TableUrl}}/delete (Delete{{.Table}}Request) returns(Delete{{.Table}}Response)
{{- end }}
}
