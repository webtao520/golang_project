package main

import (
	"bytes"
	"errors"
	"reflect"
	"strconv"
)

// 将任意值转换为json并输出到缓冲
func writeAny(buff *bytes.Buffer, value reflect.Value) error {

	switch value.Kind() {
	case reflect.String:
		// 写入带有双引号括起来的字符串
		buff.WriteString(strconv.Quote(value.String()))
	case reflect.Int:
		// 将整形转换为字符串并写入缓冲
		buff.WriteString(strconv.FormatInt(value.Int(), 10))
	case reflect.Slice:
		return writeSlice(buff, value)
	case reflect.Struct:
		return writeStruct(buff, value)
	default:
		// 遇到不认识的种类，返回错误
		return errors.New("unsupport kind: " + value.Kind().String())
	}

	return nil
}
