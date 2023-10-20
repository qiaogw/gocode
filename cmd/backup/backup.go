package backup

import (
	"fmt"
	"github.com/qiaogw/gocode/gen"
	"github.com/qiaogw/gocode/global"
	"github.com/qiaogw/gocode/pkg/dbtools"
	"github.com/qiaogw/gocode/setting"
	"github.com/qiaogw/gocode/util"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

var (
	apiPackage string
	Cmd        = &cobra.Command{
		Use:          "backup",
		Short:        "备份数据库",
		Example:      "gocode backup -p admin",
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
	pack := "admin"
	Cmd.PersistentFlags().StringVarP(&apiPackage, "package", "p", pack, "包名")
	_ = Cmd.MarkPersistentFlagRequired("package")
}

func setup() error {
	confFile := filepath.Join("dbconf", apiPackage)
	configYml := global.GetConfigFile(confFile)
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
	wd, _ := os.Getwd()
	backupDir := filepath.Join(wd, "backup", global.GenDB.Name())
	fmt.Println(util.Green(`开始备份数据库 ` + apiPackage + ` 到 ` + backupDir))
	genApp := gen.AutoCodeServiceApp
	genApp.Init()
	fmt.Printf(util.Green(fmt.Sprintf("数据库连接成功，类型为：%s,地址为：%s:%v,数据库为：%s\n",
		global.GenDB.Name(), global.GenConfig.DB.Path, global.GenConfig.DB.Port, global.GenConfig.DB.Dbname)))
	//err := dbtools.Backup(global.GenDB, "backup")
	err := dbtools.BackupDB("backup")
	if err != nil {
		log.Println(util.Red(fmt.Sprintf("备份失败：%v", err)))
	}
	fmt.Println(util.Green("Done!"))
	return err
}
