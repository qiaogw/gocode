syntax = "v1"

import (
    "replay/{{.Package}}.api"
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

// 需要登录
@server(
    group : {{.Package}}
    prefix : /{{.Package}}
    jwt: Auth
)
service {{.Package}} {
{{- range .Tables }}
    @doc(
        summary: "提取-{{.TableComment}}"
    )
	@handler Get{{.Table}} //  
	post  /{{.TableUrl}}/get (Get{{.Table}}Request) returns(CommonResponse)

    @doc(
        summary: "列表-{{.TableComment}}"
    )
	@handler List{{.Table}} //  
	post /{{.TableUrl}}/list (List{{.Table}}Request) returns(CommonResponse)

    @doc(
        summary: "创建-{{.TableComment}}"
    )
	@handler Create{{.Table}} //  
	post  /{{.TableUrl}}/create (Create{{.Table}}Request) returns(CommonResponse)

    @doc(
        summary: "更新-{{.TableComment}}"
    )
	@handler Update{{.Table}} //  
	post /{{.TableUrl}}/update (Update{{.Table}}Request) returns(CommonResponse)

    @doc(
        summary: "删除-{{.TableComment}}"
    )
	@handler Delete{{.Table}} //  
	post  /{{.TableUrl}}/delete (Delete{{.Table}}Request) returns(CommonResponse)
{{- end }}
}
