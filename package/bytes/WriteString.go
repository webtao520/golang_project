/*
	WriteString方法，把一个字符串放到缓冲器的尾部
	//func (b *Buffer) WriteString(s string)(n int,err error)
*/

package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := " world"
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String())
	buf.WriteString(s) //将string写入到buf的尾部
	fmt.Println(buf.String())
}

/*
	################ 运行结果 ####################
		PS D:\goLang\github\golang_project\package\bytes> go run  .\WriteString.go
		hello
		hello world
*/
