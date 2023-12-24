package gen

import (
	"fmt"
	"github.com/qiaogw/gocode/gen"
	"github.com/qiaogw/gocode/global"
	"github.com/qiaogw/gocode/setting"
	utils2 "github.com/qiaogw/gocode/util"
	"github.com/spf13/cobra"
)

var (
	apiPackage string
	modeGen    string
	Cmd        = &cobra.Command{
		Use:          "gen",
		Short:        "生成代码",
		Example:      "gocode gen -p admin -m rpc",
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
	}
	_, _, err := genApp.Code(m)
	return err
}
