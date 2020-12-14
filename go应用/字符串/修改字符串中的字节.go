/**
修改字符串
在 Golang 中，不能修改字符串的内容，也就是说不能通过 s[i] 这种方式修改字符串中的字符。
要修改字符串的内容，可以先将字符串的内容复制到一个可写的变量中，一般是 []byte 或 []rune 类型的变量，然后再进行修改。
如果要对字符串中的字节进行修改，就转换为 []byte 类型，如果要对字符串中的字符修改，
就转换为 []rune 类型，在转换类型的过程中会自动复制数据。
*/

//修改字符串中的字节(用 []byte)
//对于那些单字节字符来说，可以通过这种方式进行修改：

package main

import "fmt"

func main() {
	s := "hello 世界"
	for _, r := range s {
		fmt.Printf("%T\n", r)
	}
	b := []byte(s) // 转换为[]byte ,数据被自动复制
	b[5] = ','     // 把空格改为半角逗号
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", b)
	fmt.Printf("%T\n", b[5]) // uint8 byte
}

/**
PS D:\goLang\github\golang_project\go应用\字符串> go run  修改字符串中的字节.go
hello 世界
hello,世界
*/
