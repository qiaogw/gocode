package configx

import (
	"fmt"
	"github.com/qiaogw/gocode/common/gormx/logger"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type DBType string

const (
	MySQL    DBType = "mysql"
	Postgres DBType = "postgres"
	MongoDB  DBType = "mongodb"
)

// AllDBTypes 使用自定义类型定义常见数据库名称的集合
var AllDBTypes = []DBType{MySQL, Postgres}

type DbConf struct {
	Driver        string `json:",default=mysql"`
	Host          string `json:",default=mysql"` // 服务器地址
	Port          int    `json:",default=3306"`  // 数据库端口，默认 3306
	Dbname        string // 数据库名称
	Username      string // 数据库用户名
	Password      string // 数据库密码
	TimeZone      string `json:",default=Asia/Shanghai"`                                      // 数据库时区，默认 Asia/Shanghai
	SslMode       string `json:",default=disable,options=disable|enable"`                     // SSL 模式，支持 disable 或 enable，默认 disable
	Config        string `json:",default=charset%3Dutf8mb4%26parseTime%3Dtrue%26loc%3DLocal"` // 高级配置参数
	MaxIdleConns  int    `json:",default=10"`                                                 // 最大空闲连接数
	MaxOpenConns  int    `json:",default=10"`                                                 // 最大打开连接数
	LogMode       string `json:",default=dev,options=dev|test|prod|silent"`                   // 日志模式：dev、test、prod 或 silent
	LogColorful   bool   `json:",default=false"`                                              // 是否启用日志彩色输出
	SlowThreshold int64  `json:",default=1000"`                                               // 慢查询阈值（单位：毫秒）
	Schema        string `json:",default=public"`
	TablePrefix   string `json:"TablePrefix"` // 表前缀 'it_'
}

// GormLogConfigI 定义了获取 Gorm 日志配置参数的接口
type GormLogConfigI interface {
	// GetGormLogMode 返回 Gorm 的日志级别
	GetGormLogMode() gormLogger.LogLevel
	// GetSlowThreshold 返回慢 SQL 的阈值
	GetSlowThreshold() time.Duration
	// GetColorful 返回是否启用彩色日志打印
	GetColorful() bool
	Connect() (*gorm.DB, error)
	ConnectWithConfig(cfg *gorm.Config) (*gorm.DB, error)
}

func GetConnect(conf DbConf) (*gorm.DB, error) {
	switch conf.Driver {
	case string(MySQL):
		// 根据 DbConf 手动构造 Mysql 结构体（字段赋值可以直接赋值）
		m := Mysql{
			Driver:        conf.Driver,
			Host:          conf.Host,
			Port:          conf.Port,
			Dbname:        conf.Dbname,
			Username:      conf.Username,
			Password:      conf.Password,
			Config:        conf.Config,
			MaxIdleConns:  conf.MaxIdleConns,
			MaxOpenConns:  conf.MaxOpenConns,
			LogMode:       conf.LogMode,
			LogColorful:   conf.LogColorful,
			SlowThreshold: conf.SlowThreshold,
		}
		return m.Connect()
	case string(Postgres):
		p := PgSql{
			Driver:        conf.Driver,
			Host:          conf.Host,
			Port:          conf.Port,
			Dbname:        conf.Dbname,
			Username:      conf.Username,
			Password:      conf.Password,
			TimeZone:      conf.TimeZone,
			SslMode:       conf.SslMode,
			MaxIdleConns:  conf.MaxIdleConns,
			MaxOpenConns:  conf.MaxOpenConns,
			LogMode:       conf.LogMode,
			LogColorful:   conf.LogColorful,
			SlowThreshold: conf.SlowThreshold,
		}
		return p.Connect()
	default:
		return nil, fmt.Errorf("只支持 %v,不支持的数据库驱动：%s", AllDBTypes, conf.Driver)
	}
}

func GetConnectWithConfig(conf DbConf, cfg *gorm.Config) (*gorm.DB, error) {
	switch conf.Driver {
	case string(MySQL):
		// 根据 DbConf 手动构造 Mysql 结构体（字段赋值可以直接赋值）
		m := Mysql{
			Driver:        conf.Driver,
			Host:          conf.Host,
			Port:          conf.Port,
			Dbname:        conf.Dbname,
			Username:      conf.Username,
			Password:      conf.Password,
			Config:        conf.Config,
			MaxIdleConns:  conf.MaxIdleConns,
			MaxOpenConns:  conf.MaxOpenConns,
			LogMode:       conf.LogMode,
			LogColorful:   conf.LogColorful,
			SlowThreshold: conf.SlowThreshold,
		}
		return m.ConnectWithConfig(cfg)
	case string(Postgres):
		p := PgSql{
			Driver:        conf.Driver,
			Host:          conf.Host,
			Port:          conf.Port,
			Dbname:        conf.Dbname,
			Username:      conf.Username,
			Password:      conf.Password,
			TimeZone:      conf.TimeZone,
			SslMode:       conf.SslMode,
			MaxIdleConns:  conf.MaxIdleConns,
			MaxOpenConns:  conf.MaxOpenConns,
			LogMode:       conf.LogMode,
			LogColorful:   conf.LogColorful,
			SlowThreshold: conf.SlowThreshold,
			Schema:        conf.Schema,
		}
		return p.ConnectWithConfig(cfg)
	default:
		return nil, fmt.Errorf("只支持 %v,不支持的数据库驱动：%s", AllDBTypes, conf.Driver)
	}
}

// NewDefaultZeroLogger 根据配置创建一个默认的 ZeroLog 日志实例
func NewDefaultZeroLogger(cfg GormLogConfigI) gormLogger.Interface {
	newLogger := logger.NewZeroLog(
		gormLogger.Config{
			SlowThreshold:             cfg.GetSlowThreshold(), // 慢 SQL 阈值
			LogLevel:                  cfg.GetGormLogMode(),   // 日志级别
			IgnoreRecordNotFoundError: true,                   // 忽略记录未找到错误
			Colorful:                  cfg.GetColorful(),      // 是否启用彩色日志打印
		},
	)
	return newLogger
}

// NewDefaultGormLogger 根据配置创建一个默认的 Gorm 日志实例
func NewDefaultGormLogger(cfg GormLogConfigI) gormLogger.Interface {
	newLogger := gormLogger.New(
		log.New(os.Stderr, "\r\n", log.LstdFlags), // 日志输出目标，包括前缀和标准日志标志
		gormLogger.Config{
			SlowThreshold:             cfg.GetSlowThreshold(), // 慢 SQL 阈值
			LogLevel:                  cfg.GetGormLogMode(),   // 日志级别
			IgnoreRecordNotFoundError: true,                   // 忽略记录未找到错误
			Colorful:                  cfg.GetColorful(),      // 是否启用彩色日志打印
		},
	)
	return newLogger
}

// OverwriteGormLogMode 根据传入的模式字符串覆盖默认的 Gorm 日志级别
// 参数 mode 的取值及对应日志级别如下：
// "dev"    -> 返回 gormLogger.Info 级别
// "test"   -> 返回 gormLogger.Warn 级别
// "prod"   -> 返回 gormLogger.Error 级别
// "silent" -> 返回 gormLogger.Silent 级别
// 默认返回 gormLogger.Info 级别
func OverwriteGormLogMode(mode string) gormLogger.LogLevel {
	switch mode {
	case "dev":
		return gormLogger.Info
	case "test":
		return gormLogger.Warn
	case "prod":
		return gormLogger.Error
	case "silent":
		return gormLogger.Silent
	default:
		return gormLogger.Info
	}
}
