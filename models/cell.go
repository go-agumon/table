package models

// 元件
type Cell interface {
	String() string
	Length() int
}

// 默认值
const (
	Default = "__DEFAULT__"
)

const (
	AlignLeft   = iota // 左对齐
	AlignCenter        // 居中对齐
	AlignRight         // 右对齐
)
