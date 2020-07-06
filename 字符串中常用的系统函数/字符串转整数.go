package main 

import (
	"fmt"
	"strconv"
)

func main (){
	b,error := strconv.Atoi("4545")
	if error != nil{
	fmt.Println("字符串转换成整数失败")
	}else {
	fmt.Printf("%T",b)
	}
}
