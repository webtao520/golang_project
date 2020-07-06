package main 

import (
	"fmt"
)


func main (){
	str:="hello北京"
	// 字符串遍历，同时处理有中文的问题，r:=[]rune(str)
	r:=[]rune(str)
	for i:=0;i<len(r);i++ {
		fmt.Printf("字符=%c\n",r[i])
	}
	
}