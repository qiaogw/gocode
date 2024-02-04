package gen

import (
	"text/template"

	"github.com/qiaogw/gocode/global"
	"github.com/qiaogw/gocode/model"
)

const (
	modelPath       = "model"
	tempPath        = "template"
	autoPath        = "autocode_template/"
	apiPath         = "api"
	apiDescPath     = "api-desc"
	apiLogicPath    = "api-logic"
	rpcPath         = "rpc"
	initPath        = "init"
	rpcDescPath     = "rpc-desc"
	rpcLogicPath    = "rpc-logic"
	commonPath      = "common"
	internalPath    = "internal"
	svcPath         = "svc"
	configPath      = "config"
	etcPath         = "etc"
	logicPath       = "logic"
	apiResponsePath = "replay"
	webPath         = "web"
)

type tplData struct {
	template         *template.Template
	autoPackage      string
	locationPath     string
	autoCodePath     string
	autoMoveFilePath string
	tablePkg         string
}

type AutoCodeService struct {
	DB        model.Model
	Mode      string //模式(rpc、api)
	Overwrite bool   //是否覆盖
}

var AutoCodeServiceApp = new(AutoCodeService)

func (acd *AutoCodeService) Init() {
	switch global.GenDB.Name() {
	case "mysql":
		acd.DB = model.MysqlApp
		acd.DB.Init()
	case "postgres":
		acd.DB = model.PostgresApp
		acd.DB.Init()
	default:
		acd.DB = model.MysqlApp
		acd.DB.Init()
	}
}
