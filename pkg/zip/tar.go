package zip

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"github.com/qiaogw/gocode/pkg/utils"
	"io"
	"os"
	"path/filepath"
)

// DecompressWithPermissions 解压 Tar.gz 文件并保留文件权限
func DecompressWithPermissions(tarFile, dst string) error {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()

	tr := tar.NewReader(gr)
	for {
		header, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		// 修复路径，确保使用正斜杠分隔符
		header.Name = filepath.ToSlash(header.Name)
		filePath := filepath.Join(dst, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			// 创建目录并设置权限
			if err := os.MkdirAll(filePath, os.FileMode(header.Mode)); err != nil {
				return err
			}
		case tar.TypeReg, tar.TypeRegA:
			// 创建文件并设置权限，然后将文件内容复制到文件中
			file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			_, err = io.Copy(file, tr)
			file.Close()
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported file type: %s", header.Typeflag)
		}
	}
	return nil
}

// CompressDirectoryWithPermissions 压缩目录为 Tar.gz 文件并保留文件权限
func CompressDirectoryWithPermissions(sourceDir, outputFilePath string) error {
	err := utils.IsNotExistMkDir(filepath.Dir(outputFilePath))
	if err != nil {
		return fmt.Errorf("创建文件目录 %s 错误：%v", outputFilePath, err)
	}
	outputFile, err := os.Create(outputFilePath)
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

	// 获取目录的基本名称
	baseName := filepath.Base(sourceDir)

	wd, _ := os.Getwd()
	tarSrc := filepath.Join(wd, "tarTemp", baseName)
	os.Mkdir(tarSrc, 0755)
	utils.CopyDir(sourceDir, filepath.Join(wd, "tarTemp", baseName, baseName))
	fmt.Printf("sourceDir.Name : %s ->tarSrc.Name %s\n", sourceDir, tarSrc)

	err = filepath.Walk(tarSrc, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 获取相对路径
		relPath, err := filepath.Rel(tarSrc, filePath)
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
			fmt.Printf("压缩完成: %s\n", filePath)
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
	err = os.RemoveAll(tarSrc)
	if err != nil {
		return err
	}

	return nil
}
