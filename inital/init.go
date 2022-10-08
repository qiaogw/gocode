package inital

import (
	"embed"
)

var (
	//go:embed config.yaml
	ConfTpl string
	//go:embed template
	TemplateTpl embed.FS
)
