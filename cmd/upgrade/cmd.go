package upgrade

import (
	"fmt"
	"github.com/qiaogw/gocode/util"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/execx"
)

// Cmd describes a upgrade command.
var Cmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade goctl to latest version",
	RunE:  upgrade,
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
	fmt.Print(util.Green("Done!"))
	return nil
}
