package utils

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io"
	"log"
	"os"
	"path/filepath"
)

// PathExists 文件目录是否存在
// @param: path string
// @return: bool, error
func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CreateDir 批量创建文件夹
// @param: dirs ...string
// @return: err error
func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			//log.Println("create directory" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				log.Println("create directory"+v, zap.Any(" error:", err))
				return err
			}
		}
	}
	return err
}

// IsNotExistMkDir 检查文件夹是否存在
// 如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if exist := !CheckExist(src); !exist {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// CheckExist 检查文件是否存在
func CheckExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// MkDir 新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func CopyDir(src, dst string) error {
	fileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !fileInfo.IsDir() {
		return fmt.Errorf("%s is not a directory", src)
	}

	// 创建目标目录
	err = os.MkdirAll(dst, fileInfo.Mode())
	if err != nil {
		return err
	}

	// 获取源目录下的所有文件和子目录
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	// 递归复制每个文件和子目录
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

func FileWriter(file string) (io.WriteCloser, error) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	return f, err
}
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
