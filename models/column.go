package models

import (
	"table/util"
)

// 列
type Column struct {
	name         string // 名称
	defaultValue string // 默认值
	align        int    // 排列方式
	length       int    // 列名长度
}

// 创建列
func CreateColumn(name string) *Column {
	element := &Column{
		name:         name,
		defaultValue: "",
		align:        AlignCenter,
		length:       util.GetStringLength(name),
	}
	return element
}

// 获取列名
func (element *Column) String() string {
	return element.name
}

// 获取默认值
func (element *Column) Default() string {
	return element.defaultValue
}

// 设置默认值
func (element *Column) SetDefault(value string) {
	element.defaultValue = value
}

// 获取排列方式
func (element *Column) Align() int {
	return element.align
}

// 设置排列方式
func (element *Column) SetAlign(mode int) {
	switch mode {
	case AlignLeft:
		element.align = AlignLeft
	case AlignRight:
		element.align = AlignRight
	default:
		element.align = AlignCenter
	}
}

// 获取列名的长度
func (element *Column) Length() int {
	return element.length
}

// 判断列是否为完全相同
func (element *Column) Equal(other *Column) bool {
	functions := []func(other *Column) bool{
		element.nameEqual,
		element.defaultEqual,
		element.alignEqual,
		element.lengthEqual,
	}

	for _, function := range functions {
		if !function(other) {
			return false
		}
	}
	return true
}

// 判断列名是否相同
func (element *Column) nameEqual(other *Column) bool {
	return element.String() == other.String()
}

// 判断默认值是否相同
func (element *Column) defaultEqual(other *Column) bool {
	return element.Default() == other.Default()
}

// 判断排列方式是否相同
func (element *Column) alignEqual(other *Column) bool {
	return element.Align() == other.Align()
}

// 判断列长是否相同
func (element *Column) lengthEqual(other *Column) bool {
	return element.Length() == other.Length()
}
