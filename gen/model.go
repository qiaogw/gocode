package gen

import (
	"fmt"
	"github.com/qiaogw/gocode/global"
	"github.com/qiaogw/gocode/model"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// CreateModel 创建 gorm model 代码
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

// CreateModelZero 创建 zero model 代码
func (acd *AutoCodeService) CreateModelZero(table *model.Table) (err error) {
	dsn := global.GenConfig.DB.MysqlDsn()

	dir := filepath.Join(global.GenConfig.AutoCode.Root, "model")
	drivers := "mysql"
	if table.PostgreSql {
		df := global.GenConfig.DB
		dsn = fmt.Sprintf(`postgres://%s:%s@%s:%s/%s?sslmode=disable`,
			df.Username, df.Password, df.Path, df.Port, df.Dbname)
		drivers = "pg"
	}

	cmd := exec.Command("goctl", "model", drivers, "datasource",
		fmt.Sprintf("--url=%s", dsn),
		fmt.Sprintf("--table=%s", table.Name),
		fmt.Sprintf("--dir=%s", dir),
		"-c")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Errorf("error running command: %v", err)
	}
	return nil
}
