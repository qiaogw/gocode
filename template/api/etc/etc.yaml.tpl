Name: {{.Service}}-api
Host: {{.Option.System.ApiHost}}
Port: {{.Option.System.ApiPort}}


CacheRedis:
  - Host: {{.Option.Redis.Addr}}
    Pass: {{.Option.Redis.Password}}
    Type: {{.Option.Redis.DB}}

Auth:
  AccessSecret: {{.Option.Auth.AccessSecret}}
  AccessExpire: {{.Option.Auth.AccessExpire}}
  RefreshAfter: {{.Option.Auth.RefreshAfter}}

AdminRpc:
  Etcd:
    Hosts:
     {{- range  .Option.Etcd.Hosts }}
      - {{.}}
     {{- end }}
    Key: {{.Option.Etcd.Key}}

Prometheus:
  Host: 0.0.0.0
  Port: 9080
  Path: /metrics

Telemetry:
  Name: admin.api
  Endpoint: http://jaeger:14268/api/traces
  Sampler: 1.0
  Batcher: jaeger