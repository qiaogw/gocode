package global

import (
	"log"
	"os"
	"path/filepath"
)

const (
	BuildVersion = "v1.9.0"
	ProjectName  = "gocode"
	OsWindows    = "windows"
	OsMac        = "darwin"
	OsLinux      = "linux"
)

const (
	ConfigEnv         = "GVA_CONFIG"
	ConfigDefaultFile = "config.yaml"
	ConfigTestFile    = "config.test.yaml"
	ConfigDebugFile   = "config.debug.yaml"
	ConfigReleaseFile = "config.release.yaml"
)

func GetDefaultConfigFile() string {
	p, err := os.Getwd()
	if err != nil {
		log.Fatal("当前路径错误")
	}
	return filepath.Join(p, "config.yaml")
}

func GetConfigFile(f string) string {
	p, err := os.Getwd()
	if err != nil {
		log.Fatal("当前路径错误")
	}
	return filepath.Join(p, f+".yaml")
}
