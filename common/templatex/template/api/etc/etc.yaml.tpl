Name: {{.Package}}
Host: {{.Option.System.ApiHost}}
Port: {{.Option.System.ApiPort}}
Mode: dev

Database:
  DriverName: {{.Option.DB.DbType}}
  DataSource: {{.Option.DB.DataSource}}

CacheRedis:
  - Host: {{.Option.Redis.Addr}}
    Pass: {{.Option.Redis.Password}}
    Type: {{.Option.Redis.DB}}

Auth:
  AccessSecret: {{.Option.Auth.AccessSecret}}
  AccessExpire: {{.Option.Auth.AccessExpire}}
  Issuer: {{.Option.Auth.Issuer}}

{{.Service}}Rpc:
  App: {{.Package}}                          # App 标识
  Token: 6jKNZbEpYGeUMAifz10gOnmoty3TV-{{.Package}}  # Token 值
  Etcd:
    Hosts:
     {{- range  .Option.Etcd.Hosts }}
      - {{.}}
     {{- end }}
    Key: {{.Package}}.rpc


{{- if .IsFlow }}
AdminRpc:
  App: admin                          # App 标识
  Token: 6jKNZbEpYGeUMAifz10gOnmoty3TV  # Token 值
  Etcd:
    Hosts:
      {{- range  .Option.Etcd.Hosts }}
      - {{.}}
     {{- end }}
    Key: admin.rpc

FsmRpc:
  App: fsm                          # App 标识
  Token: 6jKNZbEpYGeUMAifz10gOnmoty3TV-fsm  # Token 值
  Etcd:
    Hosts:
      {{- range  .Option.Etcd.Hosts }}
      - {{.}}
     {{- end }}
    Key: fsm.rpc

GencodeRpc:
  App: gencode                          # App 标识
  Token: sub-gencode  # Token 值
  Etcd:
    Hosts:
{{- range  .Option.Etcd.Hosts }}
      - {{.}}
{{- end }}
    Key: gencode.rpc

{{- end }}
Prometheus:
  Host: 0.0.0.0
  Port: 9081
  Path: /metrics

Telemetry:
  Name: {{.Package}}.api
  Endpoint: http://jaeger:14268/api/traces
  Sampler: 1.0
  Batcher: jaeger