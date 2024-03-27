package gen

import (
	"fmt"
	"github.com/qiaogw/gocode/global"
	"github.com/qiaogw/gocode/model"

	"github.com/qiaogw/gocode/util"

	"github.com/zeromicro/go-zero/tools/goctl/pkg/golang"
	"log"
	"os"
	"strings"
)

// Code 生成代码
func (acd *AutoCodeService) Code(modeGen bool) (db model.Db, tables []model.Table, err error) {
	acd.Init()
	Check()
	fmt.Printf(util.Green(fmt.Sprintf("数据库连接成功，类型为：%s,地址为：%s:%v,数据库为：%s\n",
		global.GenDB.Name(), global.GenConfig.DB.Path, global.GenConfig.DB.Port, global.GenConfig.DB.Dbname)))
	tables, err = acd.DB.GetTables(global.GenConfig.DB.Dbname)
	if err != nil {
		log.Println(util.Red(fmt.Sprintf("获取表 err is %v", err)))
		return
	}

	db.Database = global.GenConfig.System.Name
	db.Package = strings.ToLower(db.Database)
	db.Service = util.LeftUpper(util.CamelString(db.Database))
	db.Option = global.GenConfig
	db.DriverName = global.GenDB.Name()
	dir, _ := os.Getwd()
	pkg, err := golang.GetParentPackage(dir)
	if err != nil {
		log.Println(util.Red(fmt.Sprintf("GetParentPackage err is %v", err)))
		return
	}

	db.ParentPkg = pkg + "/" + global.GenConfig.AutoCode.Pkg
	if acd.Mode == "api" {
		db.ParentPkg = global.GenConfig.AutoCode.Pkg
	}

	db.Pkg = pkg
	//fmt.Printf("Pkg:%s,ParentPkg:%s,Package:%s\n", db.Pkg, db.ParentPkg, db.Package)
	db.RpcHost = global.GenConfig.System.RpcHost
	db.RpcPort = global.GenConfig.System.RpcPort
	db.ApiHost = global.GenConfig.System.ApiHost
	db.ApiPort = global.GenConfig.System.ApiPort

	if err != nil {
		log.Println(util.Red(fmt.Sprintf("CreateConfigFile err is %v", err)))
		return
	}
	var tbList []model.Table
	for _, v := range tables {
		if !strings.HasPrefix(v.Table, global.GenConfig.DB.TablePrefix) {
			continue
		}
		columnData, err := acd.DB.GetColumn(global.GenConfig.DB.Dbname, v.Table)
		if err != nil {
			log.Println(util.Red(fmt.Sprintf("获取字段 err is %v", err)))
			continue
		}

		tb, err := columnData.Convert(v.TableComment)

		if err != nil {
			fmt.Println(util.Red(fmt.Sprintf("数据生成错误错误: %v", err)))
			continue
		}

		if tb.HasTimer {
			db.HasTimer = true
		}
		tb.ParentPkg = db.ParentPkg
		tb.Pkg = db.Pkg

		tb.IsAuth = true
		tb.IsImport = true
		tb.IsFlow = global.GenConfig.AutoCode.IsFlow
		if tb.IsFlow {
			db.IsFlow = true
		}
		tb.Dir = strings.ToLower(db.Database)
		if modeGen {
			err = acd.CreateModel(tb)
			if err != nil {
				continue
			}
		} else {
			err = acd.CreateModelZero(tb)
			if err != nil {
				continue
			}
		}

		db.Tables = append(db.Tables, tb)
		db.Email = tb.Email
		db.Author = tb.Author
		err = acd.CreateApiDesc(tb)
		if err != nil {
			log.Printf("CreateApiDesc err is %v\n", err)
		}
		tbList = append(tbList, *tb)

	}

	err = acd.CreateApi(&db)
	if err != nil {
		log.Printf("CreateApi err is %v\n", err)
		return
	}
	if acd.Mode == "rpc" {
		err = acd.CreateRpc(&db)
		if err != nil {
			log.Printf("CreateRpc err is %v\n", err)
			return
		}

	}
	err = acd.CreateRpcLogic(&db)
	if err != nil {
		log.Printf("CreateRpcLogic err is %v\n", err)
		return
	}
	err = acd.CreateWeb(&db)
	if err != nil {
		log.Printf("CreateRpcLogic err is %v\n", err)
		return
	}
	if acd.Mode == "api" {
		err = acd.CreateAdminFile()
		if err != nil {
			log.Printf("CreateAdminFile err is %v\n", err)
			return
		}
	}
	fmt.Println(util.Green("Done!"))
	return db, tbList, nil
}

func Check() {
	if len(global.GenConfig.System.RpcHost) < 1 {
		global.GenConfig.System.RpcHost = "0.0.0.0"
	}
	if global.GenConfig.System.RpcPort < 1 {
		global.GenConfig.System.RpcPort = 7000
	}
	if len(global.GenConfig.System.ApiHost) < 1 {
		global.GenConfig.System.RpcHost = "0.0.0.0"
	}
	if global.GenConfig.System.ApiPort < 1 {
		global.GenConfig.System.RpcPort = 7001
	}
}
