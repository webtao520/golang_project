/*
	WriteByte方法，将一个byte类型 / uint8别名的数据放到缓冲器的尾部

	//func (b *Buffer) WriteByte(c byte) error
*/

package main

import (
	"bytes"
	"fmt"
)

func main() {
	var s byte
	s = '?'
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String()) //把buf的内容转换为string，hello
	buf.WriteByte(s)          //将s写到buf的尾部
	fmt.Println(buf.String())
}

/*
	PS D:\goLang\github\golang_project\package\bytes> go run WriteByte.go
	hello
	hello?
*/
