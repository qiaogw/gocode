package gen

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"gocode/global"
	"gocode/models"
	"gocode/template/autocode_template/subcontract"
	"gocode/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

const (
	autoPath           = "autocode_template/"
	autocodePath       = "template/autocode_template"
	plugPath           = "template/plug_template"
	packageService     = "service/%s/enter.go"
	packageServiceName = "service"
	packageRouter      = "router/%s/enter.go"
	packageRouterName  = "router"
	packageAPI         = "api/v1/%s/enter.go"
	packageAPIName     = "api/v1"
)

type autoPackage struct {
	path string
	temp string
	name string
}

var (
	packageInjectionMap map[string]astInjectionMeta
	injectionPaths      []injectionMeta
	caser               = cases.Title(language.English)
)

func Init(Package string) {
	injectionPaths = []injectionMeta{
		{
			path: filepath.Join(global.GenConfig.AutoCode.Root,
				global.GenConfig.AutoCode.Server, global.GenConfig.AutoCode.SInitialize, "gorm.go"),
			funcName:    "MysqlTables",
			structNameF: Package + ".%s{},",
		},
		{
			path: filepath.Join(global.GenConfig.AutoCode.Root,
				global.GenConfig.AutoCode.Server, global.GenConfig.AutoCode.SInitialize, "router.go"),
			funcName:    "Routers",
			structNameF: Package + "Router.Init%sRouter(PrivateGroup)",
		},
		{
			path: filepath.Join(global.GenConfig.AutoCode.Root,
				global.GenConfig.AutoCode.Server, fmt.Sprintf(global.GenConfig.AutoCode.SApi, Package), "enter.go"),
			funcName:    "ApiGroup",
			structNameF: "%sApi",
		},
		{
			path: filepath.Join(global.GenConfig.AutoCode.Root,
				global.GenConfig.AutoCode.Server, fmt.Sprintf(global.GenConfig.AutoCode.SRouter, Package), "enter.go"),
			funcName:    "RouterGroup",
			structNameF: "%sRouter",
		},
		{
			path: filepath.Join(global.GenConfig.AutoCode.Root,
				global.GenConfig.AutoCode.Server, fmt.Sprintf(global.GenConfig.AutoCode.SService, Package), "enter.go"),
			funcName:    "ServiceGroup",
			structNameF: "%sService",
		},
	}

	packageInjectionMap = map[string]astInjectionMeta{
		packageServiceName: {
			path: filepath.Join(global.GenConfig.AutoCode.Root,
				global.GenConfig.AutoCode.Server, "service", "enter.go"),
			importCodeF:  "github.com/flipped-aurora/gin-vue-admin/server/%s/%s",
			packageNameF: "%s",
			groupName:    "ServiceGroup",
			structNameF:  "%sServiceGroup",
		},
		packageRouterName: {
			path: filepath.Join(global.GenConfig.AutoCode.Root,
				global.GenConfig.AutoCode.Server, "router", "enter.go"),
			importCodeF:  "github.com/flipped-aurora/gin-vue-admin/server/%s/%s",
			packageNameF: "%s",
			groupName:    "RouterGroup",
			structNameF:  "%s",
		},
		packageAPIName: {
			path: filepath.Join(global.GenConfig.AutoCode.Root,
				global.GenConfig.AutoCode.Server, "api/v1", "enter.go"),
			importCodeF:  "github.com/flipped-aurora/gin-vue-admin/server/%s/%s",
			packageNameF: "%s",
			groupName:    "ApiGroup",
			structNameF:  "%sApiGroup",
		},
	}
}

type injectionMeta struct {
	path        string
	funcName    string
	structNameF string // 带格式化的
}

type astInjectionMeta struct {
	path         string
	importCodeF  string
	structNameF  string
	packageNameF string
	groupName    string
}

type tplData struct {
	template         *template.Template
	autoPackage      string
	locationPath     string
	autoCodePath     string
	autoMoveFilePath string
}

type AutoCodeService struct{}

var AutoCodeServiceApp = new(AutoCodeService)

