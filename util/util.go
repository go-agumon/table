package util

import "unicode"

const (
	chineseSymbol = "！……（），。？、" // 中文符号
)

// 判断是否为中文字符
func isChinese(c int32) bool {
	if unicode.Is(unicode.Han, c) {
		return true
	}

	for _, s := range chineseSymbol {
		if c == s {
			return true
		}
	}
	return false
}

// 获取字符串长度
func GetStringLength(s string) (length int) {
	for _, c := range s {
		if isChinese(c) {
			length += 2
		} else {
			length += 1
		}
	}

	return length
}

// 获取最大值
func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
