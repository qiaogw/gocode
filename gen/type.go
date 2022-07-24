package gen

import (
	"github.com/qiaogw/gocode/global"
	"github.com/qiaogw/gocode/model"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"text/template"
)

const (
	modelPath          = "model"
	tempPath           = "template"
	autoPath           = "autocode_template/"
	packageServiceName = "service"
	packageRouterName  = "router"
	apiPath            = "api"
	rpcPath            = "rpc"
	apiLogicPath       = "api-logic"
	rpcLogicPath       = "rpc-logic"
	commonPath         = "common"
)

type autoPackage struct {
	path string
	temp string
	name string
}

type injectionMeta struct {
	path        string
	funcName    string
	structNameF string // 带格式化的
}

type astInjectionMeta struct {
	path         string
	importCodeF  string
	structNameF  string
	packageNameF string
	groupName    string
}

type tplData struct {
	template         *template.Template
	autoPackage      string
	locationPath     string
	autoCodePath     string
	autoMoveFilePath string
}

type AutoCodeService struct {
	DB model.Model
}

var AutoCodeServiceApp = new(AutoCodeService)

func (acd *AutoCodeService) Init() {
	switch global.GenDB.Name() {
	case "mysql":
		acd.DB = model.ModelMysqlApp
		acd.DB.Init()
	case "postgres":
		acd.DB = model.ModelPostgresApp
		acd.DB.Init()
	default:
		acd.DB = model.ModelMysqlApp
		acd.DB.Init()
	}
}

var (
	packageInjectionMap map[string]astInjectionMeta
	injectionPaths      []injectionMeta
	caser               = cases.Title(language.English)
)
