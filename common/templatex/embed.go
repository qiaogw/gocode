package templatex

import (
	"embed"
)

var (
	//go:embed template
	TemplateTpl embed.FS
	//go:embed template2
	TemplateApi embed.FS

	rpcPath = "template"
	apiPath = "template2"
)

func GetTpl(mode string) embed.FS {
	switch mode {
	case "rpc":
		return TemplateTpl
	case "api":
		return TemplateApi
	default:
		return TemplateTpl
	}
}
func GetTplPath(mode string) string {
	switch mode {
	case "rpc":
		return rpcPath
	case "api":
		return apiPath
	default:
		return rpcPath
	}

}
