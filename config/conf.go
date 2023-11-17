package config

type APP struct {
	System   System    `mapstructure:"system" json:"system" yaml:"system"`
	DB       GeneralDB `mapstructure:"db" json:"db" yaml:"db"`
	Redis    Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
	AutoCode Autocode  `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	Etcd     Etcd      `mapstructure:"etcd" json:"etcd" yaml:"etcd"`
	Auth     Auth      `mapstructure:"Auth" json:"Auth" yaml:"Auth"`
}

type System struct {
	Name          string `mapstructure:"name" json:"name" yaml:"name"`             // 环境值
	ApiHost       string `mapstructure:"api_host" json:"api_host" yaml:"api_host"` //
	ApiPort       int    `mapstructure:"api_port" json:"api_port" yaml:"api_port"`
	RpcHost       string `mapstructure:"rpc_host" json:"rpc_host" yaml:"rpc_host"` //
	RpcPort       int    `mapstructure:"rpc_port" json:"rpc_port" yaml:"rpc_port"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"` // 多点登录拦截
	UseRedis      bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                // 使用redis
	LimitCountIP  int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP   int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	TemplatePath  string `mapstructure:"template_path" json:"template_path" yaml:"template_path"` //模板路径
}

type Autocode struct {
	Root      string
	Pkg       string
	CoverFile []string
	WithCache bool
}

// GeneralDB  Pgsql 和 Mysql 原样使用
type GeneralDB struct {
	DataSource    string
	DbType        string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	Path          string `mapstructure:"path" json:"path" yaml:"path"`                               // 服务器地址:端口
	Port          string `mapstructure:"port" json:"port" yaml:"port"`                               //:端口
	Config        string `mapstructure:"config" json:"config" yaml:"config"`                         // 高级配置
	Dbname        string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                      // 数据库名
	Username      string `mapstructure:"username" json:"username" yaml:"username"`                   // 数据库用户名
	Password      string `mapstructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	MaxIdleConns  int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns  int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode       string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	LogZap        bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
	TablePrefix   string `mapstructure:"TablePrefix" json:"TablePrefix" yaml:"TablePrefix"`          // 表前缀 'it_'
	SingularTable bool   `mapstructure:"SingularTable" json:"SingularTable" yaml:"SingularTable"`    // 使用单数表名，启用该选项，此时，`Article` 的表名应该是 `it_article`
}

type Redis struct {
	DB       string `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
}

type Etcd struct {
	Hosts []string `mapstructure:"hosts" json:"hosts" yaml:"hosts"` // Etcd 集群
	Key   string   `mapstructure:"key" json:"key" yaml:"key"`       //rpc注册key
}

type Auth struct {
	AccessSecret string `mapstructure:"AccessSecret" json:"AccessSecret" yaml:"AccessSecret"` // Etcd 集群
	AccessExpire int64  `mapstructure:"AccessExpire" json:"AccessExpire" yaml:"AccessExpire"` //rpc注册key
	Issuer       string `mapstructure:"Issuer" json:"Issuer" yaml:"Issuer"`                   //rpc注册key
}
