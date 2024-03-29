package pathx

import (
	"fmt"
	"github.com/qiaogw/gocode/global"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTemplateDir(t *testing.T) {
	category := "foo"
	t.Run("before_have_templates", func(t *testing.T) {
		home := t.TempDir()
		RegisterGoctlHome("")
		RegisterGoctlHome(home)
		v := global.GenConfig.System.TemplatePath
		dir := filepath.Join(home, v, category)
		err := MkdirIfNotExist(dir)
		if err != nil {
			return
		}
		tempFile := filepath.Join(dir, "bar.txt")
		err = ioutil.WriteFile(tempFile, []byte("foo"), os.ModePerm)
		if err != nil {
			return
		}
		templateDir, err := GetTemplateDir(category)
		if err != nil {
			return
		}
		assert.Equal(t, dir, templateDir)
		RegisterGoctlHome("")
	})

	t.Run("before_has_no_template", func(t *testing.T) {
		home := t.TempDir()
		RegisterGoctlHome("")
		RegisterGoctlHome(home)
		dir := filepath.Join(home, category)
		err := MkdirIfNotExist(dir)
		if err != nil {
			return
		}
		templateDir, err := GetTemplateDir(category)
		if err != nil {
			return
		}
		assert.Equal(t, dir, templateDir)
	})

	t.Run("default", func(t *testing.T) {
		RegisterGoctlHome("")
		dir, err := GetTemplateDir(category)
		if err != nil {
			return
		}
		assert.Contains(t, dir, global.GenConfig.System.TemplatePath)
	})
}

func TestGetGitHome(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}
	actual, err := GetGitHome()
	if err != nil {
		return
	}

	expected := filepath.Join(homeDir, goctlDir, gitDir)
	assert.Equal(t, expected, actual)
}

func TestGetGoctlHome(t *testing.T) {
	t.Run("goctl_is_file", func(t *testing.T) {
		tmpFile := filepath.Join(t.TempDir(), "a.tmp")
		backupTempFile := tmpFile + ".old"
		err := ioutil.WriteFile(tmpFile, nil, 0o666)
		if err != nil {
			return
		}
		RegisterGoctlHome(tmpFile)
		home, err := GetGoctlHome()
		if err != nil {
			return
		}
		info, err := os.Stat(home)
		assert.Nil(t, err)
		assert.True(t, info.IsDir())

		_, err = os.Stat(backupTempFile)
		assert.Nil(t, err)
	})

	t.Run("goctl_is_dir", func(t *testing.T) {
		RegisterGoctlHome("")
		dir := t.TempDir()
		RegisterGoctlHome(dir)
		home, err := GetGoctlHome()
		assert.Nil(t, err)
		assert.Equal(t, dir, home)
	})
}
func TestRenameFilesWithPrefixAndSuffix(t *testing.T) {

	wd, err := os.Getwd()
	if err != nil {
		assert.NoError(t, err)
	}
	parentDir := filepath.Dir(wd)
	fmt.Println("Parent directory:", parentDir)
	homeDir := filepath.Join(parentDir, "templatex", "template2", "admin", "logic")

	prefix := "admin"
	suffix := ""
	err = RenameFilesWithPrefixAndSuffix(homeDir, prefix, suffix)
	if err != nil {
		assert.NoError(t, err)
	}

}
