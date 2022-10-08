Name: {{.Package}}-api
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

{{.Service}}Rpc:
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
