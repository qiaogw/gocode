package configx

import (
	"errors"
	"fmt"
	"github.com/qiaogw/gocode/common/gormx/plugins"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// PgSql 定义了 PostgreSQL 数据库的配置信息
type PgSql struct {
	Driver        string
	Host          string // 数据库主机地址
	Port          int    `json:",default=5432"` // 数据库端口，默认 5432
	Username      string // 数据库用户名
	Password      string // 数据库密码
	Dbname        string // 数据库名称
	TimeZone      string `json:",default=Asia/Shanghai"`                    // 数据库时区，默认 Asia/Shanghai
	SslMode       string `json:",default=disable,options=disable|enable"`   // SSL 模式，支持 disable 或 enable，默认 disable
	MaxIdleConns  int    `json:",default=10"`                               // 空闲连接数的最大值
	MaxOpenConns  int    `json:",default=10"`                               // 打开数据库连接的最大数
	LogMode       string `json:",default=dev,options=dev|test|prod|silent"` // 日志模式，取值范围为 dev、test、prod、silent，默认 dev
	LogColorful   bool   `json:",default=false"`                            // 是否启用日志彩色输出，默认 false
	SlowThreshold int64  `json:",default=1000"`                             // 慢 SQL 阈值（毫秒）
	Schema        string `json:",default=public"`
}

// Dsn 根据 PgSql 配置生成 PostgreSQL 的 DSN（数据源名称）
func (m *PgSql) Dsn() string {
	if len(m.Schema) < 1 {
		m.Schema = "public"
	}
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d search_path=%s sslmode=%s TimeZone=%s",
		m.Username, m.Password, m.Dbname, m.Host, m.Port, m.Schema, m.SslMode, m.TimeZone)
}

// GetGormLogMode 根据配置返回 Gorm 的日志级别
func (m *PgSql) GetGormLogMode() logger.LogLevel {
	return OverwriteGormLogMode(m.LogMode)
}

// GetSlowThreshold 返回慢 SQL 阈值，并转换为 time.Duration 类型（单位为毫秒）
func (m *PgSql) GetSlowThreshold() time.Duration {
	return time.Duration(m.SlowThreshold) * time.Millisecond
}

// GetColorful 返回是否启用日志彩色输出的配置
func (m *PgSql) GetColorful() bool {
	return m.LogColorful
}

// Connect 根据 PgSql 配置连接 PostgreSQL 数据库，返回 *gorm.DB 对象和可能出现的错误
func (m *PgSql) Connect() (*gorm.DB, error) {
	// 如果数据库名称为空，则返回错误
	if m.Dbname == "" {
		return nil, errors.New("database name is empty")
	}
	// 使用默认的 Gorm 日志配置创建 Logger
	newLogger := NewDefaultGormLogger(m)
	// 构造 postgres 配置
	pgsqlCfg := postgres.Config{
		DSN:                  m.Dsn(),
		PreferSimpleProtocol: true, // 禁用隐式预处理语句的使用
	}
	// 使用 gorm.Open 打开数据库连接
	db, err := gorm.Open(postgres.New(pgsqlCfg), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	} else {
		// 获取底层数据库连接对象，并设置最大空闲连接数和最大打开连接数
		sqldb, _ := db.DB()
		sqldb.SetMaxIdleConns(m.MaxIdleConns)
		sqldb.SetMaxOpenConns(m.MaxOpenConns)
		return db, nil
	}
}

// ConnectWithConfig 根据 PgSql 配置以及自定义的 gorm.Config 连接 PostgreSQL 数据库
// 并初始化插件，返回 *gorm.DB 对象和可能出现的错误
func (m *PgSql) ConnectWithConfig(cfg *gorm.Config) (*gorm.DB, error) {
	// 如果数据库名称为空，则返回错误
	if m.Dbname == "" {
		return nil, errors.New("database name is empty")
	}
	// 构造 postgres 配置
	pgsqlCfg := postgres.Config{
		DSN:                  m.Dsn(),
		PreferSimpleProtocol: true, // 禁用隐式预处理语句的使用
	}
	// 使用默认的 Gorm 日志配置创建 Logger
	newLogger := NewDefaultZeroLogger(m)
	cfg.Logger = newLogger
	// 使用自定义配置打开数据库连接
	db, err := gorm.Open(postgres.New(pgsqlCfg), cfg)
	if err != nil {
		return nil, err
	}

	// 初始化数据库插件
	err = plugins.InitPlugins(db)
	if err != nil {
		return nil, err
	}
	// 获取底层数据库连接对象，并设置最大空闲连接数和最大打开连接数
	sqldb, _ := db.DB()
	sqldb.SetMaxIdleConns(m.MaxIdleConns)
	sqldb.SetMaxOpenConns(m.MaxOpenConns)
	logx.Infof("数据库连接成功：%s", pgsqlCfg.DSN)
	return db, nil
}
