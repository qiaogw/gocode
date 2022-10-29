package {{.TableUrl}}

import (
	"bytes"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"{{.PKG}}/common/errx"
	"{{.ParentPkg}}/api/internal/svc"
	"{{.ParentPkg}}/api/internal/types"
	"{{.ParentPkg}}/rpc/{{.Db}}"
)

type Import{{.Table}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImport{{.Table}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Import{{.Table}}Logic {
	return &Import{{.Table}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Import{{.Table}}Logic) Import{{.Table}}(req *types.ImportRequest) (resp *types.CommonResponse, err error) {
	ioReaderData := req.UpFile.(io.Reader)
	buf := &bytes.Buffer{}
	buf.ReadFrom(ioReaderData)

	// retrieve a byte slice from bytes.Buffer
	data := buf.Bytes()
	_, err = l.svcCtx.{{.Service}}Rpc.Import{{.Table}}(l.ctx, &{{.Db}}.ExportResponse{
		Data: data,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", nil)
	}
	return &types.CommonResponse{
		Code: errx.Success,
		Msg:  "导入成功",
	}, nil
}
