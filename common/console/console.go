package console

import (
	"fmt"
	"os"
	"runtime"

	"github.com/logrusorgru/aurora"
)

// Console 定义了控制台输出接口，提供彩色输出以及带前缀的 IDEA 输出
type (
	// Console 接口封装了格式化输出的方法，
	// 默认实现为 colorConsole，用于在控制台输出彩色日志，
	// 以及 ideaConsole，用于 IntelliJ 插件下带前缀的输出。
	Console interface {
		Success(format string, a ...interface{}) // 成功信息输出
		Info(format string, a ...interface{})    // 信息输出
		Debug(format string, a ...interface{})   // 调试信息输出
		Warning(format string, a ...interface{}) // 警告信息输出
		Error(format string, a ...interface{})   // 错误信息输出
		Fatalln(format string, a ...interface{}) // 致命错误输出并退出程序
		MarkDone()                               // 标记完成
		Must(err error)                          // 错误检查，若 err 不为 nil，则输出错误并退出程序
	}

	// colorConsole 是 Console 接口的默认实现，支持彩色输出
	colorConsole struct {
		enable bool // 是否启用日志输出
	}

	// ideaConsole 用于 IDEA 日志输出，添加了固定的前缀
	ideaConsole struct{}
)

// NewConsole 根据传入的标志返回一个 Console 实例，idea 为 true 则返回 ideaConsole，否则返回 colorConsole
func NewConsole(idea bool) Console {
	if idea {
		return NewIdeaConsole()
	}
	return NewColorConsole()
}

// NewColorConsole 返回一个 colorConsole 实例，可通过可选参数设置是否启用日志输出（默认为 true）
func NewColorConsole(enable ...bool) Console {
	logEnable := true
	for _, e := range enable {
		logEnable = e
	}
	return &colorConsole{
		enable: logEnable,
	}
}

// Info 输出普通信息，当日志输出被启用时才会输出
func (c *colorConsole) Info(format string, a ...interface{}) {
	if !c.enable {
		return
	}
	msg := fmt.Sprintf(format, a...)
	fmt.Println(msg)
}

// Debug 输出调试信息，当日志输出被启用时以亮青色打印调试信息
func (c *colorConsole) Debug(format string, a ...interface{}) {
	if !c.enable {
		return
	}
	msg := fmt.Sprintf(format, a...)
	printlin(aurora.BrightCyan(msg))
}

// Success 输出成功信息，当日志输出被启用时以亮绿色打印成功信息
func (c *colorConsole) Success(format string, a ...interface{}) {
	if !c.enable {
		return
	}
	msg := fmt.Sprintf(format, a...)
	printlin(aurora.BrightGreen(msg))
}

// Warning 输出警告信息，当日志输出被启用时以亮黄色打印警告信息
func (c *colorConsole) Warning(format string, a ...interface{}) {
	if !c.enable {
		return
	}
	msg := fmt.Sprintf(format, a...)
	printlin(aurora.BrightYellow(msg))
}

// Error 输出错误信息，当日志输出被启用时以亮红色打印错误信息
func (c *colorConsole) Error(format string, a ...interface{}) {
	if !c.enable {
		return
	}
	msg := fmt.Sprintf(format, a...)
	printlin(aurora.BrightRed(msg))
}

// Fatalln 输出致命错误信息，并退出程序
func (c *colorConsole) Fatalln(format string, a ...interface{}) {
	if !c.enable {
		return
	}
	c.Error(format, a...)
	os.Exit(1)
}

// MarkDone 输出 "Done." 提示信息，表示操作完成
func (c *colorConsole) MarkDone() {
	if !c.enable {
		return
	}
	c.Success("Done.")
}

// Must 检查错误，如果 err 不为 nil，则调用 Fatalln 输出错误并退出程序
func (c *colorConsole) Must(err error) {
	if !c.enable {
		return
	}
	if err != nil {
		c.Fatalln("%+v", err)
	}
}

// NewIdeaConsole 返回一个 ideaConsole 实例
func NewIdeaConsole() Console {
	return &ideaConsole{}
}

// Info 输出普通信息（IDEA 模式）
func (i *ideaConsole) Info(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Println(msg)
}

// Debug 输出调试信息（IDEA 模式），以亮青色打印
func (i *ideaConsole) Debug(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Println(aurora.BrightCyan(msg))
}

// Success 输出成功信息（IDEA 模式），前缀为 [SUCCESS]:
func (i *ideaConsole) Success(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Println("[SUCCESS]: ", msg)
}

// Warning 输出警告信息（IDEA 模式），前缀为 [WARNING]:
func (i *ideaConsole) Warning(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Println("[WARNING]: ", msg)
}

// Error 输出错误信息（IDEA 模式），前缀为 [ERROR]:
func (i *ideaConsole) Error(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Println("[ERROR]: ", msg)
}

// Fatalln 输出致命错误信息（IDEA 模式），并退出程序
func (i *ideaConsole) Fatalln(format string, a ...interface{}) {
	i.Error(format, a...)
	os.Exit(1)
}

// MarkDone 输出 "Done." 提示信息（IDEA 模式）
func (i *ideaConsole) MarkDone() {
	i.Success("Done.")
}

// Must 检查错误，如果 err 不为 nil，则输出错误并退出程序（IDEA 模式）
func (i *ideaConsole) Must(err error) {
	if err != nil {
		i.Fatalln("%+v", err)
	}
}

// printlin 根据不同操作系统和输出类型打印消息，解决 Windows 下彩色输出问题
func printlin(msg interface{}) {
	value, ok := msg.(aurora.Value)
	if !ok {
		fmt.Println(msg)
	}

	goos := runtime.GOOS
	if goos == "windows" {
		fmt.Println(value.Value())
		return
	}

	fmt.Println(msg)
}

// defaultConsole 为默认的控制台实例，使用 colorConsole 并启用日志输出
var defaultConsole = &colorConsole{enable: true}

// Success 调用默认控制台的 Success 方法
func Success(format string, a ...interface{}) {
	defaultConsole.Success(format, a...)
}

// Info 调用默认控制台的 Info 方法
func Info(format string, a ...interface{}) {
	defaultConsole.Info(format, a...)
}

// Debug 调用默认控制台的 Debug 方法
func Debug(format string, a ...interface{}) {
	defaultConsole.Debug(format, a...)
}

// Warning 调用默认控制台的 Warning 方法
func Warning(format string, a ...interface{}) {
	defaultConsole.Warning(format, a...)
}

// Error 调用默认控制台的 Error 方法
func Error(format string, a ...interface{}) {
	defaultConsole.Error(format, a...)
}

// Fatalln 调用默认控制台的 Fatalln 方法
func Fatalln(format string, a ...interface{}) {
	defaultConsole.Fatalln(format, a...)
}

// MarkDone 调用默认控制台的 MarkDone 方法
func MarkDone() {
	defaultConsole.MarkDone()
}

// Must 调用默认控制台的 Must 方法
func Must(err error) {
	defaultConsole.Must(err)
}
