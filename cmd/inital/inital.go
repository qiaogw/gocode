package inital

import (
	"fmt"
	"github.com/qiaogw/gocode/common/pathx"
	"github.com/qiaogw/gocode/global"
	"github.com/qiaogw/gocode/inital"
	"github.com/qiaogw/gocode/util"
	"github.com/spf13/cobra"
	"log"
)

var (
	Cmd = &cobra.Command{
		Use:          "init",
		Short:        "初始化",
		Example:      "gocode init",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	// pack := "service"
	// Cmd.PersistentFlags().StringVarP(&apiPackage, "package", "p", pack, "生成包名")
	// _ = Cmd.MarkPersistentFlagRequired("package")
}

func run() {
	// 读取配置
	name := global.GetDefaultConfigFile()
	// if len(apiPackage) < 1 {
	// 	fmt.Println(util.Red("缺失必须的参数：--package 或 -p 以定义应用包名."))
	// 	return
	// }
	configYml := inital.ConfTpl
	if util.FileExist(name) {
		fmt.Println(util.Red("配置文件" + name + "已存在，请删除后重新初始化."))
		return
	}
	err := pathx.CreateFile(name, configYml, true)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
