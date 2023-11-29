package gen

import (
	"fmt"
	"github.com/qiaogw/gocode/inital"
	"github.com/qiaogw/gocode/util"
	"github.com/wxnacy/wgo/arrays"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/qiaogw/gocode/global"
)

// getNeedList 获取模板文件和构建所需目录
func (acd *AutoCodeService) getNeedList(pack, templatePath string) (dataList []tplData, fileList []string, needMkdir []string, err error) {
	// 获取 basePath 文件夹下所有tpl文件

	tplFileList, err := acd.GetAllTplFile(templatePath, nil)
	//log.Printf("pack is %s ,tplFileList is %v\n", pack, tplFileList)
	if err != nil {
		log.Printf("templatePath is %s ,err is %v\n", templatePath, err)
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
		//添加 template 函数 iota 自增
		funcMap := template.FuncMap{
			"add": add,
		}
		t1 := template.New(value.locationPath)
		t1 = t1.Funcs(funcMap)
		mi, _ := inital.TemplateTpl.ReadFile(value.locationPath)
		t2, err := t1.Parse(string(mi))
		if err != nil {
			log.Printf("templatePath is %s ,err is %v\n", templatePath, err)
			return nil, nil, nil, err
		}
		dataList[index].template = t2
	}
	// 生成文件路径，填充 autoCodePath 字段，readme.txt.tpl不符合规则，需要特殊处理
	// resource/template/web/api.js.tpl -> autoCode/web/autoCode.PackageName/api/autoCode.PackageName.js
	// resource/template/readme.txt.tpl -> autoCode/readme.txt
	pack = strings.ToLower(pack)
	// log.Printf("pack is %s ,dataList is %v\n", pack, dataList)
	for index, value := range dataList {
		trimBase := strings.TrimPrefix(value.locationPath, tempPath+"/")
		if lastSeparator := strings.LastIndex(trimBase, "/"); lastSeparator != -1 {
			origFileName := strings.TrimSuffix(trimBase[lastSeparator+1:], ".tpl")
			firstDot := strings.Index(origFileName, ".")
			fileSlice := strings.Split(origFileName, ".")
			if firstDot != -1 {
				var fileName string
				if origFileName == "config.go" || origFileName == "servicecontext.go" {
					fileName = origFileName
				} else if fileSlice[1] == "logic" {
					fileName = fileSlice[0] + pack + "logic.go"
				} else if trimBase[0:strings.Index(trimBase, "/")] == "common" {
					fileName = origFileName
				} else if origFileName == "replay.api" {
					fileName = origFileName
				} else if origFileName[firstDot:] == ".js" || origFileName[firstDot:] == ".vue" {
					fileName = "index" + origFileName[firstDot:]
				} else if origFileName[firstDot:] != ".go" {
					fileName = pack + origFileName[firstDot:]
				} else {
					//log.Printf("【pack】%+v,origFileName:%+v", pack, origFileName)
					fileName = pack + origFileName
				}
				dataList[index].autoCodePath = filepath.Join(autoPath, trimBase[:lastSeparator],
					fileName)

			}
		}

		if lastSeparator := strings.LastIndex(dataList[index].autoCodePath, string(os.PathSeparator)); lastSeparator != -1 {
			needMkdir = append(needMkdir, dataList[index].autoCodePath[:lastSeparator])
		}
	}
	for _, value := range dataList {
		fileList = append(fileList, value.autoCodePath)
	}
	return dataList, fileList, needMkdir, err
}

// GetAllTplFile 获取 pathName 文件夹下所有 tpl 文件
// @param: pathName string, fileList []string
// @return: []string, error
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

// addAutoMoveFile 生成对应的迁移文件路径
// @param: *tplData
// @return: null
func (acd *AutoCodeService) addAutoMoveFile(data *tplData) {
	base := filepath.Base(data.autoCodePath)
	fileSlice := strings.Split(data.autoCodePath, string(os.PathSeparator))
	n := len(fileSlice)
	if n <= 2 {
		return
	}
	c := global.GenConfig.AutoCode
	rPath := c.Root
	fPath := filepath.Join(rPath)
	//log.Println(util.Red(fmt.Sprintf("data is %+v\n", data)))
	switch fileSlice[1] {
	case apiLogicPath:
		fPath = filepath.Join(fPath, apiPath, internalPath, logicPath, data.tablePkg)
	case rpcLogicPath:
		fPath = filepath.Join(fPath, rpcPath, internalPath, logicPath, data.tablePkg)
	case apiDescPath:
		//log.Printf("data.autoCodePath is %s,fPath is %s\n", data.autoCodePath, fPath)
		fPath = filepath.Join(fPath, apiPath, apiDescPath)
	case commonPath:
		cp, _ := os.Getwd()
		tp := filepath.Join(fileSlice[1 : n-1]...)
		fPath = filepath.Join(cp, tp)
	case apiPath:
		fPath = filepath.Join(fPath, apiPath)
	case rpcPath:
		fPath = filepath.Join(fPath, rpcPath)
	case modelPath:
		fPath = filepath.Join(fPath, modelPath)
	case webPath:
		baseSlic := strings.Split(base, ".")
		if baseSlic[1] == "vue" {
			fPath = filepath.Join(fPath, webPath, "src", "pages", global.GenConfig.System.Name, data.tablePkg)
		} else {
			fPath = filepath.Join(fPath, webPath, "src", "api", global.GenConfig.System.Name, data.tablePkg)
		}
	default:
	}

	switch fileSlice[2] {
	case etcPath:
		fPath = filepath.Join(fPath, etcPath)
	case svcPath:
		fPath = filepath.Join(fPath, internalPath, svcPath)
	case configPath:
		fPath = filepath.Join(fPath, internalPath, configPath)
	case apiResponsePath:
		fPath = filepath.Join(fPath, apiResponsePath)
	default:
	}
	data.autoMoveFilePath = filepath.Join(fPath, base)
}

func (acd *AutoCodeService) genBefore(pack, packPath string) (dataList []tplData, err error) {
	tPath := tempPath + "/" + packPath
	dataList, _, needMkdir, err := acd.getNeedList(pack, tPath)
	if err = util.CreateDir(needMkdir...); err != nil {
		return
	}
	// 写入文件前，先创建文件夹
	if err = util.CreateDir(needMkdir...); err != nil {
		//log.Printf("err is %+v\n", err)
		return
	}
	return
}

func (acd *AutoCodeService) genAfter(dataList []tplData, pkg ...string) error {

	bf := strings.Builder{}
	for index := range dataList {
		acd.addAutoMoveFile(&dataList[index])

	}

	for _, value := range dataList { // 移动文件
		base := filepath.Base(value.autoCodePath)
		baseOk := arrays.ContainsString(global.GenConfig.AutoCode.CoverFile, base)
		// 判断目标文件是否都可以移动
		if util.FileExist(value.autoMoveFilePath) && baseOk < 0 {
			fmt.Println(util.Yellow(fmt.Sprintf("目标文件已存在:%s", value.autoMoveFilePath)))
			continue
		}
		if err := util.FileMove(value.autoCodePath, value.autoMoveFilePath); err != nil {
			return err
		}
		if len(value.autoMoveFilePath) != 0 {
			bf.WriteString(value.autoMoveFilePath)
			bf.WriteString(";")
		}
		if err := util.FmtCode(value.autoMoveFilePath); err != nil {
			return err
		}
	}
	//Init(table.Table)
	return nil
}

func add(i int) int {
	i++
	return i
}
