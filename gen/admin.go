package gen

import (
	"fmt"
	"github.com/qiaogw/gocode/common/pathx"
	"github.com/qiaogw/gocode/common/templatex"
	"github.com/qiaogw/gocode/global"
	"path/filepath"
)

// CreateAdminFile 创建 config
func (acd *AutoCodeService) CreateAdminFile() (err error) {

	templateFS := templatex.GetTpl(acd.Mode)
	// 获取 logic 目录
	src := templatex.GetTplPath(acd.Mode) + "/admin/desc"
	c := global.GenConfig.AutoCode
	rPath := c.Root
	// 获取目标目录
	dstDir := filepath.Join(rPath, apiPath, "admin-desc")
	err = pathx.CopyTpl(templateFS, src, dstDir)
	if err != nil {
		return fmt.Errorf("复制错误：%v", err)
	}
	// 获取 logic 目录
	src = templatex.GetTplPath(acd.Mode) + "/admin/logic"
	// 获取目标目录
	dstDir = filepath.Join(rPath, apiPath, internalPath, logicPath, "admin")
	err = pathx.CopyTpl(templateFS, src, dstDir)
	if err != nil {
		return fmt.Errorf("复制错误：%v", err)
	}

	// 获取 logic 目录
	src = templatex.GetTplPath(acd.Mode) + "/admin/adminmodel"
	// 获取目标目录
	dstDir = filepath.Join(rPath, modelPath, "adminmodel")
	err = pathx.CopyTpl(templateFS, src, dstDir)
	if err != nil {
		return fmt.Errorf("复制错误：%v", err)
	}
	return nil
}
