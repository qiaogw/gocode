package toolx

import (
	"bufio"
	"errors"
	"github.com/qiaogw/gocode/util"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func FileInsertInfoSeek(fileName, insertInfo, seekInfo string) (err error) {
	// 打开要操作的文件 os.O_RDWR: 可读可写
	pwd, _ := os.Getwd()
	tempPath := filepath.Join(pwd, "temp8hh8hhkjhjk")
	tempFileName := filepath.Join(tempPath, "temp.temp")
	file, err := os.OpenFile(fileName, os.O_RDWR, 0544)
	if err != nil {
		return
	}
	reader := bufio.NewReader(file)
	// 新建临时文件
	tempFile, err := os.Create(tempFileName)
	if err != nil {
		return
	}
	writer := bufio.NewWriter(tempFile)
	_ = writer.Flush()
	// 将原文件写入临时文件
	for {
		line, err := reader.ReadString('\n') // 依次读一行
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		lineStr := strings.Split(line, "\n")

		_, _ = writer.WriteString(line)
		// 判断当前行是否匹配查找字符串， 为真则插入数据
		if lineStr[0] == seekInfo {
			_, _ = writer.WriteString(insertInfo + "\n")
		}
	}
	_ = writer.Flush()

	file.Close()
	tempFile.Close()
	err = util.FileMove(tempFileName, fileName)
	if err != nil {
		return
	}
	defer func() { // 移除中间文件
		if err := os.RemoveAll(tempPath); err != nil {
			return
		}
	}()
	return
}

func FileInsertInfoLine(fileName, insertInfo string, seekLine int) (err error) {
	// 打开要操作的文件 os.O_RDWR: 可读可写
	pwd, _ := os.Getwd()
	tempPath := filepath.Join(pwd, "temp8hh8hhkjhjk")
	tempFileName := filepath.Join(tempPath, "temp.temp")
	file, err := os.OpenFile(fileName, os.O_RDWR, 0544)
	if err != nil {
		return
	}
	reader := bufio.NewReader(file)
	// 新建临时文件
	tempFile, err := os.Create(tempFileName)
	if err != nil {
		return
	}
	writer := bufio.NewWriter(tempFile)
	_ = writer.Flush()
	// 将原文件写入临时文件
	i := 0
	for {
		line, err := reader.ReadString('\n') // 依次读一行
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		//lineStr := strings.Split(line, "\n")

		_, _ = writer.WriteString(line)
		// 判断当前行是否匹配查找字符串， 为真则插入数据
		if i == seekLine {
			_, _ = writer.WriteString(insertInfo + "\n")
		}
		i++
	}
	_ = writer.Flush()

	file.Close()
	tempFile.Close()
	err = util.FileMove(tempFileName, fileName)
	if err != nil {
		return
	}
	defer func() { // 移除中间文件
		if err := os.RemoveAll(tempPath); err != nil {
			return
		}
	}()
	return
}

// WriteFileString 写入text文件内容
// coverType true 覆盖写入，false 追加写入
func WriteFileString(path, info string, coverType bool) (err error) {

	var fl *os.File
	flag := os.O_APPEND | os.O_WRONLY
	if coverType {
		flag = os.O_TRUNC | os.O_WRONLY
	}
	if util.FileExist(path) { //如果文件存在
		fl, err = os.OpenFile(path, flag, 0666) //打开文件
		if err != nil {
			err = errors.New(path + " 打开文件失败！" + err.Error())
			return
		}
	} else {
		fl, err = os.Create(path) //创建文件
		if err != nil {
			err = errors.New(path + " 创建文件失败！" + err.Error())
			return
		}
	}

	defer fl.Close()
	n, err := fl.WriteString(info + "\n")
	if err == nil && n < len(info) {
		err = errors.New(path + "写入失败！" + err.Error())
	}
	return
}

// WriteFileByte 写入Byte文件内容
// coverType true 覆盖写入，false 追加写入
func WriteFileByte(path string, info []byte, coverType bool) (err error) {
	var fl *os.File
	flag := os.O_APPEND | os.O_WRONLY
	if coverType {
		flag = os.O_APPEND | os.O_TRUNC | os.O_WRONLY
	}
	if util.FileExist(path) { //如果文件存在
		fl, err = os.OpenFile(path, flag, os.ModePerm) //打开文件
	} else {
		fl, err = os.Create(path) //创建文件
	}
	defer fl.Close()
	if err != nil {
		// err = errors.New(path + "打开文件失败！")
		return
	}
	n, err := fl.Write(info)
	if err == nil && n < len(info) {
		err = errors.New(path + "写入失败！")
	}
	return
}
