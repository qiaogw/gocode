Name: {{.Package}}-api
Host: {{.Option.System.ApiHost}}
Port: {{.Option.System.ApiPort}}
Timeout: 6000
Mode: dev


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

Auth:
  AccessSecret: {{.Option.Auth.AccessSecret}}
  AccessExpire: {{.Option.Auth.AccessExpire}}
  Issuer: {{.Option.Auth.Issuer}}

{{.Service}}Rpc:
  App: {{.Package}}                          # App 标识
  Token: sub-{{.Package}} # Token 值
  Etcd:
    Hosts:
     {{- range  .Option.Etcd.Hosts }}
      - {{.}}
     {{- end }}
    Key: {{.Package}}.rpc

Captcha:
  ImgHeight: 4
  ImgWidth: 160
  KeyLong: 80

Prometheus:
  Host: 0.0.0.0
  Port: 9081
  Path: /metrics

Telemetry:
  Name: {{.Package}}.api
  Endpoint: http://jaeger:14268/api/traces
  Sampler: 1.0
  Batcher: jaeger