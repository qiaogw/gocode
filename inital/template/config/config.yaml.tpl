# system configuration
system:
  name: {{.Package}}
  api_host: 0.0.0.0
  api_port: 8080
  rpc_host: 0.0.0.0
  rpc_port: 8000

# redis configuration
redis:
  db: node
  addr: '127.0.0.1:6379'
  password: ''

etcd:
  hosts:
    - 127.0.0.1:2379
  key: {{.Package}}.rpc

Auth:
  AccessSecret: zd-AccessSecret
  AccessExpire: 3600
  Issuer: gocode

# mysql connect configuration
# db-type: 'mysql','postgres'
#   mysql config: charset=utf8mb4&parseTime=True&loc=Local
#   postgres config: "sslmode=disable TimeZone=Asia/Shanghai"
db:
  db-type: 'mysql'
  path: 127.0.0.1
  port: "3306"
  config: charset=utf8mb4&parseTime=True&loc=Local
  db-name: {{.Package}}
  username: root
  password: "123456"
  max-idle-conns: 10
  max-open-conns: 100
  log-mode: error
  log-zap: false
  TablePrefix:   #重要，表前缀如 sys_，仅初始化含有前缀表

# 更新覆盖文件
autocode:
  coverFile:
    - servicecontext.go
