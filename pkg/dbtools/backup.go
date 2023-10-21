package dbtools

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	jsoniter "github.com/json-iterator/go"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qiaogw/gocode/global"
	utils2 "github.com/qiaogw/gocode/util"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

func BackupDB(backupDir string) error {
	dbConf := global.GenConfig.DB
	var dsn string
	switch dbConf.DbType {
	case "mysql":
		dsn = dbConf.MysqlDsn()
	case "postgres":
		dsn = dbConf.PgsqlDsn()
	default:
		dsn = dbConf.MysqlDsn()
	}
	// 连接到数据库
	db, err := sql.Open(dbConf.DbType, dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	// 查询所有表名
	tableNames, err := getTableNames(db, dbConf.DbType)
	if err != nil {
		return err
	}
	wd, _ := os.Getwd()
	backupDir = filepath.Join(wd, backupDir, dbConf.Dbname)
	utils2.IsNotExistMkDir(backupDir)
	for _, tableName := range tableNames {
		// 查询表数据filepath.Join(wd, backupDir)
		data, err := queryTableData(db, tableName)
		if err != nil {
			log.Println(utils2.Red(fmt.Sprintf("查询表 %s: 数据错误：%v", tableName, err)))
			continue
		}

		// 将数据写入 JSON 文件
		jsonFileName := filepath.Join(backupDir, tableName+".json")
		jsonFile, err := os.Create(jsonFileName)
		if err != nil {
			return err
		}
		defer jsonFile.Close()

		encoder := json.NewEncoder(jsonFile)
		if err := encoder.Encode(data); err != nil {
			return err
		}
		log.Println(utils2.Green(fmt.Sprintf("表 %s 的数据存储到 %s", tableName, jsonFileName)))
	}
	return nil
}

func getTableNames(db *sql.DB, dbType string) ([]string, error) {
	// 查询所有表名
	query := "SHOW TABLES" // MySQL
	if dbType == "postgres" {
		query = "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname = 'public'" // PostgreSQL
	} else if dbType == "sqlite3" {
		query = "SELECT name FROM sqlite_master WHERE type='table'" // SQLite
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tableNames []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tableNames = append(tableNames, tableName)
	}
	return tableNames, nil
}

func queryTableData(db *sql.DB, tableName string) (interface{}, error) {
	// 编写 SQL 查询语句
	query := fmt.Sprintf("SELECT * FROM %s", tableName)

	// 执行查询
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 将查询结果转为 JSON 格式
	var result []map[string]interface{}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
				getType := reflect.TypeOf(val)
				fmt.Println(col)
				fmt.Println(col, " 类型:", getType)

				if getType != nil && getType.String() == "time.Time" {
					v = val.(time.Time).Format("2006-01-02 15:04:05")
				}
			}
			entry[col] = v
		}
		result = append(result, entry)
	}

	return result, nil
}

func RestoreData(backupFolder string) error {
	dbConf := global.GenConfig.DB
	var dsn string
	switch dbConf.DbType {
	case "mysql":
		dsn = dbConf.MysqlDsn()
	case "postgres":
		dsn = dbConf.PgsqlDsn()
	default:
		dsn = dbConf.MysqlDsn()
	}
	// 连接到数据库
	db, err := sql.Open(dbConf.DbType, dsn)
	if err != nil {
		return err
	}
	fmt.Printf(utils2.Green(fmt.Sprintf("数据库连接成功，类型为：%s,地址为：%s:%v,数据库为：%s\n",
		global.GenDB.Name(), global.GenConfig.DB.Path, global.GenConfig.DB.Port, global.GenConfig.DB.Dbname)))

	defer db.Close()

	wd, _ := os.Getwd()
	backupFolder = filepath.Join(wd, backupFolder)
	// 打开备份文件夹
	folder, err := os.Open(backupFolder)
	if err != nil {
		return err
	}
	defer folder.Close()

	// 遍历备份文件夹中的文件
	fileInfos, err := folder.Readdir(-1)
	if err != nil {
		return err
	}

	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			// 处理每个文件
			fileName := fileInfo.Name()
			tableName := strings.TrimSuffix(fileName, ".json")

			jsonData, err := readJSONFile(backupFolder + "/" + fileName)
			if err != nil {
				log.Printf("读取 JSON 文件 %s 错误: %v\n", fileName, err)
				continue
			}
			err = insertDatasIntoTable(db, tableName, dbConf.DbType, jsonData)
			if err != nil {
				log.Printf("表  %s 数据导入错误: %v\n", tableName, err)
				continue
			}
			//for _, data := range jsonData {
			//	if err := insertDataIntoTable(db, tableName, dbConf.DbType, data); err != nil {
			//		log.Printf("表  %s 数据导入错误: %v", tableName, err)
			//		continue
			//	}
			//}
			log.Println(utils2.Green(fmt.Sprintf("数据从 %s 恢复到表 %s", fileName, tableName)))
			if dbConf.DbType == "postgres" {
				maxID, err := getMaxID(db, tableName)
				if err != nil {
					continue
				}
				// 将序列值设置为最大主键值加一
				err = setAutoIncrement(db, tableName, maxID+1)
				if err != nil {
					log.Printf("设置表 %s 的自增主键错误: %v\n", tableName, err)
					continue
				}
				log.Println(utils2.Green(fmt.Sprintf("表 %s 自增序列更新：%d", fileName, maxID+1)))
			}
		}
	}
	return nil
}

