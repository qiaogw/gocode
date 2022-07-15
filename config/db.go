package config

import (
	"gocode/global"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (m *GeneralDB) MysqlDsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *GeneralDB) GetLogMode() string {
	return m.LogMode
}

// PgsqlDsn 基于配置文件获取 dsn
func (m *GeneralDB) PgsqlDsn() string {
	return "host=" + m.Path + " user=" + m.Username + " password=" + m.Password + " dbname=" + m.Dbname + " port=" + m.Port + " " + m.Config
}

// PgsqlLinkDsn 根据 dbname 生成 dsn
func (m *GeneralDB) PgsqlLinkDsn(dbname string) string {
	return "host=" + m.Path + " user=" + m.Username + " password=" + m.Password + " dbname=" + dbname + " port=" + m.Port + " " + m.Config
}

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	m := global.GenConfig.DB
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.MysqlDsn(), // DSN data source name
		DefaultStringSize:         191,          // string 类型字段的默认长度
		SkipInitializeWithVersion: false,        // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// GormPgSql 初始化 Postgresql 数据库
func GormPgSql() *gorm.DB {
	p := global.GenConfig.DB
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.PgsqlDsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		return db
	}
}
