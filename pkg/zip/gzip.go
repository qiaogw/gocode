// Package zip
// @Description:
package zip

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"github.com/qiaogw/gocode/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

const defaultMaxSize = int64(2 * 1024 * 1024 * 1024) //默认为2G

// Compress 压缩文件和目录
// src: 要压缩的路径 dst: 压缩文件名称。默认是第一个路径加后缀 .tar.gz
func Compress(src, dst string) error {
	maxSize := defaultMaxSize
	perfix := filepath.Base(dst)
	if runtime.GOOS == "windows" {
		return ZipDirAndSplit(src, dst+".zip", perfix, maxSize)
	}
	return CompressAndSplitFiles(src, dst+".tar.gz", perfix, maxSize)
}

// CompressAndSplitFiles 压缩文件和目录并分割压缩文件
// src: 要压缩的路径 dst: 压缩文件名称 prefix: 分割文件名前缀 maxSize: 分割文件的最大大小（字节数）
func CompressAndSplitFiles(src, dst, prefix string, maxSize int64) error {
	// 解析 dst 的目录部分
	destDir := filepath.Dir(dst)

	// 检查目标目录是否存在，如果不存在，则创建
	err := utils.IsNotExistMkDir(destDir)
	if err != nil {
		return fmt.Errorf("创建文件目录 %s 错误：%v", destDir, err)
	}
	if maxSize > defaultMaxSize {
		logx.Info(" 分割文件的最大大小为1GB，你设置的maxSize：%d 已经改为1GB！", maxSize)
		maxSize = defaultMaxSize
	}

	var outputFile *os.File
	var gw *gzip.Writer
	var tw *tar.Writer
	var currentSize int64
	var partNumber int

	// 创建第一个分割文件
	err = createNewSplitFile(src, destDir, prefix, 0, &outputFile, &gw, &tw)
	if err != nil {
		return err
	}
	defer func() {
		if gw != nil {
			gw.Close()
		}
		if tw != nil {
			tw.Close()
		}
		if outputFile != nil {
			outputFile.Close()
		}
	}()

	err = filepath.Walk(src, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 获取相对路径
		relPath, err := filepath.Rel(src, filePath)
		if err != nil {
			return err
		}

		// 使用正斜杠作为路径分隔符
		relPath = filepath.ToSlash(relPath)

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}

		// 修改文件头的名称以保留目录结构
		header.Name = relPath

		// 对于目录和符号链接，将大小字段设置为零,在布尔逻辑中，OR（或运算）操作符（||）是一个短路运算符
		if info.IsDir() || info.Mode()&os.ModeSymlink != 0 {
			header.Size = 0
		}

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if !info.IsDir() {
			//fmt.Printf("压缩完成: %s\n", filePath)
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			n, err := io.Copy(tw, file)
			if err != nil {
				return fmt.Errorf("当前分割文件大小:%d,超过了最大大小%d,文件[%s,大小%d]压缩分割错误:%v",
					currentSize+n, maxSize, file.Name(), n, err)
			}

			currentSize += n

			// 如果当前分割文件大小超过了最大大小，创建一个新的分割文件
			if currentSize >= maxSize {
				tw.Close() // 关闭当前的tar写入器
				gw.Close() // 关闭当前的gzip写入器
				outputFile.Close()

				partNumber++
				currentSize = 0

				// 创建新的分割文件
				err := createNewSplitFile(src, destDir, prefix, partNumber, &outputFile, &gw, &tw)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func createNewSplitFile(src, destDir, prefix string, partNumber int, outputFile **os.File, gw **gzip.Writer, tw **tar.Writer) error {
	// 如果需要添加下标，则创建新的分割文件
	partFileName := fmt.Sprintf("%s.tar.gz", prefix)
	if partNumber > 0 {
		partFileName = fmt.Sprintf("%s_%d.tar.gz", prefix, partNumber)
	}
	//partFileName := fmt.Sprintf("%s_%d.tar.gz", prefix, partNumber)
	partFilePath := filepath.Join(destDir, partFileName)

	file, err := os.Create(partFilePath)
	if err != nil {
		return err
	}

	*outputFile = file
	*gw = gzip.NewWriter(file)
	*tw = tar.NewWriter(*gw)

	return nil
}

// CompressFilesOrFolds 压缩文件和目录
// src: 要压缩的路径 dst: 压缩文件名称。
func CompressFilesOrFolds(src, dst string) error {
	err := utils.IsNotExistMkDir(filepath.Dir(dst))
	if err != nil {
		return fmt.Errorf("创建文件目录 %s 错误：%v", dst, err)
	}
	outputFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// 创建 gzip 写入器
	gw := gzip.NewWriter(outputFile)
	defer gw.Close()

	// 创建 tar 写入器
	tw := tar.NewWriter(gw)
	defer tw.Close()

	err = filepath.Walk(src, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 获取相对路径
		relPath, err := filepath.Rel(src, filePath)
		if err != nil {
			return err
		}

		// 使用正斜杠作为路径分隔符
		relPath = filepath.ToSlash(relPath)

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}

		// 修改文件头的名称以保留目录结构
		header.Name = relPath
		//fmt.Printf("tarSrc.Name : %s ->header.Name %s\n", tarSrc, header.Name)
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if !info.IsDir() {
			//fmt.Printf("压缩完成: %s\n", filePath)
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(tw, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
