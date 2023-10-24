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
  {{- if eq .Option.DB.DbType "" }}
  db-type: mysql
  {{- else }}
  db-type: {{.Option.DB.DbType}}
  {{- end }}
{{- if eq .Option.DB.Path "" }}
  path: 127.0.0.1
{{- else }}
  path: {{.Option.DB.Path}}
{{- end }}
{{- if eq .Option.DB.Port "" }}
  port: 3306
{{- else }}
  port: {{.Option.DB.Port}}
{{- end }}
{{- if eq .Option.DB.Config "" }}
  config: charset=utf8mb4&parseTime=True&loc=Local
{{- else }}
  config: {{.Option.DB.Config}}
{{- end }}
{{- if eq .Option.DB.Dbname "" }}
  db-name: {{.Package}}
{{- else }}
  db-name: {{.Option.DB.Dbname}}
{{- end }}
{{- if eq .Option.DB.Username "" }}
  username: root
{{- else }}
  username: {{.Option.DB.Username}}
{{- end }}
{{- if eq .Option.DB.Password "" }}
  password: "123456"
{{- else }}
  password: {{.Option.DB.Password}}
{{- end }}
  max-idle-conns: 10
  max-open-conns: 100
  log-mode: error
  log-zap: false
{{- if eq .Option.DB.TablePrefix "" }}
  TablePrefix: "123456"
{{- else }}
  TablePrefix: {{.Option.DB.TablePrefix}}
{{- end }}
  #重要，表前缀如 sys_，仅初始化含有前缀表

# 更新覆盖文件
autocode:
  coverFile:
    - servicecontext.go
