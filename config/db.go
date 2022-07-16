package config

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
