package gen

import (
	"github.com/spf13/cobra"
	"gocode/config"
	"gocode/global"
)

var (
	configYml string
	apiCheck  bool
	Cmd       = &cobra.Command{
		Use:          "gen",
		Short:        "生成代码",
		Example:      "emanager start -c config/settings.yml",
		SilenceUsage: true,
	}
)

func init() {
	configFile := global.GetDefaultConfigFile()
	Cmd.PersistentFlags().StringVarP(&configYml, "config", "c", configFile, "配置文件")
}

func setup() {
	// 读取配置
	global.GenViper = config.Viper(configYml)
	global.GenDB = config.Gorm()
}
