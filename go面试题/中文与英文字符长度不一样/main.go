package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("abc张三你好啊", "你好")) // 9
	str := "abc张三你好啊"
	//  fmt.Println(strs)
	substr := "你好"
	fmt.Println(len([]rune(substr)))
	blen := strings.Index(str, substr)
	if blen >= 0 {
		prefix := []byte(str)[:blen]
		rs := []rune(string(prefix))
		blen = len(rs)
	}
	fmt.Println(blen) // 5

}

/**

func Index(s, sep string) int
子串sep在字符串s中第一次出现的位置，不存在则返回-1。

怎么输出字符串是6而不是9
fmt.Println(strings.Index("abc张三你好啊", "你好"))

中文与英文字符长度不一样,  strings.Index
只能取到字节的位置，可以这样。

byte 等同于int8，常用来处理ascii字符
rune 等同于int32,常用来处理 unicode 或 utf-8 字符， 非 ascii 字符
*/
