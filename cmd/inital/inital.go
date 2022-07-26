package inital

import (
	"github.com/qiaogw/gocode/global"
	"github.com/qiaogw/gocode/inital"
	"github.com/qiaogw/gocode/util/pathx"
	"github.com/spf13/cobra"
	"log"
)

var (
	Cmd = &cobra.Command{
		Use:          "init",
		Short:        "初始化",
		Example:      "gocode init",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	// 读取配置
	name := global.GetDefaultConfigFile()
	//category := ""
	configYml := inital.ConfTpl
	err := pathx.CreateFile(name, configYml, true)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
