package main 

import (
	"fmt"
	"strconv"
)

func main(){
	 fmt.Println("============数字转字符串============\n")
	 x:=123
	 y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) // "123 123"
	fmt.Printf("%T\n",strconv.Itoa(x))

	fmt.Println("============字符串转数字============\n")
	x,err:=strconv.Atoi("123")
	if  err!=nil{

	}else{
		fmt.Printf("%T\n",x)
	}

	// y, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits 返回字符串表示的整数值，接受正负号。

}