/*
	从缓冲器写出
	WriteTo方法，将一个缓冲器的数据写到w里，w是实现io.Writer的，比如os.File
*/

package main

import (
	"bytes"
	"os"
)

func main() {
	file, _ := os.Create("text.txt")
	buf := bytes.NewBufferString("hello world")
	buf.WriteTo(file)
	//或者使用写入，fmt.Fprintf(file,buf.String())
}
