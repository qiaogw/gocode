package gen

import (
	"fmt"
	"github.com/qiaogw/gocode/gen"
	"github.com/qiaogw/gocode/global"
	"github.com/qiaogw/gocode/model"
	"github.com/qiaogw/gocode/setting"
	"github.com/qiaogw/gocode/util"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/golang"
	"log"
	"os"
	"strings"
)

var (
	configYml  string
	apiPackage string
	Cmd        = &cobra.Command{
		Use:          "gen",
		Short:        "生成代码",
		Example:      "gocode gen -p admin -c  config.yaml",
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
	configFile := global.GetDefaultConfigFile()
	pack := "service"
	Cmd.PersistentFlags().StringVarP(&apiPackage, "package", "p", pack, "生成包名")
	Cmd.PersistentFlags().StringVarP(&configYml, "config", "c", configFile, "配置文件")
}

func setup() error {
	// 读取配置
	log.SetFlags(log.Flags() | log.Llongfile)
	global.GenViper = setting.Viper(configYml, apiPackage)
	ed, err := setting.GormInit()
	if err != nil {
		return err
	}

	global.GenDB = ed

	return nil
}

func run() error {
	//var caser = cases.Title(language.English)
	fmt.Println(util.Green(`start gen ` + configYml))
	genApp := gen.AutoCodeServiceApp
	genApp.Init()
	fmt.Printf(util.Green(fmt.Sprintf("数据库连接成功，类型为：%s,地址为：%s:%v,数据库为：%s\n",
		global.GenDB.Name(), global.GenConfig.DB.Path, global.GenConfig.DB.Port, global.GenConfig.DB.Dbname)))
	tables, err := genApp.DB.GetTables(global.GenConfig.DB.Dbname)
	if err != nil {
		log.Println(util.Red(fmt.Sprintf("获取表 err is %v", err)))
		return err
	}
	var db model.Db
	db.Database = global.GenConfig.System.Name
	db.Package = strings.ToLower(db.Database)
	db.Service = util.LeftUpper(db.Database)

	db.Option = global.GenConfig
	db.DriverName = global.GenDB.Name()
	dir, _ := os.Getwd()
	parentPkg, err := golang.GetParentPackage(dir)
	if err != nil {
		return err
	}
	db.ParentPkg = parentPkg + "/" + global.GenConfig.AutoCode.Pkg

	for _, v := range tables {
		if !strings.HasPrefix(v.Table, global.GenConfig.DB.TablePrefix) {
			continue
		}
		columnData, err := genApp.DB.GetColumn(global.GenConfig.DB.Dbname, v.Table)
		if err != nil {
			log.Println(util.Red(fmt.Sprintf("获取字段 err is %v", err)))
			continue
		}
		//log.Println("columnData len is ", len(columnData.Columns))
		tb, err := columnData.Convert(v.TableComment)
		if err != nil {
			fmt.Println(util.Red(fmt.Sprintf("数据生成错误错误: %v", err)))
			continue
		}
		if tb.HasTimer {
			db.HasTimer = true
		}
		tb.ParentPkg = db.ParentPkg
		//fmt.Printf("tb is %+v\n", v)
		err = genApp.CreateModel(tb)
		if err != nil {
			continue
		}
		//log.Printf("tb is %s,CacheKeys is  %+v\n", tb.Name, tb.CacheKeys)
		db.Tables = append(db.Tables, tb)
		db.GitEmail = tb.GitEmail
		db.GitUser = tb.GitUser
		//fmt.Printf("err is %v\n", err)
	}

	err = genApp.CreateRpc(&db)
	if err != nil {
		log.Printf("CreateRpc err is %v\n", err)
		return err
	}
	err = genApp.CreateApi(&db)
	if err != nil {
		log.Printf("CreateApi err is %v\n", err)
		//return err
	}
	err = genApp.CreateRpcLogic(&db)
	if err != nil {
		log.Printf("CreateRpcLogic err is %v\n", err)
		return err
	}
	err = genApp.CreateCommon(&db)
	if err != nil {
		log.Printf("CreateCommon err is %v\n", err)
		return err
	}
	fmt.Println(util.Green("Done!"))
	return err
}
