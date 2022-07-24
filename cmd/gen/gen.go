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
	prefix     string
	Cmd        = &cobra.Command{
		Use:          "gen",
		Short:        "生成代码",
		Example:      "gocode gen -c  config.yaml",
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
	pack := "app-service"
	Cmd.PersistentFlags().StringVarP(&configYml, "config", "c", configFile, "配置文件")
	Cmd.PersistentFlags().StringVarP(&apiPackage, "package", "p", pack, "生成包名")
}

func setup() {
	// 读取配置
	global.GenViper = setting.Viper(configYml, apiPackage)
	global.GenDB = setting.Gorm()
	log.SetFlags(log.Flags() | log.Llongfile)
}

func run() error {
	//var caser = cases.Title(language.English)
	fmt.Println(`start gen `, configYml)

	genApp := gen.AutoCodeServiceApp
	genApp.Init()
	tables, err := genApp.DB.GetTables(global.GenConfig.DB.Dbname)
	if err != nil {
		log.Fatal(err)
		return err
	}
	var db model.Db
	db.Database = global.GenConfig.System.Name
	db.Package = strings.ToLower(db.Database)
	db.Service = util.LeftUpper(db.Database)
	matchTables := make(map[string]*model.Table)
	db.Option = global.GenConfig
	db.DriverName = global.GenDB.Name()

	for _, v := range tables {
		if !strings.HasPrefix(v.Table, global.GenConfig.DB.TablePrefix) {
			continue
		}
		columnData, err := genApp.DB.GetColumn(global.GenConfig.DB.Dbname, v.Table)
		if err != nil {
			log.Printf("GetColumn err is %v\n", err)
			continue
		}
		tb, err := columnData.Convert(v.TableComment)
		if err != nil {
			log.Printf("Convert err is %v\n", err)
			continue
		}
		if tb.HasTimer {
			db.HasTimer = true
		}
		matchTables[v.Table] = tb
		//fmt.Printf("tb is %+v\n", v)
		err = genApp.CreateModel(tb)
		if err != nil {
			continue
		}
		db.Tables = append(db.Tables, tb)
		db.GitEmail = tb.GitEmail
		db.GitUser = tb.GitUser
		//fmt.Printf("err is %v\n", err)
	}
	dir, _ := os.Getwd()
	parentPkg, err := golang.GetParentPackage(dir)
	if err != nil {
		return err
	}
	db.ParentPkg = parentPkg + "/" + global.GenConfig.AutoCode.Pkg
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
	return err
}
