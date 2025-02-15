package httpclent

import (
	"bytes"
	"context"
	"net/http"
	"time"
)

// GetWithContext 使用指定的上下文 ctx 发起 GET 请求
// 参数：
//
//	ctx：上下文对象，支持超时和取消操作
//	url：请求地址
//	timeout：请求超时时间（单位：秒）
//
// 返回 ResponseWrapper 包含请求结果及可能出现的错误
func GetWithContext(ctx context.Context, url string, timeout time.Duration) ResponseWrapper {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return createRequestError(err)
	}
	return request(req, timeout)
}

// PostParamsWithContext 使用指定的上下文 ctx 发起 POST 请求，提交表单格式数据
// 参数：
//
//	ctx：上下文对象
//	url：请求地址
//	params：表单参数字符串
//	timeout：请求超时时间（单位：秒）
//
// 返回 ResponseWrapper 包含请求结果及可能出现的错误
func PostParamsWithContext(ctx context.Context, url string, params string, timeout time.Duration) ResponseWrapper {
	buf := bytes.NewBufferString(params)
	req, err := http.NewRequestWithContext(ctx, "POST", url, buf)
	if err != nil {
		return createRequestError(err)
	}
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	return request(req, timeout)
}

// PostJsonWithContext 使用指定的上下文 ctx 发起 POST 请求，提交 JSON 格式数据
// 参数：
//
//	ctx：上下文对象
//	url：请求地址
//	body：JSON 格式的请求体
//	timeout：请求超时时间（单位：秒）
//
// 返回 ResponseWrapper 包含请求结果及可能出现的错误
func PostJsonWithContext(ctx context.Context, url string, body string, timeout time.Duration) ResponseWrapper {
	buf := bytes.NewBufferString(body)
	req, err := http.NewRequestWithContext(ctx, "POST", url, buf)
	if err != nil {
		return createRequestError(err)
	}
	req.Header.Set("Content-type", "application/json")
	return request(req, timeout)
}
