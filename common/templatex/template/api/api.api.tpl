syntax = "v1"

import (
"replay/replay.api"
{{- range .Tables }}
	"api-desc/{{.TableUrl}}.api"
{{- end}}
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
	{{ else}}
		//不需要登录
	{{ end }}
	@server(
	group : {{.TableUrl}}
	prefix : /{{.TableUrl}}
	{{- if .IsAuth}}
		jwt: Auth
	{{- end }}
	)
	service {{.Db}} {
	@doc(
	summary: "提取-{{.TableComment}}"
	)
	@handler Get{{.Table}}
	post  /get (Get{{.Table}}Request) returns(Get{{.Table}}Response)

	@doc(
	summary: "列表-{{.TableComment}}"
	)
	@handler List{{.Table}}
	post /list (List{{.Table}}Request) returns(List{{.Table}}Response )

	@doc(
	summary: "创建-{{.TableComment}}"
	)
	@handler Create{{.Table}}
	post  /create (Create{{.Table}}Request) returns(CommonResponse)

	@doc(
	summary: "更新-{{.TableComment}}"
	)
	@handler Update{{.Table}}
	post /update (Update{{.Table}}Request) returns(CommonResponse)

	@doc(
	summary: "删除-{{.TableComment}}"
	)
	@handler Delete{{.Table}}
	post  /delete (Delete{{.Table}}Request) returns(CommonResponse)

	{{- if .IsImport}}
		@doc(
		summary: "导出-{{.TableComment}}"
		)
		@handler Export{{.Table}}
		post /export (List{{.Table}}Request) returns (CommonResponse)

		@doc(
		summary: "导出-{{.TableComment}}模板"
		)
		@handler ExportTemplate{{.Table}}
		post /exportTemplate (NullRequest) returns (CommonResponse)

		@doc(
		summary: "导入-{{.TableComment}}"
		)
		@handler Import{{.Table}}
		post /import (ImportRequest) returns (CommonResponse)
	{{- end }}
	}
{{- end }}