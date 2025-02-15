package filex

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ZipDir 压缩文件夹
func ZipDir(dir, zipFile string) (err error) {
	err = IsNotExistMkDir(filepath.Dir(zipFile))
	if err != nil {
		return fmt.Errorf("创建文件目录 %s 错误：%v", zipFile, err)
	}
	fz, err := os.Create(zipFile)
	if err != nil {
		return fmt.Errorf("创建ZIP文件失败: %s", err)
	}
	defer fz.Close()

	w := zip.NewWriter(fz)
	defer w.Close()

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("遍历文件夹失败: %s", err)
		}
		if !info.IsDir() {
			fDest, err := w.Create(path[len(dir)+1:])
			if err != nil {
				return fmt.Errorf("创建ZIP文件失败: %s", err)
			}
			fSrc, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("打开文件失败: %s", err)
			}
			_, err = io.Copy(fDest, fSrc)
			fSrc.Close()
			if err != nil {
				return fmt.Errorf("复制文件内容失败: %s", err)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// UnzipDir 解压文件夹
func UnzipDir(zipFile, dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		path := filepath.Join(dir, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
			continue
		}

		err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return fmt.Errorf("创建目录失败: %s", err)
		}

		fDest, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("创建文件失败: %s", err)
		}
		defer fDest.Close()

		fSrc, err := f.Open()
		if err != nil {
			return fmt.Errorf("打开ZIP文件中的文件失败: %s", err)
		}
		defer fSrc.Close()

		_, err = io.Copy(fDest, fSrc)
		if err != nil {
			return fmt.Errorf("复制文件内容失败: %s", err)
		}
	}
	return nil
}

// ZipDirAndSplit 压缩文件夹并分割压缩文件
func ZipDirAndSplit(src, dst, prefix string, maxSize int64) (err error) {
	// 提取dst的目录路径
	outputDir := filepath.Dir(dst)

	err = IsNotExistMkDir(outputDir)
	if err != nil {
		return fmt.Errorf("创建目录路径 %s 错误：%v", outputDir, err)
	}

	var zw *zip.Writer
	var outputFile *os.File
	var currentSize int64
	var partNumber int

	defer func() {
		if zw != nil {
			zw.Close()
		}
		if outputFile != nil {
			outputFile.Close()
		}
	}()

	createNewZipWriter := func() error {
		if zw != nil {
			zw.Close()
		}
		if outputFile != nil {
			outputFile.Close()
		}

		currentSize = 0

		// 如果需要添加下标，则创建新的分割文件
		partFileName := fmt.Sprintf("%s.zip", prefix)
		if partNumber > 0 {
			partFileName = fmt.Sprintf("%s_%d.zip", prefix, partNumber)
		}

		// 创建新的分割文件
		partFileName = filepath.Join(outputDir, partFileName)
		var err error
		outputFile, err = os.Create(partFileName)
		if err != nil {
			return fmt.Errorf("创建分割ZIP文件失败: %s", err)
		}

		zw = zip.NewWriter(outputFile)
		return nil
	}

	err = createNewZipWriter()
	if err != nil {
		return err
	}

	err = filepath.Walk(src, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("遍历文件夹失败: %s", err)
		}
		if !info.IsDir() {
			// 获取相对路径
			relPath, err := filepath.Rel(src, filePath)
			if err != nil {
				return fmt.Errorf("计算相对路径失败: %s", err)
			}

			// 使用正斜杠作为路径分隔符
			relPath = filepath.ToSlash(relPath)

			// 创建文件头
			header := &zip.FileHeader{
				Name:   relPath,
				Method: zip.Deflate, // 使用压缩方法
			}
			//header.Name = relPath

			// 写入文件头
			fDest, err := zw.CreateHeader(header)
			if err != nil {
				return fmt.Errorf("创建ZIP文件失败: %s", err)
			}

			// 打开源文件
			fSrc, err := os.Open(filePath)
			if err != nil {
				return fmt.Errorf("打开文件失败: %s", err)
			}
			// 复制文件内容
			n, err := io.Copy(fDest, fSrc)
			fSrc.Close()
			if err != nil {
				return fmt.Errorf("复制文件内容失败: %s", err)
			}

			currentSize += n

			// 如果当前分割文件大小超过了最大大小，创建一个新的分割文件
			if currentSize >= maxSize {
				partNumber++
				err := createNewZipWriter()
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
