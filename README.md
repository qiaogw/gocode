go-zero 代码自动生成辅助工具，支持gorm。
# 安装
GOPROXY=https://goproxy.cn/,direct go install github.com/qiaogw/gocode@latest
# 使用
 首先 gocode init，在当前目录和dbconf下生成config.yaml,根据业务自行编辑。

## gocode gen -p 服务名 即可生成遵循go-zero标准的api和rpc服务。
## 其他配置参考 go-zero 和 gorm

# 数据库
 支持mysql、postgres

## 建表要求
### 1、主键名称必须为：id, 最好为int自增
### 2、符合 gorm 标准
必须有以下字段：

 ```
type BaseModel struct {
Id int64 `json:"id" db:"id" gorm:"column:id;primaryKey;autoIncrement;comment:主键编码"`
}

type ControlBy struct {
CreateBy int64 `json:"createBy" db:"create_by" gorm:"column:create_by;index;comment:创建者"`
UpdateBy int64 `json:"updateBy" db:"update_by" gorm:"column:update_by;index;comment:更新者"`
}

type ModelTime struct {
CreatedAt time.Time      `json:"createdAt"  db:"created_at" gorm:"column:created_at;comment:创建时间"`
UpdatedAt time.Time      `json:"updatedAt"  db:"updated_at" gorm:"column:updated_at;comment:最后更新时间"`
DeletedAt gorm.DeletedAt `json:"-" db:"deleted_at" gorm:"index;comment:删除时间"` // 软删除
}
 ```

若业务要求更改以上字段，则需要修改生成的 common/golab/model.go文件。
生成的文件最好在goland中编辑，对rpc和api文件利用goctl插件进行生成，插件有bug，api下的handel有包重复import，
此时可以利用goland 进行格式化和清理优化.

## 备份业务数据
### backup
会在当前backup目录下创建以数据库为名的文件夹，将所有表导出为json文件

### restore
会将 backup 的对应数据库json文件导入的 -d 数据库（对应配置文件,文件名为数据库名称）

## Good job！