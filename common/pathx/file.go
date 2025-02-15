package pathx

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/logrusorgru/aurora"
)

// 常量定义
const (
	NL              = "\n"             // 换行符
	goctlDir        = ".goctl"         // goctl 目录名称
	gitDir          = ".git"           // git 目录名称
	autoCompleteDir = ".auto_complete" // 自动补全目录名称
	cacheDir        = "cache"          // 缓存目录名称
)

var goctlHome string // 全局变量，用于存储 goctl 的主页路径

// RegisterGoctlHome 注册 goctl 的主页路径
func RegisterGoctlHome(home string) {
	goctlHome = home
}

// CreateIfNotExist 如果文件不存在，则创建该文件；如果文件已存在，则返回错误
func CreateIfNotExist(file string) (*os.File, error) {
	_, err := os.Stat(file)
	if !os.IsNotExist(err) {
		return nil, fmt.Errorf("%s already exist", file)
	}
	return os.Create(file)
}

// RemoveIfExist 如果指定文件存在，则删除该文件
func RemoveIfExist(filename string) error {
	if !FileExists(filename) {
		return nil
	}
	return os.Remove(filename)
}

// RemoveOrQuit 如果指定文件存在，则提示用户输入确认信息，确认后删除该文件；否则直接返回
func RemoveOrQuit(filename string) error {
	if !FileExists(filename) {
		return nil
	}
	fmt.Printf("%s exists, overwrite it?\nEnter to overwrite or Ctrl-C to cancel...",
		aurora.BgRed(aurora.Bold(filename)))
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	return os.Remove(filename)
}

// FileExists 判断指定的文件是否存在，存在则返回 true，否则返回 false
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

// FileNameWithoutExt 返回文件名去除扩展名后的部分
func FileNameWithoutExt(file string) string {
	return strings.TrimSuffix(file, filepath.Ext(file))
}

// GetGoctlHome 返回 goctl 的路径
// 默认路径为 ~/.goctl，如果通过 RegisterGoctlHome 设置了用户自定义路径，则返回该路径
// 如果返回的路径存在且不是目录，则将其重命名为旧路径，并创建新的目录
func GetGoctlHome() (home string, err error) {
	defer func() {
		if err != nil {
			return
		}
		info, err := os.Stat(home)
		if err == nil && !info.IsDir() {
			os.Rename(home, home+".old")
			MkdirIfNotExist(home)
		}
	}()
	if len(goctlHome) != 0 {
		home = goctlHome
		return
	}
	home, err = GetDefaultGoctlHome()
	return
}

// GetDefaultGoctlHome 返回默认的 goctl 主页路径，即将 $HOME 与 .goctl 拼接后的结果
func GetDefaultGoctlHome() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, goctlDir), nil
}

// GetGitHome 返回 goctl 下的 git 目录路径
func GetGitHome() (string, error) {
	goctlH, err := GetGoctlHome()
	if err != nil {
		return "", err
	}
	return filepath.Join(goctlH, gitDir), nil
}

// GetAutoCompleteHome 返回 goctl 下的自动补全目录路径
func GetAutoCompleteHome() (string, error) {
	goctlH, err := GetGoctlHome()
	if err != nil {
		return "", err
	}
	return filepath.Join(goctlH, autoCompleteDir), nil
}

// GetCacheDir 返回 goctl 下的缓存目录路径
func GetCacheDir() (string, error) {
	goctlH, err := GetGoctlHome()
	if err != nil {
		return "", err
	}
	return filepath.Join(goctlH, cacheDir), nil
}

// SameFile 判断两个路径是否指向同一个文件
// 该函数考虑了文件系统大小写不敏感的情况，如 macOS 和 Windows 下的路径差异
func SameFile(path1, path2 string) (bool, error) {
	stat1, err := os.Stat(path1)
	if err != nil {
		return false, err
	}
	stat2, err := os.Stat(path2)
	if err != nil {
		return false, err
	}
	return os.SameFile(stat1, stat2), nil
}

// CreateFile 创建文件并写入内容
// 参数 file 为文件路径，content 为要写入的内容，force 为是否强制覆盖（若文件存在时）
func CreateFile(file, content string, force bool) error {
	if FileExists(file) && !force {
		return nil
	}
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	return err
}

// MustTempDir 创建一个临时目录，创建失败时直接退出程序
func MustTempDir() string {
	dir, err := os.MkdirTemp("", "")
	if err != nil {
		log.Fatalln(err)
	}
	return dir
}

// Copy 将源文件 src 复制到目标文件 dest
func Copy(src, dest string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	dir := filepath.Dir(dest)
	err = MkdirIfNotExist(dir)
	if err != nil {
		return err
	}
	w, err := os.Create(dest)
	if err != nil {
		return err
	}
	// 设置目标文件的权限
	w.Chmod(os.ModePerm)
	defer w.Close()
	_, err = io.Copy(w, f)
	return err
}

// Hash 计算指定文件的 MD5 哈希值，并以十六进制字符串返回
func Hash(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = f.Close()
	}()
	hash := md5.New()
	_, err = io.Copy(hash, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// RenameFilesAndDirWithPrefixAndSuffix 为目录下的所有文件和子目录添加前缀和后缀
// 递归修改目录内的文件和子目录名称
func RenameFilesAndDirWithPrefixAndSuffix(dir, prefix, suffix string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		if file.IsDir() {
			// 修改子目录名称
			newDirName := prefix + file.Name() + suffix
			newDirPath := filepath.Join(dir, newDirName)
			if err := os.Rename(filePath, newDirPath); err != nil {
				return err
			}
			// 递归处理子目录下的文件和子目录
			if err := RenameFilesWithPrefixAndSuffix(newDirPath, prefix, suffix); err != nil {
				return err
			}
		} else {
			// 修改文件名
			fileExt := filepath.Ext(file.Name())
			fileBase := strings.TrimSuffix(file.Name(), fileExt)
			newName := prefix + fileBase + suffix + fileExt
			newPath := filepath.Join(dir, newName)
			if err := os.Rename(filePath, newPath); err != nil {
				return err
			}
		}
	}
	return nil
}

// RenameFilesWithPrefixAndSuffix 为指定目录下的所有文件（不递归子目录）添加前缀和后缀
func RenameFilesWithPrefixAndSuffix(dir, prefix, suffix string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		if file.IsDir() {
			// 如果是子目录，则递归处理该子目录内的文件和目录
			if err := RenameFilesWithPrefixAndSuffix(filePath, prefix, suffix); err != nil {
				return err
			}
		} else {
			// 修改文件名
			fileExt := filepath.Ext(file.Name())
			fileBase := strings.TrimSuffix(file.Name(), fileExt)
			newName := prefix + fileBase + suffix + fileExt
			newPath := filepath.Join(dir, newName)
			if err := os.Rename(filePath, newPath); err != nil {
				return err
			}
		}
	}
	return nil
}

// CountFiles 遍历指定路径，统计该路径下所有文件的数量（不统计目录）
func CountFiles(path string) (int64, error) {
	var count int64
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 如果当前项是文件，则计数加一
		if !info.IsDir() {
			count++
		}
		return nil
	})
	return count, err
}
