package global

import (
	"log"
	"os"
	"path/filepath"
)

const (
	BuildVersion = "v1.10.58"
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

const (
	FormSelect       = "Select"
	FormInput        = "Input"
	FormInputNumber  = "InputNumber"
	FormInputText    = "InputText"
	FormEditor       = "Editor"
	FormRadio        = "Radio"
	FormCheckbox     = "Checkbox"
	FormToggle       = "Toggle"
	FormBtnToggle    = "BtnToggle"
	FormOptionGroup  = "OptionGroup"
	FormSlider       = "Slider"
	FormRange        = "Range"
	FormTimePick     = "TimePick"
	FormDatePick     = "DatePick"
	FormDateTimePick = "DateTimePick"
	FormFilePick     = "FilePick"

	FormOptionsFk   = "FkTable"
	FormOptionsDict = "Dict"
	FormOptionsList = "List"

	FormCol6  = "col-6 q-pb-md"
	FormCol12 = "col-12 q-pb-md"
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
