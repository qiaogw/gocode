package pathx

import (
	"embed"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func CopyTpl(srcFS embed.FS, src, dst string) error {
	// 目标文件夹路径
	destinationFolder := dst

	// 获取目录中的所有文件和子目录
	entries, err := srcFS.ReadDir(src)
	if err != nil {
		return fmt.Errorf("获取目录中的所有文件和子目录错误:%v", err)
	}

	// 逐个处理文件和目录
	for _, entry := range entries {
		if entry.IsDir() {
			// 如果是目录，递归创建目录结构
			src := path.Join(src, entry.Name())
			dst := filepath.Join(destinationFolder, entry.Name())
			err = CopyFsDir(srcFS, src, dst)
			if err != nil {
				return fmt.Errorf("复制目录错误:%v", err)
			}
		} else {
			// 如果是文件，复制文件
			src := path.Join(src, entry.Name())
			dst := filepath.Join(destinationFolder, entry.Name())
			err = CopyFsFile(srcFS, src, dst)
			if err != nil {
				return fmt.Errorf("复制文件错误:%v", err)
			}
		}
	}
	return nil
}

// CopyFsFile 复制fs文件
func CopyFsFile(srcFS embed.FS, src, dst string) error {
	data, err := srcFS.ReadFile(src)
	if err != nil {
		return fmt.Errorf("复制fs文件ReadFile错误:%v", err)
	}
	_ = MkdirIfNotExist(filepath.Dir(dst))
	err = os.WriteFile(dst, data, 0644)
	if err != nil {
		return fmt.Errorf("复制fs文件WriteFile错误:%v", err)
	}

	return nil
}

// CopyFsDir 递归复制目录
func CopyFsDir(srcFS embed.FS, src, dst string) error {
	srcInfo, err := srcFS.ReadDir(src)
	if err != nil {

		return fmt.Errorf("递归复制目录读取错误:%v", err)
	}

	err = os.MkdirAll(dst, 0755)
	if err != nil {
		return fmt.Errorf("递归复制目录创建目标错误:%v", err)
	}

	for _, entry := range srcInfo {
		entryPath := path.Join(src, entry.Name())
		targetPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = CopyFsDir(srcFS, entryPath, targetPath)
			if err != nil {
				return fmt.Errorf("递归复制目录错误:%v", err)
			}
		} else {
			err = CopyFsFile(srcFS, entryPath, targetPath)
			if err != nil {
				return fmt.Errorf("递归复制文件错误:%v", err)
			}
		}
	}

	return nil
}
