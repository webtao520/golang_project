package main 

 import (
	"fmt"
	"strconv"
)
 
 func main () {
	var a int =1234
	str:= strconv.Itoa(a) // 整型转字符串
	// str:= strconv.Atoi(a) // 字符串转整型
    fmt.Printf("str = %T, str = %v", str,str)
 }