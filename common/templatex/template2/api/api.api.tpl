syntax = "v1"

import (
"replay/replay.api"
{{- range .Tables }}
	"api-desc/{{.TableUrl}}.api"
{{- end}}
	"admin-desc/api.api"
	"admin-desc/config.api"
	"admin-desc/dept.api"
	"admin-desc/dictdata.api"
	"admin-desc/dicttype.api"
	"admin-desc/loginlog.api"
	"admin-desc/menu.api"
	"admin-desc/migration.api"
	"admin-desc/operalog.api"
	"admin-desc/post.api"
	"admin-desc/role.api"
	"admin-desc/user.api"
	"admin-desc/login.api"
)

info (
	title: "{{.Service}}"//  add title
	desc: "{{.Service}}"//  add description
	author: "{{.Author}}"
	email: "{{.Email}}"
)

{{- range .Tables }}

{{- if .IsAuth }}
//需要登录
{{- else}}
//不需要登录
{{- end }}
@server(
	group : {{.TableUrl}}
	prefix : /{{.Db}}
{{- if .IsAuth}}
	jwt: Auth
{{- end }}
)
service {{.Db}} {
	@doc(
		summary: "提取-{{.TableComment}}"
	)
	@handler Get{{.Table}}
	post  /{{.TableUrl}}/get (Get{{.Table}}Request) returns(Get{{.Table}}Response)

	@doc(
		summary: "列表-{{.TableComment}}"
	)
	@handler List{{.Table}}
	post /{{.TableUrl}}/list (List{{.Table}}Request) returns(List{{.Table}}Response )

	@doc(
		summary: "创建-{{.TableComment}}"
	)
	@handler Create{{.Table}}
	post  /{{.TableUrl}}/create (Create{{.Table}}Request) returns(CommonResponse)

	@doc(
		summary: "更新-{{.TableComment}}"
	)
	@handler Update{{.Table}}
	post /{{.TableUrl}}/update (Update{{.Table}}Request) returns(CommonResponse)

	@doc(
		summary: "删除-{{.TableComment}}"
	)
	@handler Delete{{.Table}}
	post  /{{.TableUrl}}/delete (Delete{{.Table}}Request) returns(CommonResponse)

{{- if .IsImport}}
	@doc(
		summary: "导出-{{.TableComment}}"
	)
	@handler Export{{.Table}}
	post /{{.TableUrl}}/export (List{{.Table}}Request) returns (CommonResponse)

	@doc(
	summary: "导出-{{.TableComment}}模板"
	)
	@handler ExportTemplate{{.Table}}
	post /{{.TableUrl}}/exportTemplate (NullRequest) returns (CommonResponse)

	@doc(
		summary: "导入-{{.TableComment}}"
	)
	@handler Import{{.Table}}
	post /{{.TableUrl}}/import (ImportRequest) returns (CommonResponse)
{{- end }}
}
{{- end }}

