Name: {{.Package}}-Rpc
ListenOn: {{.Option.System.RpcHost}}:{{.Option.System.RpcPort}}


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

Salt: HWVOFkGgPTryzICwd7qnJaZR9KQ2i8xe

Prometheus:
  Host: 0.0.0.0
  Port: 9080
  Path: /metrics

Telemetry:
  Name: {{.Package}}.api
  Endpoint: http://jaeger:14268/api/traces
  Sampler: 1.0
  Batcher: jaeger