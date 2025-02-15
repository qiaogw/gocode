package httpclent

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

// ResponseWrapper 封装了 HTTP 请求的返回信息，包括状态码、响应体、响应头和错误信息
type ResponseWrapper struct {
	StatusCode int         // HTTP 状态码
	Body       string      // 响应体内容
	Header     http.Header // 响应头
	Error      error       // 请求过程中产生的错误
}

// Get 发起 GET 请求
// 参数 url 为请求地址，timeout 为超时时间（单位：秒）
// 返回 ResponseWrapper 包含请求结果及可能的错误信息
func Get(url string, timeout time.Duration) ResponseWrapper {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return createRequestError(err)
	}

	return request(req, timeout)
}

// PostParams 发起 POST 请求，提交表单格式的数据
// 参数 url 为请求地址，params 为表单参数字符串，timeout 为超时时间（单位：秒）
// 返回 ResponseWrapper 包含请求结果及可能的错误信息
func PostParams(url string, params string, timeout time.Duration) ResponseWrapper {
	buf := bytes.NewBufferString(params)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return createRequestError(err)
	}
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	return request(req, timeout)
}

// PostJson 发起 POST 请求，提交 JSON 格式的数据
// 参数 url 为请求地址，body 为 JSON 字符串，timeout 为超时时间（单位：秒）
// 返回 ResponseWrapper 包含请求结果及可能的错误信息
func PostJson(url string, body string, timeout time.Duration) ResponseWrapper {
	buf := bytes.NewBufferString(body)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return createRequestError(err)
	}
	req.Header.Set("Content-type", "application/json")

	return request(req, timeout)
}

// request 执行 HTTP 请求，并返回封装后的结果
// 参数 req 为构造好的 HTTP 请求，timeout 为超时时间（单位：秒）
// 返回 ResponseWrapper 包含请求的状态码、响应体、响应头及错误信息
func request(req *http.Request, timeout time.Duration) ResponseWrapper {
	wrapper := ResponseWrapper{StatusCode: 0, Body: "", Header: make(http.Header)}
	client := &http.Client{}
	if timeout > 0 {
		// 设置客户端超时时间（将 timeout 转换为秒）
		client.Timeout = timeout * time.Second
	}
	// 设置请求头信息
	setRequestHeader(req)
	resp, err := client.Do(req)
	if err != nil {
		wrapper.Body = fmt.Sprintf("执行HTTP请求错误-%s", err.Error())
		wrapper.Error = err
		return wrapper
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		wrapper.Body = fmt.Sprintf("读取HTTP请求返回值失败-%s", err.Error())
		wrapper.Error = err
		return wrapper
	}
	wrapper.StatusCode = resp.StatusCode
	wrapper.Body = string(body)
	wrapper.Header = resp.Header

	return wrapper
}

// setRequestHeader 为请求设置默认的请求头信息
func setRequestHeader(req *http.Request) {
	req.Header.Set("User-Agent", "golang/gocron")
}

// createRequestError 构造一个请求错误的 ResponseWrapper 对象
// 参数 err 为产生的错误
// 返回一个包含错误信息的 ResponseWrapper
func createRequestError(err error) ResponseWrapper {
	errorMessage := fmt.Sprintf("HTTP请求错误-%s", err.Error())
	return ResponseWrapper{0, errorMessage, make(http.Header), err}
}
