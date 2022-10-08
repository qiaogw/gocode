package response

import (
	"github.com/pkg/errors"
	"github.com/qiaogw/gocode/common/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
	"reflect"
	"strings"
)

type Body struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

//Response 统一封装成功响应值
func Response(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
	errcode := errorx.ServerCommonError
	errmsg := "服务器开小差啦，稍后再来试一试"

	causeErr := errors.Cause(err) // err类型
	if e, ok := causeErr.(*errorx.CodeError); ok {
	// 自定义错误类型
	// 自定义CodeError
	errcode = e.GetErrCode()
	errmsg = e.GetErrMsg()
	} else {
	if gstatus, ok := status.FromError(causeErr); ok {
	// grpc err错误
	grpcCode := uint32(gstatus.Code())
	if errorx.IsCodeErr(grpcCode) {
	// 区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
	errcode = grpcCode
	errmsg = gstatus.Message()
	// 主键重复
					if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
						errcode = errorx.Duplicate
						errmsg = errorx.MapErrMsg(errcode)
					}
	}
	}
	}
	logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
	body.Code = errcode
	body.Msg = errmsg
	} else {
	body.Msg = "请求成功!"
	body.Data = resp
	if isNil(resp) {
			body.Msg = "数据为空!"
		}
	rt := reflect.TypeOf(resp)
	if rt.String() == "*types.CommonResponse" && !isNil(resp) {
	rv := reflect.ValueOf(resp)
	rt = rt.Elem()
	rv = rv.Elem()
	num := rt.NumField()
	for i := 0; i < num; i++ {
	field := rt.Field(i)
	fieldName := field.Name
	if field.Name == "Data" {
	data := rv.FieldByName(fieldName)
	body.Data = data.Interface()
	break
	}
	}
	}
	}
	httpx.OkJson(w, body)
}

func isNil(i interface{}) bool {
	defer func() {
		recover()
	}()
	vi := reflect.ValueOf(i)
	return vi.IsNil()
}
