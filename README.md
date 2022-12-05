# table

## 描述

+ `Go`的一个模块，用于生成漂亮的`ASCII`表格。

## 支持的字符集

+ `ASCII`
+ 中国文字

## 指南

* 下载库：
```bash
go get -u github.com/go-agumon/table 
```

* 创建表格：
```bash
t, err := table.Create("姓名", "性别", "年龄")
```

* 添加单行：
```bash
r1 := map[string]string{
	"姓名": "XXX",
	"性别": "男",
	"年龄": "27",
}
_ = t.AddRow(r1)

r2 := []string{"XXX", "女", "24"}
_ = t.AddRow(r2)
```

* 添加多行：
```bash
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

_ = t.AddRows(rows)
```

* 打印表格：
```bash
t.Print()
```

***
