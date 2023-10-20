package restore

import (
	"fmt"
	"github.com/qiaogw/gocode/gen"
	"github.com/qiaogw/gocode/global"
	"github.com/qiaogw/gocode/pkg/dbtools"
	"github.com/qiaogw/gocode/setting"
	"github.com/qiaogw/gocode/utils"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

var (
	src, dst string
	Cmd      = &cobra.Command{
		Use:          "restore",
		Short:        "恢复数据库",
		Example:      "gocode restore -s zero -d admin ",
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
	Cmd.PersistentFlags().StringVarP(&src, "src", "s", pack, "数据源")
	_ = Cmd.MarkPersistentFlagRequired("src")
	Cmd.PersistentFlags().StringVarP(&dst, "dst", "d", pack, "目标数据库")
	_ = Cmd.MarkPersistentFlagRequired("dst")
}

func setup() error {
	dst = filepath.Join("dbconf", dst)
	configYml := global.GetConfigFile(dst)
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
	fmt.Println(utils.Green(`开始恢复 ` + dst))
	wd, _ := os.Getwd()
	backupDir := filepath.Join(wd, "backup", src)
	fmt.Println(utils.Green(`开始恢复 ` + global.GenConfig.DB.Dbname + ` 从 ` + backupDir))
	genApp := gen.AutoCodeServiceApp
	genApp.Init()
	//err := dbtools.Backup(global.GenDB, "backup")
	err := dbtools.RestoreData(filepath.Join("backup", src))
	if err != nil {
		log.Println(utils.Red(fmt.Sprintf("恢复失败：%v", err)))
	}
	fmt.Println(utils.Green("Done!"))
	return err
}
