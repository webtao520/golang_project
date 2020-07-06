package main

import (
	"fmt"
	"strconv"
)

func main (){
	str:=strconv.FormatInt(123,2)
	//fmt.Println(str)
	fmt.Printf("123对应的二进制是= %v\n", str)
	str1:=strconv.FormatInt(123,16)
	fmt.Printf("123对应的16进制是=%v\n",str1)
}