package models

import (
	"table/util"
)

// 行
type Row struct {
	value  string // 数据
	length int    // 数据长度
}

// 创建行
func CreateRow(value string) *Row {
	element := &Row{
		value:  value,
		length: util.GetStringLength(value),
	}
	return element
}

// 创建空行
func CreateEmptyRow() *Row {
	return CreateRow("")
}

// 获取数据
func (element *Row) String() string {
	return element.value
}

// 获取数据的长度
func (element *Row) Length() int {
	return element.length
}
