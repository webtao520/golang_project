/*
	WriteRune方法，将一个 rune / int32 类型的数据放到缓冲器的尾部

	// func (b *Buffer) WriteRune(r Rune) (n int,err error)
*/

package main

import (
	"bytes"
	"fmt"
)

func main() {
	var s rune = '好'
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String())
	buf.WriteRune(s)
	fmt.Println(buf.String())
}

/*
 PS D:\goLang\github\golang_project\package\bytes> go run WriteRune.go
				hello
				hello好
*/
