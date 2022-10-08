package setting

import (
	"fmt"
	"github.com/qiaogw/gocode/global"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

// GormInit 初始化数据库并产生数据库全局变量
func GormInit() (*gorm.DB, error) {
	switch global.GenConfig.DB.DbType {
	case "mysql":
		return GormMysql()
	case "postgres":
		return GormPgSql()
	default:
		return GormMysql()
	}
}

// GormMysql 初始化Mysql数据库
func GormMysql() (*gorm.DB, error) {
	dbConf := global.GenConfig.DB
	// if dbConf.Dbname == "" {
	// 	return nil, errors.New("数据库不能为空，请修改配置文件")
	// }

	mysqlConfig := mysql.Config{
		DSN:                       dbConf.MysqlDsn(), // DSN data source name
		DefaultStringSize:         191,               // string 类型字段的默认长度
		SkipInitializeWithVersion: false,             // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   dbConf.TablePrefix,   // 表名前缀，`Article` 的表名应该是 `it_articles`
			SingularTable: dbConf.SingularTable, // 使用单数表名，启用该选项，此时，`Article` 的表名应该是 `it_article`
		},
	}); err != nil {
		fmt.Printf("gorm.Open err is %v", err)
		return nil, err
	} else {
		global.GenConfig.DB.DataSource = dbConf.MysqlDsn()
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConf.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConf.MaxOpenConns)
		return db, nil
	}
}

// GormPgSql 初始化 Postgresql 数据库
func GormPgSql() (*gorm.DB, error) {
	dbConf := global.GenConfig.DB
	// if dbConf.Dbname == "" {
	// 	return nil, errors.New("数据库不能为空，请修改配置文件")
	// }

	pgsqlConfig := postgres.Config{
		DSN:                  dbConf.PgsqlDsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   dbConf.TablePrefix,   // 表名前缀，`Article` 的表名应该是 `it_articles`
			SingularTable: dbConf.SingularTable, // 使用单数表名，启用该选项，此时，`Article` 的表名应该是 `it_article`
		},
	}); err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		global.GenConfig.DB.DataSource = dbConf.PgsqlDsn()
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConf.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConf.MaxOpenConns)
		return db, nil
	}
}
