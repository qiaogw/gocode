package configx

import (
	"errors"
	"fmt"
	"github.com/qiaogw/gocode/common/gormx/plugins"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// Mysql 定义了连接 MySQL 数据库所需的配置信息
type Mysql struct {
	Driver        string
	Host          string // 服务器地址
	Port          int    `json:",default=3306"` // 数据库端口，默认 3306
	Dbname        string // 数据库名称
	Username      string // 数据库用户名
	Password      string // 数据库密码
	Config        string `json:",default=charset%3Dutf8mb4%26parseTime%3Dtrue%26loc%3DLocal"` // 高级配置参数
	MaxIdleConns  int    `json:",default=10"`                                                 // 最大空闲连接数
	MaxOpenConns  int    `json:",default=10"`                                                 // 最大打开连接数
	LogMode       string `json:",default=dev,options=dev|test|prod|silent"`                   // 日志模式：dev、test、prod 或 silent
	LogColorful   bool   `json:",default=false"`                                              // 是否启用日志彩色输出
	SlowThreshold int64  `json:",default=1000"`                                               // 慢查询阈值（单位：毫秒）
}

// Dsn 根据 Mysql 配置生成连接 MySQL 的 DSN（数据源名称）字符串
func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + fmt.Sprintf("%d", m.Port) + ")/" + m.Dbname + "?" + m.Config
}

// GetGormLogMode 根据配置返回 Gorm 的日志级别
func (m *Mysql) GetGormLogMode() logger.LogLevel {
	return OverwriteGormLogMode(m.LogMode)
}

// GetSlowThreshold 返回慢查询阈值，并将其转换为 time.Duration 类型（单位为毫秒）
func (m *Mysql) GetSlowThreshold() time.Duration {
	return time.Duration(m.SlowThreshold) * time.Millisecond
}

// GetColorful 返回是否启用日志彩色输出的配置
func (m *Mysql) GetColorful() bool {
	return m.LogColorful
}

// Connect 根据 Mysql 配置连接 MySQL 数据库，返回 *gorm.DB 对象以及可能出现的错误
func (m *Mysql) Connect() (*gorm.DB, error) {
	// 如果数据库名称为空，则返回错误
	if m.Dbname == "" {
		return nil, errors.New("database name is empty")
	}
	// 构造 mysql 驱动所需的配置
	mysqlCfg := mysql.Config{
		DSN: m.Dsn(),
	}
	// 创建默认的 Gorm 日志实例
	newLogger := NewDefaultGormLogger(m)
	// 使用 gorm.Open 打开数据库连接
	db, err := gorm.Open(mysql.New(mysqlCfg), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	} else {
		// 获取底层数据库连接对象，并设置连接池参数
		sqldb, _ := db.DB()
		sqldb.SetMaxIdleConns(m.MaxIdleConns)
		sqldb.SetMaxOpenConns(m.MaxOpenConns)
		return db, nil
	}
}

// ConnectWithConfig 根据 Mysql 配置和自定义的 gorm.Config 连接 MySQL 数据库，初始化插件后返回 *gorm.DB 对象及可能出现的错误
func (m *Mysql) ConnectWithConfig(cfg *gorm.Config) (*gorm.DB, error) {
	// 如果数据库名称为空，则返回错误
	if m.Dbname == "" {
		return nil, errors.New("database name is empty")
	}
	// 构造 mysql 驱动所需的配置
	mysqlCfg := mysql.Config{
		DSN: m.Dsn(),
	}
	// 使用自定义配置连接数据库
	db, err := gorm.Open(mysql.New(mysqlCfg), cfg)
	if err != nil {
		return nil, err
	}

	// 初始化数据库插件
	err = plugins.InitPlugins(db)
	if err != nil {
		return nil, err
	}

	// 获取底层数据库连接对象，并设置连接池参数
	sqldb, _ := db.DB()
	sqldb.SetMaxIdleConns(m.MaxIdleConns)
	sqldb.SetMaxOpenConns(m.MaxOpenConns)
	return db, nil
}
