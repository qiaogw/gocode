package gen

import (
	"github.com/qiaogw/gocode/model"
	"log"
	"os"
)

// CreateModel 创建 model 代码
func (acd *AutoCodeService) CreateModel(table *model.Table) (err error) {
	dataList, err := acd.genBefore(table.Table, modelPath)
	//log.Printf("dataList is %+v\n", dataList)
	if err != nil {
		log.Printf("err is %v\n", err)
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
	err = acd.genAfter(dataList)
	if err != nil {
		return
	}
	return err
}
