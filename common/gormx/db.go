package gormx

import (
	"fmt"
)

// GetDsn 基于配置文件获取 dsn
func GetDsn(driver, host string, port int64, user, password, db, schema, config string) string {
	switch driver {
	case "mysql":
		return mysqlDsn(host, port, user, password, db, schema, config)
	case "postgres":
		return pgsqlDsn(host, port, user, password, db, schema, config)
	default:
		return mysqlDsn(host, port, user, password, db, schema, config)
	}
}
func mysqlDsn(host string, port int64, user, password, db, schema, config string) string {
	if len(config) < 1 {
		config = "charset=utf8&parseTime=True&loc=Local&timeout=1000ms"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		user, password, host, port, db, config)
	return dsn
}

// pgsqlDsn 基于配置文件获取 dsn
func pgsqlDsn(host string, port int64, user, password, db, schema, config string) string {
	if len(config) < 1 {
		config = "sslmode=disable TimeZone=Asia/Shanghai"
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname='%s' search_path=%s port=%d %s",
		host, user, password, db, schema, port, config)
	return dsn
}