func makeDictTypes(autoCode *models.AutoCodeStruct) {
	DictTypeM := make(map[string]string)
	for _, v := range autoCode.Fields {
		if v.DictType != "" {
			DictTypeM[v.DictType] = ""
		}
	}
	for k := range DictTypeM {
		autoCode.DictTypes = append(autoCode.DictTypes, k)
	}
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTemp
//@description: 创建代码
//@param: model.AutoCodeStruct
//@return: err error

func (acd *AutoCodeService) CreateTemp(autoCode models.AutoCodeStruct, ids ...uint) (err error) {
	makeDictTypes(&autoCode)
	for i := range autoCode.Fields {
		if autoCode.Fields[i].FieldType == "time.Time" {
			autoCode.HasTimer = true
			break
		}
	}
	// 增加判断: 重复创建struct
	//if autoCode.AutoMoveFile && AutoCodeHistoryServiceApp.Repeat(autoCode.StructName, autoCode.Package) {
	//	return RepeatErr
	//}
	dataList, _, needMkdir, err := acd.getNeedList(&autoCode)
	if err != nil {
		return err
	}
	//meta, _ := json.Marshal(autoCode)
	// 写入文件前，先创建文件夹
	if err = utils.CreateDir(needMkdir...); err != nil {
		return err
	}

	// 生成文件
	for _, value := range dataList {
		f, err := os.OpenFile(value.autoCodePath, os.O_CREATE|os.O_WRONLY, 0o755)
		if err != nil {
			return err
		}
		if err = value.template.Execute(f, autoCode); err != nil {
			return err
		}
		_ = f.Close()
	}

	defer func() { // 移除中间文件
		if err := os.RemoveAll(autoPath); err != nil {
			return
		}
	}()
	bf := strings.Builder{}
	idBf := strings.Builder{}
	injectionCodeMeta := strings.Builder{}
	for _, id := range ids {
		idBf.WriteString(strconv.Itoa(int(id)))
		idBf.WriteString(";")
	}

	Init(autoCode.Package)
	for index := range dataList {
		acd.addAutoMoveFile(&dataList[index])
	}
	// 判断目标文件是否都可以移动
	for _, value := range dataList {
		fmt.Printf("value.autoCodePath is %s, value.autoMoveFilePath is %s\n", value.autoCodePath, value.autoMoveFilePath)
		if utils.FileExist(value.autoMoveFilePath) {
			return errors.New(fmt.Sprintf("目标文件已存在:%s\n", value.autoMoveFilePath))
		}
	}
	for _, value := range dataList { // 移动文件
		//fmt.Printf("value.autoCodePath is %s, value.autoMoveFilePath is %s\n", value.autoCodePath, value.autoMoveFilePath)
		if err := utils.FileMove(value.autoCodePath, value.autoMoveFilePath); err != nil {
			return err
		}
	}
	err = injectionCode(autoCode.StructName, &injectionCodeMeta)
	if err != nil {
		return
	}
	// 保存生成信息
	for _, data := range dataList {
		if len(data.autoMoveFilePath) != 0 {
			bf.WriteString(data.autoMoveFilePath)
			bf.WriteString(";")
		}
	}

	var gormPath = filepath.Join(global.GenConfig.AutoCode.Root,
		global.GenConfig.AutoCode.Server, global.GenConfig.AutoCode.SInitialize, "gorm.go")
	var routePath = filepath.Join(global.GenConfig.AutoCode.Root,
		global.GenConfig.AutoCode.Server, global.GenConfig.AutoCode.SInitialize, "router.go")
	var imporStr = fmt.Sprintf("github.com/flipped-aurora/gin-vue-admin/server/model/%s", autoCode.Package)
	_ = ImportReference(routePath, "", "", autoCode.Package, "")
	_ = ImportReference(gormPath, imporStr, "", "", "")

	if autoCode.AutoMoveFile {
		return models.AutoMoveErr
	}
	return nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllTplFile
//@description: 获取 pathName 文件夹下所有 tpl 文件
//@param: pathName string, fileList []string
//@return: []string, error

func (acd *AutoCodeService) GetAllTplFile(pathName string, fileList []string) ([]string, error) {
	files, err := ioutil.ReadDir(pathName)
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
	return fileList, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DropTable
//@description: 删除表
//@param: tableName string, dbName string
//@return: err error, Columns []request.ColumnReq

func (acd *AutoCodeService) DropTable(tableName string) error {
	return global.GenDB.Exec("DROP TABLE " + tableName).Error
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: addAutoMoveFile
//@description: 生成对应的迁移文件路径
//@param: *tplData
//@return: null

func (acd *AutoCodeService) addAutoMoveFile(data *tplData) {
	base := filepath.Base(data.autoCodePath)
	fileSlice := strings.Split(data.autoCodePath, string(os.PathSeparator))
	n := len(fileSlice)
	if n <= 2 {
		return
	}
	if strings.Contains(fileSlice[1], "server") {
		if strings.Contains(fileSlice[n-2], "router") {
			data.autoMoveFilePath = filepath.Join(global.GenConfig.AutoCode.Root, global.GenConfig.AutoCode.Server,
				fmt.Sprintf(global.GenConfig.AutoCode.SRouter, data.autoPackage), base)
		} else if strings.Contains(fileSlice[n-2], "api") {
			data.autoMoveFilePath = filepath.Join(global.GenConfig.AutoCode.Root,
				global.GenConfig.AutoCode.Server, fmt.Sprintf(global.GenConfig.AutoCode.SApi, data.autoPackage), base)
		} else if strings.Contains(fileSlice[n-2], "service") {
			data.autoMoveFilePath = filepath.Join(global.GenConfig.AutoCode.Root,
				global.GenConfig.AutoCode.Server, fmt.Sprintf(global.GenConfig.AutoCode.SService, data.autoPackage), base)
		} else if strings.Contains(fileSlice[n-2], "model") {
			data.autoMoveFilePath = filepath.Join(global.GenConfig.AutoCode.Root,
				global.GenConfig.AutoCode.Server, fmt.Sprintf(global.GenConfig.AutoCode.SModel, data.autoPackage), base)
		} else if strings.Contains(fileSlice[n-2], "request") {
			data.autoMoveFilePath = filepath.Join(global.GenConfig.AutoCode.Root,
				global.GenConfig.AutoCode.Server, fmt.Sprintf(global.GenConfig.AutoCode.SRequest, data.autoPackage), base)
		}
	} else if strings.Contains(fileSlice[1], "web") {
		if strings.Contains(fileSlice[n-1], "js") {
			data.autoMoveFilePath = filepath.Join(global.GenConfig.AutoCode.Root,
				global.GenConfig.AutoCode.Web, global.GenConfig.AutoCode.WApi, base)
		} else if strings.Contains(fileSlice[n-2], "form") {
			data.autoMoveFilePath = filepath.Join(global.GenConfig.AutoCode.Root,
				global.GenConfig.AutoCode.Web, global.GenConfig.AutoCode.WForm, filepath.Base(filepath.Dir(filepath.Dir(data.autoCodePath))), strings.TrimSuffix(base, filepath.Ext(base))+"Form.vue")
		} else if strings.Contains(fileSlice[n-2], "table") {
			data.autoMoveFilePath = filepath.Join(global.GenConfig.AutoCode.Root,
				global.GenConfig.AutoCode.Web, global.GenConfig.AutoCode.WTable, filepath.Base(filepath.Dir(filepath.Dir(data.autoCodePath))), base)
		}
	}
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: CreateApi
//@description: 自动创建api数据,
//@param: a *model.AutoCodeStruct
//@return: err error

func (acd *AutoCodeService) AutoCreateApi(a *models.AutoCodeStruct) (ids []uint, err error) {
	apiList := []models.SysApi{
		{
			Path:        "/" + a.Abbreviation + "/" + "create" + a.StructName,
			Description: "新增" + a.Description,
			ApiGroup:    a.Abbreviation,
			Method:      "POST",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName,
			Description: "删除" + a.Description,
			ApiGroup:    a.Abbreviation,
			Method:      "DELETE",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName + "ByIds",
			Description: "批量删除" + a.Description,
			ApiGroup:    a.Abbreviation,
			Method:      "DELETE",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "update" + a.StructName,
			Description: "更新" + a.Description,
			ApiGroup:    a.Abbreviation,
			Method:      "PUT",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "find" + a.StructName,
			Description: "根据ID获取" + a.Description,
			ApiGroup:    a.Abbreviation,
			Method:      "GET",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "get" + a.StructName + "List",
			Description: "获取" + a.Description + "列表",
			ApiGroup:    a.Abbreviation,
			Method:      "GET",
		},
	}
	err = global.GenDB.Transaction(func(tx *gorm.DB) error {
		for _, v := range apiList {
			var api models.SysApi
			if errors.Is(tx.Where("path = ? AND method = ?", v.Path, v.Method).First(&api).Error, gorm.ErrRecordNotFound) {
				if err = tx.Create(&v).Error; err != nil { // 遇到错误时回滚事务
					return err
				} else {
					ids = append(ids, v.ID)
				}
			}
		}
		return nil
	})
	return ids, err
}

func (acd *AutoCodeService) getNeedList(autoCode *models.AutoCodeStruct) (dataList []tplData, fileList []string, needMkdir []string, err error) {
	// 去除所有空格
	utils.TrimSpace(autoCode)
	for _, field := range autoCode.Fields {
		utils.TrimSpace(field)
	}
	// 获取 basePath 文件夹下所有tpl文件
	tplFileList, err := acd.GetAllTplFile(autocodePath, nil)
	if err != nil {
		return nil, nil, nil, err
	}
	dataList = make([]tplData, 0, len(tplFileList))
	fileList = make([]string, 0, len(tplFileList))
	needMkdir = make([]string, 0, len(tplFileList)) // 当文件夹下存在多个tpl文件时，改为map更合理
	// 根据文件路径生成 tplData 结构体，待填充数据
	for _, value := range tplFileList {
		dataList = append(dataList, tplData{locationPath: value, autoPackage: autoCode.Package})
	}
	// 生成 *Template, 填充 template 字段
	for index, value := range dataList {
		dataList[index].template, err = template.ParseFiles(value.locationPath)
		if err != nil {
			return nil, nil, nil, err
		}
	}
	// 生成文件路径，填充 autoCodePath 字段，readme.txt.tpl不符合规则，需要特殊处理
	// resource/template/web/api.js.tpl -> autoCode/web/autoCode.PackageName/api/autoCode.PackageName.js
	// resource/template/readme.txt.tpl -> autoCode/readme.txt
	for index, value := range dataList {
		trimBase := strings.TrimPrefix(value.locationPath, autocodePath+"/")
		if trimBase == "readme.txt.tpl" {
			dataList[index].autoCodePath = autoPath + "readme.txt"
			continue
		}

		if lastSeparator := strings.LastIndex(trimBase, "/"); lastSeparator != -1 {
			origFileName := strings.TrimSuffix(trimBase[lastSeparator+1:], ".tpl")
			firstDot := strings.Index(origFileName, ".")
			if firstDot != -1 {
				var fileName string
				if origFileName[firstDot:] != ".go" {
					fileName = autoCode.PackageName + origFileName[firstDot:]
				} else {
					fileName = autoCode.HumpPackageName + origFileName[firstDot:]
				}

				dataList[index].autoCodePath = filepath.Join(autoPath, trimBase[:lastSeparator], autoCode.PackageName,
					origFileName[:firstDot], fileName)
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

// injectionCode 封装代码注入
func injectionCode(structName string, bf *strings.Builder) error {
	for _, meta := range injectionPaths {
		code := fmt.Sprintf(meta.structNameF, structName)
		if err := utils.AutoInjectionCode(meta.path, meta.funcName, code); err != nil {
			return err
		}
		bf.WriteString(fmt.Sprintf("%s@%s@%s;", meta.path, meta.funcName, code))
	}
	return nil
}

func (acd *AutoCodeService) CreateAutoCode(s *models.SysAutoCode) error {
	if s.PackageName == "autocode" || s.PackageName == "models" || s.PackageName == "example" || s.PackageName == "" {
		return errors.New("不能使用已保留的package name")
	}
	if !errors.Is(global.GenDB.Where("package_name = ?", s.PackageName).First(&models.SysAutoCode{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同PackageName")
	}
	if e := acd.CreatePackageTemp(s.PackageName); e != nil {
		return e
	}
	return global.GenDB.Create(&s).Error
}

func (acd *AutoCodeService) GetPackage() (pkgList []models.SysAutoCode, err error) {
	err = global.GenDB.Find(&pkgList).Error
	return pkgList, err
}

func (acd *AutoCodeService) DelPackage(a models.SysAutoCode) error {
	return global.GenDB.Delete(&a).Error
}

func (acd *AutoCodeService) CreatePackageTemp(packageName string) error {
	Init(packageName)
	pendingTemp := []autoPackage{{
		path: packageService,
		name: packageServiceName,
		temp: string(subcontract.Server),
	}, {
		path: packageRouter,
		name: packageRouterName,
		temp: string(subcontract.Router),
	}, {
		path: packageAPI,
		name: packageAPIName,
		temp: string(subcontract.API),
	}}
	for i, s := range pendingTemp {
		pendingTemp[i].path = filepath.Join(global.GenConfig.AutoCode.Root, global.GenConfig.AutoCode.Server, filepath.Clean(fmt.Sprintf(s.path, packageName)))
	}
	// 选择模板
	for _, s := range pendingTemp {
		err := os.MkdirAll(filepath.Dir(s.path), 0755)
		if err != nil {
			return err
		}

		f, err := os.Create(s.path)
		if err != nil {
			return err
		}

		defer f.Close()

		temp, err := template.New("").Parse(s.temp)
		if err != nil {
			return err
		}
		err = temp.Execute(f, struct {
			PackageName string `json:"package_name"`
		}{packageName})
		if err != nil {
			return err
		}
	}
	// 创建完成后在对应的位置插入结构代码
	for _, v := range pendingTemp {
		meta := packageInjectionMap[v.name]
		if err := ImportReference(meta.path, fmt.Sprintf(meta.importCodeF, v.name, packageName), fmt.Sprintf(meta.structNameF, caser.String(packageName)), fmt.Sprintf(meta.packageNameF, packageName), meta.groupName); err != nil {
			return err
		}
	}
	return nil
}

type Visitor struct {
	ImportCode  string
	StructName  string
	PackageName string
	GroupName   string
}

func (vi *Visitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.GenDecl:
		// 查找有没有import context包
		// Notice：没有考虑没有import任何包的情况
		if n.Tok == token.IMPORT && vi.ImportCode != "" {
			vi.addImport(n)
			// 不需要再遍历子树
			return nil
		}
		if n.Tok == token.TYPE && vi.StructName != "" && vi.PackageName != "" && vi.GroupName != "" {
			vi.addStruct(n)
			return nil
		}
	case *ast.FuncDecl:
		if n.Name.Name == "Routers" {
			vi.addFuncBodyVar(n)
			return nil
		}

	}
	return vi
}

func (vi *Visitor) addStruct(genDecl *ast.GenDecl) ast.Visitor {
	for i := range genDecl.Specs {
		switch n := genDecl.Specs[i].(type) {
		case *ast.TypeSpec:
			if strings.Index(n.Name.Name, "Group") > -1 {
				switch t := n.Type.(type) {
				case *ast.StructType:
					f := &ast.Field{
						Names: []*ast.Ident{
							{
								Name: vi.StructName,
								Obj: &ast.Object{
									Kind: ast.Var,
									Name: vi.StructName,
								},
							},
						},
						Type: &ast.SelectorExpr{
							X: &ast.Ident{
								Name: vi.PackageName,
							},
							Sel: &ast.Ident{
								Name: vi.GroupName,
							},
						},
					}
					t.Fields.List = append(t.Fields.List, f)
				}
			}
		}
	}
	return vi
}

func (vi *Visitor) addImport(genDecl *ast.GenDecl) ast.Visitor {
	// 是否已经import
	hasImported := false
	for _, v := range genDecl.Specs {
		importSpec := v.(*ast.ImportSpec)
		// 如果已经包含
		if importSpec.Path.Value == strconv.Quote(vi.ImportCode) {
			hasImported = true
		}
	}
	if !hasImported {
		genDecl.Specs = append(genDecl.Specs, &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(vi.ImportCode),
			},
		})
	}
	return vi
}

func (vi *Visitor) addFuncBodyVar(funDecl *ast.FuncDecl) ast.Visitor {
	hasVar := false
	for _, v := range funDecl.Body.List {
		switch varSpec := v.(type) {
		case *ast.AssignStmt:
			for i := range varSpec.Lhs {
				switch nn := varSpec.Lhs[i].(type) {
				case *ast.Ident:
					if nn.Name == vi.PackageName+"Router" {
						hasVar = true
					}
				}
			}
		}
	}
	if !hasVar {
		assignStmt := &ast.AssignStmt{
			Lhs: []ast.Expr{
				&ast.Ident{
					Name: vi.PackageName + "Router",
					Obj: &ast.Object{
						Kind: ast.Var,
						Name: vi.PackageName + "Router",
					},
				},
			},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.SelectorExpr{
					X: &ast.SelectorExpr{
						X: &ast.Ident{
							Name: "router",
						},
						Sel: &ast.Ident{
							Name: "RouterGroupApp",
						},
					},
					Sel: &ast.Ident{
						Name: caser.String(vi.PackageName),
					},
				},
			},
		}
		funDecl.Body.List = append(funDecl.Body.List, funDecl.Body.List[1])
		index := 1
		copy(funDecl.Body.List[index+1:], funDecl.Body.List[index:])
		funDecl.Body.List[index] = assignStmt
	}
	return vi
}

func ImportReference(filepath, importCode, structName, packageName, groupName string) error {
	fSet := token.NewFileSet()
	fParser, err := parser.ParseFile(fSet, filepath, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	importCode = strings.TrimSpace(importCode)
	v := &Visitor{
		ImportCode:  importCode,
		StructName:  structName,
		PackageName: packageName,
		GroupName:   groupName,
	}
	if importCode == "" {
		ast.Print(fSet, fParser)
	}

	ast.Walk(v, fParser)

	var output []byte
	buffer := bytes.NewBuffer(output)
	err = format.Node(buffer, fSet, fParser)
	if err != nil {
		log.Fatal(err)
	}
	// 写回数据
	return ioutil.WriteFile(filepath, buffer.Bytes(), 0o600)
}
