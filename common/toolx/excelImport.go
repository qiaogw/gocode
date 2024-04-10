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

type ExcelData interface {
	CreateMap(arr []string) map[string]interface{}
	ChangeTime(source string) time.Time
}

type ExcelStruct struct {
	Temp    [][]string
	Model   interface{}
	Info    []map[string]string
	Data    []map[string]interface{}
	Content []byte
	IdIsInt bool
}

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

func (excel *ExcelStruct) ReadExcelIo(tx *gorm.DB, file io.Reader) error {
	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		return err
	}
	sheets := xlsx.GetSheetList()
	rows, _ := xlsx.GetRows(sheets[0])
	//忽略标题行
	excel.Temp = rows[1:]

	err = excel.CreateMap()
	if err != nil {
		return err
	}
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

	excel.Content, err = json.Marshal(excel.Data)
	return err
}
func (excel *ExcelStruct) CreateMap() error {
	tag := GetTag(excel.Model)
	//利用反射得到字段名
	for _, v := range excel.Temp {
		//将数组  转成对应的 map
		var info = make(map[string]string)
		for i := 0; i < reflect.TypeOf(excel.Model).Elem().NumField(); i++ {
			obj := reflect.TypeOf(excel.Model).Elem().Field(i)
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

func (excel *ExcelStruct) ChangeTime(source string) time.Time {
	ChangeAfter, err := time.Parse("2006-01-02", source)
	if err != nil {
		log.Fatalf("转换时间错误:%s", err)
	}
	return ChangeAfter
}

func (excel *ExcelStruct) GenDataInt(tx *gorm.DB) (err error) {
	temp := make([]map[string]interface{}, 0)
	tag := GetTag(excel.Model)
	var id int64
	err = tx.Debug().Model(excel.Model).Unscoped().Select("max(id) as mid").
		Take(&id).Error
	if err != nil {
		id = 0
	}
	for i := 0; i < len(excel.Info); i++ {
		id++
		t := reflect.ValueOf(excel.Model).Elem()
		data := make(map[string]interface{})
		data["Id"] = id
		for k, v := range excel.Info[i] {
			field, err := tag.GetFieldByTag(k)
			if err != nil {
				continue
			}
			switch t.FieldByName(field).Kind() {
			case reflect.String:
				data[field] = v
				//t.FieldByName(field).Set(reflect.ValueOf(v))
			case reflect.Float64:
				tempV, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				data[field] = tempV
				//t.FieldByName(field).Set(reflect.ValueOf(tempV))
			case reflect.Uint64:
				reflect.ValueOf(v)
				tempV, err := strconv.ParseUint(v, 0, 64)
				if err != nil {
					return err
				}
				data[field] = tempV
				//t.FieldByName(field).Set(reflect.ValueOf(tempV))
			case reflect.Struct:
				tempV, err := time.Parse("2006-01-02", v)
				if err != nil {
					return err
				}
				data[field] = tempV
				//t.FieldByName(field).Set(reflect.ValueOf(tempV))
			default:
				continue
			}
		}
		temp = append(temp, data)

	}
	excel.Data = temp

	return nil
}

func (excel *ExcelStruct) GenDataChar(tx *gorm.DB) (err error) {
	temp := make([]map[string]interface{}, 0)
	tag := GetTag(excel.Model)
	//var id int64
	//err = tx.Debug().Model(excel.Model).Unscoped().Select("max(id) as mid").
	//	Take(&id).Error
	//if err != nil {
	//	id = 0
	//}
	for i := 0; i < len(excel.Info); i++ {
		id := uuid.New()
		t := reflect.ValueOf(excel.Model).Elem()
		data := make(map[string]interface{})
		data["Id"] = id
		for k, v := range excel.Info[i] {
			field, err := tag.GetFieldByTag(k)
			if err != nil {
				continue
			}
			switch t.FieldByName(field).Kind() {
			case reflect.String:
				data[field] = v
				//t.FieldByName(field).Set(reflect.ValueOf(v))
			case reflect.Float64:
				tempV, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				data[field] = tempV
				//t.FieldByName(field).Set(reflect.ValueOf(tempV))
			case reflect.Uint64:
				reflect.ValueOf(v)
				tempV, err := strconv.ParseUint(v, 0, 64)
				if err != nil {
					return err
				}
				data[field] = tempV
				//t.FieldByName(field).Set(reflect.ValueOf(tempV))
			case reflect.Struct:
				tempV, err := time.Parse("2006-01-02", v)
				if err != nil {
					return err
				}
				data[field] = tempV
				//t.FieldByName(field).Set(reflect.ValueOf(tempV))
			default:
				continue
			}
		}
		temp = append(temp, data)

	}
	excel.Data = temp

	return nil
}
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
