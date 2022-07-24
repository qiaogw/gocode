package gen

import (
	"fmt"
	"github.com/qiaogw/gocode/global"
	"github.com/qiaogw/gocode/inital"
	"github.com/qiaogw/gocode/util"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

// getNeedList 获取模板文件和构建所需目录
func (acd *AutoCodeService) getNeedList(pack, templatePath string) (dataList []tplData, fileList []string, needMkdir []string, err error) {
	// 去除所有空格
	// 获取 basePath 文件夹下所有tpl文件
	tplFileList, err := acd.GetAllTplFile(templatePath, nil)
	//log.Printf("templatePath is %s ,tplFileList is %v\n", templatePath, tplFileList)
	if err != nil {
		return nil, nil, nil, err
	}
	dataList = make([]tplData, 0, len(tplFileList))
	fileList = make([]string, 0, len(tplFileList))
	needMkdir = make([]string, 0, len(tplFileList)) // 当文件夹下存在多个tpl文件时，改为map更合理
	// 根据文件路径生成 tplData 结构体，待填充数据
	for _, value := range tplFileList {
		dataList = append(dataList, tplData{locationPath: value, autoPackage: pack})
	}
	// 生成 *Template, 填充 template 字段
	for index, value := range dataList {
		dataList[index].template, err = template.ParseFS(inital.TemplateTpl, value.locationPath)
		//dataList[index].template, err = template.ParseFiles(value.locationPath)
		if err != nil {
			return nil, nil, nil, err
		}
	}
	// 生成文件路径，填充 autoCodePath 字段，readme.txt.tpl不符合规则，需要特殊处理
	// resource/template/web/api.js.tpl -> autoCode/web/autoCode.PackageName/api/autoCode.PackageName.js
	// resource/template/readme.txt.tpl -> autoCode/readme.txt
	pack = strings.ToLower(pack)
	for index, value := range dataList {
		trimBase := strings.TrimPrefix(value.locationPath, tempPath+"/")
		//log.Printf("trimBase is %s\n", trimBase)

		if lastSeparator := strings.LastIndex(trimBase, "/"); lastSeparator != -1 {
			origFileName := strings.TrimSuffix(trimBase[lastSeparator+1:], ".tpl")
			firstDot := strings.Index(origFileName, ".")
			//log.Printf("pack is %s,origFileName is %s\n", pack, origFileName)
			fileSlice := strings.Split(origFileName, ".")
			//log.Printf("origFileName[firstDot:] is %s\n", origFileName[firstDot:])
			//log.Printf("fileSlice is %s\n", fileSlice)

			if firstDot != -1 {
				var fileName string
				if origFileName == "config.go" || origFileName == "servicecontext.go" {
					fileName = origFileName
				} else if fileSlice[1] == "logic" {
					//log.Printf("fileSlice[1] is %s\n", fileSlice[1])
					fileName = fileSlice[0] + pack + "logic.go"
					//log.Printf("pack is %s,origFileName is %s\n", pack, origFileName)
				} else if origFileName[firstDot:] != ".go" {
					fileName = pack + origFileName[firstDot:]
				} else {
					fileName = pack + origFileName
				}
				//log.Printf("fileName is %s\n", fileName)
				dataList[index].autoCodePath = filepath.Join(autoPath, trimBase[:lastSeparator],
					fileName)
			}
		}
		//log.Printf("dataList[index] is %+v\n", dataList[index])
		if lastSeparator := strings.LastIndex(dataList[index].autoCodePath, string(os.PathSeparator)); lastSeparator != -1 {
			needMkdir = append(needMkdir, dataList[index].autoCodePath[:lastSeparator])
		}
	}
	for _, value := range dataList {
		fileList = append(fileList, value.autoCodePath)
	}
	return dataList, fileList, needMkdir, err
}

//GetAllTplFile 获取 pathName 文件夹下所有 tpl 文件
//@param: pathName string, fileList []string
//@return: []string, error
func (acd *AutoCodeService) GetAllTplFile(pathName string, fileList []string) ([]string, error) {
	//files, err := ioutil.ReadDir(pathName)
	files, err := inital.TemplateTpl.ReadDir(pathName)
	for _, fi := range files {
		if fi.IsDir() {
			fileList, err = acd.GetAllTplFile(pathName+"/"+fi.Name(), fileList)
			if err != nil {
				return nil, err
			}
		} else {
			if strings.HasSuffix(fi.Name(), ".tpl") {
				fileList = append(fileList, pathName+"/"+fi.Name())
			}
		}
	}
	//log.Printf("pathName is %s ,fileList is %v\n", pathName, fileList)
	return fileList, err
}

//addAutoMoveFile 生成对应的迁移文件路径
//@param: *tplData
//@return: null
func (acd *AutoCodeService) addAutoMoveFile(data *tplData) {
	base := filepath.Base(data.autoCodePath)

	fileSlice := strings.Split(data.autoCodePath, string(os.PathSeparator))
	n := len(fileSlice)
	if n <= 2 {
		return
	}
	//log.Printf("fileSlice[n-2] is %s\n", fileSlice[n-2])
	if strings.Contains(fileSlice[n-2], "model_gen") {
		bn := strings.TrimSuffix(base, ".go")
		base = bn + "_gen.go"
		data.autoMoveFilePath = filepath.Join(global.GenConfig.AutoCode.Root,
			global.GenConfig.AutoCode.SModel, base)
	} else if strings.Contains(fileSlice[n-2], "model") {
		data.autoMoveFilePath = filepath.Join(global.GenConfig.AutoCode.Root,
			global.GenConfig.AutoCode.SModel, base)
	} else if strings.Contains(fileSlice[n-2], "rpc-logic") {
		//log.Printf("strings.Contains(fileSlice[n-2]: %+v\n", fileSlice[n-2:])
		data.autoMoveFilePath = filepath.Join(global.GenConfig.AutoCode.Root,
			global.GenConfig.AutoCode.RpcLogic, base)
	} else if strings.Contains(fileSlice[n-2], "api-logic") {
		//log.Printf("strings.Contains(fileSlice[n-2]: %+v\n", fileSlice[n-2:])
		data.autoMoveFilePath = filepath.Join(global.GenConfig.AutoCode.Root,
			global.GenConfig.AutoCode.ApiLogic, base)
	} else if strings.Contains(fileSlice[n-2], "common") {
		log.Printf("strings.Contains(fileSlice[n-2]: %+v\n", fileSlice[n-2:])
		data.autoMoveFilePath = filepath.Join(global.GenConfig.AutoCode.Root,
			global.GenConfig.AutoCode.Common, base)
	} else if strings.Contains(fileSlice[n-2], "rpc") {
		//log.Printf("strings.Contains(fileSlice[n-2]: %+v\n", fileSlice[n-2:])
		data.autoMoveFilePath = filepath.Join(global.GenConfig.AutoCode.Root,
			global.GenConfig.AutoCode.SRpc, base)
	} else if strings.Contains(fileSlice[n-2], "api") {

		data.autoMoveFilePath = filepath.Join(global.GenConfig.AutoCode.Root,
			global.GenConfig.AutoCode.SApi, base)
	} else {
		//log.Printf("strings.Contains(fileSlice[2]: %+v\n", fileSlice[2:])
		tp := filepath.Join(fileSlice[1:]...)
		data.autoMoveFilePath = filepath.Join(global.GenConfig.AutoCode.Root, tp)
	}
	//log.Printf("data.autoCodePath is %s;;;;;data.autoMoveFilePath is %s\n", data.autoCodePath, data.autoMoveFilePath)
}

func (acd *AutoCodeService) genBefore(pack, packPath string) (dataList []tplData, err error) {
	//tPath := filepath.Join(tempPath, packPath)
	tPath := tempPath + "/" + packPath
	dataList, _, needMkdir, err := acd.getNeedList(pack, tPath)
	injectionCodeMeta := strings.Builder{}
	err = injectionCode(pack, &injectionCodeMeta)
	//log.Printf("injectionPaths is %+v\n", injectionPaths)
	if err != nil {
		return
	}
	if err = util.CreateDir(needMkdir...); err != nil {
		return
	}
	// 写入文件前，先创建文件夹
	if err = util.CreateDir(needMkdir...); err != nil {
		log.Printf("err is %+v\n", err)
		return
	}
	return
}

func (acd *AutoCodeService) genAfter(dataList []tplData, ids ...uint) error {

	bf := strings.Builder{}
	idBf := strings.Builder{}

	for _, id := range ids {
		idBf.WriteString(strconv.Itoa(int(id)))
		idBf.WriteString(";")
	}

	for index := range dataList {
		acd.addAutoMoveFile(&dataList[index])

	}

	for _, value := range dataList { // 移动文件
		// 判断目标文件是否都可以移动
		if util.FileExist(value.autoMoveFilePath) {
			fmt.Printf("目标文件已存在:%s\n", value.autoMoveFilePath)
			continue
		}
		//fmt.Printf("value.autoCodePath is %s, value.autoMoveFilePath is %s\n", value.autoCodePath, value.autoMoveFilePath)
		if err := util.FileMove(value.autoCodePath, value.autoMoveFilePath); err != nil {
			log.Printf("err is %v\n", err)
			return err
		}
		if len(value.autoMoveFilePath) != 0 {
			bf.WriteString(value.autoMoveFilePath)
			bf.WriteString(";")
		}
	}
	//Init(table.Table)
	return nil
}
