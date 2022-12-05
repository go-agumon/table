package models

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/go-agumon/table/util"
)

// 表格
type Table struct {
	header   *Header           // 表头
	rows     []map[string]*Row // 行
	border   bool              // 边框
	sequence bool              // 序列
}

// 创建表格
func CreateTable(header *Header) *Table {
	return &Table{
		header: header,
		rows:   make([]map[string]*Row, 0),
		border: true,
	}
}

// 清除表格
func (element *Table) Clear() {
	element.header.Clear()
	element.rows = make([]map[string]*Row, 0)
}

// 获取表格长度
func (element *Table) Length() int {
	return len(element.rows)
}

// 判断表格是否为空
func (element *Table) Empty() bool {
	return element.Length() == 0
}

// 启用边框
func (element *Table) EnableBorder() {
	element.border = true
}

// 禁用边框
func (element *Table) DisableBorder() {
	element.border = false
}

// 启用序列
func (element *Table) EnableSequence() {
	element.sequence = true
}

// 禁用序列
func (element *Table) DisableSequence() {
	element.sequence = false
}

// 设置指定列的排列方式
func (element *Table) SetSingleAlign(column string, mode int) {
	for _, v := range element.header.columns {
		if v.String() == column {
			v.SetAlign(mode)
			return
		}
	}
}

// 设置表格的排列方式
func (element *Table) SetAlign(mode int) {
	for _, v := range element.header.columns {
		v.SetAlign(mode)
	}
}

// 判断是否有指定列
func (element *Table) HasColumn(column string) bool {
	for _, v := range element.header.columns {
		if v.String() == column {
			return true
		}
	}
	return false
}

// 添加列
func (element *Table) AddColumns(columns ...string) error {
	for _, column := range columns {
		err := element.header.Add(column)
		if err != nil {
			return err
		}
	}

	for _, column := range columns {
		for _, row := range element.rows {
			row[column] = CreateEmptyRow()
		}
	}
	return nil
}

// 插入列
func (element *Table) InsertColumn(column string, values []string, index int) error {
	// 向表头插入列
	if err := element.header.Insert(column, index); err != nil {
		return err
	}

	// 需要保证值的长度与行的长度一致
	if len(values) != len(element.rows) {
		return errors.New("the given values does not match the length of the rows")
	}

	// 向每一列插入数据
	for i, item := range element.rows {
		item[column] = CreateRow(values[i])
	}

	return nil
}

// 设置列的默认值
func (element *Table) SetDefault(column string, defaultValue string) {
	for _, v := range element.header.columns {
		if v.String() == column {
			v.SetDefault(defaultValue)
			break
		}
	}
}

// 删除列的默认值
func (element *Table) DropDefault(column string) {
	element.SetDefault(column, "")
}

// 获取列的默认值
func (element *Table) GetDefault(column string) string {
	for _, v := range element.header.columns {
		if v.String() == column {
			return v.Default()
		}
	}
	return ""
}

// 获取所有列的默认值
func (element *Table) GetDefaults() map[string]string {
	defaults := make(map[string]string)
	for _, v := range element.header.columns {
		defaults[v.String()] = v.Default()
	}
	return defaults
}

// 添加行 - 批量
func (element *Table) AddRows(rows []map[string]string) []map[string]string {
	failure := make([]map[string]string, 0)
	for _, row := range rows {
		err := element.AddRow(row)
		if err != nil {
			failure = append(failure, row)
		}
	}
	return failure
}

// 添加行 - 单行
func (element *Table) AddRow(row interface{}) error {
	switch v := row.(type) {
	case []string:
		return element.addRowFromSlice(v)
	case map[string]string:
		return element.addRowFromMap(v)
	default:
		return fmt.Errorf("unsupported row type: %T", row)
	}
}

// 添加行 - 切片
func (element *Table) addRowFromSlice(row []string) error {
	// 获取长度
	rowLength := len(row)
	if rowLength != element.header.Len() {
		return fmt.Errorf("the length of row(%d) does not equal the columns(%d)", rowLength, element.header.Len())
	}

	// 将列名与列进行一一对应
	item := make(map[string]*Row)
	for i := 0; i < rowLength; i++ {
		// 列值
		var value = ""
		// 判断列值是否需要设置为默认值
		if row[i] == Default {
			value = element.header.columns[i].Default()
		} else {
			value = row[i]
		}
		item[element.header.columns[i].String()] = CreateRow(value)
	}

	element.rows = append(element.rows, item)
	return nil
}

// 添加行 - Map
func (element *Table) addRowFromMap(row map[string]string) error {
	for key := range row {
		if !element.header.Exist(key) {
			return fmt.Errorf("column %s do not exist", key)
		}

		// 设置默认值
		if row[key] == Default {
			row[key] = element.header.Get(key).Default()
		}
	}

	// 若未设置，则自动取默认值
	for _, column := range element.header.columns {
		_, ok := row[column.String()]
		if !ok {
			row[column.String()] = column.Default()
		}
	}

	item := make(map[string]*Row)
	for k, v := range row {
		item[k] = CreateRow(v)
	}

	element.rows = append(element.rows, item)
	return nil
}

// 获取表头
func (element *Table) GetHeader() []string {
	header := make([]string, 0)
	for _, column := range element.header.columns {
		header = append(header, column.String())
	}
	return header
}

