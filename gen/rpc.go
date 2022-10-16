package gen

import (
	"github.com/qiaogw/gocode/model"
	"log"
	"os"
)

// CreateRpc 创建Rpc 代码
func (acd *AutoCodeService) CreateRpc(db *model.Db, ids ...uint) (err error) {
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
	return acd.genAfter(dataList, ids...)
}

// CreateRpcLogic 创建Rpc 代码
func (acd *AutoCodeService) CreateRpcLogic(db *model.Db, ids ...uint) (err error) {
	for _, v := range db.Tables {
		v.ParentPkg = db.ParentPkg
		v.PKG = db.PKG
		err = acd.createRpcLogic(v)
		if err != nil {
			log.Printf("createRpcLogic err is %v\n", err)
			continue
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
func (acd *AutoCodeService) createRpcLogic(table *model.Table, ids ...uint) (err error) {
	dataList, err := acd.genBefore(table.Table, rpcLogicPath)
	if err != nil {
		log.Printf("genBefore err is %v\n", err)
		return
	}
	// 生成文件
	for _, value := range dataList {
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
	err = acd.genAfter(dataList, ids...)
	if err != nil {
		return
	}
	return err
}

// createApiLogic 创建 model 代码
func (acd *AutoCodeService) createApiLogic(table *model.Table, ids ...uint) (err error) {
	dataList, err := acd.genBefore(table.Table, apiLogicPath)
	if err != nil {
		log.Printf("genBefore err is %v\n", err)
		return
	}
	// 生成文件
	for _, value := range dataList {
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
	err = acd.genAfter(dataList, ids...)
	if err != nil {
		return
	}
	return err
}
