package gormx

import "strconv"

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

	return user + ":" + password + "@tcp(" + host + ":" + strconv.FormatInt(port, 10) + ")/" + db + "?" + config
}

// pgsqlDsn 基于配置文件获取 dsn
func pgsqlDsn(host string, port int64, user, password, db, schema, config string) string {
	if len(config) < 1 {
		config = "sslmode=disable TimeZone=Asia/Shanghai"
	}
	return "host=" + host + " user=" + user + " password=" + password + " dbname=" + db + " port=" + strconv.FormatInt(port, 10) + " " + config
}
