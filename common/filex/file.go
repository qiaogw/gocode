package filex

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
)

// CheckExist 检查文件或目录是否不存在
// 返回 true 表示目标路径不存在，返回 false 表示目标路径存在
func CheckExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

// MkDir 新建文件夹
// 使用 os.MkdirAll 创建目标路径中所有不存在的文件夹，权限为 os.ModePerm
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// IsNotExistMkDir 检查文件夹是否存在
// 如果目标文件夹不存在，则新建该文件夹
func IsNotExistMkDir(src string) error {
	if exist := !CheckExist(src); !exist {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// SelfPath 返回当前程序的绝对路径
func SelfPath() string {
	var selfPath string
	// 如果 selfPath 为空，则获取 os.Args[0] 对应的绝对路径
	if len(selfPath) == 0 {
		selfPath, _ = filepath.Abs(os.Args[0])
	}
	return selfPath
}

// SelfDir 返回当前程序所在的目录
func SelfDir() string {
	var selfDir string
	// 如果 selfDir 为空，根据运行参数获取工作目录或根据 SelfPath 获取所在目录
	if len(selfDir) == 0 {
		if len(os.Args) > 1 {
			selfDir, _ = os.Getwd()
		} else {
			selfDir = filepath.Dir(SelfPath())
		}
	}
	return selfDir
}

// FileWriter 打开指定文件以进行写入（追加模式），如果文件不存在则创建文件
// 返回一个 io.WriteCloser 和可能出现的错误
func FileWriter(file string) (io.WriteCloser, error) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	return f, err
}

// CopyFile 复制文件
// 从源文件 src 复制内容到目标文件 dst，返回可能出现的错误
func CopyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

// CopyDir 复制目录
// 将源目录 src 中的所有文件和子目录递归复制到目标目录 dst 中，返回可能出现的错误
func CopyDir(src, dst string) error {
	// 获取源目录信息
	fileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	// 如果源路径不是目录，则返回错误
	if !fileInfo.IsDir() {
		return fmt.Errorf("%s 不是一个目录", src)
	}

	// 创建目标目录，并设置权限与源目录一致
	err = os.MkdirAll(dst, fileInfo.Mode())
	if err != nil {
		return err
	}

	// 读取源目录下所有的文件和子目录
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	// 遍历所有条目，递归复制目录或复制文件
	for _, entry := range entries {
		sourcePath := filepath.Join(src, entry.Name())
		destinationPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = CopyDir(sourcePath, destinationPath)
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(sourcePath, destinationPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

//FileMove: 文件移动供外部调用
//@param: src string, dst string(src: 源位置,绝对路径or相对路径, dst: 目标位置,绝对路径or相对路径,必须为文件夹)
//@return: err error

func FileMove(src string, dst string) (err error) {
	if dst == "" {
		return nil
	}
	src, err = filepath.Abs(src)
	if err != nil {
		log.Printf("err is %v\n", err)
		return err
	}
	dst, err = filepath.Abs(dst)
	if err != nil {
		log.Printf("err is %v\n", err)
		return err
	}
	revoke := false
	dir := filepath.Dir(dst)
Redirect:
	_, err = os.Stat(dir)
	if err != nil {
		err = os.MkdirAll(dir, 0o755)
		if err != nil {
			log.Printf("err is %v\n", err)
			return err
		}
		if !revoke {
			revoke = true
			goto Redirect
		}
	}
	return os.Rename(src, dst)
}

func DeLFile(filePath string) error {
	return os.RemoveAll(filePath)
}

//TrimSpace: 去除结构体空格
//@param: target interface (target: 目标结构体,传入必须是指针类型)
//@return: null

func TrimSpace(target interface{}) {
	t := reflect.TypeOf(target)
	if t.Kind() != reflect.Ptr {
		return
	}
	t = t.Elem()
	v := reflect.ValueOf(target).Elem()
	for i := 0; i < t.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.String:
			v.Field(i).SetString(strings.TrimSpace(v.Field(i).String()))
		}
	}
}

// FileExist 判断文件是否存在
func FileExist(path string) bool {
	fi, err := os.Lstat(path)
	if err == nil {
		return !fi.IsDir()
	}
	return !os.IsNotExist(err)
}

// FmtCode 格式化代码
func FmtCode(f string) (err error) {
	cmd := exec.Command("gofmt", "-w", f)

	var out, errBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err = cmd.Start() //如果用start则直接向后运行
	if err != nil {
		return err
	}
	go func() {
		err = cmd.Wait()
	}()
	return err
}
