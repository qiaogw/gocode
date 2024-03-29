package util

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
)

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
