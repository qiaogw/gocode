package utils

import (
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// gormSourceDir 保存 gorm 源码目录的路径
var gormSourceDir string

// gormZeroSourceDir 保存 gorm-zero 源码目录的路径
var gormZeroSourceDir string

// init 初始化函数，通过 runtime.Caller 获取当前文件路径，并确定 gorm 和 gorm-zero 的源码目录
func init() {
	_, file, _, _ := runtime.Caller(0)
	// 兼容各种操作系统的方案，获取 gorm 源码目录
	gormSourceDir = gormSourceDirPath(file)
	// 兼容各种操作系统的方案，获取 gorm-zero 源码目录
	gormZeroSourceDir = gormZeroSourceDirPath(file)
}

// gormSourceDirPath 根据传入的文件路径返回 gorm 的源码目录路径
func gormSourceDirPath(file string) string {
	dir := filepath.Dir(file)
	dir = filepath.Dir(dir)

	s := filepath.Dir(dir)
	if filepath.Base(s) != "gorm.io" {
		s = dir
	}
	// 使用斜杠格式返回目录路径，并在末尾加上斜杠
	return filepath.ToSlash(s) + "/"
}

// gormZeroSourceDirPath 根据传入的文件路径返回 gorm-zero 的源码目录路径
func gormZeroSourceDirPath(file string) string {
	dir := filepath.Dir(file)
	dir = filepath.Dir(dir)

	s := filepath.Dir(dir)
	if filepath.Base(s) != "gorm-zero" {
		s = dir
	}
	// 使用斜杠格式返回目录路径，并在末尾加上斜杠
	return filepath.ToSlash(s) + "/"
}

// FileWithLineNum 返回当前执行代码的文件名及行号
func FileWithLineNum() string {
	// 从调用栈中查找，通常第2层调用者来自 gorm 内部，因此从索引 2 开始查找
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		inGorm := strings.Contains(file, "gorm.io")
		inGormZero := strings.HasPrefix(file, gormZeroSourceDir)
		// 如果获取成功，且文件不属于 gorm 或 gorm-zero，或文件名以 _test.go 结尾，则返回文件名和行号
		if ok && (!(inGorm || inGormZero) || strings.HasSuffix(file, "_test.go")) {
			return file + ":" + strconv.FormatInt(int64(line), 10)
		}
	}

	return ""
}
