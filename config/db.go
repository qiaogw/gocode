package config

import "strconv"

func (m *GeneralDB) MysqlDsn() string {
	if len(m.Config) < 1 {
		m.Config = "charset=utf8&parseTime=True&loc=Local&timeout=1000ms"
	}
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + strconv.FormatInt(m.Port, 10) + ")/" + m.Dbname + "?" + m.Config
}

func (m *GeneralDB) GetLogMode() string {
	return m.LogMode
}

// PgsqlDsn 基于配置文件获取 dsn
func (m *GeneralDB) PgsqlDsn() string {
	if len(m.Config) < 1 {
		m.Config = "sslmode=disable TimeZone=Asia/Shanghai"
	}
	return "host=" + m.Path + " user=" + m.Username + " password=" + m.Password + " dbname=" + m.Dbname + " port=" + strconv.FormatInt(m.Port, 10) + " " + m.Config
}

// PgsqlLinkDsn 根据 dbname 生成 dsn
func (m *GeneralDB) PgsqlLinkDsn(dbname string) string {
	return "host=" + m.Path + " user=" + m.Username + " password=" + m.Password + " dbname=" + dbname + " port=" + strconv.FormatInt(m.Port, 10) + " " + m.Config
}
