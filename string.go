package goutil

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

// strings.TrimSpace
// 去掉一个字符串左右的空白串，即（0x00 - 0x20 之内的字符均为空白字符）
func Trim(s string) string {
	size := len(s)
	if size <= 0 {
		return s
	}
	l := 0
	for ; l < size; l++ {
		b := s[l]
		if !IsSpace(b) {
			break
		}
	}
	r := size - 1
	for ; r >= l; r-- {
		b := s[r]
		if !IsSpace(b) {
			break
		}
	}
	return string(s[l : r+1])
}

// 是不是空字符
func IsSpace(c byte) bool {
	if c >= 0x00 && c <= 0x20 {
		return true
	}
	return false
}

// 判断一个字符串是不是空白串，即（0x00 - 0x20 之内的字符均为空白字符）
func IsBlank(s string) bool {
	for i := 0; i < len(s); i++ {
		b := s[i]
		if !IsSpace(b) {
			return false
		}
	}
	return true
}

// 字符串转换
func String2UTF8(_str string) string {
	var utf string
	b := []byte(_str)
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		utf += string(r)
		b = b[size:]
	}
	return utf
}

// 将字符串转换成整数，如果转换失败，采用默认值
func String2Int(s string, dft int) int {
	var re, err = strconv.Atoi(s)
	if err != nil {
		return dft
	}
	return re
}

// IsText reports whether a significant prefix of s looks like correct UTF-8;
// that is, if it is likely that s is human-readable text.
func IsText(s []byte) bool {
	const max = 1024 // at least utf8.UTFMax
	if len(s) > max {
		s = s[0:max]
	}
	for i, c := range string(s) {
		if i+utf8.UTFMax > len(s) {
			// last char may be incomplete - ignore
			break
		}
		if c == 0xFFFD || c < ' ' && c != '\n' && c != '\t' && c != '\f' {
			// decoding error or control character - not a text file
			return false
		}
	}
	return true
}

// 去掉字符串首位的双引号
func TrimQuotes(s string) string {
	var content string
	if strings.HasPrefix(s, `"`) {
		content = s[1:]
	}
	if strings.HasSuffix(s, `"`) {
		content = content[:len(content)-1]
	}
	return content
}

// 字符串转Float
func String2Float(data string, reData float64) float64 {
	i, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return reData
	}
	return i
}
