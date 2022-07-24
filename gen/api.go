package gen

import (
	"github.com/qiaogw/gocode/model"
	"log"
	"os"
)

// CreateApi 创建 Api
func (acd *AutoCodeService) CreateApi(db *model.Db, ids ...uint) (err error) {
	dataList, err := acd.genBefore(db.Database, apiPath)
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

	return acd.genAfter(dataList, db.Database, ids...)
}
