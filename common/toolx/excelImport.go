package toolx

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"io"
	"log"
	"reflect"
	"strconv"
	"time"
)

// ExcelData 定义了 Excel 数据处理接口，包含创建映射和转换时间的方法
type ExcelData interface {
	CreateMap(arr []string) map[string]interface{}
	ChangeTime(source string) time.Time
}

// ExcelStruct 用于处理 Excel 导入的结构体，包括临时数据、模型、映射信息、最终数据、内容和 ID 类型标识
type ExcelStruct struct {
	Temp    [][]string               // 从 Excel 中读取的原始数据（二维字符串数组）
	Model   interface{}              // 目标数据模型（结构体类型）
	Info    []map[string]string      // 从 Excel 每行数据转换后的映射信息（字段名 -> 值）
	Data    []map[string]interface{} // 最终生成的数据列表，准备写入数据库
	Content []byte                   // 序列化后的数据内容（JSON 格式）
	IdIsInt bool                     // ID 是否为整数类型，如果为 false 则使用 UUID
}

// ReadExcel 从文件路径 file 读取 Excel 文件，并将所有 Sheet 的数据合并后存入 Temp 字段，返回 ExcelStruct 指针
func (excel *ExcelStruct) ReadExcel(file string) *ExcelStruct {
	xlsx, err := excelize.OpenFile(file)
	if err != nil {
		return nil
	}
	var rows [][]string
	sheets := xlsx.GetSheetList()
	for _, s := range sheets {
		row, _ := xlsx.GetRows(s)
		rows = append(rows, row...)
	}
	excel.Temp = rows
	return excel
}

// ReadExcelIo 从 io.Reader 中读取 Excel 文件，并处理数据转换后存入 Data 字段，同时生成 JSON 内容
// tx 为数据库事务对象，用于获取当前最大 ID
func (excel *ExcelStruct) ReadExcelIo(tx *gorm.DB, file io.Reader) error {
	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		return err
	}
	sheets := xlsx.GetSheetList()
	rows, _ := xlsx.GetRows(sheets[0])
	// 忽略标题行
	excel.Temp = rows[1:]

	// 生成每行数据的映射信息
	err = excel.CreateMap()
	if err != nil {
		return err
	}
	// 根据 ID 类型分别生成数据
	if excel.IdIsInt {
		err = excel.GenDataInt(tx)
		if err != nil {
			return err
		}
	} else {
		err = excel.GenDataChar(tx)
		if err != nil {
			return err
		}
	}

	// 将生成的数据序列化为 JSON 存入 Content 字段
	excel.Content, err = json.Marshal(excel.Data)
	return err
}

// CreateMap 利用 Excel 中的临时数据 Temp 以及模型对应的标签信息，将每一行转换为 map（字段 -> 值）存入 Info 字段
func (excel *ExcelStruct) CreateMap() error {
	tag := GetTag(excel.Model)
	// 遍历每一行数据
	for _, v := range excel.Temp {
		// 将行数据转换为 map 格式
		var info = make(map[string]string)
		// 通过反射获取模型中定义的字段
		for i := 0; i < reflect.TypeOf(excel.Model).Elem().NumField(); i++ {
			obj := reflect.TypeOf(excel.Model).Elem().Field(i)
			// 根据 tag.Field 中的字段名匹配模型字段
			for j, h := range tag.Field {
				if obj.Name == h {
					if len(v) < j+1 {
						continue
					}
					info[tag.Keys[j]] = v[j]
				}
			}
		}
		excel.Info = append(excel.Info, info)
	}
	return nil
}

// ChangeTime 将字符串 source 按 "2006-01-02" 格式解析为时间类型，解析失败则终止程序
func (excel *ExcelStruct) ChangeTime(source string) time.Time {
	ChangeAfter, err := time.Parse("2006-01-02", source)
	if err != nil {
		log.Fatalf("转换时间错误:%s", err)
	}
	return ChangeAfter
}

