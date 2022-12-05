package test

import (
	"testing"

	"github.com/go-agumon/table"
)

func TestTable(t *testing.T) {
	// 创建
	exampleTable, err := table.Create("姓名", "性别", "年龄")
	if err != nil {
		t.Errorf("unable to create table, err: %v", err)
	}

	// 添加数据
	rows := []map[string]string{
		{
			"姓名": "XXX",
			"性别": "男",
			"年龄": "27",
		},
		{
			"姓名": "XXX",
			"性别": "女",
			"年龄": "24",
		},
	}
	errs := exampleTable.AddRows(rows)
	if len(errs) != 0 {
		t.Errorf("unable to add rows to table, errs: %v", errs)
	}

	// 设置对齐方式
	exampleTable.SetAlign(table.AlignCenter)

	// 设置序列
	exampleTable.EnableSequence()

	// 打印表格
	exampleTable.Print()
}
