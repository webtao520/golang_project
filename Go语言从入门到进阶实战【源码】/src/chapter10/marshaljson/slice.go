package main

import (
	"bytes"
	"reflect"
)

// 将切片转换为json并输出到缓冲
func writeSlice(buff *bytes.Buffer, value reflect.Value) error {

	// 写入切片开始标记
	buff.WriteString("[")

	// 遍历每个切片元素
	for s := 0; s < value.Len(); s++ {
		sliceValue := value.Index(s)

		// 写入每个切片元素
		writeAny(buff, sliceValue)

		// 写入每个元素尾部逗号，最后一个字段不添加
		if s < value.Len()-1 {
			buff.WriteString(",")
		}
	}

	// 写入切片结束标记
	buff.WriteString("]")

	return nil
}
