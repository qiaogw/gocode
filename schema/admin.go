package schema

import (
	"encoding/json"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/qiaogw/gocode/schema/model"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var migrateList = []interface{}{
	model.Api{},
	model.Config{},
	model.Dept{},
	model.DictType{},
	model.DictData{},
	model.LoginLog{},
	model.Menu{},
	model.Migration{},
	model.OperaLog{},
	model.Post{},
	model.Role{},
	model.User{},
}

type tbList struct {
	Api      []model.Api
	Config   []model.Config
	Dept     []model.Dept
	DictData []model.DictData
	DictType []model.DictType
	RoleApi  []RoleApi
	Menu     []model.Menu
	UserRole []UserRole
	RoleMenu []RoleMenu
	Post     []model.Post
	Role     []model.Role
	User     []model.User
}

type tbm struct {
	Api      model.Api
	Config   model.Config
	Dept     model.Dept
	DictData model.DictData
	DictType model.DictType
	RoleApi  RoleApi
	Menu     model.Menu
	RoleMenu RoleMenu
	UserRole UserRole
	Post     model.Post
	Role     model.Role
	User     model.User
}

func NewAdmin(db *gorm.DB) error {
	err := db.AutoMigrate(migrateList...)
	if err != nil {
		return errors.New("初始化 Admin 数据库失败：" + err.Error())
	}
	loadDate(db)
	return nil
}
func loadDate(db *gorm.DB) error {

	backupFolder := filepath.Join("schema", "data")
	// 打开备份文件夹
	folder, err := os.Open(backupFolder)
	fileInfos, err := folder.Readdir(-1)
	if err != nil {
		return err
	}
	tl := tbList{
		Api:      []model.Api{},
		Config:   []model.Config{},
		Dept:     []model.Dept{},
		DictData: []model.DictData{},
		DictType: []model.DictType{},
		RoleApi:  []RoleApi{},
		Menu:     []model.Menu{},
		RoleMenu: []RoleMenu{},
		UserRole: []UserRole{},
		Post:     []model.Post{},
		Role:     []model.Role{},
		User:     []model.User{},
	}

	tbs := tbm{
		Api:      model.Api{},
		Config:   model.Config{},
		Dept:     model.Dept{},
		DictData: model.DictData{},
		DictType: model.DictType{},
		RoleApi:  RoleApi{},
		Menu:     model.Menu{},
		RoleMenu: RoleMenu{},
		UserRole: UserRole{},
		Post:     model.Post{},
		Role:     model.Role{},
		User:     model.User{},
	}

	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			// 处理每个文件
			fileName := fileInfo.Name()
			tableName := strings.TrimSuffix(fileName, ".json")

			jsonData, err := readJSONFile(backupFolder + "/" + fileName)
			if err != nil {
				log.Printf("读取 JSON 文件 %s 错误: %v\n", fileName, err)
				continue
			}
			jd, err := json.Marshal(jsonData)

			switch tableName {
			case tbs.Api.TableName():
				er := json.Unmarshal(jd, &tl.Api)
				if er != nil {
					fmt.Println("解析文件", tableName, " 错误: ", er)
				}
				er = db.Create(tl.Api).Error
				if er != nil {
					fmt.Println("创建", tableName, " 错误: ", er)
				}
			case tbs.Config.TableName():
				er := json.Unmarshal(jd, &tl.Config)
				if er != nil {
					fmt.Println("解析文件", tableName, " 错误: ", er)
				}
				er = db.Create(tl.Config).Error
				if er != nil {
					fmt.Println("创建", tableName, " 错误: ", er)
				}
			case tbs.Dept.TableName():
				er := json.Unmarshal(jd, &tl.Dept)

				if er != nil {
					fmt.Println("解析文件", tableName, " 错误: ", er)
				}
				er = db.Create(tl.Dept).Error
				if er != nil {
					fmt.Println("创建", tableName, " 错误: ", er)
				}
			case tbs.DictType.TableName():
				er := json.Unmarshal(jd, &tl.DictType)

				if er != nil {
					fmt.Println("解析文件", tableName, " 错误: ", er)
				}
				er = db.Create(tl.DictType).Error
				if er != nil {
					fmt.Println("创建", tableName, " 错误: ", er)
				}
			case tbs.DictData.TableName():
				er := json.Unmarshal(jd, &tl.DictData)

				if er != nil {
					fmt.Println("解析文件", tableName, " 错误: ", er)
				}
				er = db.Create(tl.DictData).Error
				if er != nil {
					fmt.Println("创建", tableName, " 错误: ", er)
				}
			case tbs.Menu.TableName():
				er := json.Unmarshal(jd, &tl.Menu)

				if er != nil {
					fmt.Println("解析文件", tableName, " 错误: ", er)
				}
				er = db.Create(tl.Menu).Error
				if er != nil {
					fmt.Println("创建", tableName, " 错误: ", er)
				}
			case tbs.Role.TableName():
				er := json.Unmarshal(jd, &tl.Role)
				if er != nil {
					fmt.Println("解析文件", tableName, " 错误: ", er)
				}
				er = db.Create(tl.Role).Error
				if er != nil {
					fmt.Println("创建", tableName, " 错误: ", er)
				}
			case tbs.User.TableName():
				er := json.Unmarshal(jd, &tl.User)
				if er != nil {
					fmt.Println("解析文件", tableName, " 错误: ", er)
				}
				er = db.Create(tl.User).Error
				if er != nil {
					fmt.Println("创建", tableName, " 错误: ", er)
				}
			case tbs.RoleApi.TableName():
				er := json.Unmarshal(jd, &tl.RoleApi)
				if er != nil {
					fmt.Println("解析文件", tableName, " 错误: ", er)
				}
				er = db.Create(tl.RoleApi).Error
				if er != nil {
					fmt.Println("创建", tableName, " 错误: ", er)
				}
			case tbs.RoleMenu.TableName():
				er := json.Unmarshal(jd, &tl.RoleMenu)
				if er != nil {
					fmt.Println("解析文件", tableName, " 错误: ", er)
				}
				er = db.Create(tl.RoleMenu).Error
				if er != nil {
					fmt.Println("创建", tableName, " 错误: ", er)
				}
			case tbs.UserRole.TableName():
				er := json.Unmarshal(jd, &tl.UserRole)
				if er != nil {
					fmt.Println("解析文件", tableName, " 错误: ", er)
				}
				er = db.Create(tl.UserRole).Error
				if er != nil {
					fmt.Println("创建", tableName, " 错误: ", er)
				}
			}
		}
	}
	return nil
}

func readJSONFile(fileName string) ([]map[string]interface{}, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := jsoniter.NewDecoder(file)
	var data []map[string]interface{}
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

type RoleMenu struct {
	RoleId string `gorm:"primaryKey;" json:"role_id"`
	MenuId string `gorm:"primaryKey;" json:"menu_id"`
}

func (RoleMenu) TableName() string {
	return "admin_role_menu"
}

type RoleApi struct {
	RoleId string `gorm:"primaryKey;" json:"role_id"`
	ApiId  string `gorm:"primaryKey;" json:"api_id"`
}

func (RoleApi) TableName() string {
	return "admin_role_api"
}

type UserRole struct {
	RoleId string `gorm:"primaryKey;" json:"role_id"`
	UserId string `gorm:"primaryKey;" json:"user_id"`
}

func (UserRole) TableName() string {
	return "admin_user_role"
}