// GenDataInt 根据 Info 中的映射数据生成 Data 数据，假设 ID 为整数类型
// tx 为数据库事务对象，用于获取当前最大 ID，以便生成新的连续 ID
func (excel *ExcelStruct) GenDataInt(tx *gorm.DB) (err error) {
	temp := make([]map[string]interface{}, 0)
	tag := GetTag(excel.Model)
	var id int64
	// 查询数据库中当前模型的最大 ID
	err = tx.Debug().Model(excel.Model).Unscoped().Select("max(id) as mid").
		Take(&id).Error
	if err != nil {
		id = 0
	}
	// 遍历每一行 Info 数据，生成对应的 map 数据，并自动生成 ID
	for i := 0; i < len(excel.Info); i++ {
		id++
		t := reflect.ValueOf(excel.Model).Elem()
		data := make(map[string]interface{})
		data["Id"] = id
		// 对于每个字段，根据模型中字段的类型进行转换
		for k, v := range excel.Info[i] {
			field, err := tag.GetFieldByTag(k)
			if err != nil {
				continue
			}
			switch t.FieldByName(field).Kind() {
			case reflect.String:
				data[field] = v
			case reflect.Float64:
				tempV, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				data[field] = tempV
			case reflect.Uint64:
				tempV, err := strconv.ParseUint(v, 0, 64)
				if err != nil {
					return err
				}
				data[field] = tempV
			case reflect.Struct:
				tempV, err := time.Parse("2006-01-02", v)
				if err != nil {
					return err
				}
				data[field] = tempV
			default:
				continue
			}
		}
		temp = append(temp, data)
	}
	excel.Data = temp
	return nil
}

// GenDataChar 根据 Info 中的映射数据生成 Data 数据，假设 ID 为非整数类型（例如 UUID）
func (excel *ExcelStruct) GenDataChar(tx *gorm.DB) (err error) {
	temp := make([]map[string]interface{}, 0)
	tag := GetTag(excel.Model)
	// 遍历每一行 Info 数据，生成对应的 map 数据，并生成 UUID 作为 ID
	for i := 0; i < len(excel.Info); i++ {
		id := uuid.New()
		t := reflect.ValueOf(excel.Model).Elem()
		data := make(map[string]interface{})
		data["Id"] = id
		// 根据模型字段类型转换对应的字符串值
		for k, v := range excel.Info[i] {
			field, err := tag.GetFieldByTag(k)
			if err != nil {
				continue
			}
			switch t.FieldByName(field).Kind() {
			case reflect.String:
				data[field] = v
			case reflect.Float64:
				tempV, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				data[field] = tempV
			case reflect.Uint64:
				tempV, err := strconv.ParseUint(v, 0, 64)
				if err != nil {
					return err
				}
				data[field] = tempV
			case reflect.Struct:
				tempV, err := time.Parse("2006-01-02", v)
				if err != nil {
					return err
				}
				data[field] = tempV
			default:
				continue
			}
		}
		temp = append(temp, data)
	}
	excel.Data = temp
	return nil
}

// SaveDb 将 Excel 导入的数据保存到数据库中
// tx 为数据库事务对象，reader 为 Excel 文件的 io.Reader
// 数据按批次（每批 1000 条）写入，失败时回滚事务
func (excel *ExcelStruct) SaveDb(tx *gorm.DB, reader io.Reader) (err error) {
	tx = tx.Begin()
	defer func() {
		if err != nil {
			logx.Error(err)
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = excel.ReadExcelIo(tx, reader)
	if err != nil {
		return err
	}
	// 分批写入数据，批次大小为 1000
	for i := 0; i < len(excel.Data); i += 1000 {
		end := i + 1000
		if end > len(excel.Data) {
			end = len(excel.Data)
		}
		d := excel.Data[i:end]
		err = tx.Model(excel.Model).Create(&d).Error
		if err != nil {
			return err
		}
	}
	return nil
}
