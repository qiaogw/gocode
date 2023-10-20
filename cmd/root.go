package cmd

import (
	_ "embed"
	"fmt"
	"github.com/qiaogw/gocode/cmd/backup"
	"github.com/qiaogw/gocode/cmd/restore"

	"github.com/qiaogw/gocode/cmd/gen"
	"github.com/qiaogw/gocode/cmd/inital"
	"github.com/qiaogw/gocode/cmd/upgrade"
	"github.com/qiaogw/gocode/global"
	"github.com/spf13/cobra"

	"os"
	"runtime"
	"strings"
	"text/template"
)

const (
	codeFailure = 1
	dash        = "-"
	doubleDash  = "--"
	assign      = "="
)

var (
	//go:embed usage.tpl
	usageTpl string

	rootCmd = &cobra.Command{
		Use:   "gocode",
		Short: "生成 go-zero 代码工具",
		Long:  "代码生产工具，根据数据库生成 api, zrpc, model code",
	}
)

// Execute executes the given command
func Execute() {
	os.Args = supportGoStdFlag(os.Args)
	if err := rootCmd.Execute(); err != nil {
		//log.Println(aurora.Red(err.Error()))
		os.Exit(codeFailure)
	}
}

func init() {
	// log.SetFlags(log.Flags() | log.Llongfile)
}
func supportGoStdFlag(args []string) []string {
	copyArgs := append([]string(nil), args...)
	parentCmd, _, err := rootCmd.Traverse(args[:1])
	if err != nil { // ignore it to let cobra handle the error.
		return copyArgs
	}

	for idx, arg := range copyArgs[0:] {
		parentCmd, _, err = parentCmd.Traverse([]string{arg})
		if err != nil { // ignore it to let cobra handle the error.
			break
		}
		if !strings.HasPrefix(arg, dash) {
			continue
		}

		flagExpr := strings.TrimPrefix(arg, doubleDash)
		flagExpr = strings.TrimPrefix(flagExpr, dash)
		flagName, flagValue := flagExpr, ""
		assignIndex := strings.Index(flagExpr, assign)
		if assignIndex > 0 {
			flagName = flagExpr[:assignIndex]
			flagValue = flagExpr[assignIndex:]
		}

		if !isBuiltin(flagName) {
			// The method Flag can only match the user custom flags.
			f := parentCmd.Flag(flagName)
			if f == nil {
				continue
			}
			if f.Shorthand == flagName {
				continue
			}
		}

		goStyleFlag := doubleDash + flagName
		if assignIndex > 0 {
			goStyleFlag += flagValue
		}

		copyArgs[idx] = goStyleFlag
	}
	return copyArgs
}

func isBuiltin(name string) bool {
	return name == "version" || name == "help"
}

func init() {
	cobra.AddTemplateFuncs(template.FuncMap{
		"blue":    blue,
		"green":   green,
		"rpadx":   rpadx,
		"rainbow": rainbow,
		"red":     red,
	})
	rootCmd.Version = fmt.Sprintf(
		"%s %s/%s", global.BuildVersion,
		runtime.GOOS, runtime.GOARCH)

	rootCmd.SetUsageTemplate(usageTpl)
	rootCmd.AddCommand(gen.Cmd)
	rootCmd.AddCommand(inital.Cmd)
	rootCmd.AddCommand(upgrade.Cmd)
	rootCmd.AddCommand(backup.Cmd)
	rootCmd.AddCommand(restore.Cmd)
}
