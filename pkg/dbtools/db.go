package dbtools

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/qiaogw/gocode/database/admin"
	"github.com/qiaogw/gocode/database/gencode"
	"github.com/qiaogw/gocode/global"
	"github.com/qiaogw/gocode/pkg/zip"
	"github.com/qiaogw/gocode/setting"
	"github.com/qiaogw/gocode/util"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"path/filepath"
	"time"
)

const (
	dbPingInterval = 90 * time.Second
)

var opens = map[string]func(string) gorm.Dialector{
	"mysql":    mysql.Open,
	"postgres": postgres.Open,
	"sqlite3":  sqlite.Open,
}

var tbs = [...]schema.Tabler{
	new(admin.Api),
	new(admin.Dept),
	new(admin.Config),
	new(admin.Menu),
	new(admin.User),
	new(admin.Role),
	new(admin.DictType),
	new(admin.DictData),
	new(admin.LoginLog),
	new(admin.DictData),
	new(admin.DictType),
	new(admin.Migration),
	new(admin.OperaLog),
	new(admin.Post),
	new(gencode.GenPkg),
	new(gencode.GenSource),
	new(gencode.GenTable),
	new(gencode.Column),
}

// GetDB 创建Db
func GetDB() (*gorm.DB, error) {
	ed, err := setting.GormInit()
	if err != nil {
		return nil, err
	}
	global.GenDB = ed
	logx.Infof("数据库链接成功 %s ... ", global.GenDB.Name())
	return ed, nil
}

func BackupTable(db *gorm.DB) (tb []string, err error) {
	//setting := Setting.DBSRC
	engine := db.Config.Name()
	sql := ""
	switch engine {
	case "mysql":
		sql = fmt.Sprintf("SELECT TABLE_NAME FROM INFORMATION_SCHEMA."+
			"TABLES WHERE TABLE_SCHEMA='%s'", db.Migrator().CurrentDatabase())
	case "postgres":
		sql = fmt.Sprintf("select tablename from pg_tables where schemaname='public'")
	case "sqlite3":
		sql = fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' ORDER BY name;")
	}
	err = db.Raw(sql).Find(&tb).Error
	return
}

// BackupDBDataToJson 	数据导出json
func BackupDBDataToJson(db *gorm.DB, tb interface{}, tableName, backupDir string) error {
	var ob []map[string]interface{}
	wd, _ := os.Getwd()
	backFile := filepath.Join(wd, backupDir, tableName+".json")
	util.IsNotExistMkDir(filepath.Join(wd, backupDir))
	// 查询表数据并将结果存储在通用切片中
	var data1 []map[string]interface{}
	if err := db.Table(tableName).Find(&data1).Error; err != nil {
		fmt.Printf("DumpTableToJSON err for table %s: %v\n", tableName, err)
		return err
	}

	// 使用 Delete 方法删除已软删除的数据
	db.Unscoped().Where("deleted_at IS NOT NULL").Delete(tb)
	rows, err := db.Table(tableName).Rows()

	defer rows.Close()
	if err == nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		for rows.Next() {
			db.ScanRows(rows, tb)
			var o map[string]interface{}
			es, _ := json.Marshal(tb)
			_ = json.Unmarshal(es, &o)
			ob = append(ob, o)
		}
	}
	data, _ := json.Marshal(ob)
	return os.WriteFile(backFile, data, os.ModePerm)
}

// DumpTableToJSON 将数据库表数据存储为JSON文件
func DumpTableToJSON(db *gorm.DB, tb interface{}, tableName string, filePath string) error {
	// 查询表数据
	var ob []map[string]interface{}
	var data []map[string]interface{}
	err := db.Table(tableName).Find(&data).Error
	if err == nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		for _, v := range data {
			var o map[string]interface{}
			es1, _ := json.Marshal(v)
			_ = json.Unmarshal(es1, tb)
			if tableName == "sys_role" {
				fmt.Printf("数据1:%+v\r\n", v)
				fmt.Printf("数据2:%+v\r\n", string(es1))
				fmt.Printf("数据3:%+v\r\n", tb)
			}
			es, _ := json.Marshal(tb)
			_ = json.Unmarshal(es, &o)
			ob = append(ob, o)
		}
	}
	// 将数据转换为JSON
	jsonData, err := json.Marshal(ob)
	if err != nil {
		fmt.Println("Marshal err:", err)
		return err
	}
	wd, _ := os.Getwd()
	backFile := filepath.Join(wd, filePath, tableName+".json")
	util.IsNotExistMkDir(filepath.Join(wd, filePath))
	// 写入JSON数据到文件
	file, err := os.Create(backFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}
func Migrate(dbPath string, db *gorm.DB) error {
	err := db.Migrator().AutoMigrate(tbs)

	for _, v := range tbs {
		v.TableName()
		err = InitData(dbPath, db, v, v.TableName())
	}
	return err
}

// CreateDatabaseSql 创建数据库
func CreateDatabaseSql() (sqlstring string, err error) {
	dname := global.GenDB.Name()
	switch global.GenConfig.DB.DbType {
	case "mysql":
		sqlstring = fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci;", dname)
	case "postgres":
		sqlstring = fmt.Sprintf("CREATE DATABASE %s;", dname)
	case "sqlite3":
		dns := filepath.Join(global.GenConfig.DB.Path, dname+".db")
		os.Remove(dns)
		sqlstring = "create table init (n varchar(32));drop table init;"
	default:
		logx.Error("Database driver is not allowed:", global.GenConfig.DB.DbType)
		err = errors.New("Database driver is not allowed:" + global.GenConfig.DB.DbType)
		return
	}
	return
}

func Backup(db *gorm.DB, backupDir string) error {
	for _, v := range tbs {
		err := BackupDBDataToJson(db, v, v.TableName(), backupDir)
		if err != nil {
			logx.Error("BackupDBDataToJson失败", err)
			return err
		}
	}
	return nil
}

func InitData(dataPath string, db *gorm.DB, tb interface{},
	tname string) (err error) {
	wd, _ := os.Getwd()
	operationFile := filepath.Join(wd, dataPath, tname+".json")
	jdata, _ := os.ReadFile(operationFile)
	sql := fmt.Sprintf("delete from %s", tname)
	db.Exec(sql)
	var es []interface{}
	err = json.Unmarshal(jdata, &es)
	if err != nil || len(es) < 1 {
		fmt.Println(err)
		//return err
	}
	i := 0
	for _, v := range es {
		b, _ := json.Marshal(v)
		err = json.Unmarshal(b, &tb)
		if err != nil || len(es) < 1 {
			fmt.Println(err)
			continue
		}
		err = db.Create(tb).Error
		if err != nil {
			fmt.Println(tname, "; err is", err)
		}
		i++
	}
	fmt.Printf("%s 创建成功，并有 %d 条数据初始化成功！！\n", tname, i)
	return
}

func UnzipDbData(dbDir string) {
	wd, _ := os.Getwd()
	dstFile := filepath.Join(wd, dbDir)
	srcFile := filepath.Join(wd, dbDir+".zip")
	_ = zip.UnzipDir(srcFile, dstFile)
}
