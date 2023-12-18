package {{.TableUrl}}

import (
	"context"

	"github.com/pkg/errors"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"

	"github.com/qiaogw/gocode/common/errx"
	"github.com/qiaogw/gocode/common/jwtx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/qiaogw/gocode/common/modelx"
	{{ if .HasTimer }}"github.com/qiaogw/gocode/common/timex"{{ end }}
)

type Update{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdate{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Update{{.Table}}Logic {
	return &Update{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Update{{.Table}}Logic) Update{{.Table}}(req *types.Update{{.Table}}Request) (resp *types.CommonResponse, err error) {
	userId := jwtx.GetUserIdFromCtx(l.ctx)
	// 查询{{.TableComment}}是否存在
	res, err := l.svcCtx.{{.Table}}Model.FindOne(l.ctx, req.Id)
	if err != nil {
		if err == modelx.ErrNotFound {
			return nil, errors.Wrapf(errx.NewErrCode(errx.NoData), "该{{.TableComment}}不存在，id: %v", req.Id)
		}
		return nil, errors.Wrapf(errx.NewErrCode(errx.NoData), "数据库查询 {{.TableComment}} 失败，id: %v,err:%v", req.Id,err)
	}
	{{- range  .Columns }}
		{{-  if .IsModelTime -}}
		{{- else if .IsControl -}}
        {{- else if .IsPage}}
		{{- else if .IsPk }}
		{{- else if eq .DataType "time.Time"}}
			res.{{.FieldName}}=timex.DatetimeStrToTime(req.{{.FieldName}})
		{{- else}}
			res.{{.FieldName}}=req.{{.FieldName}}
		{{- end}}
	{{- end }}
	res.UpdateBy =	userId

	err = l.svcCtx.{{.Table}}Model.Update(l.ctx, res)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DbError),
		"数据库更新 {{.TableComment}} 失败 , err:%v ,data : %+v  ", err, res)
	}

	return &types.CommonResponse{
		Code: errx.Success,
		Msg: "更新成功",
	}, nil
}
