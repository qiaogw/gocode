package util

import (
	"fmt"
	"os"
	"testing"
)

func TestFileMove(t *testing.T) {
	// 创建用于测试的临时源文件
	srcFile, err := os.CreateTemp("", "源文件.txt")
	if err != nil {
		t.Fatalf("创建源文件时发生错误：%v", err)
	}

	defer os.Remove(srcFile.Name())
	srcFile.Close()

	// 创建用于测试的临时目标目录
	//dstDir, err := os.MkdirTemp("", "目标目录")
	//
	//if err != nil {
	//	t.Fatalf("创建目标目录时发生错误：%v", err)
	//}

	//defer os.RemoveAll(dstDir)

	dstFile, err := os.CreateTemp("", "临时文件.txt")
	if err != nil {
		t.Fatalf("创建目标文件时发生错误：%v", err)
	}
	defer os.Remove(dstFile.Name())
	dstFile.Close()
	fmt.Println("目标:", dstFile.Name())
	// 调用 FileMove 函数，使用临时源和目标路径
	err = FileMove(srcFile.Name(), dstFile.Name())

	if err != nil {
		t.Fatalf("FileMove 函数失败：%v", err)
	}

	// 检查文件是否存在于目标目录中
	//dstFilePath := filepath.Join(dstDir, "目标文件.txt")
	if _, err := os.Stat(dstFile.Name()); os.IsNotExist(err) {
		t.Fatalf("文件未移动到目标目录")
	}
}

func TestDeLFile(t *testing.T) {
	// 创建用于测试的临时文件
	tempFile, err := os.CreateTemp("", "临时文件.txt")
	if err != nil {
		t.Fatalf("创建临时文件时发生错误：%v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	// 调用 DeLFile 函数，使用临时文件路径
	err = DeLFile(tempFile.Name())
	if err != nil {
		t.Fatalf("DeLFile 函数失败：%v", err)
	}

	// 检查文件是否被删除
	if _, err := os.Stat(tempFile.Name()); !os.IsNotExist(err) {
		t.Fatalf("文件未被删除")
	}
}

func TestTrimSpace(t *testing.T) {
	// 定义用于测试的样本结构体
	type TestStruct struct {
		Field1 string
		Field2 int
	}

	// 创建结构体的实例
	testStruct := &TestStruct{
		Field1: "  需要修剪  ",
		Field2: 42,
	}

	// 调用 TrimSpace 函数，使用结构体实例
	TrimSpace(testStruct)

	// 检查字符串字段是否被修剪
	if testStruct.Field1 != "需要修剪" {
		t.Fatalf("字符串字段未被修剪")
	}
}

func TestFileExist(t *testing.T) {
	// 创建用于测试的临时文件
	tempFile, err := os.CreateTemp("", "临时文件.txt")
	if err != nil {
		t.Fatalf("创建临时文件时发生错误：%v", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// 检查文件是否存在
	exists := FileExist(tempFile.Name())
	if !exists {
		t.Fatalf("对于存在的文件，FileExist 返回了 false")
	}

	// 检查函数是否正确报告不存在的文件
	exists = FileExist("不存在的文件.txt")
	if exists {
		t.Fatalf("对于不存在的文件，FileExist 返回了 true")
	}
}

func TestFmtCode(t *testing.T) {
	// 创建用于测试的临时文件
	tempFile, err := os.CreateTemp("", "临时文件.go")
	if err != nil {
		t.Fatalf("创建临时文件时发生错误：%v", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// 向文件写入一些 Go 代码
	_, err = tempFile.WriteString("package main\n\nfunc main() {\nfmt.Println(\"你好，世界！\")\n}\n")
	if err != nil {
		t.Fatalf("向临时文件写入时发生错误：%v", err)
	}

	// 调用 FmtCode 函数，使用临时文件路径
	err = FmtCode(tempFile.Name())
	if err != nil {
		t.Fatalf("FmtCode 函数失败：%v", err)
	}

	// 如果需要，可以添加其他检查以验证格式是否正确
	// 例如，可以读取文件的内容并检查其是否被正确格式化
}
