package table

import (
	"fmt"
	"table/models"
)

// 常量
const (
	Default     = models.Default     // 默认值
	AlignLeft   = models.AlignLeft   // 左对齐
	AlignCenter = models.AlignCenter // 居中对齐
	AlignRight  = models.AlignRight  // 右对齐
)

func Create(columns ...string) (*models.Table, error) {
	if len(columns) <= 0 {
		return nil, fmt.Errorf("columns length must more than zero")
	}

	// 创建表头
	header := new(models.Header)
	for _, column := range columns {
		err := header.Add(column)
		if err != nil {
			return nil, err
		}
	}

	// 创建表格
	table := models.CreateTable(header)

	return table, nil
}
