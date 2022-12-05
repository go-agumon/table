package models

// Cell 元件
type Cell interface {
	String() string
	Length() int
}

// 常量
const (
	Default = "__DEFAULT__" // 默认值
)

// 常量
const (
	AlignLeft   = iota // 左对齐
	AlignCenter        // 居中对齐
	AlignRight         // 右对齐
)
