package stringx

import (
	"bytes"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"unicode"
)

// WhiteSpace 定义了常见的空白字符，包括换行、制表符、换页、垂直制表符和空格
var WhiteSpace = []rune{'\n', '\t', '\f', '\v', ' '}

// String 提供了将原始文本转换为其他格式（如小写、蛇形命名、驼峰命名）的功能
type String struct {
	source string // 原始字符串内容
}

// From 将输入的文本转换为 String 类型并返回
func From(data string) String {
	return String{source: data}
}

// IsEmptyOrSpace 判断字符串在去除空白字符后是否为空，若为空则返回 true，否则返回 false
func (s String) IsEmptyOrSpace() bool {
	if len(s.source) == 0 {
		return true
	}
	if strings.TrimSpace(s.source) == "" {
		return true
	}
	return false
}

// Lower 调用 strings.ToLower 将字符串转换为小写形式
func (s String) Lower() string {
	return strings.ToLower(s.source)
}

// Upper 调用 strings.ToUpper 将字符串转换为大写形式
func (s String) Upper() string {
	return strings.ToUpper(s.source)
}

// ReplaceAll 调用 strings.Replace 替换字符串中所有匹配的子串
func (s String) ReplaceAll(old, new string) string {
	return strings.Replace(s.source, old, new, -1)
}

// Source 返回原始字符串内容
func (s String) Source() string {
	return s.source
}

// ToCamel 将输入文本转换为驼峰命名格式（Camel Case）
// 例如 "hello_world" 转换为 "HelloWorld"
func (s String) ToCamel() string {
	// 以 "_" 为分隔符拆分字符串，并移除分隔符
	list := s.splitBy(func(r rune) bool {
		return r == '_'
	}, true)
	var target []string
	for _, item := range list {
		target = append(target, From(item).Title())
	}
	return strings.Join(target, "")
}

// ToSnake 将输入文本转换为蛇形命名格式（Snake Case）
// 例如 "HelloWorld" 转换为 "hello_world"
func (s String) ToSnake() string {
	// 以大写字母为分隔符拆分字符串，不移除匹配的字符
	list := s.splitBy(unicode.IsUpper, false)
	var target []string
	for _, item := range list {
		target = append(target, From(item).Lower())
	}
	return strings.Join(target, "_")
}

// Title 调用 strings.Title 将字符串的每个单词的首字母转换为大写，同时将其他字母转换为小写。
func (s String) Title() string {
	if s.IsEmptyOrSpace() {
		return s.source
	}
	// 创建一个针对英文的 Title 转换器
	caser := cases.Title(language.English)
	return caser.String(s.source)
}

// UnTitle 如果字符串的首字符为字母，则将其转换为小写，其余部分保持不变；否则返回原字符串
func (s String) UnTitle() string {
	if s.IsEmptyOrSpace() {
		return s.source
	}
	r := rune(s.source[0])
	if !unicode.IsUpper(r) && !unicode.IsLower(r) {
		return s.source
	}
	return string(unicode.ToLower(r)) + s.source[1:]
}

// TitleFirstLetter 将字符串的首字符转换为大写，其余部分保持不变；如果首字符不是字母，则返回原字符串
func (s String) TitleFirstLetter() string {
	if len(s.source) == 0 {
		return s.source
	}
	// 获取首字符
	firstChar := rune(s.source[0])
	// 判断首字符是否为字母
	if !unicode.IsLetter(firstChar) {
		return s.source
	}
	// 将首字符转换为大写
	return string(unicode.ToUpper(firstChar)) + s.source[1:]
}

// splitBy 根据指定的函数 fn 拆分字符串
// remove 参数指定是否移除拆分时遇到的分隔符
func (s String) splitBy(fn func(r rune) bool, remove bool) []string {
	if s.IsEmptyOrSpace() {
		return nil
	}
	var list []string
	buffer := new(bytes.Buffer)
	for _, r := range s.source {
		if fn(r) {
			if buffer.Len() != 0 {
				list = append(list, buffer.String())
				buffer.Reset()
			}
			if !remove {
				buffer.WriteRune(r)
			}
			continue
		}
		buffer.WriteRune(r)
	}
	if buffer.Len() != 0 {
		list = append(list, buffer.String())
	}
	return list
}

// ContainsAny 判断字符串 s 是否包含任一指定的 rune 字符
func ContainsAny(s string, runes ...rune) bool {
	if len(runes) == 0 {
		return true
	}
	tmp := make(map[rune]struct{}, len(runes))
	for _, r := range runes {
		tmp[r] = struct{}{}
	}
	for _, r := range s {
		if _, ok := tmp[r]; ok {
			return true
		}
	}
	return false
}

// ContainsWhiteSpace 判断字符串 s 是否包含任一空白字符（如换行、制表符等）
func ContainsWhiteSpace(s string) bool {
	return ContainsAny(s, WhiteSpace...)
}

// IsEmptyOrSpace 判断字符串在去除空白字符后是否为空，若为空则返回 true，否则返回 false
func IsEmptyOrSpace(s string) bool {
	if len(s) == 0 {
		return true
	}
	if strings.TrimSpace(s) == "" {
		return true
	}
	return false
}

// Lower 将字符串转换为小写形式
func Lower(s string) string {
	return strings.ToLower(s)
}

// Upper 将字符串转换为大写形式
func Upper(s string) string {
	return strings.ToUpper(s)
}

// ReplaceAll 替换字符串中所有匹配的子串
func ReplaceAll(s, old, new string) string {
	return strings.Replace(s, old, new, -1)
}

// ToCamel 将输入文本转换为驼峰命名格式（Camel Case）
// 例如 "hello_world" 转换为 "HelloWorld"
func ToCamel(s string) string {
	// 以 "_" 为分隔符拆分字符串，并移除分隔符

	list := From(s).splitBy(func(r rune) bool {
		return r == '_'
	}, true)
	var target []string
	for _, item := range list {
		target = append(target, Title(item))
	}
	return strings.Join(target, "")
}

// ToSnake 将输入文本转换为蛇形命名格式（Snake Case）
// 例如 "HelloWorld" 转换为 "hello_world"
func ToSnake(s string) string {
	// 以大写字母为分隔符拆分字符串，不移除匹配的字符
	list := From(s).splitBy(unicode.IsUpper, false)
	var target []string
	for _, item := range list {
		target = append(target, Lower(item))
	}
	return strings.Join(target, "_")
}

// Title 将字符串的每个单词的首字母转换为大写，同时将其他字母转换为小写。
func Title(s string) string {
	if IsEmptyOrSpace(s) {
		return s
	}
	// 创建一个针对英文的 Title 转换器
	caser := cases.Title(language.English)
	return caser.String(s)
}

// UnTitle 如果字符串的首字符为字母，则将其转换为小写，其余部分保持不变；否则返回原字符串
func UnTitle(s string) string {
	if IsEmptyOrSpace(s) {
		return s
	}
	r := rune(s[0])
	if !unicode.IsUpper(r) && !unicode.IsLower(r) {
		return s
	}
	return string(unicode.ToLower(r)) + s[1:]
}

// TitleFirstLetter 将字符串的首字符转换为大写，其余部分保持不变；如果首字符不是字母，则返回原字符串
func TitleFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	// 获取首字符
	firstChar := rune(s[0])
	// 判断首字符是否为字母
	if !unicode.IsLetter(firstChar) {
		return s
	}
	// 将首字符转换为大写
	return string(unicode.ToUpper(firstChar)) + s[1:]
}
