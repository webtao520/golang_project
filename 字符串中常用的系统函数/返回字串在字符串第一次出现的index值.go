package main 

import (
	"fmt"
	"strings"
)

func main (){
	// 返回子串在字符串第一次出现的index值，如果没有返回-1
   index:=strings.Index("GOLANG_abcbdcabc","abc")
   fmt.Printf("index=%v\n",index)	
}