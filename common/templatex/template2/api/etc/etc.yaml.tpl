Name: {{.Package}}-api
Host: {{.Option.System.ApiHost}}
Port: {{.Option.System.ApiPort}}
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
  ResetPassword: "sub-123"

Salt: sub-admin-88732

Captcha:
  ImgHeight: 4
  ImgWidth: 160
  KeyLong: 80

NoAuthUrls:
  - "/auth/captcha"
  - "/auth/login"
  - "/admin/user/setMeRole"
