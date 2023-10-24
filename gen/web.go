package gen

import (
	"github.com/qiaogw/gocode/model"
	"log"
	"os"
)

// CreateWeb 创建 web 代码
func (acd *AutoCodeService) CreateWeb(db *model.Db) (err error) {
	for _, v := range db.Tables {
		v.ParentPkg = db.ParentPkg
		v.PKG = db.PKG
		err = acd.createWeb(v)
		if err != nil {
			log.Printf("CreateWeb err is %v\n", err)
			continue
		}
	}
	return err
}

func (acd *AutoCodeService) createWeb(table *model.Table) (err error) {
	dataList, err := acd.genBefore(table.Table, webPath)
	//log.Printf("dataList is %+v\n", dataList)
	if err != nil {
		log.Printf("err is %v\n", err)
		return
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
