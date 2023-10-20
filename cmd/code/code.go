package code

import (
	"fmt"
	"github.com/qiaogw/gocode/gen"
	"github.com/qiaogw/gocode/global"
	"github.com/qiaogw/gocode/pkg/dbtools"
	"github.com/qiaogw/gocode/pkg/utils"
	"github.com/qiaogw/gocode/setting"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/execx"
	"runtime"
)

var (
	apiPackage string
	Cmd        = &cobra.Command{
		Use:          "code",
		Short:        "生成初始化代码",
		Example:      "gocode code",
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
	apiPackage = "admin"
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
	fmt.Println(utils.Green(`start gen ` + apiPackage))
	genApp := gen.AutoCodeServiceApp
	genApp.Init()
	fmt.Printf(utils.Green(fmt.Sprintf("数据库连接成功，类型为：%s,地址为：%s:%v,数据库为：%s\n",
		global.GenDB.Name(), global.GenConfig.DB.Path, global.GenConfig.DB.Port, global.GenConfig.DB.Dbname)))
	err := dbtools.Backup(global.GenDB, "backup")
	fmt.Println(utils.Green("Done!"))
	return err
}

// Upgrade gets the latest gocode by
// go install github.com/qiaogw/gocode@latest
func upgrade(_ *cobra.Command, _ []string) error {
	cmd := `GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go install github.com/qiaogw/gocode@latest`
	if runtime.GOOS == "windows" {
		cmd = `set GOPROXY=https://goproxy.cn,direct && go install github.com/qiaogw/gocode@latest`
	}
	info, err := execx.Run(cmd, "")
	if err != nil {
		return err
	}
	fmt.Println(info, "")
	fmt.Println(utils.Green("Done!"))
	return nil
}
