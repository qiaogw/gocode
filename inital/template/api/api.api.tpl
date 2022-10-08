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
	author: "{{.GitUser}}"
	email: "{{.GitEmail}}"
)
type (
    // 用户登录
    LoginRequest {
        Mobile    string `json:"mobile"`
        Password  string `json:"password"`
        Captcha   string `json:"captcha"`   // 验证码
        CaptchaId string `json:"captchaId"` // 验证码ID
    }
    // 验证码
    CaptchaRequest  struct{}
)

service {{.Package}} {
@doc(
summary: "验证码"
)
@handler Captcha //  
post  /base/captcha (CaptchaRequest) returns(CommonResponse)

@doc(
summary: "用户登录"
)
@handler Login //  
post  /base/login (LoginRequest) returns(CommonResponse)
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
