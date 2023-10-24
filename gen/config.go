package gen

import (
	"github.com/qiaogw/gocode/common/templatex"
	"github.com/qiaogw/gocode/model"
	"os"
	"path/filepath"
	"text/template"
)

// CreateConfig 创建 config
func (acd *AutoCodeService) CreateConfig(db *model.Db) (err error) {
	pack := db.Package
	pwd, _ := os.Getwd()

	tf := "template/config/config.yaml.tpl"
	var tpl tplData
	t1 := template.New(tf)
	fi, err := templatex.TemplateTpl.ReadFile(tf)
	if err != nil {
		return
	}
	t2, err := t1.Parse(string(fi))
	if err != nil {
		return
	}
	tpl.template = t2
	mf := filepath.Join(pwd, pack+".yaml")
	f, err := os.OpenFile(mf, os.O_CREATE|os.O_WRONLY, 0o755)
	defer f.Close()
	if err != nil {
		return err
	}
	return t2.Execute(f, db)
}
