package config

type APP struct {
	Redis    Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
	System   System    `mapstructure:"system" json:"system" yaml:"system"`
	AutoCode Autocode  `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	DB       GeneralDB `mapstructure:"db" json:"db" yaml:"db"`
}

// GeneralDB 也被 Pgsql 和 Mysql 原样使用
type GeneralDB struct {
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                               // 服务器地址:端口
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                               //:端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                         // 高级配置
	Dbname       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                      // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                   // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
}

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
}

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                                  // 环境值
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`                               // 端口值
	OssType       string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`                   // Oss类型
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"` // 多点登录拦截
	UseRedis      bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                // 使用redis
	LimitCountIP  int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP   int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
}

type Autocode struct {
	TransferRestart bool   `mapstructure:"transfer-restart" json:"transfer-restart" yaml:"transfer-restart"`
	Root            string `mapstructure:"root" json:"root" yaml:"root"`
	Server          string `mapstructure:"server" json:"server" yaml:"server"`
	SApi            string `mapstructure:"server-api" json:"server-api" yaml:"server-api"`
	SPlug           string `mapstructure:"server-plug" json:"server-plug" yaml:"server-plug"`
	SInitialize     string `mapstructure:"server-initialize" json:"server-initialize" yaml:"server-initialize"`
	SModel          string `mapstructure:"server-model" json:"server-model" yaml:"server-model"`
	SRequest        string `mapstructure:"server-request" json:"server-request"  yaml:"server-request"`
	SRouter         string `mapstructure:"server-router" json:"server-router" yaml:"server-router"`
	SService        string `mapstructure:"server-service" json:"server-service" yaml:"server-service"`
	Web             string `mapstructure:"web" json:"web" yaml:"web"`
	WApi            string `mapstructure:"web-api" json:"web-api" yaml:"web-api"`
	WForm           string `mapstructure:"web-form" json:"web-form" yaml:"web-form"`
	WTable          string `mapstructure:"web-table" json:"web-table" yaml:"web-table"`
}
