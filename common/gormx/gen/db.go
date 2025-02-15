package gen

import (
	"github.com/google/uuid"
	"github.com/qiaogw/gocode/common/gormx/modelx"
)

const (
	modelPath       = "model"
	autoPath        = "autocode_template/"
	apiPath         = "api"
	apiDescPath     = "api-desc"
	apiLogicPath    = "api-logic"
	handlerPath     = "handler"
	rpcPath         = "rpc"
	TempDownPath    = "TempDown"
	rpcLogicPath    = "rpc-logic"
	commonPath      = "common"
	internalPath    = "internal"
	svcPath         = "svc"
	configPath      = "config"
	etcPath         = "etc"
	logicPath       = "logic"
	apiResponsePath = "replay"
	webPath         = "web"

	indexPri = "PRIMARY"
)

type CacheKey struct {
	Key       string
	Value     string
	Field     string
	FieldJson string
	DataType  string
}

type Database struct {
	Id       uuid.UUID `json:"id" comment:"主键" gorm:"primaryKey;column:id;size:256;comment:主键;"`
	Sort     int64     `json:"sort" comment:"排序" gorm:"column:sort;size:2;comment:排序;"`
	Name     string    `json:"name" comment:"业务名称" gorm:"column:name;size:256;comment:业务名称;"`
	Driver   string    `json:"driver" comment:"dict_value" gorm:"column:driver;size:256;comment:dict_value;"`
	Host     string    `json:"host" comment:"is_default" gorm:"column:host;size:32;comment:is_default;"`
	Port     int       `json:"port" comment:"status" gorm:"column:port;size:32;comment:status;"`
	Username string    `json:"username" comment:"用户" gorm:"column:username;size:256;comment:用户名称;"`
	Password string    `json:"password" comment:"密码" gorm:"column:password;size:256;comment:密码;"`
	Dbname   string    `json:"dbname" comment:"数据库名称" gorm:"column:dbname;size:256;comment:数据库名称;"`
	Schema   string    `json:"schema" comment:"Schema" gorm:"column:schema;size:256;comment:Schema;"`
	Config   string    `json:"config" comment:"配置" gorm:"column:config;size:256;comment:配置;"`
	Remark   string    `json:"remark" comment:"remark" gorm:"column:remark;size:256;comment:remark;"`
	CreateBy string    `json:"createBy" comment:"创建者" gorm:"column:create_by;size:256;comment:创建者;"`
	UpdateBy string    `json:"updateBy" comment:"更新者" gorm:"column:update_by;size:256;comment:更新者;"`
	modelx.ModelTime

	AppName       string `json:"appName" comment:"项目名称" gorm:"column:app_name;size:256;comment:项目名称;"`
	TablePrefix   string `json:"tablePrefix" comment:"表前缀" gorm:"column:table_prefix;size:256;comment:表前缀;"`
	Package       string `json:"package"  gorm:"column:package;size:256;comment:包名首字母小写驼峰;"`  //首字母小写驼峰
	Service       string `json:"service"  gorm:"column:service;size:256;comment:服务名首字母大写驼峰;"` //首字母大写驼峰
	FileName      string `json:"fileName"  gorm:"column:file_name;size:256;comment:文件名全小写;"`  //全小写，web生成
	ParentPackage string `json:"parentPackage"  gorm:"column:parent_package;size:256;comment:项目路径全小写;"`
	Label         string `json:"label" comment:"业务标题" gorm:"column:label;size:-5;comment:业务标题;"`
	Mode          string `json:"mode" comment:"项目模式" gorm:"column:mode;size:256;comment:项目模式;"`
	ApiHost       string `json:"apiHost" comment:"api地址" gorm:"column:api_host;type:varchar;size:255;comment:api地址;"`
	ApiPort       int64  `json:"apiPort" comment:"api端口" gorm:"column:api_port;type:integer;size:4;comment:api端口;"`
	RpcHost       string `json:"rpcHost" comment:"rpc地址" gorm:"column:rpc_host;type:varchar;size:255;comment:rpc地址;"`
	RpcPort       int64  `json:"rpcPort" comment:"rpc端口" gorm:"column:rpc_port;type:integer;size:4;comment:rpc端口;"`
	Author        string `json:"author" form:"author" db:"author" gorm:"column:author;size:255;comment:作者;"`
	Email         string `json:"email" form:"email" db:"email" gorm:"column:email;size:255;comment:作者邮箱;"`
	Dir           string `json:"dir" form:"dir" db:"dir" gorm:"column:dir;size:256;comment:项目目录;"`

	RedisHost    string `json:"redisHost" gorm:"column:redis_host;size:255;comment:Redis地址;"`
	RedisPass    string `json:"redisPass" gorm:"column:redis_pass;size:255;comment:Redis密码;"`
	RedisKey     string `json:"redisKey" gorm:"column:redis_key;size:255;comment:RedisKey;"`
	RedisType    string `json:"redisType" gorm:"column:redis_type;size:255;comment:Redis类型;"`
	EtcdHost     string `json:"etcdHost" gorm:"column:etcd_host;size:255;comment:etcd地址;"`
	AccessSecret string `json:"accessSecret" gorm:"column:access_secret;size:255;comment:认证令牌;"`

	IsFlow      bool `json:"isFlow" comment:"是否业务流程" gorm:"column:is_flow;type:bool;comment:是否业务流程;"`
	IsImport    bool `json:"isImport"  gorm:"column:is_import;size:1;comment:是否开启导入;"`
	IsAuth      bool `json:"isAuth" form:"isAuth" db:"is_auth" gorm:"column:is_auth;size:1;comment:开启用户认证;"`
	IsDataScope bool `json:"isDataScope" form:"isDataScope" db:"is_data_scope" gorm:"column:is_data_scope;size:1;comment:开启数据权限;"`
	HasTimer    bool `json:"hasTimer" form:"hasTimer" db:"has_timer" gorm:"column:has_timer;type:bool;comment:存在时间;"`
	HasCacheKey bool `json:"hasCacheKey" form:"hasCacheKey" db:"has_cache_key" gorm:"column:has_cache_key;type:bool;comment:存在非主键的唯一键;"`
	NeedValid   bool `json:"needValid" form:"needValid" db:"need_valid" gorm:"column:need_valid;type:bool;comment:需要验证;"`
	Postgres    bool `json:"postgres" form:"postgres" db:"postgres" gorm:"column:postgres;type:bool;comment:postgres;"`
	PkIsChar    bool `json:"pkIsChar" form:"postgres" db:"postgres" gorm:"column:pk_is_char;type:bool;comment:主键为uuid;"`
	HasCode     bool `json:"hasCode" gorm:"column:has_code;type:bool;comment:已经生成;"`
	HasCache    bool `json:"hasCache" gorm:"column:has_cache;type:bool;comment:需要缓存;"`

	Tables []*Table `json:"tables" gorm:"foreignKey:DatabaseId"`
}
