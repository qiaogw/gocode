package minioCli

import (
	"encoding/json"
	"fmt"
	"github.com/minio/minio-go/v7"
	"os"
	"path"
	"reflect"
	"strings"
	"time"
)

type S3Config struct {
	Endpoint        string `label:"地址"`              // 地址
	AccessKeyID     string `label:"AccessKeyID"`     // 地址
	SecretAccessKey string `label:"SecretAccessKey"` // 地址
	Region          string `label:"对象存储的region"`     // 对象存储的region
	Bucket          string `label:"对象存储的Bucket"`     // 对象存储的Bucket
	Secure          bool   `label:"true代表使用HTTPS"`   // true代表使用HTTPS
	Ignore          string `label:"隐藏文件，S3不支持空目录"`   // 地址
	Dir             string `label:"地址"`              // 地址
	CacheDir        string `label:"Cache地址"`         // 地址
	ConfigFile      string `label:"ConfigFile地址"`    // 地址
	LogFile         string `label:"LogFile地址"`       // 地址
}

type DirBody struct {
	Label    string    `json:"label"`
	Children []DirBody `json:"children"`
	Icon     string    `json:"icon"`
	Dir      string    `json:"dir"`
}

type DirInfo struct {
	Path         string     `json:"path"`
	Entries      []*Entries `json:"entries"`
	Limit        int        `json:"limit"`
	LastFileName string     `json:"lastFileName"`
}

type Entries struct {
	Mtime         time.Time         `json:"mtime"`       // time of last modification
	Crtime        time.Time         `json:"crtime"`      // time of creation (OS X only)
	Mode          os.FileMode       `json:"mode"`        // file mode
	Uid           uint32            `json:"uid"`         // owner uid
	Gid           uint32            `json:"gid"`         // group gid
	Mime          string            `json:"nime"`        // mime type
	Replication   string            `json:"replication"` // replication
	Collection    string            `json:"collection"`  // collection name
	TtlSec        int32             `json:"ttl_sec"`     // ttl in seconds
	UserName      string            `json:"user_name"`
	GroupNames    []string          `json:"group_names"`
	SymlinkTarget string            `json:"symlink_target"`
	Md5           []byte            `json:"md5"`
	FileSize      int64             `json:"file_size,omitempty"`
	Extended      map[string][]byte `json:"extended"`
	FullPath      string            `json:"full_path"`
	Chunks        []*FileChunk      `json:"chunks,omitempty"`
	IsDirectory   bool              `json:"is_directory"`
	Type          string            `json:"type"`
	Name          string            `json:"name"`
	URLDir        string            `json:"url_dir"`
	StorageClass  string            `json:"storage_class"`
	Restore       string            `json:"restore"`
}
type FileId struct {
	volumeId uint32
	fileKey  uint64
	cookie   uint64
}
type FileChunk struct {
	FileId string `json:"file_id"`
	Offset int64  `json:"offset"`
	Size   int64  `json:"size"`
	Mtime  int64  `json:"mtime"`
	ETag   string `json:"e_tag"`
	Fid    FileId `json:"fid"`
}

type Policy struct {
	Version   string `json:"Version,omitempty"`
	Statement []Statement
}
type Principal struct {
	AWS []string
}
type StringEquals struct {
	Class []string `json:"s3:x-amz-storage-class"`
}

type Condition struct {
	StringEquals StringEquals
}
type Statement struct {
	Sid       string
	Effect    string
	Principal Principal
	Action    string
	Resource  string `json:"Resource,omitempty"`
	Condition Condition
}

// ObjectInfoToBytes 将 minio.ObjectInfo 转换为 JSON 字节数组
func ObjectInfoToBytes(info minio.ObjectInfo) ([]byte, error) {
	// 将 ObjectInfo 转换为 JSON 字节流
	jsonBytes, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

// GetInfo 获取 对象的Entries信息
func GetInfo(pathStr, urls string, v fileInfo) (e Entries, err error) {
	//var e helper.Entries
	e.Name = v.Name()
	e.FileSize = v.Size()
	e.IsDirectory = v.IsDir()
	e.Mode = v.Mode()
	e.Mtime = v.ModTime()

	e.Name = strings.TrimRight(e.Name, "/")

	if pathStr != "/" {
		e.FullPath = fmt.Sprint(pathStr, "/", e.Name)
	} else {
		e.FullPath = fmt.Sprint("/", e.Name)
	}

	e.URLDir = urls + e.FullPath
	if e.IsDirectory {
		e.Type = "文件夹"
	} else {
		e.Type = path.Ext(e.Name)
		if len(e.Type) > 0 {
			e.Type = e.Type[1:len(e.Type)]
		}
	}
	sinfo := v.Sys()
	rv := reflect.ValueOf(sinfo).Type()
	if rv.String() == "s3.Object" {
		ob := sinfo.(minio.ObjectInfo)
		e.StorageClass = ob.StorageClass
		//if ob.StorageClass == "GLACIER" {
		//	svc := NewSvc(s3conf)
		//	hb, err := svc.Stat(svc.BucketName, e.FullPath)
		//	if err == nil {
		//		if hb.Restore == nil {
		//			e.Restore = "未恢复"
		//		} else {
		//			if strings.Index(*hb.Restore, `ongoing-request="false"`) < 0 {
		//				e.Restore = "恢复中"
		//			} else {
		//				e.Restore = "已恢复"
		//			}
		//		}
		//	}
		//}
	}
	//if path.Ext(e.Name) == s3conf.Ignore {
	//	err = errors.New("隐藏")
	//}
	return
}
