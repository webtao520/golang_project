package main

import (
	"bytes"
	"reflect"
)

// 将结构体序列化为json并输出到缓冲
func writeStruct(buff *bytes.Buffer, value reflect.Value) error {

	// 取值的类型对象
	valueType := value.Type()

	// 写入结构体左大括号
	buff.WriteString("{")

	// 遍历结构体的所有值
	for i := 0; i < value.NumField(); i++ {

		// 获取每个字段的字段值(reflect.Value)
		fieldValue := value.Field(i)

		// 获取每个字段的类型(reflect.StructField)
		fieldType := valueType.Field(i)

		// 写入字段名左双引号
		buff.WriteString("\"")

		// 写入字段名
		buff.WriteString(fieldType.Name)

		// 写入字段名右双引号和冒号
		buff.WriteString("\":")

		// 写入每个字段值
		writeAny(buff, fieldValue)

		// 写入每个字段尾部逗号，最后一个字段不添加
		if i < value.NumField()-1 {
			buff.WriteString(",")
		}
	}

	// 写入结构体右大括号
	buff.WriteString("}")

	return nil
}
