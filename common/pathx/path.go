package pathx

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	// pkgSep 定义包路径分隔符
	pkgSep = "/"
	// goModeIdentifier 定义 Go 模块文件的标识，即 "go.mod"
	goModeIdentifier = "go.mod"
)

// JoinPackages 使用 pkgSep 将多个包路径连接起来并返回结果
func JoinPackages(pkgs ...string) string {
	return strings.Join(pkgs, pkgSep)
}

// MkdirIfNotExist 如果指定的目录不存在，则创建该目录；否则不做任何操作
func MkdirIfNotExist(dir string) error {
	if len(dir) == 0 {
		return nil
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}

	return nil
}

// PathFromGoSrc 返回当前工作目录相对于 $GOPATH/src/{projName} 之后的路径部分
// 如果当前目录不在 $GOPATH/src/{projName} 下，则返回错误
func PathFromGoSrc(projName string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	gopath := os.Getenv("GOPATH")
	parent := path.Join(gopath, "src", projName)
	pos := strings.Index(dir, parent)
	if pos < 0 {
		return "", fmt.Errorf("%s is not in GOPATH", dir)
	}

	// 返回 parent 后面的部分（跳过一个斜杠）
	return dir[len(parent)+1:], nil
}

// FindGoModPath 查找给定目录中包含 go.mod 文件的相对路径
// 如果在递归查找过程中找到 go.mod 文件，则返回该目录相对于传入目录的相对路径以及 true，否则返回空字符串和 false
func FindGoModPath(dir string) (string, bool) {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return "", false
	}

	// 将 Windows 下的反斜杠替换为正斜杠，保证统一格式
	absDir = strings.ReplaceAll(absDir, `\`, `/`)
	var rootPath string
	tempPath := absDir
	hasGoMod := false
	for {
		if FileExists(filepath.Join(tempPath, goModeIdentifier)) {
			// 计算相对路径，去除前导斜杠
			rootPath = strings.TrimPrefix(absDir[len(tempPath):], "/")
			hasGoMod = true
			break
		}

		// 如果已到达文件系统根目录则退出循环
		if tempPath == filepath.Dir(tempPath) {
			break
		}

		tempPath = filepath.Dir(tempPath)
		if tempPath == string(filepath.Separator) {
			break
		}
	}
	if hasGoMod {
		return rootPath, true
	}
	return "", false
}

// FindProjectPath 查找项目根目录，即包含 go.mod 文件的父级目录
// 参数 loc 可以是绝对路径或相对路径，返回项目根目录及 true，如果未找到则返回空字符串和 false
func FindProjectPath(loc string) (string, bool) {
	var dir string
	// 如果 loc 是以 "/" 开头，则视为绝对路径；否则与当前工作目录拼接
	if strings.IndexByte(loc, '/') == 0 {
		dir = loc
	} else {
		wd, err := os.Getwd()
		if err != nil {
			return "", false
		}
		dir = filepath.Join(wd, loc)
	}

	// 向上遍历目录，直到找到包含 go.mod 文件的目录或到达根目录
	for {
		if FileExists(filepath.Join(dir, goModeIdentifier)) {
			return dir, true
		}
		dir = filepath.Dir(dir)
		if dir == "/" {
			break
		}
	}

	return "", false
}

// isLink 判断给定路径是否为符号链接
func isLink(name string) (bool, error) {
	fi, err := os.Lstat(name)
	if err != nil {
		return false, err
	}
	return fi.Mode()&os.ModeSymlink != 0, nil
}