func readJSONFile(fileName string) ([]map[string]interface{}, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := jsoniter.NewDecoder(file)
	var data []map[string]interface{}
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func insertDataIntoTable(db *sql.DB, tableName, dbType string, data map[string]interface{}) error {
	// 构建插入语句
	columns := make([]string, 0)
	values := make([]interface{}, 0)
	placeholders := make([]string, 0)

	for column, value := range data {
		columns = append(columns, column)
		values = append(values, value)
		placeholders = append(placeholders, "?")
	}
	columnNames := strings.Join(columns, ",")

	fmtQuery := `INSERT INTO "%s" (%s) VALUES (%s);`
	// 将占位符替换为数据库特定的占位符
	switch dbType {
	case "mysql":
		fmtQuery = "INSERT INTO `%s` (%s) VALUES (%s);"
	case "postgres":
		// 替换占位符为 $1, $2, ...
		for i := range placeholders {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
		}
	case "sqlite3":
		// 不需要替换，SQLite 使用 ? 作为占位符
	}
	valuePlaceholders := strings.Join(placeholders, ",")
	query := fmt.Sprintf(fmtQuery,
		tableName, columnNames, valuePlaceholders)

	// 执行插入操作
	stmt, err := db.Prepare(query)

	if err != nil {
		log.Printf("sql语句: %s,错误: %v \n", query, err) // 记录 SQL 查询及参数
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(values...)
	if err != nil {
		log.Printf("sql语句执行错误: %s, Values: %v,错误: %v\n", query, values, err) // 记录 SQL 查询及参数
	}
	return err
}

func getMaxID(db *sql.DB, tableName string) (int, error) {
	var maxID int
	query := fmt.Sprintf("SELECT MAX(id) FROM %s", tableName)
	err := db.QueryRow(query).Scan(&maxID)
	return maxID, err
}

func setAutoIncrement(db *sql.DB, tableName string, value int) error {
	// Escape table name to prevent SQL injection
	escapedTableName := fmt.Sprintf(`%s`, tableName)
	query := fmt.Sprintf(`SELECT setval('%s_id_seq', %d, true);`, escapedTableName, value)
	_, err := db.Exec(query)

	return err
}

func insertDatasIntoTable(db *sql.DB, tableName, dbType string, data []map[string]interface{}) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // 事务出错时回滚

	for _, rowData := range data {
		columns := make([]string, 0)
		placeholders := make([]string, 0)
		values := make([]interface{}, 0)

		for column, value := range rowData {
			columns = append(columns, column)
			placeholders = append(placeholders, "?")
			values = append(values, value)
		}

		fmtQuery := `INSERT INTO "%s" (%s) VALUES (%s);`
		// 将占位符替换为数据库特定的占位符
		switch dbType {
		case "mysql":
			fmtQuery = "INSERT INTO `%s` (%s) VALUES (%s);"
		case "postgres":
			// 替换占位符为 $1, $2, ...
			for i := range placeholders {
				placeholders[i] = fmt.Sprintf("$%d", i+1)
			}
		case "sqlite3":
			// 不需要替换，SQLite 使用 ? 作为占位符
		}
		columnNames := strings.Join(columns, ",")
		valuePlaceholders := strings.Join(placeholders, ",")

		query := fmt.Sprintf(fmtQuery,
			tableName, columnNames, valuePlaceholders)
		_, err = tx.Exec(query, values...)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
