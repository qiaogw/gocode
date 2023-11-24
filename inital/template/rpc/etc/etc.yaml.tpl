Name: {{.Package}}-rpc
ListenOn: {{.Option.System.RpcHost}}:{{.Option.System.RpcPort}}

Salt: HWVOFkGgPTryzICwd7qnJaZR9KQ2i8xe

Auth: false    //rpc认证
App: {{.Package}}                          # App 标识
Token: sub-{{.Package}} # Token 值

JwtAuth:
  AccessSecret: {{.Option.Auth.AccessSecret}}
  AccessExpire: {{.Option.Auth.AccessExpire}}
  Issuer: {{.Option.Auth.Issuer}}

Redis:                   # 指定 Redis 服务
  Key: rpc:auth:{{.Package}}     # 指定 Key 应为 hash 类型
  Host: {{.Option.Redis.Addr}}
  Pass: {{.Option.Redis.Password}}
  Type: {{.Option.Redis.DB}}

#Database:
#  DriverName: {{.Option.DB.DbType}}
#  DataSource: {{.Option.DB.DataSource}}

DbConf:
  Driver: {{.Option.DB.DbType}}
  Host: {{.Option.DB.Path}}
  Port: {{.Option.DB.Port}}
  User: {{.Option.DB.Username}}
  Password: {{.Option.DB.Password}}
  Db: {{.Option.DB.Dbname}}
  Schema:
  Config: {{.Option.DB.Config}}

CacheRedis:
  - Host: {{.Option.Redis.Addr}}
    Pass: {{.Option.Redis.Password}}
    Type: {{.Option.Redis.DB}}


Etcd:
  Hosts:
     {{- range  .Option.Etcd.Hosts }}
    - {{.}}
     {{- end }}
  Key: {{.Package}}.rpc



Prometheus:
  Host: 0.0.0.0
  Port: 9080
  Path: /metrics

Telemetry:
  Name: {{.Package}}.rpc
  Endpoint: http://jaeger:14268/api/traces
  Sampler: 1.0
  Batcher: jaeger