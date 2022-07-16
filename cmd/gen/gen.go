package gen

import (
	"fmt"
	"github.com/spf13/cobra"
	"gocode/gen"
	"gocode/global"
	"gocode/setting"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
)

var (
	configYml  string
	apiPackage string
	Cmd        = &cobra.Command{
		Use:          "gen",
		Short:        "生成代码",
		Example:      "emanager start -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	configFile := global.GetDefaultConfigFile()
	pack := "app"
	Cmd.PersistentFlags().StringVarP(&configYml, "config", "c", configFile, "配置文件")
	Cmd.PersistentFlags().StringVarP(&apiPackage, "package", "p", pack, "生成包名")
}

func setup() {
	// 读取配置
	global.GenViper = setting.Viper(configYml, apiPackage)
	global.GenDB = setting.Gorm()
}

func run() error {
	var caser = cases.Title(language.English)
	fmt.Println(`start gen `, configYml)

	//1. 读取配置
	genApp := gen.AutoCodeServiceApp
	gendb := gen.AutoCodeMysql
	tbs, err := gendb.GetTables(global.GenConfig.DB.Dbname)
	if err != nil {
		log.Fatal(err)
		return err
	}

	//fmt.Printf("tb is %+v\n", tb)
	for _, v := range tbs {
		tb := genApp.GenTable(v)
		clms, err := gendb.GetColumn(tb.TableName)
		if err != nil {
			return err
		}
		for _, o := range clms {
			fd := genApp.GenColumn(o)
			tb.Fields = append(tb.Fields, &fd)
		}
		//fmt.Printf("tb is %+v\n", v)
		tb.Pretreatment() // 处理go关键字
		tb.PackageT = caser.String(tb.Package)
		err = genApp.CreateTemp(tb)
		if err != nil {
			fmt.Printf("err is %v\n", err)
		}
		//fmt.Printf("err is %v\n", err)
	}
	return nil
}
