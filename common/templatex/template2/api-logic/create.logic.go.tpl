package {{.TableUrl}}

import (
	"context"

	"github.com/pkg/errors"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/model"
	"github.com/qiaogw/gocode/common/errx"
	"github.com/qiaogw/gocode/common/jwtx"
	"github.com/zeromicro/go-zero/core/logx"
	{{ if .HasTimer }}"github.com/qiaogw/gocode/common/timex"{{ end }}
)

type Create{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreate{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Create{{.Table}}Logic {
	return &Create{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Create{{.Table}}Logic) Create{{.Table}}(req *types.Create{{.Table}}Request) (resp *types.CommonResponse, err error) {
	userId := jwtx.GetUserIdFromCtx(l.ctx)
	{{$table:=.Table}}
	{{$tableComment:=.TableComment}}
	{{- range  .CacheKeys}}
		// 判断该{{.Field}}记录是否已经存在
		_, err := l.svcCtx.{{$table}}Model.FindOneBy{{.Field}}(l.ctx,in.{{.Field}})
		if err == nil {
			return nil, errors.Wrapf(errx.NewErrCode(errx.Duplicate), "该{{$tableComment}}已存在")
		}
	{{- end}}
	new{{.Table}} := model.{{.Table}}{
	{{- range  .Columns }}
		{{- if .IsPk }}
		{{- else if .IsModelTime -}}
		{{- else if .IsControl -}}
		{{- else}}
			{{- if eq .DataType "time.Time"}}
				{{.FieldName}}: timex.DatetimeStrToTime(req.{{.FieldName}}),
			{{- else}}
				{{- if .IsPage}}
				{{- else}}
					{{.FieldName}}: req.{{.FieldName}},
				{{- end}}
			{{- end}}
		{{- end}}
	{{- end }}
	}
	new{{.Table}}.CreateBy = userId
	_, err = l.svcCtx.{{.Table}}Model.Insert(l.ctx, &new{{.Table}})
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DbError),
		"创建 {{$tableComment}} 数据库失败 , 错误:%v ,数据 : %+v  ", err, new{{.Table}})
	}
	
	return &types.CommonResponse{
		Code : errx.Success,
		Msg: "添加成功",
	}, nil
}
