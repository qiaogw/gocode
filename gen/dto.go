package gen

import (
	"gocode/models"
	"gocode/utils"
)

// GenTable  表数据重构
func (acd *AutoCodeService) GenTable(v models.Table) (d models.AutoCodeStruct) {
	d.TableName = v.TableName
	d.PackageName = "zero"
	d.Abbreviation = utils.CamelString(v.TableName)
	d.StructName = utils.LeftUpper(d.Abbreviation)
	d.PackageName = d.Abbreviation
	d.HumpPackageName = v.TableName
	d.Description = d.Abbreviation + "表"

	//data = append(data, d)

	return d
}

// GenColumn 字段数据重构
func (acd *AutoCodeService) GenColumn(v models.Column) (d models.Field) {

	d.ColumnName = v.ColumnName
	d.DataTypeLong = v.DataTypeLong
	d.DictType = v.DataType
	d.Comment = v.ColumnComment
	d.FieldDesc = v.ColumnComment + "表"
	d.FieldJson = utils.LeftUpper(utils.CamelString(v.ColumnName))
	d.FieldName = utils.LeftUpper(utils.CamelString(v.ColumnName))
	//d.FieldType = v.DataType
	fType := getType()
	d.FieldType = fType[d.DictType]
	return d
}
