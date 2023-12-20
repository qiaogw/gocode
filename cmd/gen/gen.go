package gen

import (
	"fmt"
	"github.com/qiaogw/gocode/gen"
	"github.com/qiaogw/gocode/global"
	"github.com/qiaogw/gocode/schema"
	"github.com/qiaogw/gocode/setting"
	utils2 "github.com/qiaogw/gocode/util"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var (
	apiPackage string
	modeGen    string
	Cmd        = &cobra.Command{
		Use:          "gen",
		Short:        "生成代码",
		Example:      "gocode gen -p admin -m zero",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	pack := "config"
	Cmd.PersistentFlags().StringVarP(&apiPackage, "package", "p", pack, "包名")
	_ = Cmd.MarkPersistentFlagRequired("package")
	mode := "rpc"
	Cmd.PersistentFlags().StringVarP(&modeGen, "mode", "m", mode, "模式(rpc、api)")
	_ = Cmd.MarkPersistentFlagRequired("mode")

}

func setup() error {
	configYml := global.GetConfigFile(apiPackage)
	// 读取配置
	global.GenViper = setting.Viper(configYml)
	ed, err := setting.GormInit()
	if err != nil {
		return err
	}
	global.GenDB = ed
	return nil
}

func run() error {
	fmt.Println(utils2.Green(`start gen ` + apiPackage))
	genApp := gen.AutoCodeServiceApp
	genApp.Init()
	m := true
	genApp.Mode = "rpc"
	if modeGen == "api" {
		genApp.Mode = "api"
		schema.NewAdmin(global.GenDB)
		p, _ := os.Getwd()
		global.GenConfig.AutoCode.Pkg = "sub-admin"
		global.GenConfig.AutoCode.Root = filepath.Join(p, global.GenConfig.AutoCode.Pkg)
	}

	//genApp.Mode = "api"
	_, _, err := genApp.Code(m)
	return err
}
