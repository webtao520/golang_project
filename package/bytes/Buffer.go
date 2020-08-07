/*
一、创建缓冲期

bytes.buffer是一个缓冲byte类型的缓冲器

1、使用bytes.NewBuffer创建：参数是[]byte的话，缓冲器里就是这个slice的内容；如果参数是nil的话，就是创建一个空的缓冲器。

2、bytes.NewBufferString创建

3、bytes.Buffer{} 
*/

package main

import (
	"bytes"
	"fmt"
)

func main(){
   buf1 := bytes.NewBufferString("hello")
   buf2 := bytes.NewBuffer([]byte("hello"))
   buf3 := bytes.NewBuffer([]byte{'h','e','l','l','o'})
  // 以上三者等效,输出//hello
   buf4 := bytes.NewBufferString("")
   buf5 := bytes.NewBuffer([]byte{})
  // 以上两者等效,输出//""
   fmt.Println(buf1.String(),buf2.String(),buf3.String(),buf4,buf5,1)
}