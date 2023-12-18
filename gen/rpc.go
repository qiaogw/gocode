package gen

import (
	"github.com/qiaogw/gocode/model"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// CreateRpc 创建Rpc 代码
func (acd *AutoCodeService) CreateRpc(db *model.Db) (err error) {
	dataList, err := acd.genBefore(db.Database, rpcPath)
	if err != nil {
		return
	}
	// 生成文件
	for _, value := range dataList {
		f, err := os.OpenFile(value.autoCodePath, os.O_CREATE|os.O_WRONLY, 0o755)
		if err != nil {
			return err
		}
		if err = value.template.Execute(f, db); err != nil {
			//if err = value.template.Execute(f, db); err != nil {
			log.Printf("err is %v\n", err)
			return err
		}
		_ = f.Close()
	}

	defer func() { // 移除中间文件
		if err := os.RemoveAll(autoPath); err != nil {
			return
		}
	}()
	return acd.genAfter(dataList)
}

// CreateRpcLogic 创建Rpc 代码
func (acd *AutoCodeService) CreateRpcLogic(db *model.Db) (err error) {
	for _, v := range db.Tables {
		v.ParentPkg = db.ParentPkg
		v.Pkg = db.Pkg
		v.PackageName = db.Database
		if acd.Mode == "rpc" {
			err = acd.createRpcLogic(v)
			if err != nil {
				log.Printf("createRpcLogic err is %v\n", err)
				continue
			}
		}

		err = acd.createApiLogic(v)
		if err != nil {
			log.Printf("createApiLogic err is %v\n", err)
			continue
		}
	}
	return err
}

// createRpcLogic 创建 model 代码
func (acd *AutoCodeService) createRpcLogic(table *model.Table) (err error) {
	dataList, err := acd.genBefore(table.Table, rpcLogicPath)
	if err != nil {
		log.Printf("genBefore err is %v\n", err)
		return
	}
	if !table.IsImport {
		i := 0
		for _, v := range dataList {
			trimBase := filepath.Base(v.locationPath)
			fileSlice := strings.Split(trimBase, ".")
			importList := []string{"import", "export", "exporttemplate"}
			if !in(fileSlice[0], importList) {
				dataList[i] = v
				i++
			}
		}
		dataList = dataList[:i]
	}
	// 生成文件
	for i, value := range dataList {
		dataList[i].tablePkg = table.TableUrl
		f, err := os.OpenFile(value.autoCodePath, os.O_CREATE|os.O_WRONLY, 0o755)
		if err != nil {
			return err
		}
		if err = value.template.Execute(f, table); err != nil {
			log.Printf("err is %v\n", err)
			return err
		}
		_ = f.Close()
	}

	defer func() { // 移除中间文件
		if err := os.RemoveAll(autoPath); err != nil {
			return
		}
	}()
	err = acd.genAfter(dataList)
	if err != nil {
		return
	}
	return err
}

// createApiLogic 创建 model 代码
func (acd *AutoCodeService) createApiLogic(table *model.Table) (err error) {
	dataList, err := acd.genBefore(table.Table, apiLogicPath)
	if err != nil {
		log.Printf("genBefore err is %v\n", err)
		return
	}
	if !table.IsImport {
		i := 0
		for _, v := range dataList {
			trimBase := filepath.Base(v.locationPath)
			fileSlice := strings.Split(trimBase, ".")
			importList := []string{"import", "export", "exporttemplate"}
			if !in(fileSlice[0], importList) {
				dataList[i] = v
				i++
			}
		}
		dataList = dataList[:i]
	}

	// 生成文件
	for i, value := range dataList {
		dataList[i].tablePkg = table.TableUrl
		f, err := os.OpenFile(value.autoCodePath, os.O_CREATE|os.O_WRONLY, 0o755)
		if err != nil {
			return err
		}
		if err = value.template.Execute(f, table); err != nil {
			log.Printf("err is %v\n", err)
			return err
		}
		_ = f.Close()
	}

	defer func() { // 移除中间文件
		if err := os.RemoveAll(autoPath); err != nil {
			return
		}
	}()
	err = acd.genAfter(dataList)
	if err != nil {
		return
	}
	return err
}

func in(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}
func DeleteSlice(target interface{}, arr []interface{}) interface{} {
	j := 0
	for _, val := range arr {
		if val != target {
			arr[j] = val
			j++
		}
	}
	return arr[:j]
}
