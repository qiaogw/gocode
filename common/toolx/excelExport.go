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
	defaultSheetName = "Sheet1" // 默认的 Sheet 名称
	defaultHeight    = 25.0     // 默认行高
)

// ExcelExport 定义了 Excel 导出所需的数据结构，包括 Excel 文件对象、Sheet 名称、参数设置、数据内容及导出路径
type ExcelExport struct {
	File      *excelize.File           `json:"file"`      // Excel 文件对象
	SheetName string                   `json:"sheetName"` // Sheet 名称，可自定义
	Params    []map[string]string      `json:"params"`    // 参数设置：每个 map 包含 key、title 和 width，用于列设置
	Data      []map[string]interface{} `json:"data"`      // 数据内容：每行数据为一个 map
	Path      string                   `json:"path"`      // 导出文件的保存路径
}

// NewMyExcel 根据传入的 Sheet 名称、标签数据（TagBody）和原始数据创建一个 ExcelExport 对象
func NewMyExcel(sheetName string, tag TagBody, data interface{}) (*ExcelExport, error) {
	e := new(ExcelExport)
	e.SheetName = sheetName
	e.File = createFile(sheetName)
	// 根据 tag 中的 Keys 和 Header 构建列参数，每列默认宽度设为20
	for i, v := range tag.Keys {
		p := make(map[string]string)
		p["key"] = v
		p["title"] = tag.Header[i]
		p["width"] = "20"
		e.Params = append(e.Params, p)
	}
	// 将 data 序列化为 JSON，再反序列化为 []map[string]interface{}
	dj, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var dm []map[string]interface{}
	err = json.Unmarshal(dj, &dm)
	if err != nil {
		return nil, err
	}
	// 过滤 data 中的字段，只保留 tag 中定义的 key
	for _, v := range dm {
		st := make(map[string]interface{})
		for _, o := range tag.Keys {
			st[o], _ = v[o]
		}
		e.Data = append(e.Data, st)
	}
	return e, nil
}

// ExportToPath 将 Excel 文件导出到指定路径，并返回生成的文件路径及可能的错误
func (l *ExcelExport) ExportToPath() (string, error) {
	l.export()
	name := createFileName()
	filePath := l.Path + "/" + name
	err := l.File.SaveAs(filePath)
	return filePath, err
}

// ExportToWeb 将 Excel 文件导出为字节缓冲区，用于 Web 端下载或预览
func ExportToWeb(m, list interface{}, name string) (*bytes.Buffer, error) {
	tag := GetTag(m) // 获取 TagBody 信息，此函数需在其他地方定义
	ex, err := NewMyExcel(name, tag, list)
	if err != nil {
		return nil, err
	}
	ex.export()
	return ex.File.WriteToBuffer()
}

// ExportToWebTemplate 生成一个空模板 Excel 文件（只有表头）并返回缓冲区
func ExportToWebTemplate(m interface{}, name string) (*bytes.Buffer, error) {
	tag := GetTag(m)
	// 空数据列表
	list := make([]map[string]interface{}, 0)
	ex, err := NewMyExcel(name, tag, list)
	if err != nil {
		return nil, err
	}
	ex.export()
	return ex.File.WriteToBuffer()
}

// ExportToWebGin 使用 Gin 框架将 Excel 文件导出到浏览器进行下载
func ExportToWebGin(c *gin.Context, m, list interface{}, name string) {
	tag := GetTag(m)
	ex, err := NewMyExcel(name, tag, list)
	if err != nil {
		c.Error(err)
	}
	ex.ExportToWeb(c)
}

// ExportToWeb (方法) 将 Excel 文件导出到浏览器
// 设置响应头，指定文件类型和文件名，然后写入缓冲区内容到响应流
func (l *ExcelExport) ExportToWeb(ctx *gin.Context) {
	l.export()
	buffer, _ := l.File.WriteToBuffer()
	// 设置文件类型为 Excel 文件
	ctx.Header("Content-Type", "application/vnd.ms-excel;charset=utf8")
	// 设置下载文件名称，并进行 URL 编码
	ctx.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(createFileName()))
	_, _ = ctx.Writer.Write(buffer.Bytes())
}

// writeTop 在 Excel 文件中写入表头
func (l *ExcelExport) writeTop() {
	// 定义表头样式：字体加粗，水平和垂直居中
	style := excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	}
	topStyle, _ := l.File.NewStyle(&style)

	var word = 1
	// 逐列写入表头数据
	for _, conf := range l.Params {
		title := conf["title"]
		width, _ := strconv.ParseFloat(conf["width"], 64)
		col, _ := excelize.ColumnNumberToName(word)
		line := fmt.Sprintf("%s1", col)
		// 写入单元格标题
		_ = l.File.SetCellValue(l.SheetName, line, title)
		// 设置列宽
		_ = l.File.SetColWidth(l.SheetName, col, col, width)
		// 应用样式
		_ = l.File.SetCellStyle(l.SheetName, line, line, topStyle)
		word++
	}
}

// writeData 将数据行写入 Excel 文件
func (l *ExcelExport) writeData() {
	// 定义数据行样式：内容水平和垂直居中
	style := excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	}
	lineStyle, _ := l.File.NewStyle(&style)
	// 数据从第二行开始写入
	var j = 2
	for i, val := range l.Data {
		// 设置当前行的行高为默认值
		_ = l.File.SetRowHeight(l.SheetName, i+1, defaultHeight)
		var word = 1
		// 逐列写入数据
		for _, conf := range l.Params {
			valKey := conf["key"]
			col, _ := excelize.ColumnNumberToName(word)
			line := fmt.Sprintf("%s%v", col, j)
			// 写入单元格数据
			_ = l.File.SetCellValue(l.SheetName, line, val[valKey])
			// 设置单元格样式
			_ = l.File.SetCellStyle(l.SheetName, line, line, lineStyle)
			word++
		}
		j++
	}
	// 设置最后一行的行高
	_ = l.File.SetRowHeight(l.SheetName, len(l.Data)+1, defaultHeight)
}

// export 整体执行 Excel 导出流程，包括写入表头和数据
func (l *ExcelExport) export() {
	l.writeTop()
	l.writeData()
}

// createFile 根据传入的 Sheet 名称创建一个 Excel 文件对象
// 如果未指定 Sheet 名称，则使用默认的 Sheet 名称；否则创建指定名称的 Sheet 并删除默认 Sheet
func createFile(sheetNames ...string) *excelize.File {
	f := excelize.NewFile()
	var index int
	if len(sheetNames) < 1 {
		index, _ = f.NewSheet(defaultSheetName)
	} else {
		for _, s := range sheetNames {
			index, _ = f.NewSheet(s)
		}
		err := f.DeleteSheet(defaultSheetName)
		if err != nil {
			return nil
		}
	}
	// 设置默认激活的工作表
	f.SetActiveSheet(index)
	return f
}

// createFileName 根据当前时间和随机数生成一个唯一的 Excel 文件名
func createFileName() string {
	name := time.Now().Format("2006-01-02-15-04-05")
	//rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("excle-%v-%v.xlsx", name, rand.Int63n(time.Now().Unix()))
}
