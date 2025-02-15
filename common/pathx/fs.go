package pathx

import (
	"embed"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// CopyTpl 从嵌入文件系统 srcFS 中复制模板目录或文件到目标目录 dst。
// 参数 src 是嵌入文件系统中的源目录，dst 是复制目标目录。
func CopyTpl(srcFS embed.FS, src, dst string) error {
	// 目标文件夹路径
	destinationFolder := dst

	// 从嵌入文件系统中读取 src 目录下的所有条目（文件和子目录）
	entries, err := srcFS.ReadDir(src)
	if err != nil {
		return fmt.Errorf("获取目录中的所有文件和子目录错误:%v", err)
	}

	// 遍历每个条目，分别处理文件和目录
	for _, entry := range entries {
		if entry.IsDir() {
			// 如果条目是目录，递归复制目录及其内容
			srcPath := path.Join(src, entry.Name())
			dstPath := filepath.Join(destinationFolder, entry.Name())
			err = CopyFsDir(srcFS, srcPath, dstPath)
			if err != nil {
				return fmt.Errorf("复制目录错误:%v", err)
			}
		} else {
			// 如果条目是文件，复制该文件到目标位置
			srcPath := path.Join(src, entry.Name())
			dstPath := filepath.Join(destinationFolder, entry.Name())
			err = CopyFsFile(srcFS, srcPath, dstPath)
			if err != nil {
				return fmt.Errorf("复制文件错误:%v", err)
			}
		}
	}
	return nil
}

// CopyFsFile 复制嵌入文件系统 srcFS 中的单个文件到目标路径 dst。
// 参数 src 为源文件路径，dst 为目标文件路径。
func CopyFsFile(srcFS embed.FS, src, dst string) error {
	// 从嵌入文件系统中读取源文件数据
	data, err := srcFS.ReadFile(src)
	if err != nil {
		return fmt.Errorf("复制fs文件ReadFile错误:%v", err)
	}
	// 确保目标文件所在的目录存在（MkDirIfNotExist 应在其他地方实现）
	_ = MkdirIfNotExist(filepath.Dir(dst))
	// 将数据写入目标文件，权限设置为0644
	err = os.WriteFile(dst, data, 0644)
	if err != nil {
		return fmt.Errorf("复制fs文件WriteFile错误:%v", err)
	}
	return nil
}

// CopyFsDir 递归复制嵌入文件系统 srcFS 中的目录及其所有内容到目标目录 dst。
// 参数 src 为源目录路径，dst 为目标目录路径。
func CopyFsDir(srcFS embed.FS, src, dst string) error {
	// 读取源目录下的所有条目
	srcInfo, err := srcFS.ReadDir(src)
	if err != nil {
		return fmt.Errorf("递归复制目录读取错误:%v", err)
	}

	// 在目标位置创建目录，权限设置为0755
	err = os.MkdirAll(dst, 0755)
	if err != nil {
		return fmt.Errorf("递归复制目录创建目标错误:%v", err)
	}

	// 遍历源目录下的所有条目，分别递归处理目录和复制文件
	for _, entry := range srcInfo {
		entryPath := path.Join(src, entry.Name())
		targetPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			// 如果条目是目录，递归复制该子目录
			err = CopyFsDir(srcFS, entryPath, targetPath)
			if err != nil {
				return fmt.Errorf("递归复制目录错误:%v", err)
			}
		} else {
			// 如果条目是文件，则复制文件
			err = CopyFsFile(srcFS, entryPath, targetPath)
			if err != nil {
				return fmt.Errorf("递归复制文件错误:%v", err)
			}
		}
	}
	return nil
}
