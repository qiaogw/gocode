package logic

import (
	"context"
"github.com/pkg/errors"
	"github.com/qiaogw/gocode/common/errx"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/qiaogw/gocode/common/toolx"
	"{{.ParentPkg}}/model"

	"{{.ParentPkg}}/rpc/{{.Db}}"
	"{{.ParentPkg}}/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type Export{{.Table}}Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExport{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Export{{.Table}}Logic {
	return &Export{{.Table}}Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Export{{.Table}}Logic) Export{{.Table}}(in *{{.Db}}.ExportRequest) (*{{.Db}}.ExportResponse, error) {
	var qData model.List{{.Table}}Req
	qData.PageIndex = in.PageIndex
	qData.PageSize = in.PageSize
	qData.SearchKey = in.SearchKey
	qData.SortBY = in.SortBy
	qData.Descending = in.Descending
	list, _, err := l.svcCtx.{{.Table}}Model.FindAll(l.ctx, &qData)
	if err != nil {
		if err == modelx.ErrNotFound {
			return nil, errx.NewErrCodeMsg(errx.NoData, "{{.TableComment}}-该查询无数据")
		}
		return nil, errors.Wrapf(errx.NewErrCode(errx.DbError),
					"该{{.TableComment}}查询失败，err:%v", err)
	}
	resp, err := toolx.ExportToWeb(list[0], list, list[0].TableName())
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.ServerCommonError),
			"{{.TableComment}}导出excel 失败，data: %+v,err:%v", list,err)
	}
	return &{{.Db}}.ExportResponse{
		Data: resp.Bytes(),
	}, nil
}
