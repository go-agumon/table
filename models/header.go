package models

import (
	"fmt"
)

// 表头
type Header struct {
	columns []*Column // 列切片
}

// 获取列长
func (element *Header) Len() int {
	return len(element.columns)
}

// 获取容量
func (element *Header) Cap() int {
	return cap(element.columns)
}

// 判断列名是否存在
func (element *Header) Exist(name string) bool {
	return element.exist(name) != -1
}

// 判断列名是否存在，若存在则返回索引
func (element *Header) exist(name string) int {
	for i, column := range element.columns {
		if column.String() == name {
			return i
		}
	}
	return -1
}

// 清除表头
func (element *Header) Clear() {
	element.columns = make([]*Column, 0)
}

// 添加列
func (element *Header) Add(name string) error {
	if element.Exist(name) {
		return fmt.Errorf("column [%s] has exist", name)
	}

	column := CreateColumn(name)
	element.columns = append(element.columns, column)
	return nil
}

// 移除列
func (element *Header) Remove(name string) error {
	position := element.exist(name)
	if position == -1 {
		return fmt.Errorf("column [%s] has not exit", name)
	}

	element.columns = append(element.columns[:position], element.columns[position+1:]...)
	return nil
}

// 获取列
func (element *Header) Get(name string) *Column {
	for _, column := range element.columns {
		if column.String() == name {
			return column
		}
	}
	return nil
}

// 判断表头是否相同
func (element *Header) Equal(other *Header) bool {
	if element.Len() != other.Len() {
		return false
	}

	// 用多协程比较每一列
	channels := make(chan bool)
	for index := range element.columns {
		go func(position int) {
			if !element.columns[position].Equal(other.columns[position]) {
				channels <- false
			} else {
				channels <- true
			}
		}(index)
	}

	// 循环判断
	count := 0
	for {
		select {
		case equal := <-channels:
			count += 1
			// 当任意一个不同，则为假
			if !equal {
				return false
			}

			if count >= element.Len() {
				return true
			}
		}
	}
}
