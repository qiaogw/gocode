package inital

import (
	"fmt"
	"github.com/qiaogw/gocode/gen"
	"github.com/qiaogw/gocode/model"
	"github.com/qiaogw/gocode/utils"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var (
	apiPackage string
	Cmd        = &cobra.Command{
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
	pack := "config"
	Cmd.PersistentFlags().StringVarP(&apiPackage, "package", "p", pack, "生成包名")
	_ = Cmd.MarkPersistentFlagRequired("package")
}

func run() {
	// 读取配置
	genApp := gen.AutoCodeServiceApp
	var db model.Db
	db.Package = strings.ToLower(apiPackage)
	err := genApp.CreateConfig(&db)

	if err != nil {
		log.Printf("CreateApi err is %v\n", err)
		//return
	}
	//configYml := inital.ConfTpl
	//if util.FileExist(name) {
	//	fmt.Println(util.Red("配置文件" + name + "已存在，请删除后重新初始化."))
	//	return
	//}
	//err := pathx.CreateFile(name, configYml, true)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	fmt.Println(utils.Green("Done! init " + apiPackage + ".yaml"))
	return
}
