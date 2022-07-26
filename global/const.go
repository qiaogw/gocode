package global

import (
	"log"
	"os"
	"path/filepath"
)

const (
	BuildVersion = "1.0.0"
	// ProjectName the const value of zero
	ProjectName = "zero"
	// OsWindows represents os windows
	OsWindows = "windows"
	// OsMac represents os mac
	OsMac = "darwin"
	// OsLinux represents os linux
	OsLinux = "linux"
	// OsJs represents os js
	OsJs = "js"
	// OsIOS represents os ios
	OsIOS = "ios"
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
	//return filepath.Join(p, "config", "config.toml")
	return filepath.Join(p, "config.yaml")
}
