package gen

import (
	"fmt"
	"github.com/qiaogw/gocode/global"
	"github.com/qiaogw/gocode/model"
	utils2 "github.com/qiaogw/gocode/util"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/golang"
	"log"
	"os"
	"strings"
)

func (acd *AutoCodeService) Code() (*model.Db, *[]model.Table, error) {
	acd.Init()
	fmt.Printf(utils2.Green(fmt.Sprintf("数据库连接成功，类型为：%s,地址为：%s:%v,数据库为：%s\n",
		global.GenDB.Name(), global.GenConfig.DB.Path, global.GenConfig.DB.Port, global.GenConfig.DB.Dbname)))
	tables, err := acd.DB.GetTables(global.GenConfig.DB.Dbname)
	if err != nil {
		log.Println(utils2.Red(fmt.Sprintf("获取表 err is %v", err)))
		return nil, nil, err
	}
	var db model.Db
	db.Database = global.GenConfig.System.Name
	db.Package = strings.ToLower(db.Database)
	db.Service = utils2.LeftUpper(db.Database)

	db.Option = global.GenConfig
	db.DriverName = global.GenDB.Name()
	dir, _ := os.Getwd()
	pkg, err := golang.GetParentPackage(dir)
	if err != nil {
		log.Println(utils2.Red(fmt.Sprintf("GetParentPackage err is %v", err)))
		return nil, nil, err
	}
	db.ParentPkg = pkg + "/" + global.GenConfig.AutoCode.Pkg
	db.PKG = pkg
	if err != nil {
		log.Println(utils2.Red(fmt.Sprintf("CreateConfigFile err is %v", err)))
		return nil, nil, err
	}
	for _, v := range tables {
		if !strings.HasPrefix(v.Table, global.GenConfig.DB.TablePrefix) {
			continue
		}
		columnData, err := acd.DB.GetColumn(global.GenConfig.DB.Dbname, v.Table)
		if err != nil {
			log.Println(utils2.Red(fmt.Sprintf("获取字段 err is %v", err)))
			continue
		}
		tb, err := columnData.Convert(v.TableComment)
		if err != nil {
			fmt.Println(utils2.Red(fmt.Sprintf("数据生成错误错误: %v", err)))
			continue
		}
		if tb.HasTimer {
			db.HasTimer = true
		}
		tb.ParentPkg = db.ParentPkg
		tb.PKG = db.PKG
		err = acd.CreateModel(tb)
		if err != nil {
			continue
		}

		db.Tables = append(db.Tables, tb)
		db.Email = tb.Email
		db.Author = tb.Author
		err = acd.CreateApiDesc(tb)
		if err != nil {
			log.Printf("CreateApiDesc err is %v\n", err)
		}
	}

	err = acd.CreateRpc(&db)
	if err != nil {
		log.Printf("CreateRpc err is %v\n", err)
		return nil, nil, err
	}
	err = acd.CreateApi(&db)
	if err != nil {
		log.Printf("CreateApi err is %v\n", err)
		return nil, nil, err
	}
	err = acd.CreateRpcLogic(&db)
	if err != nil {
		log.Printf("CreateRpcLogic err is %v\n", err)
		return nil, nil, err
	}
	err = acd.CreateWeb(&db)
	if err != nil {
		log.Printf("CreateRpcLogic err is %v\n", err)
		return nil, nil, err
	}
	//err = genApp.CreateCommon(&db)
	//if err != nil {
	//	log.Printf("CreateCommon err is %v\n", err)
	//	return err
	//}
	err = acd.CreateConfigFile(&db, global.GenConfig.AutoCode.Root)
	fmt.Println(utils2.Green("Done!"))
	return &db, &tables, err
}