// 获取所有行
func (element *Table) GetRows() []map[string]string {
	rows := make([]map[string]string, 0)
	for _, row := range element.rows {
		item := make(map[string]string)
		for k, v := range row {
			item[k] = v.String()
		}
		rows = append(rows, item)
	}
	return rows
}

// 将Table转换为Slice
func (element *Table) ToStringSlice() []string {
	// 表格Slice
	tableSlice := make([]string, 0)

	// 分隔行
	splitRow := make(map[string]*Row)

	// 获取列的最大长度
	columnMaxLength := make(map[string]int)
	for _, column := range element.header.columns {
		// 假设当前列的最大长度是表头
		columnMaxLength[column.String()] = column.Length()
		for _, row := range element.rows {
			columnMaxLength[column.String()] = util.Max(columnMaxLength[column.String()], row[column.String()].Length())
		}
		splitRow[column.String()] = CreateRow("-")
	}

	// 判断是否有边框
	if element.border {
		// 添加首行
		tableSlice = append(tableSlice, element.rowsToStringSlice([]map[string]*Row{splitRow}, columnMaxLength)...)
	}

	// 添加表头
	headerRow := make(map[string]*Row)
	for _, column := range element.header.columns {
		headerRow[column.String()] = CreateRow(column.String())
	}
	tableSlice = append(tableSlice, element.rowsToStringSlice([]map[string]*Row{headerRow}, columnMaxLength)...)

	if element.border {
		// 添加分隔行
		tableSlice = append(tableSlice, element.rowsToStringSlice([]map[string]*Row{splitRow}, columnMaxLength)...)
	}
	// 打印数据行
	if !element.Empty() {
		tableSlice = append(tableSlice, element.rowsToStringSlice(element.rows, columnMaxLength)...)
	}

	// 判断是否有边框
	if element.border {
		// 添加尾行
		tableSlice = append(tableSlice, element.rowsToStringSlice([]map[string]*Row{splitRow}, columnMaxLength)...)
	}
	return tableSlice
}

func (element *Table) Print() {
	// 判断表格是否开启序列
	if element.sequence {
		values := make([]string, len(element.rows))
		for i := 1; i <= len(element.rows); i++ {
			values = append(values, strconv.Itoa(i))
		}
		_ = element.InsertColumn("序列", values, 0)
	}

	tableSlice := element.ToStringSlice()
	for _, line := range tableSlice {
		fmt.Println(line)
	}
}

func (element *Table) rowsToStringSlice(rows []map[string]*Row, columnMaxLength map[string]int) []string {
	rowSlice := make([]string, 0)

	for _, row := range rows {
		line := ""
		for i, column := range element.header.columns {
			// 获取当前列的最大长度
			itemLength := columnMaxLength[column.String()]

			if element.border {
				itemLength += 2
			}

			// 每一列的字符串
			item := ""
			if row[column.String()].String() == "-" {
				if element.border {
					item, _ = center(row[column.String()], itemLength, "-")
				}
			} else {
				switch column.Align() {
				case AlignRight:
					item, _ = right(row[column.String()], itemLength, " ")
				case AlignLeft:
					item, _ = left(row[column.String()], itemLength, " ")
				default:
					item, _ = center(row[column.String()], itemLength, " ")
				}
			}

			icon := "|"
			// 判断是否为首行
			if row[column.String()].String() == "-" {
				icon = "+"
			}

			// 判断是否有边框
			if !element.border {
				icon = " "
			}

			if i == 0 {
				item = fmt.Sprintf("%s%s%s", icon, item, icon)
			} else {
				item = fmt.Sprintf("%s%s", item, icon)
			}
			line += item
		}
		rowSlice = append(rowSlice, line)
	}
	return rowSlice
}

// 填充 - 居中对齐
func center(row *Row, length int, fillChar string) (string, error) {
	if len(fillChar) != 1 {
		err := fmt.Errorf("the fill character must be exactly one character long")
		return "", err
	}

	if row.Length() >= length {
		return row.String(), nil
	}

	result := ""
	if isEvenNumber(length - row.Length()) {
		front := ""
		for i := 0; i < ((length - row.Length()) / 2); i++ {
			front = front + fillChar
		}

		result = front + row.String() + front
	} else {
		front := ""
		for i := 0; i < ((length - row.Length() - 1) / 2); i++ {
			front = front + fillChar
		}

		behind := front + fillChar
		result = front + row.String() + behind
	}
	return result, nil
}

// 填充 - 左对齐
func left(row *Row, length int, fillChar string) (string, error) {
	if len(fillChar) != 1 {
		err := fmt.Errorf("the fill character must be exactly one character long")
		return "", err
	}

	result := row.String() + block(length-row.Length())
	return result, nil
}

// 填充 - 右对齐
func right(row *Row, length int, fillChar string) (string, error) {
	if len(fillChar) != 1 {
		err := fmt.Errorf("the fill character must be exactly one character long")
		return "", err
	}

	result := block(length-row.Length()) + row.String()
	return result, nil
}

// 填充 - 空格
func block(length int) string {
	result := ""
	for i := 0; i < length; i++ {
		result += " "
	}
	return result
}

// 判断是否为偶数
func isEvenNumber(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}
