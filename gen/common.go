package gen

import (
	"github.com/qiaogw/gocode/model"
	"log"
	"os"
)

// CreateCommon 创建 Common
func (acd *AutoCodeService) CreateCommon(db *model.Db) (err error) {
	dataList, err := acd.genBefore(db.Database, commonPath)
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
