package format

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/qiaogw/gocode/common/stringx"
	"io"
	"strings"
)

const (
	// flagGo 用于标识 Go 部分的格式
	flagGo = "GO"
	// flagZero 用于标识 Zero 部分的格式
	flagZero = "ZERO"

	// 以下定义格式转换的样式常量
	unknown style = iota // 未知样式
	title                // 首字母大写样式
	lower                // 全部小写样式
	upper                // 全部大写样式
)

// ErrNamingFormat 定义了一个不支持的格式错误
var ErrNamingFormat = errors.New("unsupported format")

type (
	// styleFormat 用于描述文件命名格式的各部分及对应的样式
	styleFormat struct {
		before    string // 格式化结果前缀
		through   string // 格式化过程中用于连接各部分的分隔符
		after     string // 格式化结果后缀
		goStyle   style  // Go 格式的样式
		zeroStyle style  // Zero 格式的样式
	}

	// style 定义了格式转换的样式类型
	style int
)

// FileNamingFormat 用于格式化文件名。你可以通过指定 go 和 zero 格式化字符来定义格式样式，
// 例如可以将蛇形命名定义为 go_zero，将驼峰命名定义为 goZero，甚至可以指定分隔符，例如 go#Zero，
// 理论上可以使用任意组合，但前提是必须满足各操作系统文件命名的规定。
// 注意：格式化基于蛇形或驼峰字符串。
func FileNamingFormat(format, content string) (string, error) {
	upperFormat := strings.ToUpper(format)
	indexGo := strings.Index(upperFormat, flagGo)
	indexZero := strings.Index(upperFormat, flagZero)
	if indexGo < 0 || indexZero < 0 || indexGo > indexZero {
		return "", ErrNamingFormat
	}
	var (
		before, through, after string
		flagGo, flagZero       string
		goStyle, zeroStyle     style
		err                    error
	)
	// 提取格式字符串中的各部分
	before = format[:indexGo]
	flagGo = format[indexGo : indexGo+2]
	through = format[indexGo+2 : indexZero]
	flagZero = format[indexZero : indexZero+4]
	after = format[indexZero+4:]

	// 根据标识获取对应的样式
	goStyle, err = getStyle(flagGo)
	if err != nil {
		return "", err
	}

	zeroStyle, err = getStyle(flagZero)
	if err != nil {
		return "", err
	}
	var formatStyle styleFormat
	formatStyle.goStyle = goStyle
	formatStyle.zeroStyle = zeroStyle
	formatStyle.before = before
	formatStyle.through = through
	formatStyle.after = after
	// 根据格式化样式对内容进行格式化
	return doFormat(formatStyle, content)
}

// doFormat 根据传入的 styleFormat 对内容进行格式化处理，返回格式化后的字符串
func doFormat(f styleFormat, content string) (string, error) {
	// 将内容拆分为多个部分
	splits, err := split(content)
	if err != nil {
		return "", err
	}
	var join []string
	// 对拆分后的各部分分别进行样式转换
	for index, split := range splits {
		if index == 0 {
			join = append(join, transferTo(split, f.goStyle))
			continue
		}
		join = append(join, transferTo(split, f.zeroStyle))
	}
	// 用指定的连接符将各部分合并，并添加前缀和后缀
	joined := strings.Join(join, f.through)
	return f.before + joined + f.after, nil
}

// transferTo 根据指定样式将字符串转换为对应格式
func transferTo(in string, style style) string {
	switch style {
	case upper:
		return strings.ToUpper(in)
	case lower:
		return strings.ToLower(in)
	case title:
		return stringx.Title(in)
	default:
		return in
	}
}

// split 将输入内容拆分为若干子串，拆分依据为下划线和大写字母的出现
func split(content string) ([]string, error) {
	var (
		list   []string
		reader = strings.NewReader(content)
		buffer = bytes.NewBuffer(nil)
	)
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				if buffer.Len() > 0 {
					list = append(list, buffer.String())
				}
				return list, nil
			}
			return nil, err
		}
		// 如果遇到下划线，则将当前缓冲区内容作为一个子串保存，并重置缓冲区
		if r == '_' {
			if buffer.Len() > 0 {
				list = append(list, buffer.String())
			}
			buffer.Reset()
			continue
		}

		// 如果遇到大写字母，并且缓冲区已有内容，则将当前缓冲区内容保存为一个子串，并重置缓冲区
		if r >= 'A' && r <= 'Z' {
			if buffer.Len() > 0 {
				list = append(list, buffer.String())
			}
			buffer.Reset()
		}
		buffer.WriteRune(r)
	}
}

// getStyle 根据格式标识返回对应的样式（lower、upper 或 title），不匹配则返回 unknown
func getStyle(flag string) (style, error) {
	compare := strings.ToLower(flag)
	switch flag {
	case strings.ToLower(compare):
		return lower, nil
	case strings.ToUpper(compare):
		return upper, nil
	case stringx.Title(compare):
		return title, nil
	default:
		return unknown, fmt.Errorf("unexpected format: %s", flag)
	}
}
