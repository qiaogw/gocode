package authx

import (
	"bytes"
	"context"
	"github.com/bitly/go-simplejson"
	"github.com/qiaogw/gocode/util"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sub-admin/admin/rpc/admin"
	"sub-admin/admin/rpc/adminclient"
)

func SwagRoute(name string, route []rest.Route) error {
	var req admin.CreateApiAllRequest
	for _, v := range route {
		var data admin.CreateApiRequest
		str := strings.Split(v.Path, "/")
		data.Method = v.Method
		data.Path = v.Path
		data.Module = str[1]
		req.Datas = append(req.Datas, &data)
	}

	err := getRoute(name, &req)
	if err != nil {
		logx.Errorf("【API-ERR】 更新API auth fail，route:%+v,err: %+v ", route, err)
	}
	rc := context.Background()

	authRpcClient := adminclient.NewAdmin(zrpc.MustNewClient(authConfig.AdminRpc))
	if err == nil {
		_, err = authRpcClient.UpdateAllApi(rc, &req)
	}
	return err
}

func getRoute(name string, in *admin.CreateApiAllRequest) error {
	pwd, _ := os.Getwd()

	if err := GenSwag(name); err != nil {
		return err
	}
	swaggerFile := filepath.Join(pwd, name, "api", "doc", "swagger.json")
	jsonFile, err := os.ReadFile(swaggerFile)
	if err != nil {
		return err
	}
	jsonData, _ := simplejson.NewFromReader(bytes.NewReader(jsonFile))
	for _, v := range in.Datas {
		if v.Method != "HEAD" &&
			!strings.Contains(v.Path, "/swagger/") &&
			!strings.Contains(v.Path, "/static/") &&
			!strings.Contains(v.Path, "/form-generator/") {
			// 根据接口方法注释里的@Summary填充接口名称，适用于代码生成器
			// 可在此处增加配置路径前缀的if判断，只对代码生成的自建应用进行定向的接口名称填充
			v.Title, _ = jsonData.Get("paths").Get(v.Path).Get(strings.ToLower(v.Method)).Get("summary").String()
		}
	}

	return err
}

func GenSwag(name string) (err error) {
	pwd, _ := os.Getwd()
	apiFile := filepath.Join(pwd, name, "api", name+".api")
	fPath := filepath.Join(pwd, name, "api", "doc")
	ok, _ := util.PathExists(fPath)
	if !ok {
		if err = os.MkdirAll(fPath, os.ModePerm); err != nil {
			return err
		}
	}
	cmd := exec.Command("goctl", "api",
		"plugin", "-p", "goctl-swagger=`swagger -filename swagger.json`", "--api", apiFile, "--dir", fPath)

	var out, errBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	//err = cmd.Start() //如果用start则直接向后运行
	err = cmd.Run()
	//outStr, errStr := string(out.Bytes()), string(errBuf.Bytes())
	//if err != nil {
	//
	//	return err
	//}
	//go func() {
	//	err = cmd.Wait()
	//}()
	return err
}
