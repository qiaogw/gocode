package {{.TableUrl}}

import (
	"context"


	"github.com/pkg/errors"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"github.com/qiaogw/gocode/common/modelx"
	"github.com/qiaogw/gocode/common/errx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/qiaogw/gocode/common/toolx"
	"{{.ParentPkg}}/model"
)

type Export{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExport{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Export{{.Table}}Logic {
	return &Export{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Export{{.Table}}Logic) Export{{.Table}}(req *types.List{{.Table}}Request) (resp []byte, err error) {
	var qData model.List{{.Table}}Req
	qData.PageIndex = req.PageIndex
	qData.PageSize = req.PageSize
	qData.SearchKey = req.SearchKey
	qData.SortBY = req.SortBY
	qData.Descending = req.Descending
	list, _, err := l.svcCtx.{{.Table}}Model.FindAll(l.ctx, &qData)
	if err != nil {
		if err == modelx.ErrNotFound {
			return nil, errx.NewErrCodeMsg(errx.NoData, "{{.TableComment}}-该查询无数据")
		}
		return nil, errors.Wrapf(errx.NewErrCode(errx.DbError),
					"该{{.TableComment}}查询失败，err:%v", err)
	}
	res, err := toolx.ExportToWeb(list[0], list, list[0].TableName())
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.ServerCommonError),
			"{{.TableComment}}导出excel 失败，data: %+v,err:%v", list,err)
	}


	return res.Bytes(), nil
}
