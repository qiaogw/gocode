package minioCli

import (
	"github.com/minio/minio-go/v7"
	"github.com/wxnacy/wgo/arrays"
	"os"
	"strings"
	"time"
)

func NewFileInfo(objectInfo minio.ObjectInfo) os.FileInfo {
	return &fileInfo{objectInfo: objectInfo}
}

type fileInfo struct {
	objectInfo minio.ObjectInfo
}

func (f *fileInfo) Name() string {
	return f.objectInfo.Key
}

func (f *fileInfo) Size() int64 {
	return f.objectInfo.Size
}

func (f *fileInfo) Mode() os.FileMode {
	return 0
}

func (f *fileInfo) ModTime() time.Time {
	return f.objectInfo.LastModified
}

func (f *fileInfo) IsDir() bool {
	return strings.HasSuffix(f.Name(), "/")
}

func (f *fileInfo) Sys() interface{} {
	return f.objectInfo
}

type FilePrefix struct {
	Key  string
	Name string
	Pid  string
	Time time.Time
	Size int64
}

func getFilePrefix(dirctory string, ld []FilePrefix) []FilePrefix {
	s := strings.Split(dirctory, "/")
	for i, v := range s {
		var td FilePrefix
		if i == 0 {
			td.Pid = "/"
		} else {
			td.Pid = ArrayToString(s[0:i])
		}
		td.Name = v + "/"
		td.Key = td.Pid + v
		if arrays.Contains(ld, td) < 0 {
			ld = append(ld, td)
			break
		}
	}
	return ld
}

// ArrayToString 数组转str，"/"为分隔符
func ArrayToString(arr []string) string {
	var result string
	for _, i := range arr { //遍历数组中所有元素追加成string
		result = result + i + "/"
	}
	return result
}
