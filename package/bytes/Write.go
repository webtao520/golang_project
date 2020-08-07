/*
	二、写入到缓冲器
	如果buffer在new的时候是空的，可以用Write在尾部写入

	1、Write方法，将一个byte类型的slice放到缓冲器的尾部

	//func (b *Buffer) Write(p []byte) (n int,err error)

*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := []byte(" world")
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String()) //hello
	buf.Write(s)              //将s这个slice添加到buf的尾部
	fmt.Println(buf.String()) //hello world
}

/*
	######################### 运行结果 ###########################
	PS D:\goLang\github\golang_project\package\bytes> go run .\Write.go
	hello
	hello world
*/
