package upgrade

import (
	"fmt"
	"github.com/qiaogw/gocode/global"
	"github.com/qiaogw/gocode/util"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/execx"
)

// Cmd describes a upgrade command.
var Cmd = &cobra.Command{
	Use:   "upgrade",
	Short: "升级至gocode最新版本",
	RunE:  upgrade,
}
var versionID string

func init() {
	ver := "latest"
	Cmd.PersistentFlags().StringVarP(&versionID, "version", "v", ver, "版本号")
}

// Upgrade gets the latest gocode by
// go install github.com/qiaogw/gocode@latest
func upgrade(_ *cobra.Command, _ []string) error {
	currentVersion := fmt.Sprintf(
		"%s %s/%s", global.BuildVersion,
		runtime.GOOS, runtime.GOARCH)
	appCmd := fmt.Sprintf("GOPROXY=https://goproxy.cn/,direct go install github.com/qiaogw/gocode@%s", versionID)
	cmd := `GO111MODULE=on ` + appCmd
	if runtime.GOOS == "windows" {
		cmd = `set ` + appCmd
	}
	fmt.Println(util.Green(fmt.Sprintf("版本升级开始, 当前版本%s", currentVersion)))
	_, err := execx.Run(cmd, "")
	if err != nil {
		fmt.Println(util.Red(fmt.Sprintf("版本升级失败：%v", err)))
		return err
	}

	fmt.Println(util.Green(fmt.Sprintf("已升级到最新版本%s", versionID)))
	return nil
}
