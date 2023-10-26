Name: {{.Package}}-rpc
ListenOn: 0.0.0.0:7200
Salt: HWVOFkGgPTryzICwd7qnJaZR9KQ2i8xe
Auth: true
App: {{.Package}}
Token: 6jKNZbEpYGeUMAifz10gOnmoty3TV-{{.Package}}

JwtAuth:
  AccessSecret: {{.Option.Auth.AccessSecret}}
  AccessExpire: {{.Option.Auth.AccessExpire}}
  Issuer: {{.Option.Auth.Issuer}}

Redis:                   # 指定 Redis 服务
  Key: rpc:auth:{{.Package}}     # 指定 Key 应为 hash 类型
  Host: {{.Option.Redis.Addr}}
  Pass: {{.Option.Redis.Password}}
  Type: {{.Option.Redis.DB}}

Database:
  DriverName: {{.Option.DB.DbType}}
  DataSource: {{.Option.DB.DataSource}}

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
  Name: {{.Package}}.api
  Endpoint: http://jaeger:14268/api/traces
  Sampler: 1.0
  Batcher: jaeger