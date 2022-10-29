package toolx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"math/rand"
	"net/url"
	"strconv"
	"time"
)

var (
	defaultSheetName = "Sheet1" //默认Sheet名称
	defaultHeight    = 25.0     //默认行高度
)

type ExcelExport struct {
	File      *excelize.File           `json:"file"`
	SheetName string                   `json:"sheetName"` //可定义默认sheet名称
	Params    []map[string]string      `json:"params"`
	Data      []map[string]interface{} `json:"data"`
	Path      string                   `json:"path"`
}

func NewMyExcel(sheetName string, tag TagBody, data interface{}) (*ExcelExport, error) {
	e := new(ExcelExport)
	e.SheetName = sheetName
	e.File = createFile(sheetName)
	for i, v := range tag.Keys {
		p := make(map[string]string)
		p["key"] = v
		p["title"] = tag.Header[i]
		p["width"] = "20"
		e.Params = append(e.Params, p)
	}
	dj, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var dm []map[string]interface{}
	err = json.Unmarshal(dj, &dm)
	if err != nil {
		return nil, err
	}
	for _, v := range dm {
		st := make(map[string]interface{})
		for _, o := range tag.Keys {
			st[o], _ = v[o]
		}
		e.Data = append(e.Data, st)
	}
	return e, nil
}

// ExportToPath 导出基本的表格
func (l *ExcelExport) ExportToPath() (string, error) {
	l.export()
	name := createFileName()
	filePath := l.Path + "/" + name
	err := l.File.SaveAs(filePath)
	return filePath, err
}

func ExportToWeb(m, list interface{}, name string) (*bytes.Buffer, error) {
	tag := GetTag(m)
	ex, err := NewMyExcel(name, tag, list)

	if err != nil {
		return nil, err
	}
	ex.export()
	return ex.File.WriteToBuffer()
}

func ExportToWebTemplate(m interface{}, name string) (*bytes.Buffer, error) {
	tag := GetTag(m)
	list := make([]map[string]interface{}, 0)
	ex, err := NewMyExcel(name, tag, list)

	if err != nil {
		return nil, err
	}
	ex.export()
	return ex.File.WriteToBuffer()
}

func ExportToWebGin(c *gin.Context, m, list interface{}, name string) {
	tag := GetTag(m)
	ex, err := NewMyExcel(name, tag, list)
	if err != nil {
		c.Error(err)
	}
	ex.ExportToWeb(c)
}

// ExportToWeb 导出到浏览器。此处使用的gin框架 其他框架可自行修改ctx
func (l *ExcelExport) ExportToWeb(ctx *gin.Context) {
	l.export()
	buffer, _ := l.File.WriteToBuffer()
	//设置文件类型
	ctx.Header("Content-Type", "application/vnd.ms-excel;charset=utf8")
	//设置文件名称
	ctx.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(createFileName()))
	_, _ = ctx.Writer.Write(buffer.Bytes())
}

// 设置首行
func (l *ExcelExport) writeTop() {
	topStyle, _ := l.File.NewStyle(`{"font":{"bold":true},"alignment":{"horizontal":"center","vertical":"center"}}`)
	var word = 1
	//首行写入
	for _, conf := range l.Params {
		title := conf["title"]
		width, _ := strconv.ParseFloat(conf["width"], 64)
		col, _ := excelize.ColumnNumberToName(word)
		line := fmt.Sprintf("%s1", col)
		//设置标题
		_ = l.File.SetCellValue(l.SheetName, line, title)
		//列宽
		_ = l.File.SetColWidth(l.SheetName, col, col, width)
		//设置样式
		_ = l.File.SetCellStyle(l.SheetName, line, line, topStyle)
		word++
	}
}

// 写入数据
func (l *ExcelExport) writeData() {
	lineStyle, _ := l.File.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"}}`)
	//数据写入
	var j = 2 //数据开始行数
	for i, val := range l.Data {
		//设置行高
		_ = l.File.SetRowHeight(l.SheetName, i+1, defaultHeight)
		//逐列写入
		var word = 1
		for _, conf := range l.Params {
			valKey := conf["key"]
			col, _ := excelize.ColumnNumberToName(word)
			line := fmt.Sprintf("%s%v", col, j)
			//设置值
			_ = l.File.SetCellValue(l.SheetName, line, val[valKey])
			//设置样式
			_ = l.File.SetCellStyle(l.SheetName, line, line, lineStyle)
			word++
		}
		j++
	}
	//设置行高 尾行
	_ = l.File.SetRowHeight(l.SheetName, len(l.Data)+1, defaultHeight)
}

func (l *ExcelExport) export() {
	l.writeTop()
	l.writeData()
}

func createFile(sheetNames ...string) *excelize.File {
	f := excelize.NewFile()
	// 创建一个默认工作表
	//SheetName := defaultSheetName
	var index int
	if len(sheetNames) < 1 {
		index = f.NewSheet(defaultSheetName)
	} else {
		for _, s := range sheetNames {
			index = f.NewSheet(s)
		}
		f.DeleteSheet(defaultSheetName)
	}
	//index := f.NewSheet(sheetName)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	return f
}

func createFileName() string {
	name := time.Now().Format("2006-01-02-15-04-05")
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("excle-%v-%v.xlsx", name, rand.Int63n(time.Now().Unix()))
}
